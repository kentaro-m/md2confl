package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Current version number
const (
	VERSION = "0.1.2"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version number",
	Long:  "Output the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%v\n", VERSION)
	},
}
