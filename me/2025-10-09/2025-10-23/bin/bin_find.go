package searching

func BinFind1(l []int, x int) int {
	links := 0

	for len(l) > 0 {
		mitte := len(l) / 2

		//vergleiche x mit dem Element in der Mitte
		//Wenn x == l [mitte], dann fertig
		if l[mitte] == x {
			return mitte + links
		}

		//Wenn x < l [mitte] dann suche im linken Teil weiter
		if x < l[mitte] {
			//lasse nur den linken Teil der Liste übrig
			// alles von Null bis exclusive mitte
			l = l[0:mitte]
		} else {
			//Wenn x > l [mitte] dann suche im rechten Teil weiter
			//lasse nur den rechten Teil der Liste übrig
			l = l[mitte+1:]
			links += mitte + 1
		}
	}

	return -1
}
