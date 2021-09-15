package main

import (
	"fmt"
	"regexp"
	"strings"
)

func changelogURL(version string) string {
	path := "https://somename.com/cli"
	r := regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	url := fmt.Sprintf("%s/releases/tag/v%s", path, strings.TrimPrefix(version, "v"))
	return url
}

func main() {
	tag := "0.3.1"
	fmt.Println(changelogURL(tag))

	tag = "v0.3.1"
	fmt.Println(changelogURL(tag))

	tag = "v0.3.1-pre.1"
	fmt.Println(changelogURL(tag))

	tag = "badbeef"
	fmt.Println(changelogURL(tag))
}
