package main

import (
	"bytes"
	"fmt"
	"io"
)

func printResult(w io.Writer, res string) {
	fmt.Fprintf(w, res)
}

func main() {
	var buf bytes.Buffer
	printResult(&buf, "diko is renewal, dictionary tool dico")
}
