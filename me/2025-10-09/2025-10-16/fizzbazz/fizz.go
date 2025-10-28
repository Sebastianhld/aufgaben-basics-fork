package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus, ersetzt aber Vielfache von 3 durch "fizz",
// jede Vielfache von 5 durch "buzz" und
// Vielfache von 3 und 5 durch "fizzbuzz".
func Fizz() {

	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Printf("fizz")
		} else if i%5 == 0 {
			fmt.Printf("buzz")

		} else {

		}
		fmt.Println(i)

	}
}
