package nope_test

import (
	"log"
	"testing"

	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
)

func testCmd(raw string) (string, error) {
	var cmd = parse.ParseCommand(raw)
	return match.GetUndoCommand(cmd)
}

func TestCd(t *testing.T) {
	undo, err := testCmd("cd ..")
	if err != nil {
		t.Errorf("Cd command failed")
	}
	log.Println(undo)
}

func TestGitAdd(t *testing.T) {
	undo, err := testCmd("git add -A")
	if err != nil {
		t.Errorf("Cd command failed")
	}
	log.Println(undo)
}
