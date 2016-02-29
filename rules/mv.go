package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
)

const (
	mvBaseCommand = "mv"
)

func isMatchMv(cmd models.Command) bool {
	return cmd.BaseCommand == mvBaseCommand
}

func getUndoMv(cmd models.Command) string {
	return ""
}
