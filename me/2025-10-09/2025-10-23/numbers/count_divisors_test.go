package numbers

import "fmt"

func ExampleCountDivisors() {
	fmt.Println(CountDivisors(1))
	fmt.Println(CountDivisors(2))
	fmt.Println(CountDivisors(3))
	fmt.Println(CountDivisors(4))
	fmt.Println(CountDivisors(5))
	fmt.Println(CountDivisors(6))
	fmt.Println(CountDivisors(7))

	// Output:
	// 1
	// 2
	// 2
	// 3
	// 2
	// 4
	// 2
}
