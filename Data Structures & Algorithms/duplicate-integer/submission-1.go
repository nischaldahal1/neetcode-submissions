func hasDuplicate(nums []int) bool {
    seen  := make(map[int]struct{})
    if (len(nums) <= 1) {
        return false
    }
    for _,i := range(nums) {
        if _, ok := seen[i]; ok {
            return true
        } else {
            seen[i] = struct{}{}
        }
    }
    return false
}