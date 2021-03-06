package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
)

func isMatchCd(cmd models.Command) bool {
	return cmd.BaseCommand == cdBaseCommand
}

func getUndoCd(cmd models.Command) (string, error) {
	return cdUndoCommand, nil
}

func init() {
	match.AddRule(isMatchCd, getUndoCd)
}
