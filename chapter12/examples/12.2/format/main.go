package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Any formats any value as a string
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
		// floating point and complex cases omitted
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // "1"
	fmt.Println(Any(d))                  // "1"
	fmt.Println(Any([]int64{x}))         // "[]int64 0xc00012c008"
	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0xc00012c020"
}
