package dto

import (
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/booking/model"
)

type CreateBookingRequest struct {
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

type BookingResponse struct {
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

func BookingToResponse(b *model.Booking) *BookingResponse {
	return &BookingResponse {
		ID         : b.ID,
		ID_car     : b.ID_car,
		ID_user    : b.ID_user,    
		Start_day  : b.Start_day,
		End_day    : b.End_day,
		Daily_cost : b.Daily_cost,
		Status     : b.Status,
		CreatedAt  : b.CreatedAt,
		UpdatedAt  : b.UpdatedAt,
	}
}
