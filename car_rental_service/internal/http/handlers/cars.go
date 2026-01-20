package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/DimaSU2020/car_rental_service/internal/http/dto"
	"github.com/DimaSU2020/car_rental_service/internal/service/cars"
	"github.com/gin-gonic/gin"
)

type CarHandlers struct {
	service cars.CarService
}

func NewCarHandlers(service cars.CarService) *CarHandlers {
	return &CarHandlers{service: service}
}

func (h *CarHandlers) List(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		writeBadRequest(c, "invalid limit")
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil || offset <= 0 {
		writeBadRequest(c, "invalid offset")
		return
	}

	items, err := h.service.List(c.Request.Context(), limit, offset)
	if err != nil {
		writeInternalError(c, err)
		return
	}

	resp := make([]*dto.CarResponse, 0, len(items))

	for _, car := range items {
		resp = append(resp, dto.CarToResponse(car)) 
	}

	writeOK(c, gin.H{"items": resp, "count": len(resp)})
}

func (h *CarHandlers) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("carID"), 10, 64)
	if err != nil || id <= 0 {
		writeBadRequest(c, "invalid car id")
		return
	}

	car, err := h.service.GetByID(c.Request.Context(), id)
	if errors.Is(err, cars.ErrCarNotFound) {
		writeNotFound(c, "car not found")
		return
	}
	if err != nil {
		writeInternalError(c, err)
		return
	}

	writeOK(c, dto.CarToResponse(car))
}

func (h *CarHandlers) Create(c *gin.Context) {
	var req dto.CreateCarRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeBadRequest(c, "invalid json body")
		return
	}

	in := cars.CreateCarInput {
		Brand        : req.Brand,
		Model        : req.Model,
		Year         : req.Year,
		DailyRentCost: req.DailyRentCost,
		Photo        : req.Photo,
	}

	car, err := h.service.Create(c.Request.Context(), in)
	if errors.Is(err, cars.ErrInvalidCarData) {
		writeUnprocessable(c, err.Error())
		return
	}

	if err != nil {
		writeInternalError(c, err)
		return
	}

	writeCreateOK(c, dto.CarToResponse(car))
}

func (h *CarHandlers) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("carID"), 10, 64)
	if err != nil || id <=0 {
		writeBadRequest(c, "invalid car id")
		return
	}

	var req dto.CreateCarRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		writeBadRequest(c, "invalid json body")
		return
	}

	in := cars.UpdateCarInput{
		ID           : id,
		Brand        : req.Brand,
		Model        : req.Model,
		Year         : req.Year,
		DailyRentCost: req.DailyRentCost,
		Photo        : req.Photo,
	}

	err = h.service.Update(c.Request.Context(), in)
	if errors.Is(err, cars.ErrCarNotFound) {
		writeNotFound(c, "car not found")
		return
	}
	if errors.Is(err, cars.ErrInvalidCarData) {
		writeUnprocessable(c, err.Error())
		return
	}
	if err != nil {
		writeInternalError(c, err)
		return
	}

	writeOK(c, gin.H{"message": "car updated"})
}

func (h *CarHandlers) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("carID"), 10, 64)
	if err != nil || id <= 0 {
		writeBadRequest(c, "invalid car id")
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if errors.Is(err, cars.ErrCarNotFound) {
		writeNotFound(c, "car not found")
		return
	}
	if err != nil {
		writeInternalError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
