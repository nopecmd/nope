package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
)

func isMatchCd(cmd models.Command) bool {
	return cmd.BaseCommand == "cd"
}

func getUndoCd(cmd models.Command) string {
	return "cd -"
}

func init() {
	match.AddRule(isMatchCd, getUndoCd)
}
