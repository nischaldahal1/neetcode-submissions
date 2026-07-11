
func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for(left < right) {
		for left < right && !isAlphabet(s[left]) {
			left++
		}
		for left < right && !isAlphabet(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		right--
		left++
	}
	return true
}

func isAlphabet(c byte) bool {
	if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')) {
			return false
		}
	return true
}
