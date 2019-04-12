package recoverable

import (
	"fmt"
	"github.com/pkg/errors"
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
		wg.ch <- errors.New(fmt.Sprintf("%v", r))
	}
}

func (wg *WaitGroup) Wait() error {
	go func() {
		wg.wg.Wait()
		close(wg.ch)
	}()
	var errs []error
	for err := range wg.ch {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return &PanicError{errs: errs}
	}
	return nil
}

type PanicError struct {
	errs []error
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("panic errors: %v", e.errs)
}

func (e *PanicError) Errors() []error {
	return e.errs
}
