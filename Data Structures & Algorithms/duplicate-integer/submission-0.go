func hasDuplicate(nums []int) bool {
    tracker  := make(map[int]bool)
    if (len(nums) <= 1) {
        return false
    }
    for _,i := range(nums) {
        if _, ok := tracker[i]; ok {
            return true
        } else {
            tracker[i] = true
        }
    }
    return false
}