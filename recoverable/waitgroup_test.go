package recoverable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaitGroup_Wait(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		as := assert.New(t)

		var wg WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		as.NotPanics(wg.Wait)
	})

	t.Run("two panics", func(t *testing.T) {
		as := assert.New(t)

		as.Panics(func() {
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
			wg.Wait()
		})
	})
}
