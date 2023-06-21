package services

import (
	"fmt"
	"gepaplexx-demos/demo-service-go/commons"
	"gepaplexx-demos/demo-service-go/logger"
	"gepaplexx-demos/demo-service-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

func ErrorLogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Error("Error-Logging endpoint called. Triggering error message")
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Error-Logging endpoint called. Triggering error message"))
		commons.CheckIfError(err)
		c.Next()
	}
}

func FreezeHandler(isHealthy *atomic.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Error("Request to freeze received. Panicking...")
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Request to panic received. Panicking..."))
		commons.CheckIfError(err)
		isHealthy.Store(false)
		c.Next()
	}
}

func GoRoutineSpawnerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		count, nok := strconv.Atoi(c.Param("count"))
		if nok != nil {
			count = 0
		}
		logger.Info("Spawning %d goroutines", count)
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte(fmt.Sprintf("Spawning %d goroutines", count)))
		commons.CheckIfError(err)
		for i := 0; i < count; i++ {
			go func() {
				for {
					time.Sleep(time.Minute * 1)
				}
			}()
		}
		c.Next()
	}
}

func RaiseErrorMetricHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Raising error metric")
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Raising error metric"))
		commons.CheckIfError(err)
		model.HttpError.WithLabelValues(c.Request.RequestURI, c.Request.Method, "500").Inc()
		c.Next()
	}
}
