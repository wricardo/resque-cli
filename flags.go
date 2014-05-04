package main

import "flag"

var (
	redis_host           = flag.String("h", "localhost", "Redis host")
	redis_port           = flag.String("p", "6379", "Redis port")
	poll_interval        = flag.String("i", "1000", "Poll interval in milliseconds")
	str_queues_to_ignore = flag.String("v", "", "Queues to ignore. Separated by coma")
	str_queues_to_watch  = flag.String("only", "", "Queues to watch. Separared by coma")
)
