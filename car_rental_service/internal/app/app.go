package app

import (
	"log"
	"net/http"
	"time"

	"github.com/DimaSU2020/car_rental_service/config"
	"github.com/DimaSU2020/car_rental_service/internal/app/server"
	dbSQLite "github.com/DimaSU2020/car_rental_service/internal/infra/db/sqlite"
	repoSQLite "github.com/DimaSU2020/car_rental_service/internal/repo/sqlite"
	"github.com/DimaSU2020/car_rental_service/internal/service"
)

func BuildHTTPServer (c *config.Config) (*http.Server, func(), error) {
	db, err := dbSQLite.Open(c.DBPath)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func () {
		if err := db.Close(); err != nil {
			log.Printf("failed to close db: %v", err)
		}
	}
	
	carRepo := repoSQLite.NewCarRepo(db)
	carSvc  := service.NewCarService(carRepo)

	userRepo := repoSQLite.NewUserRepo(db)
	userSvc  := service.NewUserService(userRepo)

	licenseRepo := repoSQLite.NewLicenseRepo(db)
	licenseSvc  := service.NewLicenseService(licenseRepo)

	bookingRepo := repoSQLite.NewBookingRepo(db)
	bookingSvc  := service.NewBookingService(bookingRepo)

	srv := server.NewServer(
		carSvc,
		userSvc,
		licenseSvc,
		bookingSvc,
	)

	readTimeout  := time.Duration(c.ReadTimeout)
	writeTimeout := time.Duration(c.WriteTimeout)

	return &http.Server{
		Addr         : c.Port,
		Handler      : srv.Router(),
		ReadTimeout  : readTimeout,
		WriteTimeout : writeTimeout,
	}, cleanup, nil
}