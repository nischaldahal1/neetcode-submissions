func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    tracker := make(map[rune]int)
    for _, char := range s {
        tracker[char]++
    }
    for _, char := range t {
        if count, ok := tracker[char]; !ok || count == 0 {
            return false
        }
        tracker[char]--
    }
    return true
}
