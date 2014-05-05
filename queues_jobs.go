package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type QueuesJobs struct {
	data       map[Queue]QueueJob
	get_values chan bool
	response   chan map[Queue]QueueJob
}

func NewQueueJobs() *QueuesJobs {
	obj := new(QueuesJobs)
	obj.data = make(map[Queue]QueueJob)
	obj.get_values = make(chan bool)
	obj.response = make(chan map[Queue]QueueJob)
	return obj
}

func (this *QueuesJobs) _store(q QueueJob) {
	this.data[q.name] = q
}

func (this *QueuesJobs) print() {
	fmt.Println("\033[2J")
	fmt.Println("\033[H")
	n := this.data
	sorted := getMapKeysSorted(n)
	for _, queue_name := range sorted {
		fmt.Println(queue_name, n[queue_name].jobs)
	}
}

func (this *QueuesJobs) SendToPoller(poller *Poller) {
	poller.queue_to_poll <- *this
}

func (this *QueuesJobs) Poll(conn redis.Conn) {
	for k, qj := range this.data {
		qj.Poll(conn)
		this.data[k] = qj
	}
}
