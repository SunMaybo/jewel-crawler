package limit

import "github.com/SunMaybo/jewel-crawler/task"

type ConcurrentLimit struct {
	total         int
	activeChannel chan byte
}

func NewConcurrentLimit(total int) *ConcurrentLimit {
	return &ConcurrentLimit{
		total:         total,
		activeChannel: make(chan byte, total),
	}
}

func (lc *ConcurrentLimit) Acquire(task task.Task, function func(task task.Task)) {
	lc.activeChannel <- byte(0)
	go function(task)
}
func (lc *ConcurrentLimit) Free() {
	<-lc.activeChannel
}
