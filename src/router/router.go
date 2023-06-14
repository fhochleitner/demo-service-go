package router

import (
	"gepaplexx/demo-service/api"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"sync/atomic"
)

var isHealthy *atomic.Value

func Initialize(config *api.Config) *gin.Engine {
	router := gin.New()
	isReady := &atomic.Value{}
	isReady.Store(false)
	isHealthy = &atomic.Value{}
	isHealthy.Store(false)

	// Global middlewares
	AddGlobalMiddleware(gin.Recovery())

	if !config.Development {
		AddGlobalMiddleware(jsonLoggerMiddleware())
		gin.SetMode(gin.ReleaseMode)
	} else {
		AddGlobalMiddleware(gin.Logger())
	}

	// Routes
	router.GET(config.MetricsPath, gin.WrapH(promhttp.Handler()))
	router.GET("/healthz", healthzMiddleware(isHealthy))
	router.GET("/readyz", readyzMiddleware(isReady))
	router.GET("/error", errorLogMiddleware())
	router.GET("/panic", panicMiddleware())
	router.GET("/", infoMiddleware(config))
	router.GET("/goroutines/:count", goRoutineSpawnerMiddleware())

	// Debugging Endpoints
	if config.Profiling {
		pprof.Register(router)
	}

	router.Use(globalMiddlewares...)
	isReady.Store(true)
	isHealthy.Store(true)
	return router
}

func AddGlobalMiddleware(middleware gin.HandlerFunc) {
	globalMiddlewares = append(globalMiddlewares, middleware)
}
