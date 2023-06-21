package services

import (
	"fmt"
	"gepaplexx-demos/demo-service-go/model"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	_ = prometheus.Register(model.TotalRequests)
	_ = prometheus.Register(model.HttpDuration)
	_ = prometheus.Register(model.ResponseStatus)
	_ = prometheus.Register(model.HttpError)
}

func CommonMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(model.HttpDuration.WithLabelValues(c.Request.RequestURI, c.Request.Method))
		c.Next()
		model.ResponseStatus.WithLabelValues(fmt.Sprintf("%d", c.Writer.Status()), c.Request.Method).Inc()
		model.TotalRequests.WithLabelValues(c.Request.RequestURI, c.Request.Method).Inc()
		timer.ObserveDuration()
	}
}
