package dto

import (
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/license/model"
)

type CreateLicenseRequest struct {
	ID             int64     `json:"id"`
	Number         string    `json:"number"`
	IssuanceDate   time.Time `json:"issuance_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UpdateLicenseRequest struct {
	ID             int64     `json:"id"`
	Number         string    `json:"number"`
	IssuanceDate   time.Time `json:"issuance_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LicenseResponse struct {
	ID             int64     `json:"id"`
	Number         string    `json:"number"`
	IssuanceDate   time.Time `json:"issuance_date"`
	ExpirationDate time.Time `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func LicenseToResponse(l *model.License) *LicenseResponse {
	return &LicenseResponse {
		ID            : l.ID,
		Number        : l.Number,
		IssuanceDate  : l.IssuanceDate,
		ExpirationDate: l.ExpirationDate,
		CreatedAt     : l.CreatedAt,
		UpdatedAt     : l.UpdatedAt,
	}
}