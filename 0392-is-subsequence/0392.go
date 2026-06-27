// Link: https://leetcode.com/problems/is-subsequence
package issubsequence

// Complexity: O(n), where n = len(t)
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for j < len(t) {
		if i < len(s) && s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
