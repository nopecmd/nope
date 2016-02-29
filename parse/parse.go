package parse

import (
	"github.com/flynn/go-shlex"
	"github.com/nopecmd/nope/models"
	"os"
	"strings"
)

func IsValidFilePath(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func GetFilePathsFromTokens(tokens []string) []string {
	var filePaths []string

	for _, token := range tokens {
		if IsValidFilePath(token) {
			filePaths = append(filePaths, token)
		}
	}
	return filePaths
}

func cleanFlag(s string) string {
	return strings.TrimPrefix(strings.TrimPrefix(s, "--"), "-")
}

func isFlag(s string) bool {
	return strings.HasPrefix(s, "-")
}

func flagMapFromTokens(tokens []string) map[string][]string {
	flags := make(map[string][]string)

	for i, token := range tokens {
		if isFlag(token) {
			var flag = cleanFlag(token)
			var flagValue = ""
			if i+1 < len(tokens) {
				flagValue = tokens[i+1]
			}

			if _, ok := flags[flag]; ok {
				flags[flag] = append(flags[flag], flagValue)
			} else {
				flags[flag] = []string{flagValue}
			}
		}
	}
	return flags
}

func ParseCommand(rawCmd string) (models.Command, error) {
	// shlex intelligently splits raw command string base upon shell style rules
	tokens, err := shlex.Split(rawCmd)
	if err != nil {
		return models.Command{}, err
	}

	return models.Command{
		RawCommandString: strings.Join(tokens, " "), // remove unneeded spaces
		BaseCommand:      tokens[0],
		Tokens:           tokens,
		Flags:            flagMapFromTokens(tokens),
	}, nil
}
