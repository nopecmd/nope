package nope_test

import (
	"os"
	"testing"

	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
	"github.com/stretchr/testify/assert"
)

func testCommand(rawCmd string, t *testing.T) string {
	cmd, err := parse.ParseCommand(rawCmd)
	if err != nil {
		t.Errorf(formatError(rawCmd, "could not parse command"))
	}

	undo, err := match.GetUndoCommand(cmd)
	if err != nil {
		t.Errorf(formatError(rawCmd, "could not match command"))
	}
	return undo
}

func TestCd(t *testing.T) {
	assert.Equal(t, testCommand("cd ..", t), "cd -", "the 'cd' command should be correctly matched")
}

func TestGitAdd(t *testing.T) {
	assert.Equal(t, testCommand("git add -A", t), "git reset", "the 'git add -A' command should be correctly matched")
}

func TestMv(t *testing.T) {
	testCommand("mv -f /hello /world", t)
}

func TestMvMultiple(t *testing.T) {
	testCommand("mv /this /should /be /multiple /commands", t)
}

func TestTouchSimple(t *testing.T) {
	var testFileName = "testfile.txt"

	if _, err := os.Create(testFileName); err != nil {
		t.Errorf("could not create file: " + testFileName)
	}
	var touchCmd = "touch " + testFileName
	assert.Equal(t, testCommand(touchCmd, t), "rm testfile.txt", "the 'touch' command should be correctly matched")

	if err := os.Remove(testFileName); err != nil {
		t.Errorf("could not remove file: " + testFileName)
	}
}
