package service

import (
	"context"
	"errors"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/license/model"
)

var ErrLicenseNotFound = errors.New("license not found")

type CreateLicenseInput struct {
	ID             int64     
	Number         string    
	IssuanceDate   time.Time 
	ExpirationDate time.Time 
	CreatedAt      time.Time 
	UpdatedAt      time.Time
}

type UpdateLicenseInput struct {
	ID             int64     
	Number         string    
	IssuanceDate   time.Time 
	ExpirationDate time.Time 
	CreatedAt      time.Time 
	UpdatedAt      time.Time
}

type LicenseRepository interface {
	List(ctx context.Context, limit, offset int) ([]*model.License, error)
	GetByID(ctx context.Context, id int64) (*model.License, error)
	Create(ctx context.Context, license *model.License) (*model.License, error)
	Update(ctx context.Context, license *model.License) error
	Delete(ctx context.Context, id int64) error
}

type LicenseService interface {
	List(ctx context.Context, limit, offset int) ([]*model.License, error)
	GetByID(ctx context.Context, id int64) (*model.License, error)
	Create(ctx context.Context, input CreateLicenseInput) (*model.License, error)
	Update(ctx context.Context, input UpdateLicenseInput) error
	Delete(ctx context.Context, id int64) error
}

type licenseService struct {
	repo LicenseRepository
}

func NewLicenseService(repo LicenseRepository) LicenseService {
	return &licenseService{repo: repo}
}

func (l *licenseService) List(ctx context.Context, limit, offset int) ([]*model.License, error) {
	limitMax := 100
	if limit <= 0 || limit > limitMax { 
		limit = 10
	}
	
	if offset < 0 { 
		offset = 0
	}

	allLicenses, err := l.repo.List(ctx, limitMax, offset)
	if err != nil {
		return nil, err
	}

	if offset > len(allLicenses) {
		return []*model.License{}, nil
	}

	endPagination := min(limit + offset, len(allLicenses))

	if len(allLicenses) == 0 {
		return []*model.License{}, nil
	}

	return allLicenses[offset:endPagination], nil
}

func (l *licenseService) GetByID(ctx context.Context, id int64) (*model.License, error) {
	return l.repo.GetByID(ctx, id)
}

func (l *licenseService) Create(ctx context.Context, input CreateLicenseInput) (*model.License, error) {
	license := model.License{
		Number        : input.Number,
		IssuanceDate  : input.IssuanceDate,
		ExpirationDate: input.ExpirationDate,
		CreatedAt     : time.Now(),
		UpdatedAt     : time.Now(),
	}

	err := license.Validate()
	if err != nil {
		return nil, err
	}
	return l.repo.Create(ctx, &license)
}

func (l *licenseService) Update(ctx context.Context, input UpdateLicenseInput) error {
	license := model.License{
		ID            : input.ID,
		Number        : input.Number,
		IssuanceDate  : input.IssuanceDate,
		ExpirationDate: input.ExpirationDate,
		CreatedAt     : input.CreatedAt,
		UpdatedAt     : time.Now(),
	}

	err := license.Validate()
	if err != nil {
		return err
	}

	return l.repo.Update(ctx, &license)
}

func (l *licenseService) Delete(ctx context.Context, id int64) error {
	return l.repo.Delete(ctx, id)
}