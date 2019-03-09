package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func printResult(w io.Writer, res string) {
	fmt.Fprintf(w, res)
}

func openDictionary(dictName string) (io.Reader, error) {
	file, err := os.Open(dictName)
	if err != nil {
		return nil, fmt.Errorf("openDictionary() failed opening dictionary file, %v", err)
	}
	return file, nil
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
	const (
		ValidArgLen  = 2
		QueryWordPos = 1
	)
	if len(os.Args) != ValidArgLen {
		log.Fatalf("diko take only 1 argument. It's a query word.")
	}
	word := os.Args[QueryWordPos]

	file, err := openDictionary(os.Getenv("DIKODICT")) // DICODICT stores location of dicionary source file (path).
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	printResult(&buf, query(word, file))
	fmt.Printf("%v\n", buf.String())
}
