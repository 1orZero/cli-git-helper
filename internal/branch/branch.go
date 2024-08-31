package branch

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/1orzero/git-helper-cli/internal/utils"
	"github.com/koki-develop/go-fzf"
	"github.com/tmc/langchaingo/llms"
)

func GenerateAndCleanBranchNames(llm llms.Model, description, username string) []string {
	branchNames, err := generateBranchNames(llm, description, username)
	if err != nil {
		utils.HandleError("Error generating branch names", err)
	}
	return cleanBranchNames(branchNames)
}

func generateBranchNames(llm llms.Model, description, username string) ([]string, error) {
	date := time.Now().Format("20060102")
	prompt := fmt.Sprintf(`Generate 10 git branch names based on this description: "%s".
	The branch names should follow this format: %s/%s-description-in-kebab-case.
	The description part should be concise and use hyphens instead of spaces.
	Each branch name should be unique.`, description, username, date)

	ctx := context.Background()
	response, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		return nil, err
	}

	return strings.Split(response, "\n"), nil
}

func cleanBranchNames(branchNames []string) []string {
	cleanedNames := make([]string, 0)
	for _, name := range branchNames {
		name = strings.TrimSpace(name)
		name = strings.TrimPrefix(name, fmt.Sprintf("%d. ", len(cleanedNames)+1))
		if name != "" {
			cleanedNames = append(cleanedNames, name)
		}
	}
	return cleanedNames
}

func SelectBranchName(branchNames []string) string {
	f, err := fzf.New(fzf.WithPrompt("Select a branch name: "))
	if err != nil {
		utils.HandleError("Error initializing fzf", err)
	}

	idx, err := f.Find(branchNames, func(i int) string {
		return branchNames[i]
	})
	if err != nil {
		utils.HandleError("Error during selection", err)
	}

	if len(idx) > 0 {
		selectedBranch := branchNames[idx[0]]
		fmt.Printf("Selected branch name: %s\n", selectedBranch)
		return selectedBranch
	}
	fmt.Println("No branch name selected")
	return ""
}
