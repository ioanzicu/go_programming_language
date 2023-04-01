// Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// operates in streaming mode - in which input is read and broken into lines as needed,
	// so in principle these programs can handle an arbitrary amount of input
	counts := make(map[string]int)
	filenames := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, files[0], filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, filenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Filename: %s - %d\t%s\n", filenames[line], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, filenames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		filenames[input.Text()] = filename
	}
}
