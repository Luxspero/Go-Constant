package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	const pi = math.Pi
	fmt.Println(pi)
	fmt.Printf("%T \n", pi)

	const (
		a = 1
		b
		c = 3
		d
	)

	fmt.Println(a, b, c, d)

	const (
		zero = iota + 5
		one
		two
	)
	fmt.Println(zero, one, two)

	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)
	fmt.Println(kb, mb, gb)
	fmt.Println(yb/zb, "\n")

	fmt.Println(math.Pi)
	fmt.Println(time.February)
	fmt.Println(time.Second)
	fmt.Println(time.UTC)
	fmt.Println(big.MaxExp)
	fmt.Println(big.MinExp)
}
