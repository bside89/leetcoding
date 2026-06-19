// Link: https://leetcode.com/problems/valid-palindrome/
package validpalindrome

import (
	"strings"
	"unicode"
)

// Complexity: O(n)
func isPalindrome(s string) bool {
	filtered := filterString(s)
	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}
	return true
}

func filterString(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			if unicode.IsLower(r) {
				builder.WriteRune(r)
			} else {
				builder.WriteRune(r + 32)
			}
		} else if unicode.IsNumber(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
