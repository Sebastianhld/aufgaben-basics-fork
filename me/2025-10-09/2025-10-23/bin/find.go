package searching

func Find(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}

//
