package commands

import (
	"github.com/nopecmd/nope/match"
	"strings"
)

func isMatchCd(cmd string) bool {
	return strings.HasPrefix(cmd, "cd ")
}

func getUndoCd(cmd string) string {
	return "cd -"
}

func init() {
	match.AddCommand(isMatchCd, getUndoCd)
}
