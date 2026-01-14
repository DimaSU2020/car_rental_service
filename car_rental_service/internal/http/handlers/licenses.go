package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/DimaSU2020/car_rental_service/internal/http/dto"
	"github.com/DimaSU2020/car_rental_service/internal/service"

	"github.com/gin-gonic/gin"
)

type LicenseHandlers struct {
	service service.LicenseService
}

func NewLicenseHandlers(service service.LicenseService) *LicenseHandlers {
	return &LicenseHandlers{service: service}
}

func (h *LicenseHandlers) List(c *gin.Context) {
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

	resp := make([]*dto.LicenseResponse, 0, len(items))
	for _, l := range items {
		resp = append(resp, dto.LicenseToResponse(l))
	}

	writeOK(c, gin.H{"items": resp, "count": len(resp)})
}

func (h *LicenseHandlers) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("licenseID"), 10, 64)
	if err != nil {
		writeBadRequest(c, "invalid id")
		return
	}

	l, err := h.service.GetByID(c.Request.Context(), id)
	if errors.Is(err, service.ErrLicenseNotFound) {
		writeNotFound(c, "license not found")
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	writeOK(c, dto.LicenseToResponse(l))
}

func (h *LicenseHandlers) Create(c *gin.Context) {
	var req dto.CreateLicenseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	license := service.CreateLicenseInput{
		Number         : req.Number,
		IssuanceDate   : req.IssuanceDate,
		ExpirationDate : req.ExpirationDate,
		CreatedAt      : req.CreatedAt,
		UpdatedAt      : req.UpdatedAt,
	}

	l, err := h.service.Create(c.Request.Context(), license)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	writeOK(c, dto.LicenseToResponse(l))
}

func (h *LicenseHandlers) Update(c *gin.Context) {
	var req dto.UpdateLicenseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	licenseID := c.Param("licenseID")
    id, err := strconv.ParseInt(licenseID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid license ID"})
        return
    }

	in := service.UpdateLicenseInput{
		ID             : id,
		Number         : req.Number,
		IssuanceDate   : req.IssuanceDate,
		ExpirationDate : req.ExpirationDate,
		CreatedAt      : req.CreatedAt,
		UpdatedAt      : req.UpdatedAt,
	}

	err = h.service.Update(c.Request.Context(), in)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	writeUpdated(c, "message: update license successful")
}

func (h *LicenseHandlers) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("licenseID"), 10, 64)
	if err != nil {
		writeBadRequest(c, "invalid id")
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if errors.Is(err, service.ErrLicenseNotFound) {
		writeNotFound(c, "license not found")
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	writeOK(c, "message: delete license successful")
}