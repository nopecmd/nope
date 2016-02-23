package commands

import (
	"github.com/nopecmd/nope/match"
)

func isMatchGitAddAll(cmd string) bool {
	return cmd == "git add -A"
}

func getUndoGitAddAll(cmd string) string {
	return "git reset"
}

func init() {
	match.AddCommand(isMatchGitAddAll, getUndoGitAddAll)
}
