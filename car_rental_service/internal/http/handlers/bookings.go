package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/DimaSU2020/car_rental_service/internal/http/dto"
	"github.com/DimaSU2020/car_rental_service/internal/service"

	"github.com/gin-gonic/gin"
)

type BookingHandlers struct {
	service service.BookingService
}

func NewBookingHandlers(service service.BookingService) *BookingHandlers {
	return &BookingHandlers{service: service}
}

func (h *BookingHandlers) List(c *gin.Context) {
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil { 
		limit = 5
	}
	
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil { 
		offset = 0
	}

	if h.service == nil {
        writeError(c, http.StatusInternalServerError, "service not initialized")
        return
    }

	items, err := h.service.List(c.Request.Context(), limit, offset)

	if err != nil {
		writeNotFound(c, err.Error())
		return
	}

	resp := make([]*dto.BookingResponse, 0, len(items))
	for _, b := range items {
		resp = append(resp, dto.BookingToResponse(b))
	}

	writeOK(c, gin.H{"items": resp, "count": len(resp)})
}

func (h *BookingHandlers) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("bookingID"), 10, 64)
	if err != nil {
		writeBadRequest(c, "invalid id")
		return
	}

	b, err := h.service.GetByID(c.Request.Context(), id)
	if errors.Is(err, service.ErrBookingNotFound) {
		writeNotFound(c, "booking not found")
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	writeOK(c, dto.BookingToResponse(b))
}

func (h *BookingHandlers) Create(c *gin.Context) {
	var req dto.CreateBookingRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	booking := service.CreateBookingInput{
		ID_car      : req.ID_car,
		ID_user     : req.ID_user,    
		Start_day   : req.Start_day,
		End_day     : req.End_day,
		Daily_cost  : req.Daily_cost,
		Status      : req.Status,
		CreatedAt   : req.CreatedAt,
		UpdatedAt   : req.UpdatedAt,
	}

	b, err := h.service.Create(c.Request.Context(), booking)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	writeOK(c, dto.BookingToResponse(b))
}
