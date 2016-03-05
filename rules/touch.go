package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/parse"
	"strings"
)

func isMatchTouch(cmd models.Command) bool {
	return cmd.BaseCommand == touchBaseCommand
}

func getUndoTouch(cmd models.Command) string {
	var filePaths = parse.GetFilePathsFromTokens(cmd.Tokens)
	return rmBaseCommand + " " + strings.Join(filePaths, " ")
}

func init() {
	match.AddRule(isMatchTouch, getUndoTouch)
}
