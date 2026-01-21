package e2e

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/DimaSU2020/car_rental_service/internal/app/server"
	dbSQLite "github.com/DimaSU2020/car_rental_service/internal/infra/db/sqlite"
	repoSQLite "github.com/DimaSU2020/car_rental_service/internal/repo/sqlite"
	"github.com/DimaSU2020/car_rental_service/internal/service/cars"
)



const testPort = ":8080"

func setupTestServer(t *testing.T) (string, func()) {
	db, err := dbSQLite.Open(":memory:")
	if err != nil {
		t.Fatal(err)
	}

	repo := repoSQLite.NewCarRepo(db)
	svc  := cars.NewCarService(repo)
	srv  := server.NewServer(svc, nil, nil, nil)

	baseURL := "http://localhost"+testPort

	httpServer := &http.Server{
		Addr        : testPort,
		Handler     : srv.Router(),
		ReadTimeout : 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		if err := httpServer.ListenAndServe(); err != nil {
			t.Logf("server failed: %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	cleanup := func () {
		if err := db.Close(); err != nil {
            t.Logf("data base close error: %v", err)
        }

		ctx, cansel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cansel()

		if err := httpServer.Shutdown(ctx); err != nil {
			t.Logf("server shutdown error: %v", err)
		}

		<-done
	}

	return baseURL, cleanup
} 