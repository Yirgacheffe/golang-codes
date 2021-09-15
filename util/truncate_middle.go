package main

import (
	"fmt"
)

func truncateMiddle(maxWidth int, t string) string {
	if len(t) <= maxWidth {
		return t
	}

	ellipsis := "..."
	if maxWidth < len(ellipsis)+2 {
		return t[0:maxWidth]
	}

	halfWidth := (maxWidth - len(ellipsis)) / 2
	return t[0:halfWidth] + ellipsis + t[len(t)-halfWidth:]
}

func main() {
	fmt.Println(truncateMiddle(10, "HDer343sdfjdk#$dfaser##4"))
}
