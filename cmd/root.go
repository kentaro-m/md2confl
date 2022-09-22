package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"github.com/kentaro-m/md2confl/confluence"
	"github.com/kentaro-m/md2confl/utils/file"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// nolint // these variables are set in build step
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

var (
	isShowVersion bool
)

var rootCmd = &cobra.Command{
	Use:           "md2confl [file path]",
	Short:         "Output the confluence wiki text",
	Long:          "Output the confluence wiki text",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          Run,
}

// Run executes md2confl command
func Run(cmd *cobra.Command, args []string) error {

	if isShowVersion {
		fmt.Fprintf(os.Stdout, "v%v\n", version)
		return nil
	}

	// Input from stdin
	if !term.IsTerminal(int(syscall.Stdin)) {
		data, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			return err
		}

		if len(bytes.TrimSpace(data)) > 0 {
			output(data)
			return nil
		}
	}

	switch {
	case len(args) == 0:
		return errors.New("file path not found")
	case len(args) > 1:
		return errors.New("too many arguments")
	}

	// Input from the file
	file := file.File{}
	if err := file.Open(args[0]); err != nil {
		return err
	}

	output(file.Data)

	return nil
}

func output(input []byte) {
	c := confluence.Confluence{}
	c.Convert(input)
	fmt.Fprint(os.Stdout, c.Contents())
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.Flags().BoolVarP(&isShowVersion, "version", "v", false, "Output the version number")
}

// Execute runs the root Cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", rootCmd.UsageString())
		os.Exit(1)
	}
}
