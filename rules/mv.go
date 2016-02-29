package rules

import (
	"fmt"
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

func buildMv(froms []string, to string) string {
	return fmt.Sprintf("%s %s %s", mvBaseCommand, strings.Join(froms, " "), to)
}

func getUndoMv(cmd models.Command) string {
	// for now, just remove all flags
	var err error
	cmd.TokensWithoutFlags, err = flags.ParseArgs(&mvFlags, cmd.Tokens[1:])

	if err != nil {
		panic(err)
	}

	var undoCommands []string
	var n = len(cmd.TokensWithoutFlags)
	var to = cmd.TokensWithoutFlags[n-1]
	for _, from := range cmd.TokensWithoutFlags[:n-1] {
		undoCommands = append(undoCommands, buildMv([]string{from}, to))
	}
	return shells.ConcatCommands(undoCommands)
}

func init() {
	match.AddRule(isMatchMv, getUndoMv)
}
