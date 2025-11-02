package forms

import "fmt"

func ExamplePerimeterSquare() {
	fmt.Println(PerimeterSquare(0))
	fmt.Println(PerimeterSquare(1))
	fmt.Println(PerimeterSquare(2))
	fmt.Println(PerimeterSquare(3))
	fmt.Println(PerimeterSquare(4))
	fmt.Println(PerimeterSquare(5))

	// Output:
	// 0
	// 4
	// 8
	// 12
	// 16
	// 20
}
