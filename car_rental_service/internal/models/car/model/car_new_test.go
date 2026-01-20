package model

import (
	"testing"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/models/helper"
	"github.com/stretchr/testify/require"
)

func TestNewCar_OK(t *testing.T) {
	car, err := NewCar(
		"Toyota",
		"Camry",
		2020,
		3000,
		"camry_1.jpg",
	)

	require.NoError(t, err)
	require.NotNil(t, car)

	require.Equal(t, "Toyota", car.Brand)
	require.Equal(t, "Camry", car.Model)
	require.Equal(t, 2020, car.Year)
	require.Equal(t, int64(3000), car.DailyRentCost)
	require.Equal(t, "camry_1.jpg", car.Photo)

	require.False(t, car.CreatedAt.IsZero())
	require.False(t, car.UpdatedAt.IsZero())
	require.WithinDuration(t, time.Now(), car.CreatedAt, time.Second)
}

func TestNewCar_Errors(t *testing.T) {
	tests := []struct {
		name      string
		brand     string
		model     string
		year      int
		rent      int64
		photo     string
		wantError error
	}{
		{
			name      : "empty brand",
			brand     : "",
			model     : "Camry",
			year      : 2020,
			rent      : 3000,
			wantError : helper.ErrEmptyBrand,
		},
		{
			name      : "empty model",
			brand     : "Toyota",
			model     : "",
			year      : 2020,
			rent      : 3000,
			wantError : helper.ErrEmptyModel,
		},
		{
			name      : "empty model",
			brand     : "Toyota",
			model     : "Camry",
			year      : 2029,
			rent      : 3000,
			wantError : helper.ErrWrongYear,
		},
		{
			name      : "empty model",
			brand     : "Toyota",
			model     : "Camry",
			year      : 2020,
			rent      : 0,
			wantError : helper.ErrWrongRentCost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			car, err := NewCar(
				tt.brand,
				tt.model,
				tt.year,
				int64(tt.rent),
				"camry_1.jpg",
			)

			require.Nil(t, car)
			require.ErrorIs(t, err, tt.wantError)
		})
	}
}