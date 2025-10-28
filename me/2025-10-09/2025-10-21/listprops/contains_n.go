package listprops

// ContainsN liefert true, falls die Liste l
// den String x mindestens n mal enthält.
func ContainsN(l []string, x string, n int) bool {
	count := 0
	for _, element := range l {
		if element == x {
			count++
		}
		if count >= n {
			return true
		}
	}
	return false
}
