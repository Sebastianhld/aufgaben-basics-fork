package numbers

import "fmt"

func ExampleMin() {
	fmt.Println(Min(1, 2))
	fmt.Println(Min(3, 17))
	fmt.Println(Min(5, -1))
	fmt.Println(Min(2, 2))

	// Output:
	// 1
	// 3
	// -1
	// 2
}

func ExampleMax() {
	fmt.Println(Max(1, 2))
	fmt.Println(Max(3, 17))
	fmt.Println(Max(5, -1))
	fmt.Println(Max(2, 2))

	// Output:
	// 2
	// 17
	// 5
	// 2
}
