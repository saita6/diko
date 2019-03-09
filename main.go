package main

import (
	"bytes"
	"fmt"
	"io"
)

func printResult(w io.Writer) {
	fmt.Fprintf(w, "diko is renewal, dictionary tool dico")
}

func main() {
	var buf bytes.Buffer
	printResult(&buf)
}
