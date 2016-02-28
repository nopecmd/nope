package parse

import (
	"github.com/flynn/go-shlex"
	"github.com/nopecmd/nope/models"
	"os"
	"strings"
)

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
	}, nil
}

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
