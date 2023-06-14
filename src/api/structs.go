package api

import "fmt"

type Config struct {
	Port        int    `mapstructure:"port"`
	MetricsPath string `mapstructure:"metrics_path"`

	Profiling   bool   `mapstructure:"profiling"`
	Development bool   `mapstructure:"development"`
	Version     string `mapstructure:"version"`
}

type SysInfo struct {
	Version      string `json:"version"`
	GoVersion    string `json:"go_version"`
	GitCommit    string `json:"git_commit"`
	BuildDate    string `json:"build_date"`
	MemLoad      uint64 `json:"mem_load"`
	CpuAvailable int    `json:"cpu_available"`
	MemAvailable uint64 `json:"mem_available"`
	GoRoutines   int    `json:"go_processes"`
	GitClean     string `json:"git_clean"`
}

func (s *SysInfo) String() string {
	return fmt.Sprintf("Version: %s\nGo Version: %s\nGit Commit: %s\nBuild Date: %s\nGit Clean: %s\nMem Load: %d\nMem Available: %d\nCpu Available: %d\nGo Routines: %d\n",
		s.Version,
		s.GoVersion,
		s.GitCommit,
		s.BuildDate,
		s.GitClean,
		s.MemLoad,
		s.MemAvailable,
		s.CpuAvailable,
		s.GoRoutines,
	)
}
