package rules

import (
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/parse"
)

func isMatchTouch(cmd models.Command) bool {
	return cmd.BaseCommand == touchBaseCommand
}

func getUndoTouch(cmd models.Command) (string, error) {
	var touchFlags struct {
		AccessAndModification string `short:"A"`
		Access                bool   `short:"a"`
		DontCreate            bool   `short:"c"`
		Force                 bool   `short:"f"`
		ChangeSymLink         bool   `short:"h"`
		Modification          bool   `short:"m"`
		SpecFile              string `short:"r"`
		Time                  string `short:"t"`
	}
	filteredTokens, err := flags.ParseArgs(&touchFlags, cmd.Tokens[1:])
	if err != nil {
		return "", err
	}
	if touchFlags.DontCreate || touchFlags.ChangeSymLink {
		return "", nil
	}

	var filePaths = parse.GetFilePathsFromTokens(filteredTokens)
	return rmBaseCommand + " " + strings.Join(filePaths, " "), nil
}

func init() {
	match.AddRule(isMatchTouch, getUndoTouch)
}
