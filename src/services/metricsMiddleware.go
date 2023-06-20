package services

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func init() {
	_ = prometheus.Register(totalRequests)
	_ = prometheus.Register(httpDuration)
	_ = prometheus.Register(responseStatus)
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{
		"path",
		"method",
	},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{
		"status",
		"method",
	},
)

var httpDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{
		"path",
		"method",
	},
)

func CommonMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(c.Request.RequestURI, c.Request.Method))
		c.Next()
		responseStatus.WithLabelValues(string(rune(c.Writer.Status())), c.Request.Method).Inc()
		totalRequests.WithLabelValues(c.Request.RequestURI, c.Request.Method).Inc()
		timer.ObserveDuration()
	}
}
