package main

import "fmt"

// use GoTest for run GoTest
// use GoCoverage for see Coverage
// use GoCoverageClear for clear Coverage mark

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Printf("fib(0)=%d\n", fib(0))
	fmt.Printf("fib(1)=%d\n", fib(1))
	fmt.Printf("fib(2)=%d\n", fib(2))
	fmt.Printf("fib(3)=%d\n", fib(3))
	fmt.Printf("fib(4)=%d\n", fib(4))
	fmt.Printf("fib(5)=%d\n", fib(5))
}
