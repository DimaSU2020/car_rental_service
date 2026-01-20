package test_data

import "github.com/DimaSU2020/car_rental_service/internal/models/helper"

var Tests = []struct{
		Name         string
		Fixture      string
		ExpectError  bool
		ErrorСontent error
	}{
		{
			Name         : "valid car",
			Fixture      : "valid_car.json",
			ExpectError  : false,
			ErrorСontent : nil,	
		},
		{
			Name         : "empty brand",
			Fixture      : "empty_brand.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrEmptyBrand,	
		},
		{
			Name         : "empty model",
			Fixture      : "empty_model.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrEmptyModel,	
		},
		{
			Name         : "wrong rent",
			Fixture      : "wrong_rent.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrWrongRentCost,	
		},
		{
			Name         : "year too early",
			Fixture      : "year_too_early.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrWrongYear,	
		},
		{
			Name         : "year too late",
			Fixture      : "year_too_late.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrWrongYear,	
		},
		{
			Name         : "year cant be negative",
			Fixture      : "year_cant_be_negative.json",
			ExpectError  : true,
			ErrorСontent : helper.ErrWrongYear,	
		},
	}