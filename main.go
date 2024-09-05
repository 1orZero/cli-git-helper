package main

import (
	"github.com/1orzero/git-helper-cli/internal/branch"
	"github.com/1orzero/git-helper-cli/internal/config"
	"github.com/1orzero/git-helper-cli/internal/openai"
	"github.com/1orzero/git-helper-cli/internal/utils"
)

func main() {
	config := config.LoadConfig()
	description := utils.GetDescription()
	llm := openai.InitializeOpenAIClient(config.APIEndpoint, config.APISecret)
	branchNames := branch.GenerateAndCleanBranchNames(llm, description, config.Username)
	selectedBranch := branch.SelectBranchName(branchNames)
	utils.CopyToClipboard(selectedBranch)
	utils.Output(selectedBranch)
}
