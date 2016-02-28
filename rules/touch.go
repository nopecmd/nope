package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
)

const (
	touchBaseCommand = "touch"
)

func isMatchTouch(cmd models.Command) bool {
	return cmd.BaseCommand == touchBaseCommand
}

func getUndoTouch(cmd models.Command) string {
	// do things here
	return ""
}

func init() {
	match.AddRule(isMatchTouch, getUndoTouch)
}
