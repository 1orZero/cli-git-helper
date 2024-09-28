# This is a WIP project

# git-helper-cli

A CLI tool to automate git related tasks by using OpenAI's completion API.

## Roadmap

- [x] Generate branch names based on a description
- [x] Allow user to specify the branch format (by passing the format to the prompt
- [x] Generate a commit message based on the current directory staged files

## Usage

```bash
git-helper-cli
```

## Configuration

You can configure the tool by setting the following environment variables:

- `GIT_HELPER_USERNAME`: Your username (used to generate the branch names,)
- `GIT_HELPER_OPENAI_API_ENDPOINT`: Your OpenAI API endpoint
- `GIT_HELPER__OPENAI_API_SECRET`: Your OpenAI API secret
