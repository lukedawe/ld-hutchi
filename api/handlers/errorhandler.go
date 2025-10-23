package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Production aware error handler
func HandleError(c *gin.Context, status int, err error, productionErrorMessage string) {
	log.Printf("API Error [%d]: %s - Request: %s %s",
		status, err.Error(), c.Request.Method, c.Request.URL.Path)

	if gin.IsDebugging() {
		c.AbortWithError(status, err)
		return
	}
	if productionErrorMessage == "" {
		c.AbortWithStatusJSON(status, gin.H{"error": "Request failed"})
	}
	c.AbortWithStatusJSON(status, gin.H{"error": productionErrorMessage})
}
