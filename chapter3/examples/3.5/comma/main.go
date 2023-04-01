package main

import "fmt"

func main() {
	// comma inserts commas in a non-negative decimal integer string
	fmt.Println(comma("12345")) // 12,345
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
