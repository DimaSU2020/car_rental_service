package helper

import "errors"

var (
	ErrEmptyBrand        = errors.New("brand cannot be empty")
	ErrEmptyModel        = errors.New("model cannot be empty")
	ErrWrongRentCost     = errors.New("rent cannot be empty or negative")
	ErrWrongYear         = errors.New("year must be between 1900 and 2025")
	ErrEmptyPhoto        = errors.New("photo connot be empty")
	ErrTooShortName      = errors.New("name cant be less than 2 symbols")
	ErrTooLongName       = errors.New("name cant be more than 20 symbols")
	ErrEmptyName         = errors.New("name cannot be empty")
	ErrWrongEmail        = errors.New("wrong email")
	ErrEmptyEmail        = errors.New("email cannot be empty")
	ErrWrongID           = errors.New("wrong ID format")
	ErrUserWrongRole     = errors.New("wrong role of user")
	ErrEmptyNumber       = errors.New("number cannot be empty")
	ErrIssuanceDate      = errors.New("IssuanceDate cant be less than date Now")
	ExpirationDate       = errors.New("ExpirationDate cant be less than 10 years from IssuanceDate")
	ErrInvalidStatus     = errors.New("Such status does not exist")
	ErrStartBefore       = errors.New("start_day should be before end_day")
	ErrWrongRentalPeriod = errors.New("Wrong rental period")
	ErrBookingInPast     = errors.New("booking in past")
)