// Write a function expand(s string, f func(string) string) string that
// replaces each substring ‘‘$foo’’ within s by the text returned by f("foo").
package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {

	s := "$this is very $interesting thing?"

	res := expand(s, strings.ToUpper)
	if res != "THIS is very INTERESTING thing?" {
		t.Errorf("not matched: %s", res)
	}

	s = "$THIS IS VERY $INTERESTING THING?"

	res = expand(s, strings.ToLower)
	if res != "this IS VERY interesting THING?" {
		t.Errorf("not matched: %s", res)
	}
}
