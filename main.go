package main

import (
	"fmt"
	"os"

	"github.com/1orzero/git-helper-cli/internal/app"
)

func main() {
	app := app.NewApp()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
