package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(horst []int) int {
	if len(horst) == 0 {
		return 0
	}
	if len(horst) == 1 {
		return horst[0] //eckige Klammer ist Position in einer Liste

	}
	minRest := MinListRecursive(horst[1:]) //:= um eine neue Variable zum ersten Mal erwähnen
	//in der neuen variable sind jetzt alle Elemente von 1- Ende drin
	if minRest > horst[0] {
		return horst[0]
	}
	return minRest

}
