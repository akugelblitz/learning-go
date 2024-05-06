package main

import (
	"fmt"
	"time"
)

func greet(val string) {
	fmt.Println(val)
}

func slowGreet(val string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	greet(val)

	doneChan <- true
}

func main() {
	// go greet("hi")
	// go greet("hi 2")
	doneChan := make(chan bool)
	go slowGreet("hi 3", doneChan)
	// go greet("hi 4")
	<-doneChan
}
