func findTheLongestBalancedSubstring(s string) int {
	var count0, count1, answer int
	var lastChar string
	for i := 0; i < len(s); i++ {
		if s[i:i+1] == "0" {
			if s[i:i+1] != lastChar {
				answer = max(answer, min(count0, count1))
				count0 = 0
				count1 = 0
			}
			count0++
		} else {
			if s[i:i+1] != lastChar {
				count1 = 0
			}
			count1++
		}
		lastChar = s[i : i+1]
	}

	answer = max(answer, min(count0, count1))
	return answer * 2
}