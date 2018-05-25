package interviews

// FindFirstNotRepeatingChar returns the first not repeating char in the string
func FindFirstNotRepeatingChar(s string) byte {
	if s == "" {
		return 0
	}
	cnt := [256]int{}
	strArr := []byte(s)
	for _, c := range strArr {
		cnt[c]++
	}

	for _, c := range strArr {
		if cnt[c] == 1 {
			return c
		}
	}
	return 0
}

// DeleteChar deletes the chars of s1 that are in s2, and returns result
func DeleteChar(s1, s2 string) string {
	if s2 == "" {
		return s1
	}
	if s1 == "" {
		return ""
	}

	table := [256]bool{}
	for _, c := range []byte(s2) {
		table[c] = true
	}

	result := []byte(s1)
	for i := 0; i < len(result); {
		if table[result[i]] {
			// delete char i at i
			result = append(result[:i], result[i+1:]...)
		} else {
			i++
		}
	}
	return string(result)
}

// DeleteRepeatedChar deletes the repeated char, and returns result
func DeleteRepeatedChar(s string) string {
	if s == "" {
		return ""
	}
	table := [256]bool{}
	result := []byte(s)
	for i := 0; i < len(result); {
		if !table[result[i]] {
			table[result[i]] = true
			i++
		} else {
			// delete char at i
			result = append(result[:i], result[i+1:]...)
		}
	}
	return string(result)
}
