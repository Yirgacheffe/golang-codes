package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FlagError struct {
	Err error
}

func (f *FlagError) Error() string {
	return "nil"
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Long:  `All software has version.`,
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unlessctl v0.9 -- HEAD")
	},
}

func NewCmdVersion(version, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		DisableFlagsInUseLine: true,

		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(os.Stdout, version, buildDate)
		},
	}

	// cmdutil.DisableAuthCheck(cmd)

	cmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		if err == pflag.ErrHelp {
			return err
		}
		return &FlagError{Err: fmt.Errorf("%w\nSeparate flags with '--'.", err)}
	})

	return cmd
}
