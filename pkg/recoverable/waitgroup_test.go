package recoverable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaitGroup_Wait(t *testing.T) {
	as := assert.New(t)

	t.Run("no panic", func(t *testing.T) {
		var wg WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		as.NoError(wg.Wait())
	})

	t.Run("two panics", func(t *testing.T) {
		var wg WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			panic("panic 1")
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			panic("panic 2")
		}()
		err := wg.Wait()
		as.Error(err)
		for _, e := range err.(*PanicError).Errors() {
			t.Logf("%+v", e)
		}
	})
}
