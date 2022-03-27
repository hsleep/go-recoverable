package recoverable

// Runner will run f function, and rerun f if panic occurs in f
func Runner(f func(), panicHandler func(r interface{})) {
	running := true
	for running {
		func() {
			defer func() {
				if r := recover(); r != nil {
					panicHandler(r)
				}
			}()
			f()
			running = false
		}()
	}
}
