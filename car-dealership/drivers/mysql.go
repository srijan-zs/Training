package drivers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // sql driver
)

// Connection is the driver function which connects to mysql database of the service
func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "car_dealership:srastogi@zopsmart@tcp(localhost:3306)/car_dealership")
	if err != nil {
		return nil, err
	}

	return db, nil
}
