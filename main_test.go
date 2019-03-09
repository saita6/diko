package main

import (
	"bytes"
	"testing"
)

func TestPrintResult(t *testing.T) {
	var buf bytes.Buffer
	printResult(&buf)

	got := buf.String()
	want := "diko is renewal, dictionary tool dico"
	if got != want {
		t.Fatalf("printResult() want=%s, got=%s", want, got)
	}

}
