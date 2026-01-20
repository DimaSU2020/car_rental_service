package cars

import (
	"context"
	"errors"
	"fmt"

	"github.com/DimaSU2020/car_rental_service/internal/models/car/model"
	"github.com/DimaSU2020/car_rental_service/internal/repo"
)

type CreateCarInput struct {
    Brand         string   
    Model         string
    Year          int
    DailyRentCost int64
	Photo         string
}

type UpdateCarInput struct {
	ID            int64
    Brand         string   
    Model         string
    Year          int
    DailyRentCost int64
	Photo         string
}

type CarRepository interface {
	List(ctx context.Context, limit, offset int) ([]*model.Car, error)
	GetByID(ctx context.Context, id int64) (*model.Car, error)
	Create(ctx context.Context, car *model.Car) (*model.Car, error)
	Update(ctx context.Context, car *model.Car) error
	Delete(ctx context.Context, id int64) error
}

type CarService interface {
	List(ctx context.Context, limit, offset int) ([]*model.Car, error)
	GetByID(ctx context.Context, id int64) (*model.Car, error)
	Create(ctx context.Context, input CreateCarInput) (*model.Car, error)
	Update(ctx context.Context, input UpdateCarInput) error
	Delete(ctx context.Context, id int64) error
}

type carService struct {
	repo CarRepository
}

var (
	ErrCarNotFound     = errors.New("car not found")
	ErrInvalidCarData  = errors.New("invalid car data")
)

func NewCarService(repo CarRepository) CarService {
    return &carService{repo: repo}
}

func (s *carService) List(ctx context.Context, limit, offset int) ([]*model.Car, error) {
	limitMax := 100
	if limit <= 0 || limit > limitMax { 
		limit = 10
	}
	
	if offset < 0 { 
		offset = 0
	}

	return s.repo.List(ctx, limit, offset)
}

func (s *carService) Create(ctx context.Context, input CreateCarInput) (*model.Car, error) {
	car, err := model.NewCar(
        input.Brand,
		input.Model,
        input.Year,
        input.DailyRentCost,
		input.Photo,
	)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidCarData, err.Error())
	}

	return s.repo.Create(ctx, car)
}

func (s *carService) GetByID(ctx context.Context, id int64) (*model.Car, error) {
	car, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, repo.ErrNotFound) {
		return nil, ErrCarNotFound
	}
	if err != nil {
		return nil, err
	}

    return car, nil
}

func (s *carService) Update(ctx context.Context, input UpdateCarInput) error {
	genCar, err := s.repo.GetByID(ctx, input.ID)
	if errors.Is(err, repo.ErrNotFound) {
		return ErrCarNotFound
	}
	if err != nil {
		return err
	}

	car := *genCar

	if err := car.UpdateCar(
		input.Brand,
		input.Model,
        input.Year,
		input.DailyRentCost,
		input.Photo,
	); err != nil {
		return err
	}

	return s.repo.Update(ctx, &car)
}

func (s *carService) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if errors.Is(err, repo.ErrNotFound) {
		return nil
	}
	return err
}
