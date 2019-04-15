# Go-recoverable

## WaitGroup

It has similar interface with `sync.WaitGroup`.
But it will call `recover()` function in `WaitGroup.Done()` function to recover panic in goroutine.
And `WaitGroup.Wait()` function will raise panic if any panic occurs in `WaitGroup.Done()`.
The panic by raised `WaitGroup.Wait()` can be recovered by caller function.

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
		as.NotPanics(wg.Wait)
	})

	t.Run("two panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				if sr, ok := r.([]error); ok {
					for _, err := range sr {
						fmt.Printf("%+v\n", err)
					}
				}
			}
		}()
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
}
```