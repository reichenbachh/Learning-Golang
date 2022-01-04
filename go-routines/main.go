package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func printChannel(str string, out chan string) {
	defer close(out) // defer closing channel to end of execution
	for i := 0; i <= 10; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
		out <- str //pulling data out of channel
	}
}

func main() {
	//to create a go routine you just prefix the keyword "go" before a function
	//go printHello()

	//Using wait groups to handle concurrency issue with main function closing down
	// var wg sync.WaitGroup

	// wg.Add(10) //Number of processes function should wait for
	// for i := 0; i < 10; i++ {
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 		wg.Done() // decrements number of functions by 1
	// 	}(i)
	// }
	// wg.Wait() // tells thread to wait for go routine to end

	//using channels

	out := make(chan string) // used to create a channel from which data can be pulled from
	go printChannel("Channel", out)
}
