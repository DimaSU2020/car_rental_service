package cars

import (
	"context"

	"github.com/DimaSU2020/car_rental_service/internal/models/car/model"
	"github.com/DimaSU2020/car_rental_service/internal/repo"
)

type fakeCarRepo struct {
	cars map[int64]*model.Car

	createErr    error
	getErr       error
	updateErr    error
	updateCalled bool
	lastUpdated  *model.Car
	deleteCalled bool
	deleteErr    error
}

func NewFakeCarRepo() *fakeCarRepo {
	return &fakeCarRepo{
		cars: make(map[int64]*model.Car),
	}
}

func (f *fakeCarRepo) List(ctx context.Context, limit, offset int) ([]*model.Car, error) {
	res := make([]*model.Car, 0, len(f.cars))

	for _, c := range f.cars {
		res = append(res, c)
	}

	return res, nil
}

func (f *fakeCarRepo) GetByID(ctx context.Context, id int64) (*model.Car, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}

	car, ok := f.cars[id]
	if !ok {
		return nil, repo.ErrNotFound
	}

	return car, nil
}

func (f *fakeCarRepo) Create(ctx context.Context, car *model.Car) (*model.Car, error) {
	if f.createErr != nil {
		return nil, f.createErr
	}

	car.ID = int64(len(f.cars)+1)
	f.cars[car.ID] = car

	return car, nil
}


func (f *fakeCarRepo) Update(ctx context.Context, car *model.Car) error {
	f.updateCalled = true
	f.lastUpdated = car

	if f.updateErr != nil {
		return f.updateErr
	}
	
	f.cars[car.ID] = car
	return nil
}

func (f *fakeCarRepo) Delete(ctx context.Context, id int64) error {
	f.deleteCalled = true

	if f.deleteErr != nil {
		return f.deleteErr
	}

	if _, ok := f.cars[id]; !ok {
		return repo.ErrNotFound
	}

	delete(f.cars, id)
	return nil
}