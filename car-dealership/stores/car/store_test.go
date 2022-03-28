package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/filters"
	"github.com/zopsmart/GoLang-Interns-2022/models"
	"github.com/zopsmart/GoLang-Interns-2022/types/brand"
	"github.com/zopsmart/GoLang-Interns-2022/types/fuel"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// testInitializer function creates a mock store for tests to call
func testInitializer(t *testing.T) (*sql.DB, sqlmock.Sqlmock, store) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Errorf("error in sql mock")
	}

	s := New(db)

	return db, mock, s
}

// TestStore_Create function checks the test cases for Create method of store layer for car
func TestStore_Create(t *testing.T) {
	db, mock, s := testInitializer(t)
	defer db.Close()

	id := uuid.New()

	input := models.Car{ID: id, Name: "X-200", Year: 2020, Brand: brand.BMW, Fuel: fuel.Petrol, Engine: models.Engine{ID: id}}

	cases := []struct {
		desc string
		id   uuid.UUID
		err  error
	}{
		{"Success", input.ID, nil},
		{"Query error", uuid.Nil, errors2.DB{Err: errors.New("query error")}},
	}

	mock.ExpectExec("INSERT INTO cars (id, name, year, brand, fuel, engine_id) VALUES (?,?,?,?,?,?)").
		WithArgs(input.ID, input.Name, input.Year, input.Brand, input.Fuel, input.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO cars (id, name, year, brand, fuel, engine_id) VALUES (?,?,?,?,?,?)").
		WithArgs(input.ID, input.Name, input.Year, input.Brand, input.Fuel, input.ID).
		WillReturnError(errors.New("query error"))

	for _, tc := range cases {
		err := s.Create(context.TODO(), &input)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// TestStore_GetAll function checks the test cases for GetAll method of store layer for car
func TestStore_GetAll(t *testing.T) {
	db, mock, s := testInitializer(t)

	defer db.Close()

	cars := []models.Car{
		{ID: uuid.New(), Name: "ABC", Year: 2020, Brand: brand.Mercedes, Fuel: fuel.Diesel},
	}

	filter := filters.Car{Brand: "Mercedes"}

	mock.ExpectQuery("SELECT * FROM cars WHERE brand = ?").WithArgs(filter.Brand).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "year", "brand", "fuel", "engine_id"}).
			AddRow(cars[0].ID.String(), cars[0].Name, cars[0].Year, []byte("mercedes"), []byte("diesel"), cars[0].Engine.ID.String()))

	mock.ExpectQuery("SELECT * FROM cars WHERE brand = ?").WithArgs(filter.Brand).
		WillReturnError(errors.New("query error"))

	mock.ExpectQuery("SELECT * FROM cars WHERE brand = ?").WithArgs(filter.Brand).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "year", "brand", "fuel", "engine_id", "extra"}).
			AddRow(cars[0].ID.String(), cars[0].Name, cars[0].Year, []byte("mercedes"), []byte("diesel"), cars[0].Engine.ID.String(), "scan error"))

	mock.ExpectQuery("SELECT * FROM cars WHERE brand = ?").WithArgs(filter.Brand).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "year", "brand", "fuel", "engine_id"}).
			AddRow(cars[0].ID.String(), cars[0].Name, cars[0].Year, []byte("mercedes"), []byte("diesel"), cars[0].Engine.ID).
			RowError(0, errors.New("row error")))

	cases := []struct {
		desc   string
		output []models.Car
		err    error
	}{
		{"Success", cars, nil},
		{"Query error", nil, errors2.DB{Err: errors.New("query error")}},
		{"Scan error", nil, errors2.DB{Err: fmt.Errorf("sql: expected %d destination arguments in Scan, not %d", 7, 6)}},
		{"Row error", nil, errors2.DB{Err: errors.New("row error")}},
	}

	for _, tc := range cases {
		output, err := s.GetAll(context.TODO(), filter)

		assert.Equal(t, tc.output, output, tc.desc)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// TestStore_GetByID function checks the test cases for GetByID method of store layer for car
func TestStore_GetByID(t *testing.T) {
	db, mock, s := testInitializer(t)

	defer db.Close()

	id := uuid.New()

	input := models.Car{ID: id, Name: "ABC", Year: 2020, Brand: brand.Mercedes, Fuel: fuel.Diesel}

	mock.ExpectQuery("SELECT * FROM cars WHERE id = ?").WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "year", "brand", "fuel", "engine_id"}).
			AddRow(input.ID.String(), input.Name, input.Year, []byte("mercedes"), []byte("diesel"), input.Engine.ID.String()))

	mock.ExpectQuery("SELECT * FROM cars WHERE id = ?").WithArgs(id.String()).WillReturnError(errors.New("query error"))

	mock.ExpectQuery("SELECT * FROM cars WHERE id = ?").WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "year", "brand", "fuel"}).
			RowError(8, errors2.EntityNotFound{Entity: "car", ID: id}))

	cases := []struct {
		desc   string
		output models.Car
		err    error
	}{
		{"Success", input, nil},
		{"Query error", models.Car{}, errors2.DB{Err: errors.New("query error")}},
		{"No row error", models.Car{}, errors2.EntityNotFound{Entity: "car", ID: id}},
	}

	for _, tc := range cases {
		output, err := s.GetByID(context.TODO(), id)

		assert.Equal(t, tc.output, output, tc.desc)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// TestStore_Update function checks the test cases for Update method of store layer for car
func TestStore_Update(t *testing.T) {
	db, mock, s := testInitializer(t)

	defer db.Close()

	id := uuid.New()

	input := models.Car{ID: id, Name: "X-200", Year: 2020, Brand: brand.BMW, Fuel: fuel.Petrol, Engine: models.Engine{ID: id}}

	cases := []struct {
		desc string
		err  error
	}{
		{"Success", nil},
		{"Failure", errors2.DB{Err: errors.New("update failed")}},
		{"No row affected", errors2.DB{Err: errors.New("zero rows affected")}},
	}

	mock.ExpectExec("UPDATE cars SET name=?, year=?, brand=?, fuel=?, engine_id=? WHERE id=?").
		WithArgs(input.Name, input.Year, input.Brand, input.Fuel, input.Engine.ID, input.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE cars SET name=?, year=?, brand=?, fuel=?, engine_id=? WHERE id=?").
		WithArgs(input.Name, input.Year, input.Brand, input.Fuel, input.Engine.ID, input.ID).
		WillReturnError(errors.New("update failed"))

	mock.ExpectExec("UPDATE cars SET name=?, year=?, brand=?, fuel=?, engine_id=? WHERE id=?").
		WithArgs(input.Name, input.Year, input.Brand, input.Fuel, input.Engine.ID, input.ID).
		WillReturnError(errors.New("zero rows affected"))

	for _, tc := range cases {
		_, err := s.Update(context.TODO(), id, &input)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// TestStore_Delete function checks the test cases for Delete method of store layer for car
func TestStore_Delete(t *testing.T) {
	db, mock, s := testInitializer(t)

	defer db.Close()

	id := uuid.New()

	cases := []struct {
		desc string
		err  error
	}{
		{"Success", nil},
		{"Failure", errors2.DB{Err: errors.New("delete failed")}},
		{"No row affected", errors2.DB{Err: errors.New("zero rows affected")}},
	}

	mock.ExpectExec("DELETE FROM cars WHERE id = ?").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("DELETE FROM cars WHERE id = ?").WithArgs(id).WillReturnError(errors.New("delete failed"))

	mock.ExpectExec("DELETE FROM cars WHERE id = ?").WithArgs(id).WillReturnError(errors.New("zero rows affected"))

	for _, tc := range cases {
		err := s.Delete(context.TODO(), id)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}
