package main

import "strings"

func countBeef(text string) map[string]int {
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ToLower(text)

	beef := make(map[string]int)
	for _, name := range strings.Fields(text) {
		beef[name] += 1
	}
	return beef
}
