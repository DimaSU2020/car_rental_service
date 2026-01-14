package test_data

import "github.com/DimaSU2020/car_rental_service/internal/models/helper"



var Tests = []struct{
		Name         string
		Fixture      string
		ExpectError  bool
		ErrorСontent error
	}{
		{
			Name         : "valid license",
			Fixture      : "valid_license.json",
			ExpectError  : false,
			ErrorСontent : nil,	
		},
		{
			Name         : "empty number",
			Fixture      : "empty_number.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrEmptyNumber,	
		},
		{
			Name         : "wrong issuance date",
			Fixture      : "wrong_issuance_date.json",
			ExpectError  : false,
			ErrorСontent : helper.ErrIssuanceDate,
		},
		{
			Name         : "wrong expiration date",
			Fixture      : "wrong_expiration_date.json",
			ExpectError  : false,
			ErrorСontent : helper.ExpirationDate,
		},
	}