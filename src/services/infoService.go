package services

import (
	"gepaplexx-demos/demo-service-go/commons"
	"gepaplexx-demos/demo-service-go/logger"
	"gepaplexx-demos/demo-service-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"runtime/debug"
)

func InfoHandler(cfg *model.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("SysInfo endpoint called")
		var memstats = &runtime.MemStats{}
		runtime.ReadMemStats(memstats)

		commit, builddate, modified := "unknown", "unknown", "unknown"
		if info, ok := debug.ReadBuildInfo(); ok {
			commit, builddate, modified = getBuildInformation(info)
		}

		applicationInfo := model.SysInfo{
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
		commons.CheckIfError(err)
		c.Next()
	}
}

func getBuildInformation(info *debug.BuildInfo) (string, string, string) {
	var commit, builddate, modified string

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
	return commit, builddate, modified
}
