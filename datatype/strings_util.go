package main

import (
	"fmt"
	"regexp"
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

	s := "Hello world!"
	fmt.Println(s[2:4])
	fmt.Println(string(s[1]))

	sr := strings.Split(s, " ")
	fmt.Println(strings.Contains(s, sr[1]))

	re := strings.NewReplacer("C", "C++", "float", "int")
	ss := "a float in C"
	res := re.Replace(ss)
	fmt.Println(res)

	// ....
	sss := "xyzxdyddfyzdydjydyzdd"
	t := regexp.MustCompile(`[y]`)
	v := t.Split(sss, -1)
	fmt.Println(v) // regexp to split a string
}
