package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version   = "dev"
	buildDate = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Printf("Version %s\\nBuild Date: %s\\n", version, buildDate)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
