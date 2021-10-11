package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	exitFail = 1
)

func run(args []string, stdout io.Writer) error {
	if len(args) < 2 {
		return errors.New("no names")
	}

	for _, name := range args[1:] {
		fmt.Fprintf(stdout, "Hi %s\n", name)
	}
	return nil
}

func runFalgs(args []string, stdout io.Writer) error {

	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	var (
		verbose = flags.Bool("v", false, "verbose")
		format  = flags.String("f", "Hi %s", "greeting format")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	fmt.Println(stdout, verbose)
	fmt.Println(stdout, format)

	return nil
}

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}
