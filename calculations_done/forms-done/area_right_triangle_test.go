package forms

import "fmt"

func ExampleAreaRightTriangle() {
	fmt.Println(AreaRightTriangle(3, 4))
	fmt.Println(AreaRightTriangle(2.5, 10))
	fmt.Println(AreaRightTriangle(1.5, 25))

	// Output:
	// 6
	// 12.5
	// 18.75
}
