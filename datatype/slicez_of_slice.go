package main

import (
	"fmt"
	"io"
	"strings"
)

type tableField struct {
	Text      string
	Translate func(string) string
}

type ttyTablePrinter struct {
	out      io.Writer
	maxWidth int
	rows     [][]tableField
}

var rows = make([][]tableField, 2)

func main() {
	fmt.Println(rows)
	rowI := len(rows) - 1

	field := tableField{Text: "row 1 or 2?"}
	rows[rowI] = append(rows[rowI], field)

	fmt.Println(rows)

	field.Text += strings.Repeat(" k", 3)
	fmt.Println(field.Text)
}
