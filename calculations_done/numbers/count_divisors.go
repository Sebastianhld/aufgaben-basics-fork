package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	count := 0

	// Prüfe alle Zahlen von 1 bis n
	for i := 1; i <= n; i++ {
		// Wenn n durch i teilbar ist (Rest 0), dann ist i ein Teiler
		if n%i == 0 {
			count++
		}
	}

	return count
}
