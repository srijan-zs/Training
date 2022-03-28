package errors

import (
	"fmt"
	"strings"
)

// MissingParam struct is defined for the missing parameter encountered errors
type MissingParam struct {
	Params []string
}

// Error method specifies the format of the error returned according to its types
func (m MissingParam) Error() string {
	return fmt.Sprintf("missing parameters %s", strings.Join(m.Params, ", "))
}
