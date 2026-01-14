package model

import (
	"strings"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/helper"
)

type License struct{
	ID             int64     `json:"id"`
	Number         string    `json:"number"`
	IssuanceDate   time.Time `json:"issuance_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (l *License) Validate() error {

	if strings.TrimSpace(l.Number) == "" {
		return helper.ErrEmptyNumber
	}

	if l.IssuanceDate.After(time.Now()) {
		return helper.ErrIssuanceDate
	}

	if l.ExpirationDate.Before(l.IssuanceDate.AddDate(10, 0, 0)) {
		return helper.ExpirationDate
	}

	return nil
}
