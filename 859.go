func buddyStrings(s string, goal string) bool {
	if len(s) <= 1 || len(s) != len(goal) {
		return false
	}
	diff := 0
	diffChars := []byte{}
	charCount := make(map[byte]int)
	charMuti := false
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff++
			if diff <= 2 {
				diffChars = append(diffChars, s[i], goal[i])
			}
		}
		charCount[s[i]]++
		if charCount[s[i]] > 1 {
			charMuti = true
		}
	}
	return (diff == 0 && charMuti) || (diff == 2 && diffChars[0] == diffChars[3] && diffChars[1] == diffChars[2])
}