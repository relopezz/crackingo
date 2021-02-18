package first

// UniqueCharacters :
// Implement an algorithm to determine if a string has
// all unique characters. What about if you cannot
// use additional data structure */
func UniqueCharacters(a string) bool {
	b := map[byte]int{}
	for i := 0; i < len(a); i++ {
		if _, ok := b[a[i]]; ok {
			return false
		}
		b[a[i]] = i
	}
	return true
}

// UniqueCharactersASCII
func UniqueCharactersASCII(a string) bool {
	if len(a) > 128 {
		return false
	}
	var b [128]bool
	for i := 0; i < len(a); i++ {
		if ok := b[a[i]]; ok {
			return false
		}
		b[a[i]] = true
	}
	return true
}

// UniqueCharactersUsingVector
// Assuming it is lower case ASCII
func UniqueCharactersUsingVector(a string) bool {
	flag := 0
	for i := 0; i < len(a); i++ {
		// Get distance from the first character (a = 0 => z=>25)
		if (flag & (1 << (a[i] - 'a'))) > 0 {
			return false
		}
		flag |= (1 << (a[i] - 'a'))
	}
	return true
}
