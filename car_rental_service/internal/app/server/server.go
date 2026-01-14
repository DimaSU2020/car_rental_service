package server

import (
	"github.com/DimaSU2020/car_rental_service/internal/http/handlers"
	"github.com/DimaSU2020/car_rental_service/internal/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	carService       service.CarService
	userService      service.UserService
	licenseService   service.LicenseService
	bookingService   service.BookingService
	router           *gin.Engine
}

func NewServer(
	carService       service.CarService, 
	userService      service.UserService,
	licenseService   service.LicenseService,
	bookingService   service.BookingService,
) *Server {

	s := &Server{
		carService     : carService,
		userService    : userService,
		licenseService : licenseService,
		bookingService : bookingService,
	}

	s.SetupRouter()

	return s
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) SetupRouter() {
	r := gin.New()
	r.Use(gin.Recovery())

	carHandlers     := handlers.NewCarHandlers(s.carService)
	userHandlers    := handlers.NewUserHandlers(s.userService)
	licenseHandlers := handlers.NewLicenseHandlers(s.licenseService)
	bookingHandlers := handlers.NewBookingHandlers(s.bookingService)

	v1 := r.Group("/v1")
	{
		cars := v1.Group("/cars")
		cars.GET("/",             carHandlers.List)
		cars.POST("/",           carHandlers.Create)
		cars.GET("/:carID",      carHandlers.GetByID)
		cars.PATCH("/:carID",    carHandlers.Update)
		cars.DELETE("/:carID",   carHandlers.Delete)
		
		users := v1.Group("/users")
		users.POST("/",                 userHandlers.Create)
		users.GET("/",                  userHandlers.List)
		users.GET("/:userID",           userHandlers.GetByID)
		users.PATCH("/:userID",         userHandlers.Update)
		users.DELETE("/:userID",        userHandlers.Delete)
		users.POST("/:userID/license",  licenseHandlers.Create)
		users.GET("/:userID/license",   licenseHandlers.GetByID)
		users.PATCH("/:userID/license", licenseHandlers.Update)
		users.DELETE("/:userID/license",licenseHandlers.Delete)

		bookings := v1.Group("/bookings")
		bookings.GET("/",             bookingHandlers.List)
		bookings.POST("/",            bookingHandlers.Create)
		bookings.GET("/:bookingID",   bookingHandlers.GetByID)
		// bookings.GET("/available",    bookingHandlers.IsAvailable)
		// bookings.GET("/user/:userId", bookingHandlers.GetByUserID)
	}

	s.router = r
}
