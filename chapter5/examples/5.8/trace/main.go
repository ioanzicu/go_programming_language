package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bigSlowOperation()

	// Observe function results
	_ = double2(4)
	fmt.Println(triple(4)) // "12"
}

// Deferred functions run after return statements
// have updated the function’s result variables.

func bigSlowOperation() {
	defer trace("bigSlowOperation")() //don't forget the extra parentheses
	// ...lots of work is going on here...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) int {
	return x + x
}

// By naming its result variable and adding a defer statement,
// we can make the function print its arguments
// and results each time it is called.
func double2(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}

func triple(x int) (result int) {
	defer func() {
		result += x
	}()
	return double2(x)
}
