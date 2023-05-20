/*
Make display safe to use on cyclic data structures by bounding
the number of steps it takes before abandoning the recursion.
(In Section 13.3, we’ll see another way to detect cycles.)
*/
package display

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0)
}

func display(path string, v reflect.Value, level int) {
	if level > 5 {
		fmt.Printf("%s = %s", path, formatAtom(v))
		return
	}
	level++
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		/*
			The logic is the same for both. The Len method returns the number of elements
			of a slice or array value, and Index(i) retrieves the element at index i, also as a
			reflect.Value; it panics if i is out of bounds. These are analogous to the built-in len(a)
			and a[i] operations on sequences. The display function recursively invokes itself on each
			element of the sequence, appending the subscript notation "[i]" to the path.

			Although reflect.Value has many methods, only a few are safe to call on any given value.
			For example, the Index method may be called on values of kind Slice, Array, or String, but
			panics for any other kind.
		*/
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), level)
		}
	case reflect.Struct:
		/*
			The NumField method reports the number of fields in the struct, and Field(i)
			returns the value of the i-th field as a reflect.Value. The list of fields includes ones
			promoted from anonymous fields. To append the field selector notation ".f" to the path, we
			must obtain the reflect.Type of the struct and access the name of its i-th field.
		*/
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), level)
		}
	case reflect.Map:
		/*
			The MapKeys method returns a slice of reflect.Values, one per map key. As usual
			when iterating over a map, the order is undefined. MapIndex(key) returns the value
			corresponding to key. We append the subscript notation "[key]" to the path.
			(We’re cutting a corner here. The type of a map key isn’t restricted to the types formatAtom handles best;
			arrays, structs, and interfaces can also be valid map keys.
			Extending this case to print the key in full is Exercise 12.1.)
		*/
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatKey(key)), v.MapIndex(key), level)
		}
	case reflect.Ptr:
		/*
			The Elem method returns the variable pointed to by a pointer, again as a
			reflect.Value. This operation would be safe even if the pointer value is nil,
			in which case the result would have kind Invalid, but we use IsNil to detect
			nil pointers explicitly so we can print a more appropriate message. We prefix
			the path with a "*" and parenthesize it to avoid ambiguity.
		*/
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), level)
		}
	case reflect.Interface:
		/*
			Again, we use IsNil to test whether the interface is nil,
			and if not, we retrieve its dynamic value using v.Elem()
			and print its type and value.
		*/
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), level)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

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
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func formatKey(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		b := &bytes.Buffer{}
		b.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i != 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(b, "%s: %s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		b.WriteByte('}')
		return b.String()
	case reflect.Array:
		b := &bytes.Buffer{}
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(formatAtom(v.Index(i)))
		}
		return b.String()
	default:
		return formatAtom(v)
	}
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", strangelove)

	/*
		go run main.go

		Display strangelove (main.Movie):
		strangelove.Title = "Dr. Strangelove"
		strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
		strangelove.Year = 1964
		strangelove.Color = false
		strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
		strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
		strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
		strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
		strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
		strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
		strangelove.Oscars[0] = "Best Actor (Nomin.)"
		strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
		strangelove.Oscars[2] = "Best Director (Nomin.)"
		strangelove.Oscars[3] = "Best Picture (Nomin.)"
		strangelove.Sequel = nil
	*/

	Display("os.Stderr", os.Stderr)

	/*
		Display os.Stderr (*os.File):
		(*(*os.Stderr).file).pfd.fdmu.state = 0
		(*(*os.Stderr).file).pfd.fdmu.rsema = 0
		(*(*os.Stderr).file).pfd.fdmu.wsema = 0
		(*(*os.Stderr).file).pfd.Sysfd = 2
		(*(*os.Stderr).file).pfd.pd.runtimeCtx = uintptr value
		(*(*os.Stderr).file).pfd.iovecs = nil
		(*(*os.Stderr).file).pfd.csema = 0
		(*(*os.Stderr).file).pfd.isBlocking = 1
		(*(*os.Stderr).file).pfd.IsStream = true
		(*(*os.Stderr).file).pfd.ZeroReadIsEOF = true
		(*(*os.Stderr).file).pfd.isFile = true
		(*(*os.Stderr).file).name = "/dev/stderr"
		(*(*os.Stderr).file).dirinfo = nil
		(*(*os.Stderr).file).nonblock = false
		(*(*os.Stderr).file).stdoutOrErr = true
		(*(*os.Stderr).file).appendMode = false
	*/

	Display("rV", reflect.ValueOf(os.Stderr))

	/*
		Display rV (reflect.Value):
		(*rV.typ).size = uintptr value
		(*rV.typ).ptrdata = uintptr value
		(*rV.typ).hash = 871609668
		(*rV.typ).tflag = 9
		(*rV.typ).align = 8
		(*rV.typ).fieldAlign = 8
		(*rV.typ).kind = 54
		(*rV.typ).equal = func(unsafe.Pointer, unsafe.Pointer) bool 0x402c20
		(*(*rV.typ).gcdata) = 1
		(*rV.typ).str = 4101
		(*rV.typ).ptrToThis = 0
		rV.ptr = unsafe.Pointer value
		rV.flag = reflect.flag value
	*/

	var i interface{} = 3

	Display("i", i)
	/*
		Display calls reflect.ValueOf(i), which returns a value of kind Int.
		As we mentioned in Section 12.2, reflect.ValueOf always returns
		a Value of a concrete type since it extracts the contents of an interface value.
	*/
	/*
		Display i (int):
		i = 3
	*/

	Display("&i", &i)
	/*
		Display calls reflect.ValueOf(&i), which returns a pointer to i, of
		kind Ptr. The switch case for Ptr calls Elem on this value, which returns a Value representing
		the variable i itself, of kind Interface. A Value obtained indirectly, like this one, may rep-
		resent any value at all, including interfaces. The display function calls itself recursively and
		this time, it prints separate components for the interface’s dynamic type and value.
	*/
	/*
		Display &i (*interface {}):
		(*&i).type = int
		(*&i).value = 3
	*/

	// a struct that points to itself
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)
	/*
		c.Value = 42
		(*c.Tail).Value = 42
		(*(*c.Tail).Tail).Value = 42
		(*(*(*c.Tail).Tail).Tail).Value = 42
		...ad infinitum...
	*/
}
