package cars

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/DimaSU2020/car_rental_service/internal/models/car/model"
	"github.com/stretchr/testify/require"
)

func TestCarService_Create_Ok(t *testing.T) {
	repo := NewFakeCarRepo()
	svc  := NewCarService(repo)
	
	car, err := svc.Create(context.Background(), CreateCarInput{
		Brand:         "Toyota",
		Model:         "Camry",
		Year:          2020,
		DailyRentCost: 3000,
		Photo:         "camry.jpg",
	})

	require.NoError(t, err)
	require.NotNil(t, car)

	require.Equal(t, "Toyota", car.Brand)
	require.Equal(t, "Camry", car.Model)
	require.Equal(t, 2020, car.Year)
	require.Equal(t, int64(3000), car.DailyRentCost)
	require.Equal(t, int64(1), car.ID)
}

func TestCarService_Create_Invalid(t *testing.T) {
	repo := NewFakeCarRepo()
	svc  := NewCarService(repo)

	car, err := svc.Create(context.Background(), CreateCarInput{
		Brand:         "",
		Model:         "Camry",
		Year:          2020,
		DailyRentCost: 3000,
	})

	require.Nil(t, car)
	require.Error(t, err)
}

func TestCarService_GetByID_OK(t *testing.T) {
	repo := NewFakeCarRepo()
	svc  := NewCarService(repo)

	car, err := svc.Create(context.Background(), CreateCarInput{
		Brand:         "Toyota",
		Model:         "Camry",
		Year:          2020,
		DailyRentCost: 3000,
		Photo:         "camry.jpg",
	})

	car, err = svc.GetByID(context.Background(), 1)

	require.NotNil(t, car)
	require.ErrorIs(t, err, nil)
}

func TestCarService_GetByID_NotFound(t *testing.T) {
	repo := NewFakeCarRepo()
	svc  := NewCarService(repo)

	car, err := svc.GetByID(context.Background(), 11)

	require.Nil(t, car)
	require.ErrorIs(t, err, ErrCarNotFound)
}

func TestCarService_GetByID_RepoError(t *testing.T) {
	repo := NewFakeCarRepo()
	repo.getErr = errors.New("db connection lost")
	svc  := NewCarService(repo)

	car, err := svc.GetByID(context.Background(), 5)

	require.Nil(t, car)
	require.Error(t, err)
	require.ErrorIs(t, err, repo.getErr)
}

func TestCarService_Update_OK(t *testing.T) {
	repo := NewFakeCarRepo()

	car, _ := model.NewCar("Toyota", "Camry", 2020, 3000, "toyota.jpg")
	car.ID = 1
	repo.cars[1] = car

	svc := NewCarService(repo)

	err := svc.Update(context.Background(), UpdateCarInput{
		ID: 1,
		Brand: "Honda",
		Model: "Accord",
		Year: 2021,
		DailyRentCost: 5000,
		Photo: "honda.jpg",
	})

	require.NoError(t, err)
	require.True(t, repo.updateCalled)

	updated := repo.lastUpdated
	require.NotNil(t, updated)

	require.Equal(t, "Honda", updated.Brand)
	require.Equal(t, "Accord", updated.Model)
	require.Equal(t, 2021, updated.Year)
	require.Equal(t, int64(5000), updated.DailyRentCost)
	require.Equal(t, "honda.jpg", updated.Photo)
}

func TestCarService_Update_NotFound(t *testing.T) {
	repo := NewFakeCarRepo()
	svc  := NewCarService(repo)

	car, err := svc.GetByID(context.Background(), 11)

	require.Nil(t, car)
	require.ErrorIs(t, err, ErrCarNotFound)
}

func TestCarService_Update_InvalidInput(t *testing.T) {
	repo := NewFakeCarRepo()

	car, _ := model.NewCar("Toyota", "Camry", 2020, 3000, "toyota.jpg")
	car.ID = 1
	repo.cars[1] = car

	svc := NewCarService(repo)

	err := svc.Update(context.Background(), UpdateCarInput{
		ID: 1,
		Brand: "Honda",
		Model: "",
		Year: 2021,
		DailyRentCost: 5000,
		Photo: "honda.jpg",
	})

	require.Error(t, err)
	require.False(t, repo.updateCalled)
	fmt.Println("Error update fake repo:", repo.updateErr)
	fmt.Println("Error service:", err)
	require.NotEqual(t, "Honda", car.Brand)
	require.NotEqual(t, 2021, car.Year)
	require.NotEqual(t, int64(5000), car.DailyRentCost)
	require.NotEqual(t, "honda.jpg", car.Photo)
}

func TestCarService_Update_RepoError_OnGet(t *testing.T) {
	repo := NewFakeCarRepo()
	dbError := errors.New("db connection lost")
	repo.getErr = dbError
	svc := NewCarService(repo)

	err := svc.Update(context.Background(), UpdateCarInput{
		ID: 1,
		Brand: "Honda",
		Model: "",
		Year: 2021,
		DailyRentCost: 5000,
		Photo: "honda.jpg",
	})

	require.ErrorIs(t, err, dbError)
	require.False(t, repo.updateCalled)
}

func TestCarService_Update_RepoError_OnUpdate(t *testing.T) {
	repo := NewFakeCarRepo()

	car, _ := model.NewCar("Toyota", "Camry", 2020, 3000, "toyota.jpg")
	car.ID = 1
	repo.cars[1] = car
	repo.updateErr = errors.New("db connection lost")

	svc := NewCarService(repo)

	err := svc.Update(context.Background(), UpdateCarInput{
		ID:            1,
		Brand:         "Honda",
		Model:         "Accord",
		Year:          2021,
		DailyRentCost: 5000,
		Photo:         "honda.jpg",
	})

	require.ErrorIs(t, err, repo.updateErr)
	require.True(t, repo.updateCalled)

	require.Equal(t, "Toyota", car.Brand)
	require.Equal(t, "Camry", car.Model)
	require.Equal(t, 2020, car.Year)
	require.Equal(t, int64(3000), car.DailyRentCost)
	require.Equal(t, "toyota.jpg", car.Photo)
}

func TestCarService_Delete_OK(t *testing.T) {
	repo := NewFakeCarRepo()

	car, _ := model.NewCar("Toyota", "Camry", 2020, 3000, "toyota.jpg")
	car.ID = 1
	repo.cars[1] = car

	svc := NewCarService(repo)
	
	err := svc.Delete(context.Background(), 1)

	require.NoError(t, err)
	require.True(t, repo.deleteCalled)
	require.Empty(t, repo.cars)
}

func TestCarService_Delete_RepoError(t *testing.T) {
	repo := NewFakeCarRepo()
	repo.deleteErr = errors.New("db unavailable")

	svc := NewCarService(repo)

	err := svc.Delete(context.Background(), 1) 

	require.EqualError(t, err, "db unavailable")
	require.True(t, repo.deleteCalled)
}

func TestCarService_Delete_Idempotent(t *testing.T) {
	repo := NewFakeCarRepo()

	car, _ := model.NewCar("Toyota", "Camry", 2020, 3000, "toyota.jpg")
	car.ID = 1
	repo.cars[1] = car

	svc := NewCarService(repo)

	err := svc.Delete(context.Background(), 1)
	require.NoError(t, err)

	err = svc.Delete(context.Background(), 1)
	require.NoError(t, err)
}


