package cmd

import (
	"github.com/Halalins/backend/internal/bootstrap"
	"github.com/spf13/cobra"
)

// workerCmd represents the worker command.
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Starts a worker process",
	Long: `The worker command starts a worker process that performs specific tasks or computations.
This command is used in conjunction with other commands or tools to distribute workloads and improve overall efficiency.

Examples:

Start a worker to process incoming data streams and perform real-time analytics.
Run a worker to handle background jobs in a distributed job queue system.
`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartWorker()
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
