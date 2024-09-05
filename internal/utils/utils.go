package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func HandleError(message string, err error) {
	fmt.Printf("%s: %v\n", message, err)
	os.Exit(1)
}

func CopyToClipboard(text string) {
	if text == "" {
		return
	}
	err := clipboard.WriteAll(text)
	if err != nil {
		fmt.Printf("Error copying to clipboard: %v\n", err)
	} else {
		fmt.Println("Branch name copied to clipboard!")
	}
}

func Output(text string) {
	fmt.Println(text)
}

func GetDescription() string {
	fmt.Print("Enter branch description: ")
	reader := bufio.NewReader(os.Stdin)
	description, err := reader.ReadString('\n')
	if err != nil {
		HandleError("Error reading input", err)
	}
	description = strings.TrimSpace(description)
	if description == "" {
		fmt.Println("Description cannot be empty. Please try again.")
		os.Exit(1)
	}
	return description
}
