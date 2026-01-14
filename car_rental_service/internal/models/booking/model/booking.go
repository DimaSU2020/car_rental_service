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

type Status struct {
	Done     string
	Canceled string
}

var Statuses = map[string]string{
	"done"     : "done",
	"canceled" : "canceled",
}

func (b *Booking) Validate() error {

	maxRentalPeriod := 28

	if b.Start_day.After(b.End_day) {
		return helper.ErrStartBefore
	}

	if b.End_day.After(b.Start_day.AddDate(0,0,maxRentalPeriod)) {
		return helper.ErrWrongRentalPeriod
	}

	if b.Daily_cost <= 0 {
		return helper.ErrWrongRentCost
	}

	_, exists := Statuses[b.Status]
	if !exists {
		return helper.ErrInvalidStatus
	}

	return nil
}