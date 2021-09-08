package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	kiloToMiles = 0.621371
	milesToKilo = 1.60934
)

func Convert(from, to string) (string, error) {
	var result float64
	switch {
	case strings.HasSuffix(from, "mi"):
		miles, err := strconv.ParseFloat(from[:len(from)-2], 64)
		if err != nil {
			return "", err
		}

		switch to {
		case "km":
			result = miles * milesToKilo
		case "m":
			result = miles * milesToKilo * 1000
		case "mi":
			result = miles
		}
	}
	return strconv.FormatFloat(result, 'f', -1, 64), nil
}

func main() {
	res, _ := Convert("50mi", "km")
	fmt.Println(res)
}
