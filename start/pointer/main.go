package main

import "fmt"

func main() {
	var i, j int
	var p *int
	i = 1
	j = 3

	fmt.Printf("i = %d\n", i)
	fmt.Printf("address of i %v\n", &i)
	fmt.Printf("address of j %v\n", &j)
	p = &i
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)
	p = &j
	fmt.Printf("p = %v\n", p)
	fmt.Printf("*p = %v\n", *p)
}
