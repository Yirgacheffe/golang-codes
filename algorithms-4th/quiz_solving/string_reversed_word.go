package main

import (
	"fmt"
	"strings"
)

func trimLeft(s string) string {
	i := 0
	for i < len(s) {
		if s[i] != ' ' {
			break
		}
		i++
	}
	return s[i:]
}

func trimRight(s string) string {
	i := len(s) - 1
	for i > 0 {
		if s[i] != ' ' {
			break
		}
		i--
	}
	return s[0 : i+1]
}

func reversedWord(s string) {
	var words []string
	var word []byte

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word = append(word, s[i])
			if i == len(s)-1 {
				words = append(words, string(word))
				word = word[:0]
			}
		} else if (len(word) != 0) && (s[i] == ' ') {
			words = append(words, string(word))
			word = word[:0]
		}
	}

	var sb strings.Builder

	for i := len(words); i > 0; i-- {
		sb.WriteString(words[i-1])
		if i != 1 {
			sb.WriteString(" ")
		}
	}

	fmt.Println(sb.String())
}

func main() {
	s := "  coffee is   not ok for me "

	a := strings.Trim(s, " ")
	d := strings.TrimSpace(s)

	b := trimLeft(s)
	c := trimRight(b)

	if a == b {
		fmt.Println("Not correct")
	}

	if a != c {
		fmt.Println("Not correct")
	}

	fmt.Println(d)
	reversedWord(c)
}
