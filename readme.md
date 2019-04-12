# Go-recoverable

## WaitGroup

It has similar interface with `sync.WaitGroup`.
But it will call `recover()` function in `wg.Done()` function to recover panic in goroutine.
And `wg.Wait()` function will return `PanicError` if panic occurs.

**DO NOT re-use `WaitGroup` after `wg.Wait()` is called.**

## Example
```go
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
```