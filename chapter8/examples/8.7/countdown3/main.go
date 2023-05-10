package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	/*
		we start a goroutine that tries to read a single
		byte from the standard input and, if it succeeds,
		sends a value on a channel called abort.
	*/
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Launch the rocket!!!")
}
