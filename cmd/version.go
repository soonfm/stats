package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Set by cmd.Execute("version")
var VERSION string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build:", VERSION)
		os.Exit(-1)
	},
}
