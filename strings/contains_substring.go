package strings

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	if len(t) == 0 {
		return true
	}
	if len(t) > len(s) {
		return false
	}
	//wir haben die 2 Sonerfälle definiert
	for i := 0; i <= len(s)-len(t); i++ {
		//bestimmt, wo wir anfangen die Schleife ab s zu überprüfen, bei Hallo und all bei Position 2
		match := true
		for j := 0; j < len(t); j++ {
			if s[i+j] != t[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
