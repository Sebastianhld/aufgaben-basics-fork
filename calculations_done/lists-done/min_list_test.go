package lists

import "fmt"

func ExampleMinList() {
	fmt.Println(MinList([]int{1, 2, 3, 4, 5}))
	fmt.Println(MinList([]int{5, 4, 3, 2, 1}))
	fmt.Println(MinList([]int{1, 7, 2, -1, 5, -3}))
	fmt.Println(MinList([]int{}))

	// Output:
	// 1
	// 1
	// -3
	// 0
}
