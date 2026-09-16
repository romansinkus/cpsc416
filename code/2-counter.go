package main

import (
      "fmt"
      "time"
)

var counter = 0

func increment(arg string) {

	fmt.Println(arg, ": begin")

	for i := 0; i < 1e7; i++ {
		counter = counter + 1
	}

	fmt.Println(arg, ": end")
}

func main() {
	fmt.Printf("begin (counter=%d)\n", counter)
	go increment("A")
	go increment("B")
	time.Sleep(10 * time.Second)
	fmt.Printf("end (counter=%d)\n", counter)
}