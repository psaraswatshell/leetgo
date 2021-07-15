package leetcode

func longestpalindromesubstring(s string) string {
	var (
		maxLen int
		maxStr string
	)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					maxStr = s[i : j+1]
				}
			}
		}
	}
	return maxStr
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func ManacherAlgorithm(s string) string {
	if len(s) == 0 {
		return ""
	}
	newString := make([]rune, 0)
	newString = append(newString, '#')
	for _, v := range s {
		newString = append(newString, v)
		newString = append(newString, '#')
	}

	dp := make([]int, len(newString))
	center, right, maxLen, start := 0, 0, 0, 0
	for i := 0; i < len(newString); i++ {
		mirror := 2*center - i
		if right > i {
			dp[i] = min(dp[mirror], right-i)
		}
		r1 := i + dp[i] + 1
		l1 := i -(dp[i]+1)
		for r1 < len(newString) && l1 >= 0 && newString[r1] == newString[l1] {
			r1++
			l1--
			dp[i]++
		}
		if dp[i]+ i > right {
			center = i
			right = dp[i] + i
			if dp[i] > maxLen {
				maxLen = dp[i]
				start = (i-maxLen)/2
			}
		}
	}
	return s[start : start+maxLen]	

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
