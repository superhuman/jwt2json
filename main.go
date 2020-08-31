package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), ".")
	b, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	if err := json.Indent(buf, b, "", "  "); err != nil {
		panic(err)
	}
	if _, err := io.Copy(os.Stdout, buf); err != nil {
		panic(err)
	}
}
