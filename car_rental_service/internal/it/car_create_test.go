package it

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DimaSU2020/car_rental_service/internal/http/handlers"
	dbSQLite "github.com/DimaSU2020/car_rental_service/internal/infra/db/sqlite"
	repoSQLite "github.com/DimaSU2020/car_rental_service/internal/repo/sqlite"
	"github.com/DimaSU2020/car_rental_service/internal/service/cars"
	"github.com/gin-gonic/gin"
)


func TestIT_Car_Create(t *testing.T) {
	db, err := dbSQLite.Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := repoSQLite.NewCarRepo(db)
	svc  := cars.NewCarService(repo)
	handler := handlers.NewCarHandlers(svc)
	router := gin.Default()
	router.POST("/v1/cars/", handler.Create)

	testCases := []struct {
		name       string
		body       string
		wantStatus int
	} {
		{
			name: "valid request",
			body: `{
				"brand": "Toyota",
				"model": "Camry",
				"year": 2020,
				"rent": 3000,
				"photo": "toyota.jpg"
			}`,
			wantStatus: http.StatusCreated,
		},
		{
			name: "missing brand",
			body: `{
				"brand": "",
				"model": "Camry",
				"year": 2020,
				"rent": 3000,
				"photo": "toyota.jpg"
			}`,
			wantStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/v1/cars/", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Fatalf("expected %d, got %d w/ %s", tt.wantStatus, w.Code, w.Body.String())
			}
		})
	}
}