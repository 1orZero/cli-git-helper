package cli

import (
	"strings"

	"github.com/1orzero/git-helper-cli/internal/branch"
	"github.com/1orzero/git-helper-cli/internal/config"
	"github.com/1orzero/git-helper-cli/internal/openai"
	"github.com/1orzero/git-helper-cli/internal/utils"
	ucli "github.com/urfave/cli/v2"
)

func Commands() []*ucli.Command {
	return []*ucli.Command{
		{
			Name:      "generate-branch-name",
			Usage:     "Generate a branch name based on the issue type and title",
			Action:    generateBranchName,
			ArgsUsage: "<description>",
		},
	}
}

func generateBranchName(c *ucli.Context) error {
	cfg, err := config.LoadConfig(c.String("config"))
	if err != nil {
		return ucli.Exit("Error loading config: "+err.Error(), 1)
	}
	description := strings.Join(c.Args().Slice(), " ")
	if description == "" {
		return ucli.Exit("Description is required", 1)
	}

	// Initialize OpenAI client
	llm := openai.InitializeOpenAIClient(cfg.API.APIEndpoint, cfg.API.APISecret)

	branchNames := branch.GenerateAndCleanBranchNames(llm, description, cfg)

	// Select a branch name from the list (using fzf)
	selectedBranch := branch.SelectBranchName(branchNames)
	utils.CopyToClipboard(selectedBranch)

	Output(selectedBranch)
	return nil
}
