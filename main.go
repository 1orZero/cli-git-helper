package main

import (
	"github.com/1orzero/git-helper-cli/internal/branch"
	"github.com/1orzero/git-helper-cli/internal/config"
	"github.com/1orzero/git-helper-cli/internal/openai"
	"github.com/1orzero/git-helper-cli/internal/utils"
)

func main() {
	// Load config from environment variables
	config := config.LoadConfig()

	// Get description from user input
	description := utils.GetDescription()

	// Initialize OpenAI client
	llm := openai.InitializeOpenAIClient(config.APIEndpoint, config.APISecret)

	// Generate 10 branch names based on the description
	branchNames := branch.GenerateAndCleanBranchNames(llm, description, config.Username)

	// Select a branch name from the list (using fzf)
	selectedBranch := branch.SelectBranchName(branchNames)
	utils.CopyToClipboard(selectedBranch)

	// Output the selected branch name
	utils.Output(selectedBranch)
}
