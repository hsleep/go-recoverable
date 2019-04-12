# Go-recoverable

## WaitGroup

It has similar interface with `sync.WaitGroup`.
But it will call `recover()` function in `WaitGroup.Done()` function to recover panic in goroutine.
And `WaitGroup.Wait()` function will return `PanicError` if panic occurs.

**DO NOT reuse `WaitGroup` after `WaitGroup.Wait()` is called.**

## PanicError

When panic occurs in a goroutine, deferred `wg.Done()` function recover the panic and send an error to internal channel.
`WaitGroup.Wait()` function return a `PanicError` if the channel receive an error.
If multiple panic errors occur, all errors can be accessed by `PanicError.Errors()` function.

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