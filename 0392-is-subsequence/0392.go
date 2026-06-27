package issubsequence

// Example 1:
// Input: s = "abc", t = "ahbgdc"
// Output: true

// Example 2:
// Input: s = "axc", t = "ahbgdc"
// Output: false

// s -> nil
// t -> nil

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	result := false
	for i < len(s) {
		result = false
		for j < len(t) {
			if s[i] == t[j] {
				i++
				j++
				result = true
				break
			}
			j++
		}
		if result == false {
			return false
		}
	}
	return true
}
