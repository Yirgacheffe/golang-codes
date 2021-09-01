package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test8jsdfsf", "d"))
	fmt.Println(strings.Join([]string{"138", "8743", "3117"}, "-"))
	fmt.Println(strings.Repeat("8", 3))
	fmt.Println(strings.Replace("aaaaaaa", "a", "b", 5))
	fmt.Println(strings.Split("138-0983-3275", "-"))
	fmt.Println(strings.ToLower("EjsdfTJAJDFAENdkei"))
	fmt.Println(strings.EqualFold("go", "GO"))
}
