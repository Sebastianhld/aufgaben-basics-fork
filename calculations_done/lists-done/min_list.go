package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]           // starte mit dem ersten Element
	for _, n := range nums { // gehe jedes Element durch
		if n < min { // wenn ein kleineres gefunden wird
			min = n // merke es als neues Minimum
		}
	}
	return min // am Ende: kleinstes Element
}
