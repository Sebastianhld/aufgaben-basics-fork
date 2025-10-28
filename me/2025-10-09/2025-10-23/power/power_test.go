package power

import "fmt"

func ExamplePow2() {
	fmt.Println(Pow2(0))
	fmt.Println(Pow2(1))
	fmt.Println(Pow2(2))
	fmt.Println(Pow2(8))
	fmt.Println(Pow2(10))

	// Output:
	// 1
	// 2
	// 4
	// 256
	// 1024
}
