package main

import (
	"codespaces-test/async"
	"fmt"
	"time"
)

// Runs async and returns a channel
func asyncFunc() chan int {
	r := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		// writes to channel
		// once async function is done
		r <- 1
		fmt.Println("Done ...")
	}()

	return r
}

func main() {
	fmt.Println("Start ...")
	r := async.Exec(func() interface{} {
		return asyncFunc()
	})
	fmt.Println("Async task runnin ...")
	val := r.Await()
	fmt.Println(val)
}
