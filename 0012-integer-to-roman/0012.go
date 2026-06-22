// Link: https://leetcode.com/problems/integer-to-roman/
package integertoroman

import "strings"

type RomanPair struct {
	Value  int
	Symbol string
}

// Complexity: O(1)
func intToRoman(num int) string {
	pairings := []RomanPair{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var result strings.Builder
	for _, pair := range pairings {
		for num >= pair.Value {
			result.WriteString(pair.Symbol)
			num -= pair.Value
		}
	}
	return result.String()
}
