package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type ListOpts struct {
	io.Writer
	HTTPClient func() (*http.Client, error)
}

var orgName string

func NewCmdList(runF func(*ListOpts) error) *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:  "list",
		Args: cobra.ExactArgs(1),
		// Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(args[0])
			org, _ := cmd.Flags().GetString("org")
			fmt.Println("org:", org)
			if runF != nil {
				return runF(opts)
			} else {
				return listRun(opts)
			}
		},
		Annotations: map[string]string{
			"help:environment": `HOST: make the request to...`,
		},
	}

	cmd.Flags().StringVarP(&orgName, "org", "o", "", "List secrets...")
	return cmd
}

func listRun(opts *ListOpts) error {

	fmt.Println(orgName)

	url := fmt.Sprintf("%s%s?per_page=%d", "http://127.0.0.1/", "users/keys", 50)
	fmt.Println(url)
	return nil
}
