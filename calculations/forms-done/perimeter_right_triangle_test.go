package forms

import "fmt"

func ExamplePerimeterRightTriangle() {
	fmt.Println(PerimeterRightTriangle(3, 4))
	fmt.Println(PerimeterRightTriangle(2.5, 10))
	fmt.Println(PerimeterRightTriangle(1.5, 25))

	// Output:
	// 12
	// 22.80776406404415
	// 51.54495957273639
}
