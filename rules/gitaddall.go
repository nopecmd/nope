package rules

import (
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
)

const (
	gitAddAllFullCommand = "git add -A"
	gitAddAllUndoCommand = "git reset"
)

func isMatchGitAddAll(cmd models.Command) bool {
	return cmd.RawCommandString == gitAddAllFullCommand
}

func getUndoGitAddAll(cmd models.Command) (string, error) {
	return gitAddAllUndoCommand, nil
}

func init() {
	match.AddRule(isMatchGitAddAll, getUndoGitAddAll)
}
