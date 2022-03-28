package errors

import "fmt"

// DB struct is defined for database errors
type DB struct {
	Err error
}

// Error method specifies the format of the error returned according to its types
func (db DB) Error() string {
	return fmt.Sprintf("database error: %s", db.Err)
}
