// The du3 command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size() // in bytes
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

/*
go run du3/main.go -v $HOME /usr  /bin /etc
127693 files 8.7 GB
253989 files 18.2 GB
377313 files 18.8 GB
499757 files 20.7 GB
638624 files 29.5 GB
793079 files 33.6 GB
*/

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {

	// Cancel traversal when input is detected
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	// ...start background goroutine...

	// Print the results.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64

loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish
			for range fileSizes {
				// Do nothing
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was close
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
