package router

import (
	"gepaplexx/demo-service/api"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Initialize(config *api.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Handle("GET", config.MetricsPath, gin.WrapH(promhttp.Handler()))

	if config.Profiling {
		pprof.Register(router)
	}

	if !config.Development {
		router.Use(jsonLoggerMiddleware())
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
	}

	return router
}
