package main

import "testing"

func TestDecode(t *testing.T) {
	testCases := []struct {
		encoded  string
		expected string
	}{
		{
			encoded:  "LLRR=",
			expected: "210122",
		},
		{
			encoded:  "==RLL",
			expected: "000210",
		},
		{
			encoded:  "=LLRR",
			expected: "221012",
		},
		{
			encoded:  "RRL=R",
			expected: "012001",
		},
		{
			encoded:  "LLLLL",
			expected: "543210",
		},
		{
			encoded:  "LLLRR",
			expected: "321012",
		},
		{
			encoded:  "LL=RR",
			expected: "210012",
		},
		{
			encoded:  "RR=LL",
			expected: "012210",
		},
		{
			encoded:  "LL=LL",
			expected: "432210",
		},
		{
			encoded:  "RL=LL",
			expected: "032210",
		},
		{
			encoded:  "RL=RL",
			expected: "010010",
		},
		{
			encoded:  "LR=LR",
			expected: "101101",
		},
		{
			encoded:  "LR==R",
			expected: "101112",
		},
		{
			encoded:  "RRRRR",
			expected: "012345",
		},
		{
			encoded:  "RRRLL",
			expected: "012310",
		},
		{
			encoded:  "=====",
			expected: "000000",
		},
		{
			encoded:  "=LRL=",
			expected: "110100",
		},
		{
			encoded:  "=LLL=",
			expected: "332100",
		},
		{
			encoded:  "=RRR=",
			expected: "001233",
		},
	}

	for _, tc := range testCases {
		t.Run("decode", func(t *testing.T) {
			result := decode(tc.encoded)
			if result != tc.expected {
				t.Errorf("encoded %s, expected %s, but got %s", tc.encoded, tc.expected, result)
			}
		})
	}
}
