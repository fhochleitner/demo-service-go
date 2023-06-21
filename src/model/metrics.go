package model

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var TotalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{
		"path",
		"method",
	},
)

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{
		"status",
		"method",
	},
)

var HttpDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{
		"path",
		"method",
	},
)

var HttpError = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_response_error",
		Help: "Number of errors in HTTP requests.",
	}, []string{
		"path",
		"method",
		"status",
	},
)
