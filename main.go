package main

import (
	"bufio"
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

func query(word string, dict io.Reader) (res string) {
	s := bufio.NewScanner(dict)

	for s.Scan() {
		if s.Text() == word {
			s.Scan() // Skip matched word line. Next line has the meaning.
			return s.Text()
		}
	}

	return "NotFound"
}

func main() {
	var buf bytes.Buffer
	printResult(&buf, "diko is renewal, dictionary tool dico")
}
