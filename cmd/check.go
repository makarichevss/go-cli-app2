package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the health of specified URL(s)",
	Long:  `Performs a health check by sending a request to the specified URL(s) and reports the status.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, url := range args {
			checkURL(url, threshold, retries)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func checkURL(url string, threshold float64, retries int) {
	var resp *http.Response
	var err error
	var duration time.Duration

	for attempt := 0; attempt <= retries; attempt++ {
		start := time.Now()
		resp, err = http.Get(url)
		duration = time.Since(start)

		if err == nil && resp != nil {
			resp.Body.Close()
			break
		}

		if attempt < retries {
			fmt.Fprintf(os.Stderr, "Attempt %d failed, retrying...", attempt+1)
			l.Warn("failed retrying", "url", url, "attempts", attempt+1)
			time.Sleep(time.Second * 2) // Backoff
		}
	}

	if err != nil {
		l.Error("fetching error", "url", url, "retries", retries, "err", err)
		return
	}

	if resp != nil {
		defer resp.Body.Close()

		if duration.Seconds() > threshold && verbose {
			l.Error("exceeding threshold", "url", url, "duration", duration, "threshold", threshold)
		}
		l.Info("Successful check", "url", url, "status code", resp.StatusCode, "duration", duration)
	}
}
