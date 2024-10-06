package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/1orzero/git-helper-cli/internal/utils"
)

func GetDescription() string {
	fmt.Print("Enter branch description: ")
	reader := bufio.NewReader(os.Stdin)
	description, err := reader.ReadString('\n')
	if err != nil {
		utils.HandleError("Error reading input", err)
	}
	description = strings.TrimSpace(description)
	if description == "" {
		fmt.Println("Description cannot be empty. Please try again.")
		os.Exit(1)
	}
	return description
}
