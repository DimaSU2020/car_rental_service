package dto

import (
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/user/model"
)

type CreateUserRequest struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Birthday  time.Time     `json:"birthday"`
    IsAdmin   bool          `json:"is_admin"`
	LicenseId *int64         `json:"license_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type UpdateUserRequest struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Birthday  time.Time     `json:"birthday"`
    IsAdmin   bool          `json:"is_admin"`
	LicenseId *int64         `json:"license_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type UserResponse struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Birthday  time.Time     `json:"birthday"`
    IsAdmin   bool          `json:"is_admin"`
	LicenseId *int64         `json:"license_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func UserToResponse(u *model.User) *UserResponse {
	return &UserResponse {
		ID          : u.ID,
		Name        : u.Name,
		Email       : u.Email,
		Birthday    : u.Birthday,
		IsAdmin     : u.IsAdmin,
		LicenseId   : u.LicenseId,
		CreatedAt   : u.CreatedAt,
		UpdatedAt   : u.UpdatedAt,
	}
}