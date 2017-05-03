package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/kentaro-m/md2confl/util"
	"github.com/kentaro-m/md2confl/confluence"
)

var previewCmd = &cobra.Command{
	Use:   "preview [file path]",
	Short: "Shows a confluence wiki text converted from markdown.",
	Long: "Shows a confluence wiki text converted from markdown.",
	RunE: preview,
}

func preview(cmd *cobra.Command, args []string) error {
	data := util.ReadFile(args[0])
	output := confluence.Convert(data)
	fmt.Println(output)
	return nil
}

func init() {
	RootCmd.AddCommand(previewCmd)
}
