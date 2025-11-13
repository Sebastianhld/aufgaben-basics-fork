package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		count := 0

		for j := 0; j < len(s); j++ {
			if s[j] == char {
				count++
			}
		}
		for k := 0; k < 2; k++ {
			result += string(char)
		}
	}
	return result
}
