package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	fmt.Println(unsafe.Sizeof(float64(0))) // 8

	/*
		Typical 32-bit platform:
		Sizeof(x)   = 16    Alignof(x) = 4
		Sizeof(x.a) = 1     Alignof(x.a) = 1    Offsetof(x.a) = 0
		Sizeof(x.b) = 2     Alignof(x.b) = 2    Offsetof(x.b) = 2
		Sizeof(x.c) = 12    Alignof(x.c) = 4    Offsetof(x.c) = 4
	*/

	// 64-bit platform
	fmt.Printf("Sizeof(x)   = %d \tAlignof(x)    = %d\n", unsafe.Sizeof(x), unsafe.Alignof(x))                                                // Sizeof(x)   = 32        Alignof(x)    = 8
	fmt.Printf("Sizeof(x.a) = %d \tAlignof(x.a)  = %d \tOffsetof(x.a) = %d\n", unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a)) // Sizeof(x.a) = 1         Alignof(x.a)  = 1       Offsetof(x.a) = 0
	fmt.Printf("Sizeof(x.b) = %d \tAlignof(x.b)  = %d \tOffsetof(x.b) = %d\n", unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b)) // Sizeof(x.b) = 2         Alignof(x.b)  = 2       Offsetof(x.b) = 2
	fmt.Printf("Sizeof(x.c) = %d \tAlignof(x.c)  = %d \tOffsetof(x.c) = %d\n", unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c)) // Sizeof(x.c) = 24        Alignof(x.c)  = 8       Offsetof(x.c) = 8
}
