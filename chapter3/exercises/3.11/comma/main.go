package main

import (
	"fmt"
	"strings"
)

func main() {
	// comma inserts commas in a non-negative decimal integer string
	fmt.Println(formatter("+1"))       // +1
	fmt.Println(formatter("+12"))      // +12
	fmt.Println(formatter("+123"))     // +123
	fmt.Println(formatter("+1234"))    // +1,234
	fmt.Println(formatter("+12345"))   // +12,345
	fmt.Println(formatter("+123456"))  // +123,456
	fmt.Println(formatter("+1234567")) // +1,234,567
	fmt.Println("\n")

	fmt.Println(formatter("-1"))       // -1
	fmt.Println(formatter("-12"))      // -12
	fmt.Println(formatter("-123"))     // -123
	fmt.Println(formatter("-1234"))    // -1,234
	fmt.Println(formatter("-12345"))   // -12,345
	fmt.Println(formatter("-123456"))  // -123,456
	fmt.Println(formatter("-1234567")) // -1,234,567
	fmt.Println("\n")

	fmt.Println(formatter("-1.1"))                // -1.1
	fmt.Println(formatter("-1.11"))               // -1.11
	fmt.Println(formatter("-1.11111"))            // -1.11111
	fmt.Println(formatter("-1.1111111111"))       // -1.1111111111
	fmt.Println(formatter("-12.1111111111"))      // -12.1111111111
	fmt.Println(formatter("-123.1111111111"))     // -123.1111111111
	fmt.Println(formatter("-1234.1111111111"))    // -1,234.1111111111
	fmt.Println(formatter("-12345.1111111111"))   // -12,345.1111111111
	fmt.Println(formatter("-123456.1111111111"))  // -123,456.1111111111
	fmt.Println(formatter("-1234567.1111111111")) // -1,234,567.1111111111
}

func formatter(s string) string {
	start, end := 0, len(s)

	plus := strings.HasPrefix(s, "+")
	minus := strings.HasPrefix(s, "-")
	if plus || minus {
		start++ // not include the sign
	}

	dotIdx := strings.LastIndex(s, ".")
	if dotIdx != -1 {
		end = dotIdx // mark the dot index the end of the string
	}

	result := comma(s[start:end]) // get commas

	// add sign
	if plus {
		result = "+" + result
	} else if minus {
		result = "-" + result
	}

	// add floating part
	if dotIdx != -1 {
		return result + s[dotIdx:]
	}
	return result
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

/*
12345
n = 5

comma(s[:n-3]) + "," + s[n-3:]
	 	:5-3				   5-3:
	  n <= 3				 n <= 3
		"12"				  "345"

1234567890
n = 10

comma(s[:n-3]) + "," + s[n-3:]		"1,234,567" + "," + "890" = "1,234,567,890"
	 	:10-3			10-3:

comma("1234567")
n = 7

comma(s[:n-3]) + "," + s[n-3:]		"1,234" + "," + "567"
	  s[:4]            s[4:]

comma("1234")
n = 4

comma(s[:n-3]) + "," + s[n-3:]
comma(s[:1]) + "," + s[1:]			"1" + "," + "234" = "1,234"
*/
