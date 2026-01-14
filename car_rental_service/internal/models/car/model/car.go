package model

import (
	"strings"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/helper"
)

type Car struct {
    ID            int64     `json:"id"`
    Brand         string     `json:"brand"`
    Model         string    `json:"model"` 
    Year          int       `json:"year"`     
    DailyRentCost int64     `json:"rent"`
    Photo         string    `json:"photo"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (c *Car) Validate() error {
    if strings.TrimSpace(c.Brand) == "" {
		return helper.ErrEmptyBrand
	}

    if strings.TrimSpace(c.Model) == "" {
		return helper.ErrEmptyModel
	}

    if c.DailyRentCost <= 0 {
        return helper.ErrWrongRentCost
    }

	if c.Year < 1900 || c.Year > 2025 {
		return helper.ErrWrongYear
	}

    return nil
}