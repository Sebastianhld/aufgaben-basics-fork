package listprops

// ContainsOnly liefert true, falls die Liste l
// ausschließlich den String x enthält.
func ContainsOnly(l []string, x string) bool {
	// Wenn die Liste leer ist kommt "enthält nur x", wenn leer ist true
	if len(l) == 0 {
		return true
	}
	for _, v := range l {
		if v != x {

			return false
			//wenn wir Element finden, was nicht x entspricht, dann false
		}
	}
	return true
	//wenn for if nicht eintritt, dann return true
}
