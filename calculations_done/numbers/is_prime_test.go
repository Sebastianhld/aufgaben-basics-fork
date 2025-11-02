package numbers

import "fmt"

func ExampleIsPrime() {
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))
	fmt.Println(IsPrime(7))

	// Output:
	// false
	// true
	// true
	// false
	// true
	// false
	// true
}
