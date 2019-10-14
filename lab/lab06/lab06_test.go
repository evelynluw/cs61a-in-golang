package main

import (
	"fmt"
)

func ExampleAdder() {
	adder1 := makeAdderInc(5)
	adder2 := makeAdderInc(6)
	fmt.Println(adder1(2))
	fmt.Println(adder1(2))
	fmt.Println(adder1(10))
	for i := 1; i <= 3; i++ {
		fmt.Print(adder1(i), ",")
	}
	fmt.Print("\n")
	fmt.Println(adder2(5))
	// Output:
	// 7
	// 8
	// 17
	// 9,11,13,
	// 11
}
