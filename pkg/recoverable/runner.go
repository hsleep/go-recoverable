package recoverable

// Runner will run f function, and rerun f if panic occurs in f
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
