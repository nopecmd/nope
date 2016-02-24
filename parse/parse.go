package parse

import (
	"github.com/nopecmd/nope/models"
	"strings"
)

func ParseCommand(rawCmd string) models.Command {
	var tokens = strings.Fields(rawCmd)

	return models.Command{
		RawCommandString: rawCmd,
		BaseCommand:      tokens[0],
		Tokens:           tokens,
	}
}
