package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	VERSION = "0.1.2"
)

var RootCmd = &cobra.Command{
	Use:           "md2confl",
	Short:         "md2confl - Convert markdown text to confluence wiki",
	Long:          "md2confl - Convert markdown text to confluence wiki",
	SilenceErrors: true,
	SilenceUsage:  true,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of md2confl",
	Long:  "Print the version number of md2confl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%v\n", VERSION)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v\n", err)
	}
}

func init() {
	cobra.EnableCommandSorting = false
	RootCmd.AddCommand(versionCmd)
}
