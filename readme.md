# Go-recoverable

## WaitGroup

It has similar interface with `sync.WaitGroup`.
But it will call `recover()` function in `WaitGroup.Done()` function to recover panic in goroutine.
And `WaitGroup.Wait()` function will return `[]error` if panic occurs.

**DO NOT reuse `WaitGroup` after `WaitGroup.Wait()` is called.**

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
		as.Empty(wg.Wait())
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
		errs := wg.Wait()
		as.NotEmpty(errs)
		for _, e := range errs {
			t.Logf("%+v", e)
		}
	})
}
```