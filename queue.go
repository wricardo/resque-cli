package main

import "fmt"

type Queue string

func (q Queue) String() string {
	return fmt.Sprintf("%s", string(q))
}

func (this *Queue) getCompleteName() string {
	return "resque:queue:" + fmt.Sprintf("%s", string(*this))
}
