// Link: https://leetcode.com/problems/maximum-number-of-balloons
package maximumnumberofbaloons

// Complexity: O(n)
func maxNumberOfBalloons(text string) int {
	frequencies := map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0}
	for _, b := range text {
		if _, ok := frequencies[b]; ok {
			frequencies[b]++
		}
	}
	frequencies['l'] /= 2
	frequencies['o'] /= 2
	minCount := frequencies['b']
	for _, v := range frequencies {
		if v < minCount {
			minCount = v
		}
	}
	return minCount
}
