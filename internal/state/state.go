package state

import (
	"github.com/1orzero/git-helper-cli/internal/config"
	"github.com/tmc/langchaingo/llms"
)

type AppState struct {
	Config *config.Config
	LLM    *llms.Model
}
