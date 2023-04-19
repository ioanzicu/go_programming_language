package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

/*
$ go build gopl.io/ch7/sleep
$ ./sleep

OR

go run main.go
Sleeping for 1s...

go run main.go -period 1.5h
Sleeping for 1h30m0s...

go run main.go -period "1 day"
invalid value "1 day" for flag -period: parse error
*/

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
