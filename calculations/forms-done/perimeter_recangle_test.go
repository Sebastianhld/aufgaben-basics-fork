package forms

import "fmt"

func ExamplePerimeterRectangle() {
	fmt.Println(PerimeterRectangle(3, 4))
	fmt.Println(PerimeterRectangle(2.5, 10))
	fmt.Println(PerimeterRectangle(1.5, 25))

	// Output:
	// 14
	// 25
	// 53
}
