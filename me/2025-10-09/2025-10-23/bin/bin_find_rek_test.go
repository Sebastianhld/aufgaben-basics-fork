package searching

import "fmt"

func ExampleBinFindRek() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}
	fmt.Println(BinFindRek(l1, 1))
	fmt.Println(BinFindRek(l1, 277))
	fmt.Println(BinFindRek(l1, 3))

	// Output:
	// 0
	// 7
	//-1
}
