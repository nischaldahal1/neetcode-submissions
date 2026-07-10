func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
	for index, num := range(nums) {
		if _, ok := seen[num]; ok {
			return []int{seen[num],index}
		}
		seen[target-num] = index
	}
	return []int{-1,-1}
}
