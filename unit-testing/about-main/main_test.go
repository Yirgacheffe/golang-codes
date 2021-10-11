package main

import (
	"bytes"
	"strings"

	"github.com/matryer/is"
	"testing"
)

func Test_WithNames(t *testing.T) {
	is := is.New(t)

	args := []string{"greeter", "David", "Kat", "Jon", "Natalie", "Mark"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.NoErr(err)

	out := stdout.String()
	is.True(strings.Contains(out, "Hi David"))
	is.True(strings.Contains(out, "Hi Kat"))
	is.True(strings.Contains(out, "Hi Jon"))
	is.True(strings.Contains(out, "Hi Natalie"))
	is.True(strings.Contains(out, "Hi Mark"))
}

func Test_NoNames(t *testing.T) {
	is := is.New(t)

	args := []string{"greeter"}
	var stdout bytes.Buffer

	err := run(args, &stdout)
	is.True(err != nil)

	// testing flags
	// err := runWithFlag([]string{"greeter", "-v", "-debug=true", "-another=2"})
}
