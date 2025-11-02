package forms

import "fmt"

func ExampleHypotenuse() {
	fmt.Println(Hypotenuse(1, 1))
	fmt.Println(Hypotenuse(3, 4))

	// Output:
	// 1.4142135623730951
	// 5
}
