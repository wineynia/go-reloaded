package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Goreloaded(input string) string {
	words := strings.Fields(input)

	for i := 1; i < len(words); i++ {
		if words[i] == "(bin)" {
			num, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error converting binary to decimal:", err)
				os.Exit(1)
			}

			words[i-1] = strconv.Itoa(int(num))
			words[i] = ""
		}

		if words[i] == "(hex)" {
			num, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error converting hexadecimal to decimal:", err)
				os.Exit(1)
			}

			words[i-1] = strconv.Itoa(int(num))
			words[i] = ""
		}

		if words[i] == "(low)" {
			words[i-1] = strings.ToLower(words[i-1])
			words[i] = ""
		}

		if words[i] == "(up)" {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""
		}
		if words[i] == "(cap)" {
			words[i-1] = strings.ToUpper(string(words[i-1][0])) + strings.ToLower(words[i-1][1:])
			words[i] = ""
		}

		if words[i] == "(cap," && i < len(words)-1 {
			n := StrToNum(words[i+1])
			for j := i - n; j < i; j++ {
				words[j] = strings.ToUpper(string(words[j][0])) + strings.ToLower(words[j][1:])
			}
			words[i] = ""
			words[i+1] = ""
		}

		if words[i] == "(low," && i < len(words)-1 {
			n := StrToNum(words[i+1])
			for j := i - n; j < i; j++ {
				words[j] = strings.ToLower(words[j])
			}
			words[i] = ""
			words[i+1] = ""
		}

		if words[i] == "(up," && i < len(words)-1 {
			n := StrToNum(words[i+1])
			for j := i - n; j < i; j++ {
				words[j] = strings.ToUpper(words[j])
			}
			words[i] = ""
			words[i+1] = ""
		}

	}
	output := PunQuoVow(words)
	return strings.Join(output, " ")
}

func StrToNum(s string) (n int) {
	strconv.Atoi(s[:len(s)-1])
	n, _ = strconv.Atoi(s[:len(s)-1])

	return
}

func PunQuoVow(s []string) (pqv []string) {
	for _, a := range s {
		if a != "" {
			pqv = append(pqv, a)
		}
	}
	return
}
