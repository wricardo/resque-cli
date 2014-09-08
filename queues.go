package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Queues []Queue

func (this Queues) PrintCountJobs(conn redis.Conn) error {
	for _, q := range this {
		jobs, err := q.CountJobs(conn)
		if err != nil {
			return err
		}
		fmt.Println(q+":", jobs)
	}
	return nil
}

func GetQueues(conn redis.Conn) (Queues, error) {
	to_return := make(Queues, 0)
	queues, err := redis.Strings(conn.Do("smembers", "resque:queues"))
	if err != nil {
		return to_return, err
	}
	for _, q := range queues {
		to_return = append(to_return, Queue(q))
	}
	return to_return, nil
}

type Queue string

func (this Queue) CountJobs(conn redis.Conn) (int, error) {
	return redis.Int(conn.Do("llen", "resque:queue:"+this))
}
