package recoverable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunner(t *testing.T) {
	as := assert.New(t)
	var i int
	Runner(func() {
		for i < 3 {
			panic("panic")
		}
	}, func(r interface{}) {
		i++
		t.Log(r)
	})
	as.Equal(3, i)
}
