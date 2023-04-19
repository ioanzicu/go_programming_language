// The LimitReader function in the io package accepts an io.Reader r and a
// number of bytes n, and returns another Reader that reads from r but
// reports an end-of-file condition after n bytes. Implement it.
//
// func LimitReader(r io.Reader, n int64) io.Reader

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type ReaderLimit struct {
	reader io.Reader
	limit  int64
}

func (r *ReaderLimit) Read(p []byte) (n int, err error) {
	if r.limit <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > r.limit {
		p = p[0:r.limit]
	}

	n, err = r.reader.Read(p)
	r.limit -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &ReaderLimit{
		reader: r,
		limit:  n,
	}
}

func main() {
	s := "hi Ioan"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 3)
	n, _ := b.ReadFrom(r)
	if n != 3 {
		fmt.Println("Expected n = 3, got", n)
	}
	if b.String() != "hi " {
		fmt.Printf("Expected \"hi \", got %+#v", b.String())
	}
}
