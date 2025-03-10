package main

import (
	"maps"
	"testing"
)

func TestCountBeef(t *testing.T) {
	text := "Corned beef t-bone frankfurter ut meatball mollit strip steak ball tip.  Filet mignon aliquip ut est, Corned t-bone ham mollit nulla ribeye."
	expected := map[string]int{
		"corned":      2,
		"beef":        1,
		"t-bone":      2,
		"frankfurter": 1,
		"meatball":    1,
		"mollit":      2,
		"strip":       1,
		"steak":       1,
		"ball":        1,
		"tip":         1,
		"filet":       1,
		"mignon":      1,
		"aliquip":     1,
		"ut":          2,
		"est":         1,
		"ham":         1,
		"nulla":       1,
		"ribeye":      1,
	}
	result := countBeef(text)
	if !maps.Equal(result, expected) {
		t.Errorf("text %s, expected %v, but got %v", text, expected, result)
	}
}
