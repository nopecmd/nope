package main

import (
	"os"
	"testing"

	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
	"github.com/nopecmd/nope/utils"
	"github.com/stretchr/testify/assert"
)

func getTestPrompt(cmd string) string {
	return "the '" + cmd + "' command should be correctly matched"
}

func testParseCommand(rawCmd string, t *testing.T) (string, error) {
	cmd, err := parse.ParseCommand(rawCmd)
	if err != nil {
		t.Errorf(utils.FormatError(rawCmd, "could not parse command"))
	}

	return match.GetUndoCommand(cmd)
}

func testCommand(rawCmd string, t *testing.T) string {
	undo, err := testParseCommand(rawCmd, t)
	if err != nil {
		t.Errorf(utils.FormatError(rawCmd, "could not match command"))
	}

	return undo
}

func TestCd(t *testing.T) {
	assert.Equal(t, testCommand("cd ..", t), "cd -", getTestPrompt("cd"))
}

func TestGitAdd(t *testing.T) {
	assert.Equal(t, testCommand("git add -A", t), "git reset", getTestPrompt("git add"))
}

func TestMv(t *testing.T) {
	testCommand("mv -f /hello /world", t)
}

func TestMvMultiple(t *testing.T) {
	testCommand("mv /this /should /be /multiple /commands", t)
}

func TestTouch(t *testing.T) {
	var testFileName = "testfile.txt"
	var rmCmd = "rm testfile.txt"

	if _, err := os.Create(testFileName); err != nil {
		t.Errorf("could not create file: " + testFileName)
	}
	var touchCmd = "touch " + testFileName
	assert.Equal(t, testCommand(touchCmd, t), rmCmd, getTestPrompt("touch"))

	var touchCmdDontCreate = "touch -c whatever another"
	assert.Equal(t, testCommand(touchCmdDontCreate, t), "", getTestPrompt("touch"))

	var touchCmdSymLink = "touch -r dontdelete " + testFileName
	assert.Equal(t, testCommand(touchCmdSymLink, t), rmCmd, getTestPrompt("touch"))

	if err := os.Remove(testFileName); err != nil {
		t.Errorf("could not remove file: " + testFileName)
	}
}

func TestTouchBadFlags(t *testing.T) {
	var touchCmdWithBadFlags = "touch -z thing"
	_, err := testParseCommand(touchCmdWithBadFlags, t)
	assert.NotNil(t, err, "touch should not allow z flag")
}

func TestMkdir(t *testing.T) {
	var testDir1 = "dog"
	var testDir2 = "cat"
	var testDir3 = "turtle"
	var testSubDir = testDir2 + "/" + testDir3

	if err := os.MkdirAll(testDir1, 0777); err != nil {
		t.Errorf("could not create directory" + testDir1)
	}
	if err := os.MkdirAll(testSubDir, 0777); err != nil {
		t.Errorf("could not create directory" + testSubDir)
	}

	var mkdirCmd = "mkdir -v dog cat/turtle"
	var expected = "rm -rf dog cat/turtle"
	assert.Equal(t, testCommand(mkdirCmd, t), expected, getTestPrompt("mkdir"))

	if err := os.Remove(testDir1); err != nil {
		t.Errorf("could not remove direcotry " + testDir1)
	}
	if err := os.RemoveAll(testDir2); err != nil {
		t.Errorf("could not remove direcotry " + testDir2)
	}
}
