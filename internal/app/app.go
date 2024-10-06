package app

import (
	"github.com/1orzero/git-helper-cli/internal/cli"
	ucli "github.com/urfave/cli/v2"
)

func NewApp() *ucli.App {
	return &ucli.App{
		Name:     "git-helper",
		Usage:    "A CLI tool to help with git workflows",
		Flags:    cli.GlobalFlags(),
		Commands: cli.Commands(),
	}
}
