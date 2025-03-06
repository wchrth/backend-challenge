package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestMaxPath(t *testing.T) {
	jsonFile, err := os.ReadFile("files/hard.json")
	if err != nil {
		t.Fatalf("unable to read file: %v", err)
	}

	var hard [][]int
	err = json.Unmarshal([]byte(jsonFile), &hard)
	if err != nil {
		t.Fatalf("unable to unmarshal json data: %v", err)
	}

	testCases := []struct {
		triangle [][]int
		expected int
	}{
		{
			triangle: hard,
			expected: 7273,
		},
		{
			triangle: [][]int{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			expected: 237,
		},
		{
			triangle: [][]int{
				{1},
				{2, 3},
				{4, 5, 6},
			},
			expected: 10,
		},
		{
			triangle: [][]int{
				{5},
				{9, 6},
				{4, 6, 8},
				{0, 7, 1, 5},
			},
			expected: 27,
		},
		{
			triangle: [][]int{
				{1},
			},
			expected: 1,
		},
		{
			triangle: [][]int{},
			expected: 0,
		},
		{
			triangle: [][]int{{}},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run("maxPath", func(t *testing.T) {
			result := maxPath(tc.triangle)
			if result != tc.expected {
				t.Errorf("triangle %v, expected %d, but got %d", tc.triangle, tc.expected, result)
			}
		})
	}
}
