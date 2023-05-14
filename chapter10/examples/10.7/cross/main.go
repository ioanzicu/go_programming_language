//go:build linux || darwin
// +build linux darwin

package main

// go doc go/build

import (
	"fmt"
	"runtime"
)

/*
go run main.go
linux amd64

GOARCH=arm go run main.go
*/

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
