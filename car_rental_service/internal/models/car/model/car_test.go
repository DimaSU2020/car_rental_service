package model

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/DimaSU2020/car_rental_service/internal/models/fixture"
	"github.com/DimaSU2020/car_rental_service/internal/models/car/test_data"
)


func TestModelCar(t *testing.T) {
	for _, test := range test_data.Tests {
		tt := test
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			data := fixture.LoadFixture(t, tt.Fixture)

			var car Car

			err := json.Unmarshal(data, &car)
			if err != nil {
				t.Error("received an unmarshalling error")
			}

			err = car.Validate()

			if tt.ExpectError && err == nil {
				t.Error("expected error, got nil")
			}

			if !errors.Is(err, tt.ErrorСontent) {
				t.Errorf("got = %v; want = %v", err, tt.ErrorСontent)
			}
		})
	}
}