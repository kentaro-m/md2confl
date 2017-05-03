package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/atotto/clipboard"
	"github.com/kentaro-m/md2confl/util"
	"github.com/kentaro-m/md2confl/confluence"
)

var copyCmd = &cobra.Command{
	Use:   "copy [file path]",
	Short: "Copy a confluence wiki text converted from markdown.",
	Long: "Copy a confluence wiki text converted from markdown.",
	RunE: copy,
}

func copy(cmd *cobra.Command, args []string) error {
	data := util.ReadFile(args[0])
	output := confluence.Convert(data)
	fmt.Println("Copied to clipboard!")
  return clipboard.WriteAll(output)
}

func init() {
	RootCmd.AddCommand(copyCmd)
}
