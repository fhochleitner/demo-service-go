package model

type Config struct {
	Port        int    `mapstructure:"port"`
	MetricsPath string `mapstructure:"metricspath"`
	Profiling   bool   `mapstructure:"profiling"`
	Development bool   `mapstructure:"development"`
	Version     string `mapstructure:"version"`
}
