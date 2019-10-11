package main

import (
	"fmt"
	"math"
)

// uh.. no generic Abs because golang is simplistic
func iAbs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func intersection(st int, ave int) int {
	return (st+ave)*(st+ave+1)/2 + ave
}

func street(inter int) int {
	return w(inter) - avenue(inter)
}

func avenue(inter int) int {
	// cast, cast, cast
	return inter - (int(math.Pow(float64(w(inter)), 2))+w(inter))/2
}

func w(z int) int {
	return int((math.Pow(8*float64(z)+1, 0.5) - 1) / 2)
}

func taxicab(a int, b int) int {
	return iAbs(street(a)-street(b)) + iAbs(avenue(a)-avenue(b))
}

func main() {
	timeSquare := intersection(46, 7)
	essABagel := intersection(51, 3)
	fmt.Println(taxicab(timeSquare, essABagel))
	fmt.Println(taxicab(essABagel, timeSquare))
}
