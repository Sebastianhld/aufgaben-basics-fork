package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {
	// TODO
	if len(nums) == 0 {
		return 0
		//Liste leer=0
	}
	if len(nums) == 1 {
		return nums[0]
		//Liste hat nur ein Element, gib Position 0
	}
	minRest := MinListRecursive(nums[1:])
	//Liste 1 Ende

	if nums[0] < minRest {
		//wenn Position 0 kleiner als 1-Rest ist, gib Position 0
		return nums[0]
	}
	return minRest
	//ansonsten nimm die kleinere Position
}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
