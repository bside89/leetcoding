// Link: https://leetcode.com/problems/weighted-word-mapping
package weightedwordmapping

import "strings"

// Complexity: O(m * n), where m = len(words), n = len(words[i])
func mapWordWeights(words []string, weights []int) string {
	var result strings.Builder
	resultMap := map[int]rune{
		0:  'z',
		1:  'y',
		2:  'x',
		3:  'w',
		4:  'v',
		5:  'u',
		6:  't',
		7:  's',
		8:  'r',
		9:  'q',
		10: 'p',
		11: 'o',
		12: 'n',
		13: 'm',
		14: 'l',
		15: 'k',
		16: 'j',
		17: 'i',
		18: 'h',
		19: 'g',
		20: 'f',
		21: 'e',
		22: 'd',
		23: 'c',
		24: 'b',
		25: 'a',
	}
	for _, s := range words {
		sum := 0
		for _, r := range s {
			i := int(r - 'a')
			weight := weights[i]
			sum += weight
		}
		result.WriteRune(resultMap[sum%26])
	}

	return result.String()
}
