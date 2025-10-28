package numbers

// Berechnet das Minimum von zwei Zahlen.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Berechnet das Maximum von zwei Zahlen.
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//wir brauchen keine extra Bedingung für a = b, weils ja dann automatisch zu return b springt.
