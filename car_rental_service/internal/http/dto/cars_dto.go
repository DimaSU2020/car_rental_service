package dto

import (
	"github.com/DimaSU2020/car_rental_service/internal/models/car/model"
)

type CreateCarRequest struct {
    Brand         string    `json:"brand"`
    Model         string    `json:"model"` 
    Year          int       `json:"year"`     
    DailyRentCost int64     `json:"rent"`
    Photo         string    `json:"photo"`
}

type CarResponse struct {
    ID            int64     `json:"id"`
    Brand         string    `json:"brand"`
    Model         string    `json:"model"` 
    Year          int       `json:"year"`     
    DailyRentCost int64     `json:"rent"`
    Photo         string    `json:"photo"`
}

func CarToResponse(c *model.Car) *CarResponse {
	return &CarResponse {
		ID            : c.ID,
		Brand         : c.Brand,
		Model         : c.Model,
		Year          : c.Year,
		DailyRentCost : c.DailyRentCost,
		Photo         : c.Photo,
	}
}