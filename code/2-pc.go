package main

// import (
//       "fmt"
//       "time"
// )

var buffer int
var count = 0

func put(item int) {
	if count != 0 {
		panic("count != 0")
	}
	count = 1
	buffer = item
}

func get() int {
	if count != 1 {
		panic("count != 1")
	}
	count = 0
	return buffer
}

var mutex = sync.Mutex{}
var empty = sync.Cond

// TODO: Implement Producer
func producer() {
	for i in [0, loops]
}

// TODO: Implement Consumer
func consumer() {
	while (1):
		print()
}


func main() {
	fmt.Println("main: begin")
	ch_buffer := make (chan int, 2)
	go producer(ch_buffer, 10)
	go consumer(ch_buffer, )
}

// NOTES:
// Count helps us in the case that we miss a signal
// Whenever we use condition variables, we need to be careful. In this case, we are adding an additional variable
// to double check.
// Channels are very useful.
