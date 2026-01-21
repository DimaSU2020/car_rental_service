package bookings

import (
	"context"
	"errors"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/booking/model"
	"github.com/DimaSU2020/car_rental_service/internal/models/helper"
)

type CreateBookingInput struct {
    ID            int64     
    ID_car        int64     
    ID_user       int64     
    Start_day     time.Time  
    End_day       time.Time 
    Daily_cost    int64     
	Status        string    
	CreatedAt     time.Time 
	UpdatedAt     time.Time 
}

type UpdateBookingInput struct {
    ID            int64     
    ID_car        int64     
    ID_user       int64     
    Start_day     time.Time  
    End_day       time.Time 
    Daily_cost    int64     
	Status        string    
	CreatedAt     time.Time 
	UpdatedAt     time.Time 
}

type CheckBookingInput struct { 
    ID_car        int64     
    ID_user       int64     
    Start_day     time.Time  
    End_day       time.Time  
	Status        string    
}

type BookingRepository interface {
	List(ctx context.Context, limit, offset int) ([]*model.Booking, error)
	GetByID(ctx context.Context, id int64) (*model.Booking, error)
	Create(ctx context.Context, booking *model.Booking) (*model.Booking, error)
	ExistsOverlappingBooking(ctx context.Context, carID int64, from, to time.Time) (bool, error)
}

type BookingService interface {
	List(ctx context.Context, limit, offset int) ([]*model.Booking, error)
	GetByID(ctx context.Context, id int64) (*model.Booking, error)
	Create(ctx context.Context, input CreateBookingInput) (*model.Booking, error)
	IsCarAvailable(ctx context.Context, carID int64, from, to time.Time) (bool, error)
}

type bookingService struct {
	repo BookingRepository
}

func NewBookingService(repo BookingRepository) BookingService {
    return &bookingService{repo: repo}
}

func (b *bookingService) List(ctx context.Context, limit, offset int) ([]*model.Booking, error) {
	limitMax := 100
	if limit <= 0 || limit > limitMax { 
		limit = 10
	}
	
	if offset < 0 { 
		offset = 0
	}

	allBookings, err := b.repo.List(ctx, limitMax, offset)
	if err != nil {
		return nil, err
	}

	if offset > len(allBookings) {
		return []*model.Booking{}, nil
	}
	
	endPagination := min(limit + offset, len(allBookings))

	return allBookings[offset:endPagination], nil
}

func (b *bookingService) GetByID(ctx context.Context, id int64) (*model.Booking, error) {
	f, err := b.repo.GetByID(ctx, id)
	if err != nil || errors.Is(err, ErrBookingNotFound) {
		return nil, err
	}
    return f, nil
}

func (b *bookingService) Create(ctx context.Context, input CreateBookingInput) (*model.Booking, error) {
	now := time.Now()
	booking := model.Booking{
		ID_car      : input.ID_car,
		ID_user     : input.ID_user,    
		Start_day   : input.Start_day,
		End_day     : input.End_day,
		Daily_cost  : input.Daily_cost,
		Status      : model.BookingStatusDone,
		CreatedAt   : now,
		UpdatedAt   : now,
    }

	err := booking.Validate()
	if err != nil {
		return nil, err
	}

	return b.repo.Create(ctx, &booking)
}

func (b *bookingService) IsCarAvailable(ctx context.Context, carID int64, from, to time.Time) (bool, error) {
	if !from.Before(to) {
		return false, helper.ErrStartBefore
	}

	exists, err := b.repo.ExistsOverlappingBooking(ctx, carID, from, to)
	if err != nil {
		return false, err
	}

	return !exists, nil
}

var ErrBookingNotFound = errors.New("booking not found")
