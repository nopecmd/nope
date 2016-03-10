package rules

import (
	"fmt"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/shells"
	"github.com/nopecmd/nope/utils"
)

var mvFlags struct {
	Interactive bool `short:"i"`
	Force       bool `short:"f"`
}

func isMatchMv(cmd models.Command) bool {
	return cmd.BaseCommand == mvBaseCommand
}

func buildMv(oldFrom string, oldTo string) (string, error) {
	var oldFromTokens = strings.Split(oldFrom, "/")
	var oldToTokens = strings.Split(oldTo, "/")

	wasRenamed, err := utils.WasRenamed(oldFromTokens, oldToTokens)
	if err != nil {
		return "", err
	}

	if !wasRenamed {
		oldToTokens = append(oldToTokens, oldFromTokens[len(oldFromTokens)-1])
		oldFromTokens = oldFromTokens[:len(oldFromTokens)-1]
		if len(oldFromTokens) == 0 {
			oldFromTokens = []string{"."}
		}
	}

	var newFrom = strings.Join(oldToTokens, "/")
	var newTo = strings.Join(oldFromTokens, "/")

	return fmt.Sprintf("%s %s %s", mvBaseCommand, newFrom, newTo), nil
}

func getUndoMv(cmd models.Command) (string, error) {
	filteredTokens, err := flags.ParseArgs(&mvFlags, cmd.TokensWithoutBase)
	if err != nil {
		return "", err
	}

	var undoCommands []string
	var tokensLen = len(filteredTokens)
	var oldTo = utils.CleanPath(filteredTokens[tokensLen-1])

	for _, oldFrom := range filteredTokens[:tokensLen-1] {
		moveBack, err := buildMv(utils.CleanPath(oldFrom), oldTo)
		if err != nil {
			return "", err
		}
		undoCommands = append(undoCommands, moveBack)
	}

	return shells.ConcatCommands(undoCommands), nil
}

func init() {
	match.AddRule(isMatchMv, getUndoMv)
}
