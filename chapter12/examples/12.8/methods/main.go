package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Print prints the method set of the value x
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methodType.String(), "func"))
	}
}

func main() {
	Print(time.Hour)
	/*
		type time.Duration
		func (time.Duration) Hours() float64
		func (time.Duration) Microseconds() int64
		func (time.Duration) Milliseconds() int64
		func (time.Duration) Minutes() float64
		func (time.Duration) Nanoseconds() int64
		func (time.Duration) Round(time.Duration) time.Duration
		func (time.Duration) Seconds() float64
		func (time.Duration) String() string
		func (time.Duration) Truncate(time.Duration) time.Duration
	*/
	Print(new(strings.Replacer))
	/*
		type *strings.Replacer
		func (*strings.Replacer) Replace(string) string
		func (*strings.Replacer) WriteString(io.Writer, string) (int, error)
	*/
}
