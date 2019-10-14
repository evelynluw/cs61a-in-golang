package main

import "fmt"

func makeAdderInc(n int) func(int) int {
	inc := -1
	return func(d int) int {
		inc++
		return n + d + inc
	}
}

// try to make a test case

func main() {
	fmt.Println("test")
}
