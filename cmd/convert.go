package cmd

import (
	"errors"

	"github.com/kentaro-m/md2confl/confluence"
	"github.com/kentaro-m/md2confl/utils/file"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert [file path]",
	Short: "Output the confluence wiki text",
	Long:  "Output the confluence wiki text",
	RunE:  convert,
}

func convert(cmd *cobra.Command, args []string) error {

	switch {
	case len(args) == 0:
		return errors.New("File path not found")
	case len(args) > 1:
		return errors.New("Too many arguments")
	}

	file := file.File{}
	if err := file.Open(args[0]); err != nil {
		return err
	}

	confluence := confluence.Confluence{}
	confluence.Convert(file.Data)
	confluence.Show()

	return nil
}
