package recoverable

import (
	"sync"
)

type WaitGroup struct {
	wg   sync.WaitGroup
	ch   chan interface{}
	once sync.Once
}

func (wg *WaitGroup) Add(n int) {
	wg.once.Do(func() {
		wg.ch = make(chan interface{})
	})
	wg.wg.Add(n)
}

func (wg *WaitGroup) Done() {
	defer wg.wg.Done()
	if r := recover(); r != nil {
		wg.ch <- r
	}
}

func (wg *WaitGroup) Wait() []error {
	go func() {
		wg.wg.Wait()
		close(wg.ch)
	}()
	var errs []error
	for r := range wg.ch {
		errs = append(errs, ToError(r))
	}
	return errs
}
