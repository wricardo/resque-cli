package main

import (
	"github.com/garyburd/redigo/redis"
)

type Poller struct {
	queue_to_poll chan QueuesJobs
	redis_pool    *redis.Pool
	out           chan QueuesJobs
}

func NewPoller(redis_pool *redis.Pool) *Poller {
	p := new(Poller)
	p.queue_to_poll = make(chan QueuesJobs)
	p.redis_pool = redis_pool
	p.out = make(chan QueuesJobs)
	return p
}

func (this *Poller) Poll() {
	conn := this.redis_pool.Get()
	defer conn.Close()
	go func() {
		for {
			select {
			case q := <-this.queue_to_poll:
				q.Poll(conn)
				this.out <- q
			}
		}
	}()
}
