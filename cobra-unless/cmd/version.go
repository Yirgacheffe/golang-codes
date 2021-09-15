package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

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
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(os.Stdout, version, buildDate)
		},
	}
	// cmdutil.DisableAuthCheck(cmd)
	return cmd
}
