package recoverable

import (
	"fmt"
)

func toError(r interface{}) error {
	return fmt.Errorf("%v", r)
}
