package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func printResult(w io.Writer, res string) {
	fmt.Fprintf(w, res)
}

func newDictionary(dict string) io.Reader {
	return strings.NewReader(dict)
}

func main() {
	var buf bytes.Buffer
	printResult(&buf, "diko is renewal, dictionary tool dico")
}
