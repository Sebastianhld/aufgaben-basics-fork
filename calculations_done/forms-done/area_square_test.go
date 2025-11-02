package forms

import "fmt"

func ExampleAreaSquare() {
	fmt.Println(AreaSquare(0))
	fmt.Println(AreaSquare(1))
	fmt.Println(AreaSquare(2))
	fmt.Println(AreaSquare(3))
	fmt.Println(AreaSquare(4))
	fmt.Println(AreaSquare(5))

	// Output:
	// 0
	// 1
	// 4
	// 9
	// 16
	// 25
}
