package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func writeBadRequest(c *gin.Context, msg string) {
    c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "bad_request", "message": msg}})
}

func writeNotFound(c *gin.Context, msg string) {
    c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "not_found", "message": msg}})
}

func writeOK(c *gin.Context, payload any) { 
	c.JSON(http.StatusOK, payload) 
}

func writeCreateOK(c *gin.Context, payload any) {
	c.JSON(http.StatusCreated, payload)
}

func writeError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func writeUpdated(c *gin.Context, payload any) {
	c.JSON(http.StatusNoContent, payload)
}

func writeInternalError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
}

func writeUnprocessable(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": msg})
}

func writeConflict(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": msg})
}
