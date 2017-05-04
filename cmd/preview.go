package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kentaro-m/md2confl/confluence"
	"github.com/kentaro-m/md2confl/utils/file"
)

var previewCmd = &cobra.Command{
	Use:   "preview [file path]",
	Short: "Show a confluence wiki text converted from markdown.",
	Long: "Show a confluence wiki text converted from markdown.",
	RunE: preview,
}

func preview(cmd *cobra.Command, args []string) error {
	file := file.File{}
	if err := file.Open(args[0]); err != nil {
		return err
	}

	confluence := confluence.Confluence{}
	confluence.Convert(file.Data)
	confluence.Preview()

	return nil
}

func init() {
	RootCmd.AddCommand(previewCmd)
}
