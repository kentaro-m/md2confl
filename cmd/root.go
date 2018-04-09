package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"github.com/kentaro-m/md2confl/confluence"
	"github.com/kentaro-m/md2confl/utils/file"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// Current version number
const (
	VERSION = "0.1.2"
)

var (
	version bool
	help    bool
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

	if cmd.Flag("version").Changed {
		fmt.Fprintf(os.Stdout, "v%v\n", VERSION)
		return nil
	}

	// Input from stdin
	if !terminal.IsTerminal(int(syscall.Stdin)) {
		data, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			return err
		}

		confluence := confluence.Confluence{}
		confluence.Convert(data)
		confluence.Show()

		return nil
	}

	switch {
	case len(args) == 0:
		return errors.New("File path not found")
	case len(args) > 1:
		return errors.New("Too many arguments")
	}

	// Input from the file
	file := file.File{}
	if err := file.Open(args[0]); err != nil {
		return err
	}

	confluence := confluence.Confluence{}
	confluence.Convert(file.Data)
	confluence.Show()

	return nil
}

func init() {
	cobra.EnableCommandSorting = false
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Output the version number")
	rootCmd.Flags().BoolVarP(&help, "help", "h", false, "Output usage information")
}

// Execute runs the root Cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v\n", err)
		fmt.Fprintf(os.Stdout, "%v\n", rootCmd.UsageString())
	}
}
