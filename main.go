package main

import (
	"flag"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool             *redis.Pool
	queues_to_ignore []string
	queues_to_watch  []string
)

func initGlobals() {
	pool = newRedisPool(*redis_host + ":" + *redis_port)
	queues_to_ignore = strings.Split(*str_queues_to_ignore, ",")
	if *str_queues_to_watch != "" {
		queues_to_watch = strings.Split(*str_queues_to_watch, ",")
	}
}

func main() {
	flag.Parse()
	initGlobals()

	poller := NewPoller(pool)
	poller.Poll()

	queue_jobs := NewQueueJobs()
	queue_jobs.store(poller.out)
	queue_jobs.print(time.Tick(1 * time.Second))

	for {
		queues := getQueues(pool)
		queues.SendToPoller(poller)
		Sleep(*poll_interval)
	}

}

func Sleep(poll_interval string) {
	pi, _ := strconv.Atoi(poll_interval)
	time.Sleep(time.Millisecond * time.Duration(pi))
}

func getQueues(pool *redis.Pool) Queues {
	conn := pool.Get()
	defer conn.Close()
	var tmp []string
	if len(queues_to_watch) > 0 {
		tmp = make([]string, len(queues_to_watch))
		for i := 0; i < len(queues_to_watch); i++ {
			tmp[i] = "resque:queue:" + queues_to_watch[i]
		}
	} else {
		tmp, _ = redis.Strings(conn.Do("smembers", "resque:queues"))
	}
	return sliceStringsToSliceQueues(removeIgnoredQueues(tmp))
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
