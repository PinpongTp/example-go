package main

import (
	"fmt"
)

var c = make(chan string)

func stuff() {
	done := make(chan bool)
	println("enter")
	defer func() {
		println("exit")
	}()
	go func() {
		println("    wait start")
		<-done
		println("    wait end")
	}()
	defer func() {
		println("  notify start")
		done <- true
		println("  notify end")
	}()
}

func pong() {
	var str string
	for {
		fmt.Println("...")
		str = <-c
		fmt.Println(str)
		if str == "ping" {
			fmt.Println("pong")
		} else {
			fmt.Println("error")
		}
	}
}

func main() {
	stuff()
	fmt.Println(" -- new -- ")
	go pong()
	fmt.Println("start")
	c <- "ping"
	c <- "test"
	c <- "ping"
	c <- "ping"
	c <- "ping"
	c <- "ping"
	fmt.Println("end")
}
