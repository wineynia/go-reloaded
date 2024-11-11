package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	fileB, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	ModifiedText := Goreloaded(string(fileB))

	os.WriteFile(os.Args[2], []byte(ModifiedText), 0644)

	testRunes := []rune{'a', '!', '?', '1', ';', ' '}
	for _, r := range testRunes {
		if IsPunct(r) {
			fmt.Printf("'%c' is a punctuation character.\n", r)
		} else {
			fmt.Printf("'%c' is not a punctuation character.\n", r)
		}
	}

	testRunes = []rune{' ', '\'', '"'}
	for _, r := range testRunes {
		if IsPunctQuote(r) {
			fmt.Printf("'%c' is a punctuation or quotation character.\n", r)
		} else {
			fmt.Printf("'%c' is not a punctuation or quotation character.\n", r)
		}
	}
}
