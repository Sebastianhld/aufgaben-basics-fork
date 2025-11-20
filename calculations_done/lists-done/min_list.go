package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	min := nums[0]           //Position 0
	for _, n := range nums { //ganze Liste durchgehen
		if n < min {
			min = n
		}
	}
	return min
}

//func MinList (nums [] int) int
// nums []int → die Funktion bekommt eine Liste (Slice) von int-Werten,
// nums ist der Name der Variablen
//[]int = Type: Slice von ganzen Zahlen
