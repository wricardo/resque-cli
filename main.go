package main

import (
	"github.com/codegangsta/cli"
	"github.com/garyburd/redigo/redis"
	"log"
	"os"
	"time"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "resque-cli"
	app.Usage = "Monitor your resque queues from the cli"
	app.Commands = []cli.Command{
		{
			Name:  "watch",
			Usage: "options for task templates",
			Subcommands: []cli.Command{
				{
					Name:   "queues",
					Usage:  "Watch how many jobs you have on the resque queues",
					Action: watchResqueQueues,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "host",
							Value: "localhost",
							Usage: "Redis host",
						},
						cli.StringFlag{
							Name:  "port",
							Value: "6379",
							Usage: "Redis port",
						},
						cli.StringFlag{
							Name:  "i",
							Value: "1s",
							Usage: "Refresh interval. ex: 500ms|1s|1m",
						},
					},
				},
			},
		},
	}

	app.Run(os.Args)
}

func watchResqueQueues(c *cli.Context) {
	pool := newRedisPool(c.String("host") + ":" + c.String("port"))
	i, err := time.ParseDuration(c.String("i"))
	if err != nil {
		log.Fatalln("Invalid parameter \"i\"")
	}
	ticker := time.NewTicker(i)

	for _ = range ticker.C {
		clearScreen()
		conn := pool.Get()
		defer conn.Close()
		queues, err := GetQueues(conn)
		if err != nil {
			log.Fatalln(err)
		}
		queues.PrintCountJobs(conn)
	}
}

func newRedisPool(server string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func clearScreen(){
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
}
