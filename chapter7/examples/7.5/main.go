package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil>"
	// w.Write([]byte("salut")) // panic: nil pointer dereference

	w = os.Stdout
	fmt.Printf("%T\n", w) // "*os.File"
	// (*os.File).Write([]byte("salut"))
	// like
	// os.Stdout.Write([]byte("hello"))
	w.Write([]byte("salut")) // "salut"

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)    // "*bytes.Buffer"
	w.Write([]byte("salut")) // writes "hello" to the bytes.Buffer
	fmt.Println(w.(*bytes.Buffer).String())
	w = nil
}
