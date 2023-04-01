// Write an in-place function that squashes each run of adjacent Unicode spaces
// (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	// https://pkg.go.dev/unicode#example-IsSpace
	// '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
	// fmt.Printf("%t\n", unicode.IsSpace(' '))
	// fmt.Printf("%t\n", unicode.IsSpace('\n'))
	// fmt.Printf("%t\n", unicode.IsSpace('\t'))
	// r := bytes.Runes([]byte("ssss"))
	// fmt.Printf("%v\n", r)

	t := "Hello\t\n\v\fWorld!"
	fmt.Printf("%+#v\n", t)
	res := squashAdjUnicodeSpaceToAschii([]byte(t))
	fmt.Printf("%s\n", res)

	fmt.Println("\n------------\n")

	t2 := "Hello\t\n\v\fWorld!"
	fmt.Printf("%+#v\n", t2)
	res2 := squashAdjUnicodeSpaceToAschii2([]byte(t2))
	fmt.Printf("%s\n", res2)
}

func squashAdjUnicodeSpaceToAschii(b []byte) []byte {
	k := 0
	runes := bytes.Runes(b)

	for i := 0; i < len(runes); {
		j := i + 1
		// move to the index that is not adjiacent duplicate
		for ; j < len(runes) && unicode.IsSpace(runes[j]); j++ {
		}

		runes[k] = runes[i]
		k++
		if i != j-1 {
			runes[k] = []rune(" ")[0]
			k++
		}
		i = j // update i to the index where are non duplicated item starts
	}

	return runesToUTF8(runes[:k])
}

func squashAdjUnicodeSpaceToAschii2(b []byte) []byte {
	runes := bytes.Runes(b)
	out := runes[:0]
	for i := 0; i < len(runes); {
		j := i + 1
		// move to the index that is not adjiacent duplicate
		for ; j < len(runes) && unicode.IsSpace(runes[j]); j++ {
		}

		// append non duplicated values
		out = append(out, runes[i])
		if i != j-1 {
			out = append(out, []rune(" ")[0])
		}
		i = j // update i to the index where are non duplicated item starts
	}
	return runesToUTF8(out)
}

func runesToUTF8(rs []rune) []byte {
	return []byte(string(rs))
}
