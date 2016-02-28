package nope_test

import (
	"log"
	"os"
	"testing"

	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
)

func testCommand(rawCmd string, t *testing.T) {
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
	testCommand("cd ..", t)
}

func TestGitAdd(t *testing.T) {
	testCommand("git add -A", t)
}

func TestTouchSimple(t *testing.T) {
	var testFileName = "testfile.txt"

	if _, err := os.Create(testFileName); err != nil {
		t.Errorf("could not create file: " + testFileName)
	}
	var touchCmd = "touch " + testFileName
	testCommand(touchCmd, t)

	if err := os.Remove(testFileName); err != nil {
		t.Errorf("could not remove file: " + testFileName)
	}
}
