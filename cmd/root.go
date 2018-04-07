package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "md2confl",
	Short:         "md2confl - Convert markdown text to confluence wiki text",
	Long:          "md2confl - Convert markdown text to confluence wiki text",
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute runs the root Cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v\n", err)
		fmt.Fprintf(os.Stdout, "%v\n", rootCmd.UsageString())
	}
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(versionCmd)
}
