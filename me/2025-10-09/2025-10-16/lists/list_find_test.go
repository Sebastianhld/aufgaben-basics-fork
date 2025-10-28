package listsearch

import "fmt"

func Find(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	// TODO
	return -1
}
func ExampleFind() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23}

	pos1 := Find(l1, 25)
	pos2 := Find(l1, 200)

	fmt.Println(pos1)
	fmt.Println(pos2)
	// Output:
	// 2
	// -1
}
