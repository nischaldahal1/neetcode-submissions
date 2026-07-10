func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _,s := range strs {
		hashVal := getHash(s)
		anagramMap[hashVal] = append(anagramMap[hashVal],s)
	}
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	} 

	return result
}

func getHash(str string) [26]int {
	str = strings.ToLower(str)
	var arr [26]int // properly initializes all 26 slots to 0
	for _, ch := range str {
		// Ensure we only count valid lowercase letters to prevent panics/out-of-bounds
		if ch >= 'a' && ch <= 'z' {
			arr[ch-'a']++
		}
	}
	return arr
}
