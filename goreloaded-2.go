package main

import (
	"unicode"
)

func IsPunct(r rune) bool {
	punct := []rune{'.', ',', '!', '?', ';'}

	for _, p := range punct {
		if r == p {
			return true
		}
	}
	return unicode.IsPunct(r)
}

func IsPunctQuote(r rune) bool {
	punctquote := []rune{'\'', '"'}

	for _, q := range punctquote {
		if r == q {
			return true
		}
	}
	return unicode.IsPunct(r)
}
