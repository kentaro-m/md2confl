package cmd

import (
	"github.com/spf13/cobra"
	"github.com/kentaro-m/md2confl/utils/util"
	"github.com/kentaro-m/md2confl/utils/file"
	"github.com/kentaro-m/md2confl/confluence"
)

var copyCmd = &cobra.Command{
	Use:   "copy [file path]",
	Short: "Copy a confluence wiki text converted from markdown.",
	Long: "Copy a confluence wiki text converted from markdown.",
	RunE: copy,
}

func copy(cmd *cobra.Command, args []string) error {
	file := file.File{}
	if err := file.Open(args[0]); err != nil {
		return err
	}

  confluence := confluence.Confluence{}
	confluence.Convert(file.Data)
	return util.Copy(confluence.Contents)
}

func init() {
	RootCmd.AddCommand(copyCmd)
}
