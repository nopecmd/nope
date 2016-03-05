package rules

import (
	"github.com/jessevdk/go-flags"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/parse"
	"strings"
)

func isMatchMkdir(cmd models.Command) bool {
	return cmd.BaseCommand == mkdirBaseCommand
}

func getUndoMkdir(cmd models.Command) (string, error) {
	var mkdirFlags struct {
		Mode         string `short:"m"`
		CreateInterm bool   `short:"p"`
		Verbose      bool   `short:"v"`
	}
	filteredTokens, err := flags.ParseArgs(&mkdirFlags, cmd.Tokens[1:])
	if err != nil {
		return "", err
	}
	var dirPaths = parse.GetFilePathsFromTokens(filteredTokens)

	return rmBaseCommand + " -rf " + strings.Join(dirPaths, " "), nil
}

func init() {
	match.AddRule(isMatchMkdir, getUndoMkdir)
}
