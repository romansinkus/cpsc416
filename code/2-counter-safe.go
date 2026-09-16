package main

import (
      "fmt"
      "time"
	  "sync"
)

type safe_counter_t struct {
	mu sync.Mutex
	value int
}


var safe_counter = safe_counter_t{}
func safe_increment(arg string) {
	fmt.Println(arg, ": begin")
	for i := 0; i < 1e7; i++ {
		safe_counter.mu.Lock()
		safe_counter.value++
		safe_counter.mu.Unlock()
	}
	fmt.Println(arg, ": end")
}

func main() {
	fmt.Printf("begin (counter=%d)\n", safe_counter.value)
	go safe_increment("A")
	go safe_increment("B")
	time.Sleep(20 * time.Second)
	fmt.Printf("end (counter=%d)\n", safe_counter.value)
}