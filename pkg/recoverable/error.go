package recoverable

import (
	"fmt"
	"github.com/pkg/errors"
)

func ToError(r interface{}) error {
	return errors.New(fmt.Sprintf("%v", r))
}
