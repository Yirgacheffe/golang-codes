package main

import (
	"fmt"
	"os"
	"unless/cmd"
	// "github.com/MakeNowJust/heredoc/v2"
)

func main() {
	fmt.Println(os.Args[1:])
	cmd.Execute()
	/*
	   	fmt.Println(heredoc.Doc(`
	   	Secrets can be set at the repository or organization level for use in GitHub Actions.
	   	Run "gh help secret set" to learn how to get started.
	   `))
	*/
}
