/*
Copyright © 2023 Felix Hochleitner <felix.hochleitner@gepardec.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"gepaplexx/demo-service/logger"
	"gepaplexx/demo-service/router"
	"gepaplexx/demo-service/utils"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "starts the demo service server",
	Long: `mini application that serves various endpoints for demo purposes.
	For example:
		/healthz
		/livez
		/readyz
		/metrics
		/error
		/panic
		/ping
	and debugging endpoints:
`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.
	serverCmd.PersistentFlags().StringVar(&Config.MetricsPath, "metrics-path", "/metrics", "path to metrics endpoint")
}

func serve() {
	logger.Info("starting server on localhost:%d", Config.Port)
	router := router.Initialize(&Config)

	err := router.Run(fmt.Sprintf(":%d", Config.Port))
	utils.CheckIfError(err)

}
