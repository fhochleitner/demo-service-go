package services

import (
	"gepaplexx-demos/demo-service-go/commons"
	"gepaplexx-demos/demo-service-go/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

func HealthzHandler(isHealthz *atomic.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Healthz endpoint called")
		if isHealthz == nil || !isHealthz.Load().(bool) {
			respondUnavailable(c)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Application is healthy"))
		commons.CheckIfError(err)
		c.Next()
	}
}

func ReadyzHandler(isReady *atomic.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Readyz endpoint called")
		if isReady == nil || !isReady.Load().(bool) {
			respondUnavailable(c)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Application is ready"))
		commons.CheckIfError(err)
		c.Next()
	}
}

func respondUnavailable(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusServiceUnavailable)
	_, err := c.Writer.Write([]byte(http.StatusText(http.StatusServiceUnavailable) + ": Application is not available"))
	commons.CheckIfError(err)
	c.AbortWithStatus(http.StatusServiceUnavailable)
}
