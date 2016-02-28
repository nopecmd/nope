package nope_test

import (
	"log"
	"testing"

	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
)

func formatError(rawCmd string, msg string) string {
	return rawCmd + " command failed: " + msg
}

func testCmd(rawCmd string, t *testing.T) {
	cmd, err := parse.ParseCommand(rawCmd)
	if err != nil {
		t.Errorf(formatError(rawCmd, "could not parse command"))
	}

	undo, err := match.GetUndoCommand(cmd)
	if err != nil {
		t.Errorf(formatError(rawCmd, "could not match command"))
	}
	log.Println(undo)
}

func TestCd(t *testing.T) {
	testCmd("cd ..", t)
}

func TestGitAdd(t *testing.T) {
	testCmd("git add -A", t)
}
