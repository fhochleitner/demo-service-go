package cmd

import (
	"fmt"
	"gepaplexx-demos/demo-service-go/logger"
	"gepaplexx-demos/demo-service-go/model"
	"gepaplexx-demos/demo-service-go/services"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"sync/atomic"
)

var isHealthy *atomic.Value = new(atomic.Value)
var isReady *atomic.Value = new(atomic.Value)
var engine *gin.Engine

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts the demo service server",
	Long: `mini application that serves various endpoints for demo purposes.
	For example:
		# TODO add all endpoints to documentation
`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func serve() {
	logger.Info("starting server on localhost: %d", Config.Port)
	isHealthy.Store(false)
	isReady.Store(false)
	engine = gin.New()

	configureRouter(&Config)
	registerRoutes()

	isReady.Store(true)
	isHealthy.Store(true)

	err := engine.Run(fmt.Sprintf(":%d", Config.Port))
	if err != nil {
		logger.Error("error starting server: %v", err)
	}
}

func configureRouter(cfg *model.Config) {
	engine.Use(gin.Recovery())
	engine.Use(services.CommonMetricsMiddleware())

	if cfg.Development == false {
		gin.SetMode(gin.ReleaseMode)
		engine.Use(services.JsonLoggerMiddleware())
	}

	if cfg.Development == true {
		engine.Use(gin.Logger())
		gin.SetMode(gin.DebugMode)
	}

	if cfg.Profiling == true {
		pprof.Register(engine)
	}
}

func registerRoutes() {
	engine.GET(Config.MetricsPath, gin.WrapH(promhttp.Handler()))
	engine.GET("/healthz", services.HealthzHandler(isHealthy))
	engine.GET("/readyz", services.ReadyzHandler(isReady))
	//engine.GET("/", services.InfoHandler(&Config))
	engine.GET("/error", services.ErrorLogHandler())
	engine.GET("/freeze", services.FreezeHandler(isHealthy))
	engine.GET("/goroutines/:count", services.GoRoutineSpawnerHandler())

	//e.GET("/goroutines/:count", services.GoRoutineSpawnerMiddleware())
	//e.GET("/jokes", services.JokesMiddleware())
	//e.GET("/jokes/random", services.JokesMiddleware())
	//e.Run(":8080")
}
