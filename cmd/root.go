package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "md2confl",
	Short: "md2confl - Convert markdown text to confluence wiki",
	Long: "md2confl - Convert markdown text to confluence wiki",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init()  {
	cobra.EnableCommandSorting = false
}
