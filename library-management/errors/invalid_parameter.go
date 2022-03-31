package errors

import (
	"fmt"
	"strings"
)

type InvalidParam struct {
	Params []string
}

func (i InvalidParam) Error() string {
	return fmt.Sprintf("invalid parameters %s", strings.Join(i.Params, ", "))
}
