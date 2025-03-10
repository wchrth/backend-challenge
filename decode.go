package main

import "strconv"

func decode(encoded string) string {
	result := make([]int, len(encoded)+1)

	for i, ch := range encoded {
		switch ch {
		case 'L':
			for j := i; j >= 0; j-- {
				if encoded[j] == 'L' && result[j] <= result[j+1] {
					result[j] = result[j+1] + 1
				}
				if encoded[j] == '=' {
					result[j] = result[j+1]
				}
				if encoded[j] == 'R' {
					break
				}
			}
		case 'R':
			result[i+1] = result[i] + 1
		case '=':
			result[i+1] = result[i]
		}
	}
	resultStr := ""
	for _, n := range result {
		resultStr += strconv.Itoa(n)
	}

	return resultStr
}
