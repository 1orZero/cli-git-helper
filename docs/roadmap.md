# Roadmap

## v0.1.0 (Current)

- [x] Generate branch names based on a description
- [x] Allow user to pass description as an argument

## v0.1.1

- [x] Config file support
  - Support for format of the branch name
  - Support for OpenAI API endpoint and secret
  - Support branch format (by passing the format to the prompt in the config file)
- [ ] Allow user to pass the config file as an argument

## v0.1.2

- [ ] Help menu
- [ ] Feature list menu

## v0.1.3

- [ ] Can generate a commit message
- [ ] Generate a commit message based on the current directory staged files
- [ ] Generate a commit message based on the current directory staged files and unstaged files (include untracked files)

## v0.1.4

- [ ] Support for other LLM providers

## v0.1.5

- [ ] Create Pull Request action, open a browser to create a PR
- [ ] Generate the PR title based on the current branch name and the commit message
- [ ] Generate the PR description based on the commit message and the difference between the current branch and the base branch
