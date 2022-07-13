package main

// go routine
// when blocking event (sleep) well run next go routine
// this programe have 5 go routine

import (
	"fmt"
	"time"
)

func f0() {
	fmt.Println("f0")
}

func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2_1")
	time.Sleep(time.Second * 1)
	fmt.Println("f2_2")
}

func f3() {
	fmt.Println("f3")
}

func f4() {
	// go f2()
	fmt.Println("f4")
}

func f5() {
	fmt.Println("f5")
}

func main() {
	defer f0()
	f1()
	go f2()
	go f3()
	go f4()
	fmt.Println("main")
	time.Sleep(time.Second * 2)
	go f5()
	defer f0()
	go f5()
}
