package main

type Queues []Queue

func (this *Queues) SendToPoller(poller *Poller) {
	for _, queue := range *this {
		poller.queue_to_poll <- queue
	}
}
