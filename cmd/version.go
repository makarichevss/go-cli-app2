package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	buildDate = time.Now().Format("2006-01-02")
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Printf("Version: %s\nBuild Date: %s\n", version, buildDate)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
