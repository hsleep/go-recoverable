package recoverable

import (
	"sync"
)

type WaitGroup struct {
	wg   sync.WaitGroup
	ch   chan error
	once sync.Once
}

func (wg *WaitGroup) Add(n int) {
	wg.once.Do(func() {
		wg.ch = make(chan error)
	})
	wg.wg.Add(n)
}

func (wg *WaitGroup) Done() {
	defer wg.wg.Done()
	if r := recover(); r != nil {
		wg.ch <- toError(r)
	}
}

func (wg *WaitGroup) Wait() {
	go func() {
		wg.wg.Wait()
		close(wg.ch)
	}()
	var errs []error
	for err := range wg.ch {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		panic(errs)
	}
}
