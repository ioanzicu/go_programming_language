/*
Define a generic archive file-reading function capable of reading
ZIP files (archive/zip) and POSIX tar files (archive/tar). Use a
registration mechanism similar to the one described above so that
support for each file format can be plugged in using blank imports.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	uncp "uncompressor"
	_ "uncompressor/tar"
	_ "uncompressor/zip"
)

func printArchive(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := uncp.Open(f)
	if err != nil {
		return fmt.Errorf("open archive reader: %s", err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return fmt.Errorf("printing: %s", err)
	}
	return nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: extract FILE ...")
	}
	exitCode := 0
	for _, filename := range os.Args[1:] {
		err := printArchive(filename)
		if err != nil {
			log.Print(err)
			exitCode = 2
		}
	}
	os.Exit(exitCode)
}
