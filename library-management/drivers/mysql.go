package drivers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "library_management:srastogi@zopsmart@tcp(localhost:3307)/library_management")
	if err != nil {
		return nil, err
	}

	return db, nil
}
