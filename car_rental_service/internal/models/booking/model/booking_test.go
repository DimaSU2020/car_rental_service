package model

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/DimaSU2020/car_rental_service/internal/models/booking/test_data"
	"github.com/DimaSU2020/car_rental_service/internal/models/fixture"
)

func TestBooking(t *testing.T) {
	for _, test := range test_data.Tests {
		tt := test
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			data := fixture.LoadFixture(t, tt.Fixture)

			var booking Booking

			err := json.Unmarshal(data, &booking)
			if err != nil {
				t.Error("received an unmarshalling error")
			}

			err = booking.Validate()

			if tt.ExpectError && err == nil {
				t.Error("expected error, got nil")
			}

			if !errors.Is(err, tt.ErrorСontent) {
				t.Errorf("got = %v; want = %v", err, tt.ErrorСontent)
			}
		})
	}
}