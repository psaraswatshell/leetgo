package leetcode

func lenghtOfLongestSubstring1(s string) int {
	left, right, result:= 0, 0, 0
	mapIndices := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok:= mapIndices[s[left]]; ok && idx>= right {
			right = idx + 1
		}
		mapIndices[s[left]] = left
		left++
		result = max(result, left-right)
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

