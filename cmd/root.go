package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go-cli-app2/logger"
)

var (
	logfile     string
	l           *slog.Logger
	threshold   float64
	retries     int
	silent      bool
	verbose     bool
	versionFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-app2",
	Short: "A tool for monitoring health status and responsiveness of web applications",
	Long: `The healthcheck command is designed to assess the health and
responsiveness of specified web applications. It sends HTTP requests
to URLs provided by the user, evaluating whether the services are
accessible and how quickly they respond. This command supports both
immediate, one-off checks and continuous monitoring, allowing users
to specify intervals for ongoing health assessments. With additional
flags for customization, users can tailor the command to meet various
monitoring needs, from simple uptime checks to detailed performance
analysis.`,
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			printVersion()
			os.Exit(0)
		}
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		silent, _ = cmd.Flags().GetBool("silent")
		verbose, _ := cmd.Flags().GetBool("verbose")
		l = logger.New(logfile, silent, verbose)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&logfile, "logfile", "healthcheck.log", "log file path")
	rootCmd.PersistentFlags().Float64Var(
		&threshold, "threshold", 0.5, "Threshold for slow response")
	rootCmd.PersistentFlags().IntVar(
		&retries, "r", 3, "Number of retries")
	rootCmd.PersistentFlags().BoolVar(
		&silent, "s", false, "Silent mode")
	rootCmd.PersistentFlags().BoolVar(
		&verbose, "v", false, "Verbose mode")
	rootCmd.Flags().BoolVar(&versionFlag, "version", false, "Print version")
}

func Execute() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
