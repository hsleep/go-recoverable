package recoverable

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunner(t *testing.T) {
	as := assert.New(t)
	var i int
	Runner(func() {
		for i < 3 {
			panic("panic")
		}
	}, func(e error) {
		i++
		fmt.Println(e)
	})
	as.Equal(3, i)
}
