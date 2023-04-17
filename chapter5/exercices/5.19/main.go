// Use panic and recover to write a function that
// contains no return statement yet returns a non-zero value.
package main

import "fmt"

func main() {
	someFunc()
	fmt.Println()
	fmt.Println("Returned value by panic:", someFunc2())
}

func someFunc() {
	fmt.Println("Start someFunc")

	defer func() (x int) {
		defer func() {
			fmt.Println("Returned value: ", x)
		}()

		if p := recover(); p != nil {
			fmt.Println("OMG, I hava panic attack!!!")
			// "expected" panic
			return 1
		}
		return x
	}()

	panic("Start to panic...")
}

func someFunc2() (r int) {
	fmt.Println("Start someFunc2")

	defer func() {
		if p := recover(); p != nil {
			fmt.Println("OMG, I hava panic attack, again!!!")
			// "expected" panic
			r = 1
		}
	}()

	panic("Start to panic, againg...")
}
