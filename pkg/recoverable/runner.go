package recoverable

// run f function, and if panic occurs rerun f
func Runner(f func(), panicHandler func(e error)) {
	running := true
	for running {
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicHandler(toError(r))
				}
			}()
			f()
			running = false
		}()
	}
}
