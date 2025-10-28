package searching

import "fmt"

func ExampleBinFind1() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}
	//fmt.Println(BinFind1(l1, 3))
	fmt.Println(BinFind1(l1, 8))
	fmt.Println(BinFind1(l1, 277))
	fmt.Println(BinFind1(l1, 3))

}

// Output
// 2
// 7
//-1
