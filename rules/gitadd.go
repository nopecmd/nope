package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
)

func isMatchGitAddAll(cmd models.Command) bool {
	return cmd.RawCommandString == "git add -A"
}

func getUndoGitAddAll(cmd models.Command) string {
	return "git reset"
}

func init() {
	match.AddRule(isMatchGitAddAll, getUndoGitAddAll)
}
