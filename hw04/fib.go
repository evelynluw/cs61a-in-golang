package main

// use math/big for HUGE numbers
import (
	"fmt"
	"math/big"
	"time"
)

func fib() func() *big.Int {
	a, b := big.NewInt(-1), big.NewInt(1)
	return func() *big.Int {
		a.Add(a, b)
		a, b = b, a
		return b
	}
}

// fibonacci with timing
func main() {
	f := fib()
	start := time.Now()
	for i := 0; i < 1000; i++ {
		fmt.Println(f())
	}
	elapsed := time.Since(start)
	fmt.Printf("it took %s", elapsed)
}
