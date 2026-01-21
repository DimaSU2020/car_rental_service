package it

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/DimaSU2020/car_rental_service/internal/http/handlers"
	dbSQLite "github.com/DimaSU2020/car_rental_service/internal/infra/db/sqlite"
	repoSQLite "github.com/DimaSU2020/car_rental_service/internal/repo/sqlite"
	"github.com/DimaSU2020/car_rental_service/internal/service/cars"
	"github.com/gin-gonic/gin"
)

func TestIT_Car_GetByID(t *testing.T) {
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
	router.GET("/v1/cars/:carID", handler.GetByID)

	var created struct {
		ID int64 `json:"id"`
	}

	var body = `{
			"brand": "Toyota",
			"model": "Camry",
			"year": 2020,
			"rent": 3000,
			"photo": "toyota.jpg"
	 	}`

	t.Run("create car", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/v1/cars/", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code == http.StatusCreated {
			_ = json.Unmarshal(w.Body.Bytes(), &created)
		} else {
			t.Fatalf("create failed: %d %s", w.Code, w.Body.String())
		}
	})

	t.Run("valid get car by id", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/v1/cars/"+strconv.FormatInt(created.ID, 10),
			nil,
		)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
		}
	})

	t.Run("invalid get car by id - not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/cars/999", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Fatalf("expected 404, got %d: %s", w.Code, w.Body.String())
		}
	})
}