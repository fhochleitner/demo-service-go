package cmd

import (
	"gepaplexx/demo-service/logger"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of demo-service",
	Long:  `All software has versions. This is demo-service's`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("application version: %s", Version)
	},
}
