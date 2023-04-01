// Write const declarations for KB, MB, up through YB as compactly as you can.e
package main

import (
	"fmt"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)

// const (
// 	base = 1000
// 	KB   = base      // 1000
// 	MB   = KB * base // 1000000
// 	GB   = MB * base // 1000000000
// 	TB   = GB * base // 1000000000000
// 	PB   = TB * base // 1000000000000000
// 	EB   = PB * base // 1000000000000000000
// 	ZB   = EB * base // 1000000000000000000000
// 	YB   = ZB * base // 1000000000000000000000000
// )

// RUNE
// const (
// 	base uint64 = 'Ϩ'       // Unicode code point = 10000
// 	KB          = base      // 1000
// 	MB          = KB * base // 1000000
// 	GB          = MB * base // 1000000000
// 	TB          = GB * base // 1000000000000	(exceeds 1 << 32)
// 	PB          = TB * base // 1000000000000000
// 	EB          = PB * base // 1000000000000000000
// 	// ZB          = EB * base // 1000000000000000000000 	(exceeds 1 << 64)
// 	// YB          = ZB * base // 1000000000000000000000000
// )

const (
	base = 1e1
	KB   = base      // 1000
	MB   = KB * base // 1000000
	GB   = MB * base // 1000000000
	TB   = GB * base // 1000000000000	(exceeds 1 << 32)
	PB   = TB * base // 1000000000000000
	EB   = PB * base // 1000000000000000000
	ZB   = EB * base // 1000000000000000000000 	(exceeds 1 << 64)
	YB   = ZB * base // 1000000000000000000000000
)

func main() {
	fmt.Printf("KiB\t= %g\n", KB)
	fmt.Printf("MiB\t= %g\n", MB)

	fmt.Printf("GiB\t= %g\n", GB)
	fmt.Printf("TiB\t= %g\n", TB)

	fmt.Printf("PiB\t= %g\n", PB)
	fmt.Printf("EiB\t= %g\n", EB)

	fmt.Printf("ZiB\t= %g\n", ZB)
	fmt.Printf("YiB\t= %g\n", YB)
}

// func main() {
// 	fmt.Printf("KiB\t= %d\n", KB)
// 	fmt.Printf("MiB\t= %d\n", MB)

// 	fmt.Printf("GiB\t= %d\n", GB)
// 	fmt.Printf("TiB\t= %d\n", TB)

// 	fmt.Printf("PiB\t= %d\n", PB)
// 	fmt.Printf("EiB\t= %d\n", EB)

// 	// fmt.Printf("ZiB\t= %d\n", ZB)
// 	// fmt.Printf("YiB\t= %d\n", YB)
// }
