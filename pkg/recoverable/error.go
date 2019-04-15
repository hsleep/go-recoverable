package recoverable

import (
	"fmt"
	"github.com/pkg/errors"
)

func toError(r interface{}) error {
	return errors.New(fmt.Sprintf("%v", r))
}
