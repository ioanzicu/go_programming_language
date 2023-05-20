/*
Define a deep comparison function that considers numbers (of any type)
equal if they differ by less than one part in a billion.
*/
package equal

import (
	"fmt"
	"reflect"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

// Equal reports whether x and y are deeply equal.
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

const multiplier = 1000 * 1000 * 1000

func numbersEqual(x, y float64) bool {
	if x == y {
		return true
	}
	var diff float64
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}
	d := diff * multiplier
	if d < x && d < y {
		return true
	}
	return false
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	if x.Type() != y.Type() {
		return false
	}

	// check cycle
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical reference
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return numbersEqual(float64(x.Int()), float64(y.Int()))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return numbersEqual(float64(x.Uint()), float64(y.Uint()))

	case reflect.Float32, reflect.Float64:
		return numbersEqual(float64(x.Float()), float64(y.Float()))

	case reflect.Complex64, reflect.Complex128:
		realEqual := numbersEqual(float64(real(x.Complex())), float64(real(y.Complex())))
		imagEqual := numbersEqual(float64(imag(x.Complex())), float64(imag(y.Complex())))
		return realEqual && imagEqual

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, key := range x.MapKeys() {
			if !equal(x.MapIndex(key), y.MapIndex(key), seen) {
				return false
			}
		}
		return true
	}

	panic("unreachable")
}

func main() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))        // "true"
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))      // "false"
	fmt.Println(Equal([]string(nil), []string{}))             // "true"
	fmt.Println(Equal(map[string]int(nil), map[string]int{})) // "true"

	// Circular linked lists a -> b -> a and c -> c.
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a)) // "true"
	fmt.Println(Equal(b, b)) // "true"
	fmt.Println(Equal(c, c)) // "true"
	fmt.Println(Equal(a, b)) // "false"
	fmt.Println(Equal(a, c)) // "false
}
