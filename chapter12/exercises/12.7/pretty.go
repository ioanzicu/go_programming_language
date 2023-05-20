package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

const margin = 80

// Marshal encodes a Go value in S-expression form
func MarshalIndent(v interface{}) ([]byte, error) {
	p := printer{width: margin}
	if err := pretty(&p, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return p.Bytes(), nil
}

type token struct {
	kind rune // one of "s ()" (string, blank, start, end)
	str  string
	size int
}

type printer struct {
	tokens []*token // FIFO buffer
	stack  []*token //stack of open ' ' and '(' tokens
	rtotal int      // total number of spaces needed to print stream

	bytes.Buffer
	indents []int
	width   int // remaining space
}

func (p *printer) string(str string) {
	tok := &token{
		kind: 's',
		str:  str,
		size: len(str),
	}
	if len(p.stack) == 0 {
		p.print(tok)
	} else {
		p.tokens = append(p.tokens, tok)
		p.rtotal += len(str)
	}
}

func (p *printer) pop() (top *token) {
	last := len(p.stack) - 1
	top, p.stack = p.stack[last], p.stack[:last]
	return
}

func (p *printer) begin() {
	if len(p.stack) == 0 {
		p.rtotal = 1
	}
	t := &token{
		kind: '(',
		size: -p.rtotal,
	}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t) // push
	p.string("(")
}

func (p *printer) end() {
	p.string(")")
	p.tokens = append(p.tokens, &token{kind: ')'})
	x := p.pop()
	x.size += p.rtotal
	if x.kind == ' ' {
		p.pop().size += p.rtotal
	}
	if len(p.stack) == 0 {
		for _, tok := range p.tokens {
			p.print(tok)
		}
		p.tokens = nil
	}
}

func (p *printer) space() {
	last := len(p.stack) - 1
	x := p.stack[last]
	if x.kind == ' ' {
		x.size += p.rtotal
		p.stack = p.stack[:last] // pop
	}
	t := &token{
		kind: ' ',
		size: -p.rtotal,
	}
	p.tokens = append(p.tokens, t)
	p.stack = append(p.stack, t)
	p.rtotal++
}

func (p *printer) print(t *token) {
	switch t.kind {
	case 's':
		p.WriteString(t.str)
		p.width -= len(t.str)
	case '(': // push
		p.indents = append(p.indents, p.width)
	case ')': // pop
		p.indents = p.indents[:len(p.indents)-1]
	case ' ':
		if t.size > p.width {
			p.width = p.indents[len(p.indents)-1] - 1
			fmt.Fprintf(&p.Buffer, "\n%*s", margin-p.width, "")
		} else {
			p.WriteByte(' ')
			p.width--
		}
	}
}

func (p *printer) stringf(format string, args ...interface{}) {
	p.string(fmt.Sprintf(format, args...))
}

func pretty(p *printer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		p.string("nil")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p.stringf("%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		p.stringf("%d", v.Uint())

	case reflect.String:
		p.stringf("%q", v.String())

	case reflect.Ptr:
		return pretty(p, v.Elem())

	case reflect.Float32, reflect.Float64:
		p.stringf("%4.4f", v.Float())

	case reflect.Complex64, reflect.Complex128:
		p.stringf("#C(%4.4f %4.4f)", real(v.Complex()), imag(v.Complex()))

	case reflect.Bool:
		if v.Bool() {
			p.string("t")
		} else {
			p.string("nil")
		}

	case reflect.Array, reflect.Slice: // (value ...)
		p.begin()
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				p.space()
			}
			if err := pretty(p, v.Index(i)); err != nil {
				return err
			}
		}
		p.end()

	case reflect.Struct: // ((name value) ...)
		p.begin()
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				p.space()
			}
			p.begin()
			p.string(v.Type().Field(i).Name)
			if err := pretty(p, v.Field(i)); err != nil {
				return err
			}
			p.end()
		}
		p.end()

	case reflect.Map: // ((key value) ...)
		p.begin()
		for i, key := range v.MapKeys() {
			if i > 0 {
				p.space()
			}
			p.begin()
			if err := pretty(p, key); err != nil {
				return err
			}
			p.space()
			if err := pretty(p, v.MapIndex(key)); err != nil {
				return err
			}
			p.end()
		}
		p.end()

	default: // chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
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

	out, err := MarshalIndent(strangelove)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(string(out))

	/*
		go run main.go

		((Title"Dr. Strangelove")
		 (Subtitle"How I Learned to Stop Worrying and Love the Bomb") (Year1964)
		 (Colornil)
		 (Actor(("Brig. Gen. Jack D. Ripper" "Sterling Hayden")
		        ("Maj. T.J. \"King\" Kong" "Slim Pickens")
		        ("Dr. Strangelove" "Peter Sellers")
		        ("Grp. Capt. Lionel Mandrake" "Peter Sellers")
		        ("Pres. Merkin Muffley" "Peter Sellers")
		        ("Gen. Buck Turgidson" "George C. Scott")))
		 (Oscars("Best Actor (Nomin.)" "Best Adapted Screenplay (Nomin.)"
		         "Best Director (Nomin.)" "Best Picture (Nomin.)")) (Sequelnil))
	*/
}
