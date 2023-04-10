package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide url")
	}
	url := os.Args[1]

	log.SetPrefix("wait: ")
	log.SetFlags(0)

	if err := WaitForServer(url); err != nil {
		// fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		// os.Exit(1)

		log.Fatalf("Site is down: %v\n", err)
	}
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("sever %s failed to respond after %s", url, timeout)
}
