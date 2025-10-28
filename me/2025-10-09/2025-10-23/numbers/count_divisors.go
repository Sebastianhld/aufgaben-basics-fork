package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	if n < 1 {
		return 0
	}

	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
