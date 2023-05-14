/*
Write a program with two goroutines that send messages back
and forth over two unbuffered channels in ping-pong fashion.
How many communications per second can the program sustain?
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func pingpong1() {
	in := make(chan string)
	out := make(chan string)

	go func(in chan string, out chan string) {
		for {
			in <- "ping"
			time.Sleep(1 * time.Second)
			fmt.Println(<-out)
		}
	}(in, out)

	go func(in chan string, out chan string) {
		for {
			fmt.Println(<-out)
			in <- "pong"
			time.Sleep(1 * time.Second)
		}
	}(out, in)

	time.Sleep(5 * time.Second)
}

func pingpong2() {
	q := make(chan int)
	var i int64
	start := time.Now()
	go func() {
		q <- 1
		for {
			i++
			q <- <-q
		}
	}()
	go func() {
		for {
			q <- <-q
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(float64(i)/float64(time.Since(start))*1e9, "rounds per second")
}

func main() {
	// pingpong1()
	pingpong2()
}
