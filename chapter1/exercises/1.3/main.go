// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the
// time package, and Section 11.4 shows how to write benchmark tests for systematic per-
// formance evaluation.)
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("Program name:", os.Args[0])
	start := time.Now()

	// += - will create a new variable at each iteratiosn and copy the old content, the old variable will be garbage collected
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	secs := time.Since(start)
	println("With +=", fmt.Sprintf("Time %.2s", secs))

	// strings.Join
	start = time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))

	secs = time.Since(start)
	println("String join", fmt.Sprintf("Time %.4s", secs))
}

/*
go run chapter1/excercises/1.3/main.go wawadw awd ada  j kjkjh kjhkjhkjhkj h uf yt ytcy fyt fyt f ytfyt fyt fyt dyt dtr d trd 6d 654 65re6 7f6 d86 d65 d75 cyc y jhkj k h  j h h kjh u fu fiut f fdy c hj fckyu fyu gfi i v

Program name: /tmp/go-build3029052093/b001/exe/main

wawadw awd ada j kjkjh kjhkjhkjhkj h uf yt ytcy fyt fyt f ytfyt fyt fyt dyt dtr d trd 6d 654 65re6 7f6 d86 d65 d75 cyc y jhkj k h j h h kjh u fu fiut f fdy c hj fckyu fyu gfi i v

With += Time 14

wawadw awd ada j kjkjh kjhkjhkjhkj h uf yt ytcy fyt fyt f ytfyt fyt fyt dyt dtr d trd 6d 654 65re6 7f6 d86 d65 d75 cyc y jhkj k h j h h kjh u fu fiut f fdy c hj fckyu fyu gfi i v

String join Time 3.32
*/
