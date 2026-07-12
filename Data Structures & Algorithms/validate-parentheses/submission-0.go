func isValid(s string) bool {
    brac_map := map[rune]rune {
	'[' : ']',
	'{' : '}',
	'(' : ')',
	}
	tracer := []rune{}

	for _, ch := range s {
		if val,ok := brac_map[ch]; ok {
			tracer = append(tracer,val)
		} else {
			if len(tracer) == 0 || tracer[len(tracer) - 1] != ch {
				return false
			}
			tracer = tracer[:len(tracer) - 1]
		}
	}
	return len(tracer) == 0
}
