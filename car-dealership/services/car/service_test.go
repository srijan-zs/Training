package car

import (
	"context"
	"testing"

	errors2 "github.com/zopsmart/GoLang-Interns-2022/errors"
	"github.com/zopsmart/GoLang-Interns-2022/filters"
	"github.com/zopsmart/GoLang-Interns-2022/models"
	"github.com/zopsmart/GoLang-Interns-2022/services"
	"github.com/zopsmart/GoLang-Interns-2022/types/brand"
	"github.com/zopsmart/GoLang-Interns-2022/types/fuel"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// initializeTest generates the mocks for creating mock calls to test the defined methods
func initializeTest(t *testing.T) (*services.MockCar, *services.MockEngine, service) {
	ctrl := gomock.NewController(t)

	mockCar := services.NewMockCar(ctrl)
	mockEngine := services.NewMockEngine(ctrl)

	svc := New(mockCar, mockEngine)

	return mockCar, mockEngine, svc
}

// engine is created as a model for testing functions to use as a sample data
var engine = models.Engine{ // nolint:gochecknoglobals // variable declared as global to reduce redundancy in tests
	Displacement: 2000,
	NCylinders:   4,
}

// car is created as a model for testing functions to use as a sample data
var car = models.Car{ // nolint:gochecknoglobals // variable declared as global to reduce redundancy in tests
	Name:   "D-280",
	Year:   2019,
	Brand:  brand.BMW,
	Fuel:   fuel.Petrol,
	Engine: engine,
}

// TestService_Create function checks the test cases for Create method of service layer
func TestService_Create(t *testing.T) {
	id := uuid.New()

	car.ID = id
	engine.ID = id

	mockCar, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Create(gomock.Any(), gomock.Any()).Return(uuid.Nil, nil)
	mockCar.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)

	mockCar.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(car, nil)
	mockEngine.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(engine, nil)

	resp, err := svc.Create(context.TODO(), &car)
	car.Engine = engine

	assert.Equal(t, nil, err)

	assert.Equal(t, &car, resp)
}

// Test_CreateInvalidCar function checks the test cases for invalid car data
func Test_CreateInvalidCar(t *testing.T) {
	_, _, svc := initializeTest(t)

	resp, err := svc.Create(context.TODO(), &models.Car{})

	assert.Equal(t, errors2.InvalidParam{Params: []string{"name"}}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: invalid car model\nGot: %v\n Expected: %v", resp, nil)
	}
}

// Test_CreateInvalidEngine function checks the test cases for invalid engine data
func Test_CreateInvalidEngine(t *testing.T) {
	car1 := car
	car1.Engine = models.Engine{Displacement: -2500}

	_, _, svc := initializeTest(t)

	resp, err := svc.Create(context.TODO(), &car1)

	assert.Equal(t, errors2.InvalidParam{Params: []string{"displacement"}}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: invalid engine parameter\nGot: %v\n Expected: %v", resp, nil)
	}
}

// Test_CreateEngineDBError function check the case for database error while creating engine
func Test_CreateEngineDBError(t *testing.T) {
	_, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Create(gomock.Any(), gomock.Any()).Return(uuid.Nil, errors2.DB{})

	resp, err := svc.Create(context.TODO(), &car)

	assert.Equal(t, errors2.DB{}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: db error when creating engine\nGot: %v\n Expected: %v", resp, &car)
	}
}

// Test_CreateCarVerificationError function checks the case where created car is not verified in the database
func Test_CreateCarVerificationError(t *testing.T) {
	mockCar, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Create(gomock.Any(), gomock.Any()).Return(uuid.Nil, nil)
	mockCar.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	mockCar.EXPECT().GetByID(gomock.Any(), car.ID).Return(models.Car{}, errors2.DB{})

	resp, err := svc.Create(context.TODO(), &car)

	assert.Equal(t, errors2.DB{}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: create failed\nGot: %v\n Expected: %v", resp, &car)
	}
}

// Test_CreateEngineVerificationError function checks the case where created engine is not verified in the database
func Test_CreateEngineVerificationError(t *testing.T) {
	mockCar, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Create(gomock.Any(), gomock.Any()).Return(uuid.Nil, nil)
	mockCar.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
	mockCar.EXPECT().GetByID(gomock.Any(), car.ID).Return(car, nil)
	mockEngine.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Engine{}, errors2.DB{})

	resp, err := svc.Create(context.TODO(), &car)

	assert.Equal(t, errors2.DB{}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: create failed\nGot: %v\n Expected: %v", resp, &car)
	}
}

// Test_CreateCarDBError function checks the case for database error while creating car
func Test_CreateCarDBError(t *testing.T) {
	mockCar, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Create(gomock.Any(), gomock.Any()).Return(uuid.Nil, nil)
	mockCar.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors2.DB{})

	resp, err := svc.Create(context.TODO(), &car)

	assert.Equal(t, errors2.DB{}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: create failed\nGot: %v\n Expected: %v", resp, &car)
	}
}

// TestService_GetAll function checks the test cases for read all cars with specified brand and engine
func TestService_GetAll(t *testing.T) {
	cars := []models.Car{car}

	mockCarStore, mockEngineStore, svc := initializeTest(t)

	mockCarStore.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(cars, nil)
	mockEngineStore.EXPECT().GetByID(gomock.Any(), car.ID).Return(engine, nil)

	output, err := svc.GetAll(context.TODO(), filters.Car{Brand: "mercedes", IncludeEngine: true})

	assert.Equal(t, nil, err)

	assert.Equal(t, cars, output)
}

// Test_GetAllWithEngineDBError function checks the case for database error while reading car with engine being included
func Test_GetAllWithEngineDBError(t *testing.T) {
	cars := []models.Car{car}

	mockCarStore, mockEngineStore, svc := initializeTest(t)

	mockCarStore.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(cars, nil)
	mockEngineStore.EXPECT().GetByID(gomock.Any(), car.ID).Return(engine, errors2.DB{})

	output, err := svc.GetAll(context.TODO(), filters.Car{Brand: "mercedes", IncludeEngine: true})

	assert.Equal(t, errors2.DB{}, err)

	assert.Equal(t, cars, output)
}

// Test_GetAllWithoutEngineDBError function checks the case for database error while reading car with engine not being included
func Test_GetAllWithoutEngineDBError(t *testing.T) {
	mockCarStore, _, svc := initializeTest(t)

	mockCarStore.EXPECT().GetAll(gomock.Any(), gomock.Any()).Return(nil, errors2.DB{})

	output, err := svc.GetAll(context.TODO(), filters.Car{Brand: "mercedes", IncludeEngine: false})

	assert.Equal(t, errors2.DB{}, err)

	if output != nil {
		t.Errorf("[TEST] Failed \nDesc: get failed, DB error\nGot: %v\n Expected: %v", output, nil)
	}
}

// TestService_GetByID function checks the test cases for extracting the car of a specific id
func TestService_GetByID(t *testing.T) {
	mockCarStore, mockEngineStore, svc := initializeTest(t)

	mockCarStore.EXPECT().GetByID(gomock.Any(), car.ID).Return(car, nil)
	mockEngineStore.EXPECT().GetByID(gomock.Any(), car.ID).Return(engine, nil)

	car.Engine = engine

	output, err := svc.GetByID(context.TODO(), car.ID)

	assert.Equal(t, nil, err)

	assert.Equal(t, car, output)
}

// Test_GetByIDInvalidCar function check the case for extracting the non-existing car
func Test_GetByIDInvalidCar(t *testing.T) {
	mockCar, _, svc := initializeTest(t)

	mockCar.EXPECT().GetByID(gomock.Any(), car.ID).Return(car, errors2.EntityNotFound{})

	car.Engine = engine
	resp, err := svc.GetByID(context.TODO(), car.ID)

	assert.Equal(t, errors2.EntityNotFound{}, err)

	assert.Equal(t, models.Car{}, resp)
}

// Test_GetByIDInvalidEngine function check the case for extracting the non-existing engine
func Test_GetByIDInvalidEngine(t *testing.T) {
	mockCar, mockEngine, svc := initializeTest(t)

	mockCar.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(car, nil)
	mockEngine.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(models.Engine{}, errors2.EntityNotFound{})

	car.Engine = engine
	resp, err := svc.GetByID(context.TODO(), car.Engine.ID)

	assert.Equal(t, errors2.EntityNotFound{}, err)

	assert.Equal(t, models.Car{}, resp)
}

// TestService_Update function checks the test cases for Update method of service layer
func TestService_Update(t *testing.T) {
	mockCarStore, mockEngineStore, svc := initializeTest(t)

	mockEngineStore.EXPECT().Update(gomock.Any(), car.ID, &car.Engine).Return(&car.Engine, nil)
	mockCarStore.EXPECT().Update(gomock.Any(), car.ID, &car).Return(&car, nil)

	output, err := svc.Update(context.TODO(), car.ID, &car)

	assert.Equal(t, nil, err)

	assert.Equal(t, &car, output)
}

// Test_UpdateInvalidParam function checks the case for invalid data given for modification
func Test_UpdateInvalidParam(t *testing.T) {
	car1 := car
	car1.Year = 2030

	_, _, svc := initializeTest(t)

	resp, err := svc.Update(context.TODO(), car1.ID, &car1)

	assert.Equal(t, errors2.InvalidParam{Params: []string{"year"}}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: update failed, parameter is invalid\nGot: %v\n Expected: %v", resp, car)
	}
}

// Test_UpdateInvalidEngine function checks the case for updating the non-existing engine
func Test_UpdateInvalidEngine(t *testing.T) {
	_, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Update(gomock.Any(), gomock.Any(), &engine).Return(&car.Engine, errors2.EntityNotFound{})

	resp, err := svc.Update(context.TODO(), car.ID, &car)

	assert.Equal(t, errors2.EntityNotFound{}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: update failed, engine is invalid\nGot: %v\n Expected: %v", resp, car)
	}
}

// Test_UpdateInvalidCar function checks the case for updating the non-existing car
func Test_UpdateInvalidCar(t *testing.T) {
	mockCar, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Update(gomock.Any(), gomock.Any(), &engine).Return(&car.Engine, nil)
	mockCar.EXPECT().Update(gomock.Any(), gomock.Any(), &car).Return(&car, errors2.EntityNotFound{})

	resp, err := svc.Update(context.TODO(), car.ID, &car)

	assert.Equal(t, errors2.EntityNotFound{}, err)

	if resp != nil {
		t.Errorf("[TEST] Failed \nDesc: update failed, car is invalid\nGot: %v\n Expected: %v", resp, car)
	}
}

// TestService_Delete function checks the test cases for Delete method of service layer
func TestService_Delete(t *testing.T) {
	id := uuid.New()

	mockCarStore, mockEngineStore, svc := initializeTest(t)

	mockEngineStore.EXPECT().Delete(gomock.Any(), id).Return(nil)
	mockCarStore.EXPECT().Delete(gomock.Any(), id).Return(nil)

	err := svc.Delete(context.TODO(), id)

	assert.Equal(t, nil, err)
}

// Test_DeleteInvalidCar function checks the case for deleting the non-existing car
func Test_DeleteInvalidCar(t *testing.T) {
	id := uuid.New()

	mockCar, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Delete(gomock.Any(), id).Return(nil)
	mockCar.EXPECT().Delete(gomock.Any(), id).Return(errors2.EntityNotFound{})

	err := svc.Delete(context.TODO(), id)

	assert.Equal(t, errors2.EntityNotFound{}, err)
}

// Test_DeleteInvalidEngine function checks the case for deleting the non-existing engine
func Test_DeleteInvalidEngine(t *testing.T) {
	id := uuid.New()

	_, mockEngine, svc := initializeTest(t)

	mockEngine.EXPECT().Delete(gomock.Any(), id).Return(errors2.EntityNotFound{})

	err := svc.Delete(context.TODO(), id)

	assert.Equal(t, errors2.EntityNotFound{}, err)
}

// Test_validCar function checks the test cases for invalid car data
func Test_validCar(t *testing.T) {
	cases := []struct {
		desc  string
		input models.Car
		err   error
	}{
		{"invalid name", models.Car{Name: ""}, errors2.InvalidParam{Params: []string{"name"}}},
		{"invalid year", models.Car{Name: "ABC", Year: 1780}, errors2.InvalidParam{Params: []string{"year"}}},
		{"invalid brand", models.Car{Name: "ABC", Year: 2020, Brand: 789}, errors2.InvalidParam{Params: []string{"brand"}}},
		{"invalid fuel", models.Car{Name: "ABC", Year: 2020, Brand: brand.Mercedes, Fuel: 123}, errors2.InvalidParam{Params: []string{"fuel"}}},
	}

	for _, tc := range cases {
		err := validCar(&tc.input)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}

// Test_validEngine function checks the test cases for invalid engine data
func Test_validEngine(t *testing.T) {
	engine1 := models.Engine{Displacement: 1000, NCylinders: 2, Range: 600}
	engine2 := models.Engine{Displacement: -500, NCylinders: -4, Range: -300}
	engine3 := models.Engine{Displacement: 0, NCylinders: 0, Range: 0}
	engine4 := models.Engine{Displacement: 0, NCylinders: 0, Range: -250}
	engine5 := models.Engine{Displacement: 1000, NCylinders: -5}

	validEngine1 := models.Engine{Displacement: 2000, NCylinders: 6}
	validEngine2 := models.Engine{Displacement: 0, NCylinders: 0, Range: 300}

	cases := []struct {
		desc  string
		input models.Engine
		err   error
	}{
		{"positive values", engine1, errors2.InvalidParam{Params: []string{"displacement", "nCylinder", "range"}}},
		{"negative values", engine2, errors2.InvalidParam{Params: []string{"displacement"}}},
		{"zero values", engine3, errors2.InvalidParam{Params: []string{"displacement", "nCylinder", "range"}}},
		{"negative range", engine4, errors2.InvalidParam{Params: []string{"range"}}},
		{"negative nCylinder", engine5, errors2.InvalidParam{Params: []string{"nCylinder"}}},
		{"valid for non electric", validEngine1, nil},
		{"valid for electric", validEngine2, nil},
	}

	for _, tc := range cases {
		err := validEngine(tc.input)

		assert.Equal(t, tc.err, err, tc.desc)
	}
}
