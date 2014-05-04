package main

import (
	"fmt"
	"time"
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

func (this *QueuesJobs) store(qj chan QueueJob) {
	go func() {
		for {
			select {
			case q := <-qj:
				this.data[q.name] = q
			case <-this.get_values:
				this.response <- this.data
			}
		}
	}()
}

func (this *QueuesJobs) print(tick <-chan time.Time) {
	go func() {
		for _ = range tick {
			fmt.Println("\033[2J")
			fmt.Println("\033[H")
			this.get_values <- true
			n := <-this.response
			sorted := getMapKeysSorted(n)
			for _, queue_name := range sorted {
				fmt.Println(queue_name, n[queue_name].jobs)
			}
		}
	}()
}
