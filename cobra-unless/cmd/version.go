package cmd

import (
	"fmt"

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
