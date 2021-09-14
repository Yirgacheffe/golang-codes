package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var fooCmd = newFooCommand()

func init() {
	rootCmd.AddCommand(fooCmd)
	fooCmd.Flags().BoolP("float", "f", false, "Include Float Number")
}

func newFooCommand() *cobra.Command {
	return &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {

			f, _ := cmd.Flags().GetBool("float")
			if !f {
				intFoo(args)
			} else {
				floatFoo(args)
			}
		},
		Use:   "foo",
		Short: "Command foo",
		Long:  "This is a new Command Foo",
		// Args: cobra.MinimumNArgs(2),
	}
}

func intFoo(args []string) {
	var sum int
	for _, v := range args {
		temp, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of numbers %s is %d\n", args, sum)
}

func floatFoo(args []string) {
	var sum float64
	for _, v := range args {
		temp, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of numbers %s is %f\n", args, sum)
}
