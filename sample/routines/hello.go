package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	defer func() { fmt.Println("done") }()
	f("direct")
	go func(msg string) { fmt.Println(msg) }("going")
	for i := 0; i < 50; i++ {
		go f("goroutine")
	}
}
