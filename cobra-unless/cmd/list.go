package main

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
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if runF != nil {
				return runF(opts)
			} else {
				return listRun(opts)
			}
		},
	}

	cmd.Flags().StringVarP(&orgName, "org", "o", "", "List secrets...")
	return cmd
}

func listRun(opts *ListOpts) error {
	url := fmt.Sprintf("%s%s?per_page=%d", "http://127.0.0.1/", "users/keys", 50)
	fmt.Println(url)

	return nil
}
