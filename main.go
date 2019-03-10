package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

// addWordToSrc adds a new word/meaning pair to a specified source file of dictionary.
func addWordToSrc(word string, meaning string, dictName string) {
	file, err := os.OpenFile(dictName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatalf("addWord() failed at adding new word, %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(word); err != nil {
		log.Fatalf("addWordToSrc() failed %v, word=%v", err, word)
	}
	if _, err := file.Write([]byte{'\n'}); err != nil {
		log.Fatalf("addWordToSrc() failed at adding new line '\n' %v", err)
	}
	if _, err := file.WriteString(meaning); err != nil {
		log.Fatalf("addWordToSrc() failed %v, meaning=%v", err, meaning)
	}
	if _, err := file.Write([]byte{'\n'}); err != nil {
		log.Fatalf("addWordToSrc() failed at adding new line '\n' %v", err)
	}
}

func main() {
	var AddWordMode *bool = flag.Bool("a", false, "mode: add new 'word:meaning'")
	flag.Parse()

	if *AddWordMode {
		const (
			ValidArgLen    = 1
			ArgPos         = 0
			WordMeaningSep = ":"
			WordPos        = 0
			MeaningPos     = 1
		)

		args := flag.Args()
		if len(args) != ValidArgLen {
			log.Fatalf("diko, -a option needs a word & meaning pair")
		}

		arg := strings.Split(args[ArgPos], WordMeaningSep)
		word := arg[WordPos]
		meaning := arg[MeaningPos]

		addWordToSrc(word, meaning, os.Getenv("DIKODICT"))
		return
	}

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
