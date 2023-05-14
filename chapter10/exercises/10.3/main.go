/*
Using fetch http://gopl.io/ch1/helloworld?go-get=1, find out which
service hosts the code samples for this book. (HTTP requests from go get include the go-get
parameter so that servers can distinguish them from ordinary browser requests.)
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
 go run main.go http://gopl.io/ch1/helloworld?go-get=1 | grep go-import
<meta name="go-import" content="gopl.io git https://github.com/adonovan/gopl.io">
*/

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
