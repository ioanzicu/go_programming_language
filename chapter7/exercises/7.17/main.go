// Extend xmlselect so that elements may be selected not just by name,
// but by their attributes too, in the manner of CSS, so that, for instance,
// an element like <div id="page" class="wide"> could be selected
// by a matching id or class as well as its name.

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"text/scanner"
)

/*
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run main.go div div h2
html body div div h2: 1 Introduction
html body div div h2: 2 Documents
html body div div h2: 3 Logical Structures
html body div div h2: 4 Physical Structures
html body div div h2: 5 Conformance
html body div div h2: 6 Notation
html body div div h2: A References
html body div div h2: B Definitions for Character Normalization
html body div div h2: C Expansion of Entity and Character References (Non-Normative)
html body div div h2: D Deterministic Content Models (Non-Normative)
html body div div h2: E Autodetection of Character Encodings (Non-Normative)
html body div div h2: F W3C XML Working Group (Non-Normative)
html body div div h2: G W3C XML Core Working Group (Non-Normative)
html body div div h2: H Production Notes (Non-Normative)
html body div div h2: I Suggestions for XML Names (Non-Normative)
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("There are no selectors provided")
		os.Exit(0) // no selectors, nothing to do
	}

	sels, err := parseSelectors(strings.Join(os.Args[2:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	xmlselect(os.Stdout, os.Stdin, sels)
}

type attribute struct {
	Name  string
	Value string
}

func attrMatch(selAttrs []attribute, xmlAttrs []xml.Attr) bool {
SelectorAttribute:
	for _, sa := range selAttrs {
		for _, xa := range xmlAttrs {
			if sa.Name == xa.Name.Local && sa.Value == xa.Value || sa.Value == "" {
				continue SelectorAttribute
			}
		}
		return false
	}
	return true
}

type selector struct {
	tag   string
	attrs []attribute
}

func (s selector) String() string {
	b := &bytes.Buffer{}
	b.WriteString(s.tag)
	for _, attr := range s.attrs {
		switch attr.Value {
		case "":
			fmt.Fprintf(b, "[%s]", attr.Name)
		default:
			fmt.Fprintf(b, `[%s="%s"]`, attr.Name, attr.Value)
		}
	}
	return b.String()
}

type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

// describe returns a string describing the current token, for use in errors
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token)) // any other rune
}

func (lex *lexer) eatWhitespace() int {
	i := 0
	for lex.token == ' ' || lex.token == '\t' {
		lex.next()
		i++
	}
	return i
}

func (lex *lexer) next() {
	lex.token = lex.scan.Scan()
}

func (lex *lexer) text() string {
	return lex.scan.TokenText()
}

type lexPanic string

func parseSelectors(input string) (_ []selector, err error) {
	defer func() {
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			// unexpected panic: resume state of panic.
			panic(x)
		}
	}()

	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanStrings
	lex.scan.Whitespace = 0 // handle whitespace ourselves
	lex.next()              // initial lookahead

	selectors := make([]selector, 0)
	for lex.token != scanner.EOF {
		selectors = append(selectors, parseSelector(lex))
	}
	return selectors, nil
}

func parseSelector(lex *lexer) selector {
	var sel selector
	lex.eatWhitespace()
	if lex.token != '[' {
		if lex.token != scanner.Ident {
			panic(lexPanic(fmt.Sprintf("got %s, want ident", lex.describe())))
		}
		sel.tag = lex.text()
		lex.next() // consume tag ident
	}
	for lex.token == '[' {
		sel.attrs = append(sel.attrs, parseAttr(lex))
	}
	return sel
}

func parseAttr(lex *lexer) attribute {
	var attr attribute
	lex.next() // consume first square bracket`[`
	if lex.token != scanner.Ident {
		panic(lexPanic(fmt.Sprintf("got %s, want ident", lex.describe())))
	}
	attr.Name = lex.text()
	lex.next() // consume ident
	// No value given for the attribute
	if lex.token != '=' {
		if lex.token != ']' {
			panic(lexPanic(fmt.Sprintf("got %s, want ']'", lex.describe())))
		}
		lex.next()
		return attr
	}
	lex.next()
	switch lex.token {
	case scanner.Ident:
		attr.Value = lex.text()
	case scanner.String:
		attr.Value = strings.Trim(lex.text(), `"`)
	default:
		panic(lexPanic(fmt.Sprintf("got %s, want ident or string", lex.describe())))
	}
	lex.next() // consume value
	if lex.token != ']' {
		panic(lexPanic(fmt.Sprintf("got %s, want ']'", lex.describe())))
	}
	lex.next()
	return attr
}

func isSelected(stack []xml.StartElement, sels []selector) bool {
	if len(stack) < len(sels) {
		return false
	}
	start := len(sels) - len(stack)
	stack = stack[start:]
	for i := 0; i < len(sels); i++ {
		sel := sels[i]
		el := stack[i]
		if sel.tag != "" && sel.tag != el.Name.Local {
			return false
		}
		if !attrMatch(sel.attrs, el.Attr) {
			return false
		}
	}
	return true
}

func xmlselect(w io.Writer, r io.Reader, sels []selector) {
	dec := xml.NewDecoder(r)
	var stack []xml.StartElement //stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if isSelected(stack, sels) {
				fmt.Fprintf(w, "%s\n", tok)
			}
		}
	}
}
