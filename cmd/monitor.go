package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"
)

var (
	interval time.Duration
)

var monitorCmd = &cobra.Command{
	Use:   "monitor [urls]",
	Short: "Monitor the health of specified URL(s) over time",
	Long:  `Continuously monitors the health of the specified URL(s) at the specified interval`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		monitorURL(ctx, args)
	},
}

func init() {
	monitorCmd.Flags().DurationVar(&interval, "interval", 2*time.Second, "Interval between checks")
	rootCmd.AddCommand(monitorCmd)
}

func monitorURL(ctx context.Context, urls []string) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		for _, url := range urls {
			checkURL(ctx, url, threshold, retries)
		}
	}
}
