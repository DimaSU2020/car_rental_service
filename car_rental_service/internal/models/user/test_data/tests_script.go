package test_data

import "github.com/DimaSU2020/car_rental_service/internal/models/helper"

var Tests = []struct{
		Name         string
		Fixture      string
		ExpectError  bool
		ErrorСontent error
	}{
		{
			Name         : "valid user",
			Fixture      : "valid_user.json",
			ExpectError  : false,
			ErrorСontent : nil,	
		},
		{
			Name         : "empty name",
			Fixture      : "empty_name.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrEmptyName,	
		},
		{
			Name         : "too short name",
			Fixture      : "too_short_name.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrTooShortName,	
		},
		{
			Name         : "too long name",
			Fixture      : "too_long_name.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrTooLongName,
		},
		{
			Name         : "empty email",
			Fixture      : "empty_email.json",
			ExpectError  : false,
			ErrorСontent : helper.ErrEmptyEmail,
		},
		{
			Name         : "wrong email",
			Fixture      : "wrong_email.json",
			ExpectError  : false,
			ErrorСontent : helper.ErrWrongEmail,
		},
	}