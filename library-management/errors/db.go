package errors

import "fmt"

type DB struct {
	Err error
}

func (db DB) Error() string {
	return fmt.Sprintf("database error: %s", db.Err)
}
