package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus, ersetzt aber Vielfache von x durch "fizz",
// jede Vielfache von y durch "buzz" und
// Vielfache von x und y int) {
func FizzImproved(x, y, n int) {
	for i := 1; i <= 30; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Println("fizzbuzz")
		} else if i%x == 0 {
			fmt.Println("fizz")
		} else if i%y == 0 {
			fmt.Println("buzz")

		} else {
			fmt.Println(i)
		}
	}
}
