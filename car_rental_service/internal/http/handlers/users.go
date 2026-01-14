package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/DimaSU2020/car_rental_service/internal/http/dto"
	"github.com/DimaSU2020/car_rental_service/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct {
	service service.UserService
}

func NewUserHandlers(service service.UserService) *UserHandlers {
	return &UserHandlers{service: service}
}

func (h *UserHandlers) List(c *gin.Context) {
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

	resp := make([]*dto.UserResponse, 0, len(items))
	for _, u := range items {
		resp = append(resp, dto.UserToResponse(u))
	}

	writeOK(c, gin.H{"items": resp, "count": len(resp)})
}

func (h *UserHandlers) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		writeBadRequest(c, "invalid id")
		return
	}

	u, err := h.service.GetByID(c.Request.Context(), id)
	if errors.Is(err, service.ErrUserNotFound) {
		writeNotFound(c, "user not found")
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	writeOK(c, dto.UserToResponse(u))
}

func (h *UserHandlers) Create(c *gin.Context) {
	var req dto.CreateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := service.CreateUserInput{
		Name     : req.Name,
		Email    : req.Email,
		Birthday : req.Birthday,
		IsAdmin  : req.IsAdmin,
		LicenseId: req.LicenseId,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}

	u, err := h.service.Create(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	writeOK(c, dto.UserToResponse(u))
}

func (h *UserHandlers) Update(c *gin.Context) {
	var req dto.UpdateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("userID")
    id, err := strconv.ParseInt(userID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
        return
    }

	in := service.UpdateUserInput{
		ID        : id,
		Name      : req.Name,
		Email     : req.Email,
		Birthday  : req.Birthday,
		IsAdmin   : req.IsAdmin,
		LicenseId : req.LicenseId,
		UpdatedAt : req.UpdatedAt,
	}

	err = h.service.Update(c.Request.Context(), in)
	
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	writeUpdated(c, "message: update user successful")
}

func (h *UserHandlers) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("userID"), 10, 64)
	if err != nil {
		writeBadRequest(c, "invalid id")
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if errors.Is(err, service.ErrUserNotFound) {
		writeNotFound(c, "user not found")
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	writeOK(c, "message: delete user successful")
}