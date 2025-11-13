package strings

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			count++
		}
	}
	return count

}

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {

	result := 0
	for _, char := range s {
		if char == c {
			result++
		}
	}
	return result

}
