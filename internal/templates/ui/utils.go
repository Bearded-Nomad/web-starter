package ui

import "strings"

func TwMerge(classes ...string) string {
	var filtered []string
	for _, c := range classes {
		if c != "" {
			filtered = append(filtered, c)
		}
	}
	return strings.Join(filtered, " ")
}
