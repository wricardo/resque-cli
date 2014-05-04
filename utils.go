package main

func sliceStringsToSliceQueues(strings []string) Queues {
	queues := make(Queues, len(strings))
	for x := 0; x < len(strings); x++ {
		queues[x] = Queue(strings[x])
	}
	return queues
}

func removeIgnoredQueues(queues []string) []string {
	for x := 0; x < len(queues); x++ {
		if stringInSlice(queues[x], queues_to_ignore) {
			queues = append(queues[:x], queues[x+1:]...)
		}
	}
	return queues
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getMapKeysSorted(to_sort map[Queue]QueueJob) []Queue {
	a := make([]Queue, len(to_sort))
	var x int
	x = 0
	for q, _ := range to_sort {
		a[x] = q
		x = x + 1
	}
	for itemCount := len(a) - 1; ; itemCount-- {
		hasChanged := false
		for index := 0; index < itemCount; index++ {
			if a[index] > a[index+1] {
				a[index], a[index+1] = a[index+1], a[index]
				hasChanged = true
			}
		}
		if hasChanged == false {
			break
		}
	}
	return a
}
