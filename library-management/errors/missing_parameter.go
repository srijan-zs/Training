package errors

import (
	"fmt"
	"strings"
)

type MissingParam struct {
	Params []string
}

func (m MissingParam) Error() string {
	return fmt.Sprintf("missing parameters %s", strings.Join(m.Params, ", "))
}
