package rules

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/shells"
)

var mvFlags struct {
	Interactive bool `short:"i"`
	Force       bool `short:"f"`
}

func isMatchMv(cmd models.Command) bool {
	return cmd.BaseCommand == mvBaseCommand
}

func buildMv(from string, to string) string {
	return fmt.Sprintf("%s %s %s", mvBaseCommand, from, to)
}

func getUndoMv(cmd models.Command) (string, error) {
	filteredTokens, err := flags.ParseArgs(&mvFlags, cmd.TokensWithoutBase)
	if err != nil {
		return "", err
	}

	var undoCommands []string
	var tokensLen = len(filteredTokens)
	var to = filteredTokens[tokensLen-1]
	for _, from := range filteredTokens[:tokensLen-1] {
		undoCommands = append(undoCommands, buildMv(to, from))
	}
	return shells.ConcatCommands(undoCommands), nil
}

func init() {
	match.AddRule(isMatchMv, getUndoMv)
}
