package main

import (
	"sync"
)

type (
	Fn   func(...interface{}) interface{}
	Work struct {
		Max    int
		F      Fn
		TaskCh chan Fn
		once   *sync.Once
	}
)

// fill in goroutine pool
func Pool(max int) *Work {
	work := &Work{
		Max:    max,
		TaskCh: make(chan Fn, max),
		once:   &sync.Once{},
	}
	return work
}

//Receive  value  chan
func (w *Work) NewPool(args ...interface{}) {
	for i := 0; i < w.Max; i++ {
		w.makePool(args)
	}
}

func (w *Work) makePool(args interface{}) {
	go func() {
		for task := range w.TaskCh {
			task(args)
		}
	}()
}

//Back into that the pool
func (w *Work) Put(args ...interface{}) {
	w.makePool(args)
}

//Execute Call
func (w *Work) Do(f Fn) {
	w.TaskCh <- f
}

//Shut down the goroutine pool
func (w *Work) Close() {
	w.once.Do(func() {
		close(w.TaskCh)
	})
}
