// The du2 command computes the disk usage of the files in a directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size() // in bytes
		}
	}
}

// dirents returns the entries of directory dir
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

/*
go run du2/main.go -v $HOME /usr  /bin /etc
127693 files 8.7 GB
253989 files 18.2 GB
377313 files 18.8 GB
499757 files 20.7 GB
638624 files 29.5 GB
793079 files 33.6 GB
*/

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree.
	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
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
		case size, ok := <-fileSize:
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
