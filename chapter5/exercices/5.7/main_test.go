package main

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettify(t *testing.T) {
	input := `
<html>
<body>
    <div>
        <h1>This I magic!!!</h1>
        <p id="item123" class="text">It is True</p>
    </div>
</body>
</html>
`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}

	stdout := new(bytes.Buffer)
	prettyPrint(doc)
	_, err = html.Parse(bytes.NewReader(stdout.Bytes()))

	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
