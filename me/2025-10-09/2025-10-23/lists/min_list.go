package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0] //erstes Element als STartwert
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

//for
//suche kleinste zahl
//gib kleinste zahl aus
//else 0
