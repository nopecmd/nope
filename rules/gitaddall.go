package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
)

func isMatchGitAddAll(cmd models.Command) bool {
	return cmd.RawCommandString == gitAddAllFullCommand
}

func getUndoGitAddAll(cmd models.Command) string {
	return gitAddAllUndoCommand
}

func init() {
	match.AddRule(isMatchGitAddAll, getUndoGitAddAll)
}
