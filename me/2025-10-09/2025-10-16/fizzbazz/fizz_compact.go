package fizzbuzz

import "fmt"

// Fizz gibt die Zahlen von 1 bis 30 aus, ersetzt aber Vielfache von x durch "fizz",
// jede Vielfache von y durch "buzz" und
// Vielfache von x und y int) {
func FizzCompact() {
	for i := 1; i <= 30; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}
		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()
	}
}
