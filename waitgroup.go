package maxworkers

import (
	"sync"
)

type WaitGroup struct {
	maxParallel uint64
	concurrency chan bool
	waitgroup   sync.WaitGroup
}

func NewGroup(maxParallel uint64) *WaitGroup {
	return &WaitGroup{
		maxParallel: maxParallel,
		concurrency: make(chan bool, maxParallel),
		waitgroup:   sync.WaitGroup{},
	}
}

func (wg *WaitGroup) Run(fn func()) {
	wg.concurrency <- true
	wg.waitgroup.Add(1)
	go func() {
		defer func() { wg.waitgroup.Done(); <-wg.concurrency }()
		fn()
	}()
}

func (wg *WaitGroup) Wait() {
	wg.waitgroup.Wait()
}
