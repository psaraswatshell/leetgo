package leetcode

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if _, ok := mapping[diff]; ok {
			return []int{i, mapping[diff]}
		}
		mapping[num] = i
	}
	return []int{}
}