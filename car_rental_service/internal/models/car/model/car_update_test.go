package model

import (
	"testing"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/helper"
	"github.com/stretchr/testify/require"
)


func TestUpdateCar_OK(t *testing.T) {
	car, err := NewCar(
		"Toyota",
		"Camry",
		2020,
		3000,
		"camry_1.jpg",
	)

	require.NoError(t, err)
	require.NotNil(t, car)

	oldCreatedAt := car.CreatedAt
	oldUpdatedAt := car.UpdatedAt

	time.Sleep(time.Millisecond)

	err = car.UpdateCar(
		"Honda",
		"Accord",
		2021,
		int64(6000),
		"honda_1.jpg",
	)

	require.Equal(t, "Honda", car.Brand)
	require.Equal(t, "Accord", car.Model)
	require.Equal(t, 2021, car.Year)
	require.Equal(t, int64(6000), car.DailyRentCost)
	require.Equal(t, "honda_1.jpg", car.Photo)

	require.Equal(t, oldCreatedAt, car.CreatedAt)
	require.True(t, car.UpdatedAt.After(oldUpdatedAt))

}

func TestUpdateCar_Errors(t *testing.T) {
	tests := []struct {
		name      string
		brand     string
		model     string
		year      int
		rent      int64
		wantError error
	}{
		{
			name:      "empty brand",
			brand:     "",
			model:     "Camry",
			year:      2020,
			rent:      5000,
			wantError: helper.ErrEmptyBrand,
		},
		{
			name:      "empty model",
			brand:     "Toyota",
			model:     "",
			year:      2020,
			rent:      5000,
			wantError: helper.ErrEmptyModel,
		},
		{
			name:      "invalid rent cost",
			brand:     "Toyota",
			model:     "Camry",
			year:      2020,
			rent:      -1,
			wantError: helper.ErrWrongRentCost,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			car, err := NewCar(
				"Toyota",
				"Camry",
				2020,
				3000,
				"camry_1.jpg",
			)

			require.NoError(t, err)
			require.NotNil(t, car)

			err = car.UpdateCar(
				tt.brand,
				tt.model,
				tt.year,
				int64(tt.rent),
				"photo.jpg",
			)

			require.ErrorIs(t, err, tt.wantError)
		})
	}
}