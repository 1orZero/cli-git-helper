package cli

import (
	"strings"

	"github.com/1orzero/git-helper-cli/internal/branch"
	"github.com/1orzero/git-helper-cli/internal/state"
	"github.com/1orzero/git-helper-cli/internal/utils"
	ucli "github.com/urfave/cli/v2"
)

func Commands(appState *state.AppState) []*ucli.Command {
	return []*ucli.Command{
		{
			Name:      "generate-branch-name",
			Usage:     "Generate a branch name based on the issue type and title",
			Action:    generateBranchName(appState),
			ArgsUsage: "<description>",
		},
	}
}

func generateBranchName(appState *state.AppState) func(c *ucli.Context) error {
	return func(c *ucli.Context) error {
		description := strings.Join(c.Args().Slice(), " ")
		if description == "" {
			return ucli.Exit("Description is required", 1)
		}

		branchNames := branch.GenerateAndCleanBranchNames(*appState.LLM, description, *appState.Config)

		// Select a branch name from the list (using fzf)
		selectedBranch := branch.SelectBranchName(branchNames)
		utils.CopyToClipboard(selectedBranch)

		Output(selectedBranch)
		return nil
	}
}
