// Write a version of du that computes and periodically displays
// separate totals for each of the root directories.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type RootSize struct {
	root int
	size int64
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

/*
go run main.go -v $HOME /usr  /bin /etc

136207 files 9.7 GB under /home/ioan
98633 files 3.9 GB under /usr
79 files 0.1 GB under /bin
1220 files 0.0 GB under /etc
0 files 0.0 GB under -v
540621 files 21.3 GB under /home/ioan
184306 files 8.1 GB under /usr
1123 files 0.4 GB under /bin
2664 files 0.0 GB under /etc
0 files 0.0 GB under -v
608253 files 24.4 GB under /home/ioan
209754 files 9.4 GB under /usr
1816 files 0.6 GB under /bin
2664 files 0.0 GB under /etc

*/

func main() {
	// Determine the initial directories.
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Cancel traversal when input is detected
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Traverse each root of the file tree in parallel.
	rootSizes := make(chan RootSize)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, i, rootSizes)
	}
	go func() {
		n.Wait()
		close(rootSizes)
	}()

	// Print the results periodically
	tick := time.Tick(500 * time.Millisecond)
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))

loop:
	for {
		select {
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish
			for range rootSizes {
				// Do nothing
			}
			return
		case rs, ok := <-rootSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[rs.root]++
			nbytes[rs.root] += rs.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes) // final totals
		}
	}
	printDiskUsage(roots, nfiles, nbytes) // final totals
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%d files %.1f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, root int, rootSizes chan<- RootSize) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, root, rootSizes)
		} else {
			rootSizes <- RootSize{root, entry.Size()} // in bytes
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20) // concurrency limiting counting semaphore

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token
	case <-done:
		return nil // cancelled
	}
	defer func() { <-sema }() // release token

	// read directoyr
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return Readdir may return partial results
	}
	return entries
}
