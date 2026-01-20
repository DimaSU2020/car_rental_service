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

func NewCar(
	brand string,
	model string,
	year int,
	dailyRentCost int64,
	photo string,
) (*Car, error) {

	c := &Car{
		Brand: strings.TrimSpace(brand),
		Model: strings.TrimSpace(model),
		Year: year,
		DailyRentCost: dailyRentCost,
		Photo: photo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Car) UpdateCar(
	brand string,
	model string,
	year int,
	dailyRentCost int64,
	photo string,
) error {

	c.Brand = strings.TrimSpace(brand)
	c.Model = strings.TrimSpace(model)
	c.Year  = year
	c.DailyRentCost = dailyRentCost
	c.Photo = photo
	c.UpdatedAt = time.Now()

	return c.Validate()
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

	if c.Year < 1900 || c.Year > time.Now().Year()+1 {
		return helper.ErrWrongYear
	}

    return nil
}