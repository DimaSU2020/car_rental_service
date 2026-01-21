package model

import (
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/helper"
)


type Booking struct {
    ID            int64     `json:"id"`
    ID_car        int64     `json:"id_car"`
    ID_user       int64     `json:"id_user"` 
    Start_day     time.Time `json:"start_day"`    
    End_day       time.Time `json:"end_day"`
    Daily_cost    int64     `json:"daily_cost"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

const (
	BookingStatusDone     = "done"
	BookingStatusCanceled = "canceled"
)

var AllowedStatuses = map[string]struct{}{
	BookingStatusDone:     {},
	BookingStatusCanceled: {},
}

func (b *Booking) Validate() error {
	maxRentalPeriod := 28
	today := time.Now().Truncate(24 * time.Hour)

	if !b.Start_day.Before(b.End_day) {
		return helper.ErrStartBefore
	}

	if b.Start_day.Before(today) {
		return helper.ErrBookingInPast
	}

	if b.End_day.After(b.Start_day.AddDate(0,0,maxRentalPeriod)) {
		return helper.ErrWrongRentalPeriod
	}

	if b.Daily_cost <= 0 {
		return helper.ErrWrongRentCost
	}

	if _, ok := AllowedStatuses[b.Status]; !ok {
		return helper.ErrInvalidStatus
	}

	return nil
}