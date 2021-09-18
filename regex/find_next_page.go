package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

var linkRE = regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"`)

func findNextPage(resp *http.Response) (string, bool) {
	for _, m := range linkRE.FindAllStringSubmatch(resp.Header.Get("Link"), -1) {
		if len(m) >= 2 && m[2] == "next" {
			return m[1], true
		}
	}
	return "", false
}

type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func main() {

	resp := &http.Response{
		Header: http.Header{
			"Link": []string{`<https://api.github.com/issues?page=3>; rel="last"`},
		},
	}

	fmt.Println(findNextPage(resp))

	r := `
		{
			"data": {
				"products": {
					"pageInfo": {
						"hasNextPage": false,
						"endCursor": "THE_END"
					}
				}
			}
		}
	`

	dec := json.NewDecoder(strings.NewReader(r))
	var stack []json.Delim

	// loop:
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}

		fmt.Println(t)

		switch tt := t.(type) {
		case json.Delim:
			{
				switch tt {
				case '{', '[':
					stack = append(stack, tt)
				case ']', '}':
					stack = stack[:len(stack)-1]
				}
			}
		default:
			// isKey := len(stack) > 0 && stack[len(stack)-1] == '{' && idx%2 == 0
			// idx++
		}

	}

}
