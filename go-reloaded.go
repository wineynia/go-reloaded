package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function to process directives in the input string
func Goreloaded(input string) string {
	// Split the input string into words for easy processing
	words := strings.Fields(input)

	// Iterate through each word to identify and process directives
	for i := 1; i < len(words); i++ {
		// Convert previous word from binary to decimal if "(bin)" directive is found
		if words[i] == "(bin)" {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error converting binary to decimal:", err)
				os.Exit(1)
			}
			// Replace previous word with its decimal value
			words[i-1] = strconv.Itoa(int(num))
			words[i] = "" // Clear the directive word
		}

		// Convert previous word from hexadecimal to decimal if "(hex)" directive is found
		if words[i] == "(hex)" {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error converting hexadecimal to decimal:", err)
				os.Exit(1)
			}
			// Replace previous word with its decimal value
			words[i-1] = strconv.Itoa(int(num))
			words[i] = "" // Clear the directive word
		}

		// Convert previous word to lowercase if "(low)" directive is found
		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words[i] = "" // Clear the directive word
		}

		// Convert previous word to uppercase if "(up)" directive is found
		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = "" // Clear the directive word
		}

		// Capitalize the first letter of the previous word if "(cap)" directive is found
		if words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
			words[i] = "" // Clear the directive word
		}

		// Capitalize a specified number of previous words if "(cap, <number>)" directive is found
		if words[i] == "(cap," && i < len(words)-1 {
			n := StrToNum(words[i+1])
			for j := i - n; j < i; j++ {
				words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(words[j][1:])
			}
			words[i] = ""   // Clear the directive word
			words[i+1] = "" // Clear the number word
		}

		// Convert a specified number of previous words to lowercase if "(low, <number>)" directive is found
		if words[i] == "(low," && i < len(words)-1 {
			n := StrToNum(words[i+1])
			for j := i - n; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}
			words[i] = ""   // Clear the directive word
			words[i+1] = "" // Clear the number word
		}

		// Convert a specified number of previous words to uppercase if "(up, <number>)" directive is found
		if words[i] == "(up," && i < len(words)-1 {
			n := StrToNum(words[i+1])
			for j := i - n; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}
			words[i] = ""   // Clear the directive word
			words[i+1] = "" // Clear the number word
		}

		// Adjust "a" to "an" if the following word starts with a vowel
		if (words[i] == "a" || words[i] == "A") && i < len(words)-1 {
			nextWord := words[i+1]
			if IsVowel(rune(nextWord[0])) {
				words[i] += "n"
			}
		}
	}

	// Remove empty indexes from the list of words
	output := RemoveEmptyIndex(words)
	// Join the words into a single string with spaces
	ModifiedText := strings.Join(output, " ")
	// Adjust punctuation spacing
	ModifiedText = Punctuation(ModifiedText)
	// Adjust quotes around words
	ModifiedText = Quote(ModifiedText)
	return ModifiedText
}

// Checks if a character is a vowel
func IsVowel(r rune) bool {
	vowel := []rune{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}
	for _, p := range vowel {
		if r == p {
			return true
		}
	}
	return false
}

// Checks if a character is a punctuation mark
func IsPunct(r rune) bool {
	punct := []rune{'.', ',', '!', '?', ';', ':'}
	for _, p := range punct {
		if r == p {
			return true
		}
	}
	return false
}

// Adjusts spaces around punctuation in the text
func Punctuation(s string) string {
	var newString string
	for _, char := range s {
		if newString != "" {
			length := len(newString)
			lastChar := newString[length-1]
			if IsPunct(char) {
				if lastChar == ' ' {
					newString = newString[:length-1] // Remove space before punctuation
				}
			} else {
				if IsPunct(rune(lastChar)) && char != ' ' {
					newString += " " // Add space after punctuation
				}
			}
		}
		newString += string(char)
	}
	return newString
}

// Places quotation marks without spaces in the middle of a word
func Quote(s string) string {
	var newString string
	count := 0 // Track opening and closing quotes
	for _, char := range s {
		if newString != "" {
			length := len(newString)
			lastChar := newString[length-1]
			if char == '\'' {
				count++
				if count%2 == 0 { // Closing quote
					if lastChar == ' ' {
						newString = newString[:length-1] // Remove space before closing quote
					}
				}
			} else if char == ' ' && lastChar == '\'' && count%2 != 0 { // Ignore space after opening quote
				continue
			}
		}
		newString += string(char)
	}
	return newString
}

// Converts string ending with ',' to an integer
func StrToNum(s string) int {
	n := 0
	strconv.Atoi(s[:len(s)-1])
	n, _ = strconv.Atoi(s[:len(s)-1])
	return n
}

// Removes empty strings from the list
func RemoveEmptyIndex(s []string) []string {
	var rei []string
	for _, str := range s {
		if str != "" {
			rei = append(rei, str)
		}
	}
	return rei
}
