// File: internal/app/app.go
package app

import (
	"github.com/1orzero/git-helper-cli/internal/cli"
	"github.com/1orzero/git-helper-cli/internal/config"
	"github.com/1orzero/git-helper-cli/internal/openai"
	"github.com/1orzero/git-helper-cli/internal/state"
	ucli "github.com/urfave/cli/v2"
)

func NewApp() *ucli.App {
	appState := &state.AppState{}

	app := &ucli.App{
		Name:  "git-helper",
		Usage: "A CLI tool to help with git workflows",
		Flags: cli.GlobalFlags(),
		Before: func(c *ucli.Context) error {
			cfg, err := config.LoadConfig(c.String("config"))
			if err != nil {
				return ucli.Exit("Error loading config: "+err.Error(), 1)
			}
			appState.Config = &cfg
			llm := openai.InitializeOpenAIClient(cfg.API.APIEndpoint, cfg.API.APISecret)
			appState.LLM = &llm
			return nil
		},
	}

	app.Commands = cli.Commands(appState)

	return app
}
