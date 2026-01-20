package test_data

import "github.com/DimaSU2020/car_rental_service/internal/models/helper"


var Tests = []struct{
	Name         string
	Fixture      string
	ExpectError  bool
	ErrorСontent error
}{
	{
		Name         : "valid booking",
		Fixture      : "valid_booking.json",
		ExpectError  : false,
		ErrorСontent : nil,	
	},
	{
		Name         : "wrong cost",
		Fixture      : "wrong_cost.json",
		ExpectError  : true,
		ErrorСontent : helper.ErrWrongRentCost,	
	},
	{
		Name         : "invalid status",
		Fixture      : "invalid_status.json",
		ExpectError  : true,
		ErrorСontent : helper.ErrInvalidStatus,	
	},
	{
		Name         : "wrong period",
		Fixture      : "wrong_period.json",
		ExpectError  : true,
		ErrorСontent : helper.ErrWrongRentalPeriod,	
	},
}