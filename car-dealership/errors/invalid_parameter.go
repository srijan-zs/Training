package errors

import (
	"fmt"
	"strings"
)

// InvalidParam struct is defined for the invalid parameter encountered errors
type InvalidParam struct {
	Params []string
}

// Error method specifies the format of the error returned according to its types
func (i InvalidParam) Error() string {
	return fmt.Sprintf("invalid parameters %s", strings.Join(i.Params, ", "))
}
