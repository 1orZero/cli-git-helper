package openai

import (
	"github.com/1orzero/git-helper-cli/internal/utils"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func InitializeOpenAIClient(apiEndpoint, apiSecret string) llms.Model {
	llm, err := openai.New(
		openai.WithBaseURL(apiEndpoint),
		openai.WithToken(apiSecret),
	)
	if err != nil {
		utils.HandleError("Error initializing OpenAI client", err)
	}
	return llm
}
