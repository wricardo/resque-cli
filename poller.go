package main

import (
	"github.com/garyburd/redigo/redis"
)

type Poller struct {
	queue_to_poll chan Queue
	redis_pool    *redis.Pool
	out           chan QueueJob
}

func NewPoller(redis_pool *redis.Pool) *Poller {
	p := new(Poller)
	p.queue_to_poll = make(chan Queue)
	p.redis_pool = redis_pool
	p.out = make(chan QueueJob)
	return p
}

func (this *Poller) Poll() {
	conn := this.redis_pool.Get()
	defer conn.Close()
	go func() {
		for {
			select {
			case q := <-this.queue_to_poll:
				llen, err := redis.Int(conn.Do("llen", "resque:queue:"+q))
				if err != nil {
					panic(err)
				}
				this.out <- QueueJob{
					name: q,
					jobs: llen,
				}
			}
		}
	}()
}
