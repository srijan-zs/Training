package engine

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/models"

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

// TestStore_Create function checks the test cases for Create method of store layer for engine
func TestStore_Create(t *testing.T) {
	db, mock, s := testInitializer(t)

	defer db.Close()

	input := models.Engine{ID: uuid.Nil, Displacement: 2000, NCylinders: 4}

	cases := []struct {
		desc string
		err  error
	}{
		{"Success", nil},
		{"Query error", errors2.DB{Err: errors.New("query error")}},
	}

	mock.ExpectExec("INSERT INTO engines (id, displacement, nCylinder, extent) VALUES (?,?,?,?)").
		WithArgs(sqlmock.AnyArg(), input.Displacement, input.NCylinders, input.Range).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO engines (id, displacement, nCylinder, extent) VALUES (?,?,?,?)").
		WithArgs(sqlmock.AnyArg(), input.Displacement, input.NCylinders, input.Range).
		WillReturnError(errors.New("query error"))

	for _, tc := range cases {
		_, err := s.Create(context.TODO(), &input)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// TestStore_GetByID function checks the test cases for GetByID method of store layer for engine
func TestStore_GetByID(t *testing.T) {
	db, mock, s := testInitializer(t)

	defer db.Close()

	id := uuid.New()

	cases := []struct {
		desc   string
		output models.Engine
		err    error
	}{
		{"Success", models.Engine{}, nil},
		{"Query error", models.Engine{}, errors2.DB{Err: errors.New("query error")}},
		{"No row error", models.Engine{}, errors2.DB{Err: errors.New("sql: no rows in result set")}},
	}

	output := models.Engine{}

	mock.ExpectQuery("SELECT * FROM engines WHERE id = ?").WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "displacement", "nCylinder", "extent"}).
			AddRow(output.ID, output.Displacement, output.NCylinders, output.Range))

	mock.ExpectQuery("SELECT * FROM engines WHERE id = ?").WithArgs(id.String()).WillReturnError(errors.New("query error"))

	mock.ExpectQuery("SELECT * FROM engines WHERE id = ?").WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "displacement", "nCylinder", "extent"}).
			RowError(4, errors.New("sql: no rows in result set")))

	for _, tc := range cases {
		output, err := s.GetByID(context.TODO(), id)

		assert.Equal(t, tc.output, output, tc.desc)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// TestStore_Update function checks the test cases for Update method of store layer for engine
func TestStore_Update(t *testing.T) {
	db, mock, s := testInitializer(t)

	defer db.Close()

	id := uuid.New()

	input := models.Engine{ID: id, Range: 450}

	cases := []struct {
		desc string
		err  error
	}{
		{"Success", nil},
		{"Failure", errors2.DB{Err: errors.New("update failed")}},
		{"No row affected", errors2.DB{Err: errors.New("zero rows affected")}},
	}

	mock.ExpectExec("UPDATE engines SET displacement=?, nCylinder=?, extent=? WHERE id = ?").
		WithArgs(input.Displacement, input.NCylinders, input.Range, input.ID.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE engines SET displacement=?, nCylinder=?, extent=? WHERE id = ?").
		WithArgs(input.Displacement, input.NCylinders, input.Range, input.ID.String()).
		WillReturnError(errors.New("update failed"))

	mock.ExpectExec("UPDATE engines SET displacement=?, nCylinder=?, extent=? WHERE id = ?").
		WithArgs(input.Displacement, input.NCylinders, input.Range, input.ID.String()).
		WillReturnError(errors.New("zero rows affected"))

	for _, tc := range cases {
		_, err := s.Update(context.TODO(), id, &input)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// TestStore_Delete function checks the test cases for Delete method of store layer for engine
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

	mock.ExpectExec("DELETE FROM engines WHERE id = ?").WithArgs(id.String()).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("DELETE FROM engines WHERE id = ?").WithArgs(id.String()).WillReturnError(errors.New("delete failed"))

	mock.ExpectExec("DELETE FROM engines WHERE id = ?").WithArgs(id.String()).WillReturnError(errors.New("zero rows affected"))

	for _, tc := range cases {
		err := s.Delete(context.TODO(), id)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}
