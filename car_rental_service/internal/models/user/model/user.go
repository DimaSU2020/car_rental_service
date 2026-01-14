package model

import (
	"regexp"
	"strings"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/helper"
)

type User struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Birthday  time.Time     `json:"birthday"`
    IsAdmin   bool          `json:"is_admin"`
	LicenseId *int64        `json:"license_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (u *User) Validate() error {
	patternForMail := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	
	if strings.TrimSpace(u.Name) == "" {
		return helper.ErrEmptyName
	}
	if len(u.Name) < 2 {
		return helper.ErrTooShortName
	}

	if len(u.Name) > 20 {
		return helper.ErrTooLongName
	}

	if strings.TrimSpace(u.Email) == "" {
		return helper.ErrEmptyEmail
	}

	matched, _ := regexp.MatchString(patternForMail, u.Email)
	if !matched {
		return helper.ErrWrongEmail
	}

	return nil
}
