package service

import (
	"context"
	"errors"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/car/model"
)

type CreateCarInput struct {
	ID            int64
    Brand         string   
    Model         string
    Year          int
    DailyRentCost int64
	Photo         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UpdateCarInput struct {
	ID            int64
    Brand         string   
    Model         string
    Year          int
    DailyRentCost int64
	Photo         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
	ErrCarAlreadyExist = errors.New("car already exist")
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

	allCars, err := s.repo.List(ctx, limitMax, offset)
	if err != nil {
		return nil, err
	}

	if offset > len(allCars) {
		return []*model.Car{}, nil
	}
	
	endPagination := min(limit + offset, len(allCars))

	return allCars[offset:endPagination], nil
}

func (s *carService) GetByID(ctx context.Context, id int64) (*model.Car, error) {
	f, err := s.repo.GetByID(ctx, id)
	if err != nil || errors.Is(err, ErrCarNotFound) {
		return nil, err
	}
    return f, nil
}

func (s *carService) Create(ctx context.Context, input CreateCarInput) (*model.Car, error) {
	c := model.Car{
        Brand         : input.Brand,
		Model         : input.Model,
        Year          : input.Year,
        DailyRentCost : input.DailyRentCost,
		Photo         : input.Photo,
		CreatedAt     : input.CreatedAt,
		UpdatedAt     : input.UpdatedAt,
    }

	err := c.Validate()
	if err != nil {
		return nil, err
	}

	return s.repo.Create(ctx, &c)
}

func (s *carService) Update(ctx context.Context, input UpdateCarInput) error {
	car := model.Car{
		ID            : input.ID,
		Brand         : input.Brand,
		Model         : input.Model,
        Year          : input.Year,
		Photo         : input.Photo,
        DailyRentCost : input.DailyRentCost,
		UpdatedAt     : input.UpdatedAt,
    }

	err := car.Validate()
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, &car)
}

func (s *carService) Delete(ctx context.Context, id int64) error {
	_, err := s.repo.GetByID(ctx, id)
	if err != nil || errors.Is(err, ErrCarNotFound) {
		return err
	}
	return s.repo.Delete(ctx, id)
}
