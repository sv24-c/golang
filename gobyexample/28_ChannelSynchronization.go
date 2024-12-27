package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)

	go worker(done)

	<-done
}

func worker(done chan bool) {
	fmt.Println("done calling")
	time.Sleep(time.Second)
	done <- true
}
