package router

import (
	"encoding/json"
	"fmt"
	"gepaplexx/demo-service/api"
	"gepaplexx/demo-service/consts"
	"gepaplexx/demo-service/logger"
	"gepaplexx/demo-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/debug"
	"strconv"
	"sync/atomic"
	"time"
)

var globalMiddlewares []gin.HandlerFunc

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("27.01.2003 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

func healthzMiddleware(isHealthz *atomic.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Healthz endpoint called")
		if isHealthz == nil || !isHealthz.Load().(bool) {
			c.Writer.WriteHeader(http.StatusServiceUnavailable)
			_, err := c.Writer.Write([]byte(http.StatusText(http.StatusServiceUnavailable) + ": Application is not healthy"))
			utils.CheckIfError(err)
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Application is healthy"))
		utils.CheckIfError(err)
		c.Next()
	}
}

func readyzMiddleware(isReady *atomic.Value) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Readyz endpoint called")
		if isReady == nil || !isReady.Load().(bool) {
			c.Writer.WriteHeader(http.StatusServiceUnavailable)
			_, err := c.Writer.Write([]byte(http.StatusText(http.StatusServiceUnavailable) + "\n"))
			utils.CheckIfError(err)
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Application is ready"))
		utils.CheckIfError(err)
		c.Next()
	}
}

func errorLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Error("Error-Logging endpoint called. Triggering error message")
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Error-Logging endpoint called. Triggering error message"))
		utils.CheckIfError(err)
		c.Next()
	}
}

func panicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Error("Request to panic received. Panicking...")
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte("Request to panic received. Panicking..."))
		utils.CheckIfError(err)
		isHealthy.Store(false)
		c.Next()
	}
}

func infoMiddleware(cfg *api.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("SysInfo endpoint called")
		var memstats = &runtime.MemStats{}
		runtime.ReadMemStats(memstats)

		commit, builddate, modified := "unknown", "unknown", "unknown"
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "vcs.revision" {
					commit = setting.Value
				}
				if setting.Key == "vcs.timestamp" {
					builddate = setting.Value
				}
				if setting.Key == "vcs.modified" {
					modified = setting.Value
				}
			}
		}

		applicationInfo := api.SysInfo{
			Version:      cfg.Version,
			GoVersion:    runtime.Version(),
			GitCommit:    commit,
			BuildDate:    builddate,
			MemLoad:      memstats.HeapInuse / 1024 / 1024,
			CpuAvailable: runtime.NumCPU(),
			MemAvailable: memstats.HeapSys / 1024 / 1024,
			GoRoutines:   runtime.NumGoroutine(),
			GitClean:     modified,
		}

		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte(applicationInfo.String()))
		utils.CheckIfError(err)
		c.Next()
	}
}

func goRoutineSpawnerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		count, nok := strconv.Atoi(c.Param("count"))
		if nok != nil {
			count = 0
		}
		logger.Info("Spawning %d goroutines", count)
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.Write([]byte(fmt.Sprintf("Spawning %d goroutines", count)))
		utils.CheckIfError(err)
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

func jokesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("Jokes endpoint called")
		c.Writer.WriteHeader(http.StatusOK)
		if c.Request.RequestURI == "/jokes/random" {
			_, err := c.Writer.Write([]byte(consts.Jokes[rand.Intn(len(consts.Jokes)-1)]))
			utils.CheckIfError(err)
		} else {
			for _, val := range consts.Jokes {
				_, err := c.Writer.Write([]byte(val + "\n"))
				utils.CheckIfError(err)
			}
		}
		c.Next()
	}
}

func prometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(c.Request.RequestURI, c.Request.Method))
		c.Next()
		responseStatus.WithLabelValues(string(c.Writer.Status()), c.Request.Method).Inc()
		totalRequests.WithLabelValues(c.Request.RequestURI, c.Request.Method).Inc()
		timer.ObserveDuration()
	}
}
