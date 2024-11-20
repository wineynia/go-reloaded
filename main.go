package main

import (
	"fmt"
	"os"
)

func main() {

	//checking if the number of argument is only 3
	if len(os.Args) != 3 {
		fmt.Println("Error: wrong numbers of arguments.")
		return
	}

	//reading sample.txt
	fileB, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	//fix the text file
	ModifiedText := Goreloaded(string(fileB))

	//Writing the modified text to result.txt
	err = os.WriteFile(os.Args[2], []byte(ModifiedText), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
