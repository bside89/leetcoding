// Link: https://leetcode.com/problems/roman-to-integer/
package romantointeger

// Complexity: O(n)
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	total := 0
	length := len(s)

	for i := range length {
		currentValue := romanMap[s[i]]
		if i+1 < length && currentValue < romanMap[s[i+1]] {
			total -= currentValue // Subtract (Ex: IV -> -1 + 5 = 4)
		} else {
			total += currentValue
		}
	}

	return total
}
