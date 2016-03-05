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

func getUndoMv(cmd models.Command) string {
	// for now, just remove all flags
	var err error
	cmd.TokensWithoutFlags, err = flags.ParseArgs(&mvFlags, cmd.Tokens[1:])

	if err != nil {
		panic(err)
	}

	var undoCommands []string
	var tokensLen = len(cmd.TokensWithoutFlags)
	var to = cmd.TokensWithoutFlags[tokensLen-1]
	for _, from := range cmd.TokensWithoutFlags[:tokensLen-1] {
		undoCommands = append(undoCommands, buildMv(to, from))
	}
	return shells.ConcatCommands(undoCommands)
}

func init() {
	match.AddRule(isMatchMv, getUndoMv)
}
