package rules

import (
	"github.com/jessevdk/go-flags"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/shells"
	"strings"
)

var mvFlags struct {
	Interactive bool `short:"i"`
	Force       bool `short:"f"`
}

const (
	mvBaseCommand = "mv"
)

func isMatchMv(cmd models.Command) bool {
	return cmd.BaseCommand == mvBaseCommand
}

func getUndoMv(cmd models.Command) string {
	// for now, just remove all flags
	var err error
	cmd.TokensWithoutFlags, err = flags.ParseArgs(&mvFlags, cmd.Tokens[1:])

	if err != nil {
		panic(err)
	}

	var undoCommands []string
	var dest = cmd.TokensWithoutFlags[len(cmd.TokensWithoutFlags)-1]
	for _, from := range cmd.TokensWithoutFlags[0 : len(cmd.TokensWithoutFlags)-1] {
		var commandTokens = []string{
			mvBaseCommand,
			dest,
			from,
		}
		undoCommands = append(undoCommands, strings.Join(commandTokens, " "))
	}
	return shells.ConcatCommands(undoCommands)
}

func init() {
	match.AddRule(isMatchMv, getUndoMv)
}
