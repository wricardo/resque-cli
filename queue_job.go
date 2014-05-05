package main

import "github.com/garyburd/redigo/redis"

type QueueJob struct {
	name Queue
	jobs int
}

func NewQueueJob(name Queue) *QueueJob {
	i := new(QueueJob)
	i.name = name
	i.jobs = 0
	return i
}

func (this *QueueJob) Poll(conn redis.Conn) {
	llen, err := redis.Int(conn.Do("llen", this.name.getCompleteName()))
	if err != nil {
		panic(err)
	}
	this.jobs = llen
}
