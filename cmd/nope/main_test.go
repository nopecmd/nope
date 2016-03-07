package main

import (
	"fmt"
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
		fmt.Println(err)
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

func TestMvDirs(t *testing.T) {
	var testFrom = "hello"
	var testTo = "world"

	// rename
	utils.CreateDir(testTo, t)

	var mvCommand = fmt.Sprintf("mv %s %s", testFrom, testTo)
	var expected = fmt.Sprintf("mv %s %s", testTo, testFrom)

	assert.Equal(t, testCommand(mvCommand, t), expected, getTestPrompt("mv"))

	//move
	var subDir = fmt.Sprintf("%s/%s", testTo, testFrom)
	utils.CreateDir(subDir, t)

	expected = fmt.Sprintf("mv %s .", subDir)
	assert.Equal(t, testCommand(mvCommand, t), expected, getTestPrompt("mv"))

	utils.RemoveDir(testTo, t)
}

func TestMvFiles(t *testing.T) {
	var testToDir = "bucket"
	var testToFile = "dog"
	var testToPath = fmt.Sprintf("%s/%s", testToDir, testToFile)

	var testFromDir = "house/room/bathroom"
	var testFromFile = "cat"
	var testFromPath = fmt.Sprintf("%s/%s", testFromDir, testFromFile)

	utils.CreateDir(testToDir, t)
	utils.CreateFile(testToPath, t)

	// rename
	var mvCommand = fmt.Sprintf("mv %s %s", testFromPath, testToPath)
	var expected = fmt.Sprintf("mv %s %s", testToPath, testFromPath)

	assert.Equal(t, testCommand(mvCommand, t), expected, getTestPrompt("mv"))

	utils.RemoveDir(testToDir, t)
}

func TestMvMultiple(t *testing.T) {
	testCommand("mv /this /should /be /multiple /commands", t)
}

func TestTouch(t *testing.T) {
	var testFileName = "testfile.txt"
	var rmCmd = "rm testfile.txt"

	utils.CreateFile(testFileName, t)

	var touchCmd = "touch " + testFileName
	assert.Equal(t, testCommand(touchCmd, t), rmCmd, getTestPrompt("touch"))

	var touchCmdDontCreate = "touch -c whatever another"
	assert.Equal(t, testCommand(touchCmdDontCreate, t), "", getTestPrompt("touch"))

	var touchCmdSymLink = "touch -r dontdelete " + testFileName
	assert.Equal(t, testCommand(touchCmdSymLink, t), rmCmd, getTestPrompt("touch"))

	utils.RemoveFile(testFileName, t)
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
	var testDir4 = "ferret"

	var testSubDir1 = testDir2 + "/" + testDir3
	var testSubDir2 = testDir2 + "/" + testDir4

	utils.CreateDir(testDir1, t)
	utils.CreateDir(testSubDir1, t)

	var mkdirCmd = "mkdir dog cat/turtle"
	var expected = "rm -rf dog cat/turtle"
	assert.Equal(t, testCommand(mkdirCmd, t), expected, getTestPrompt("mkdir"))

	var mkdirCmdWithP = "mkdir -p dog cat/turtle"
	var expectedWithP = "rm -rf dog cat"
	assert.Equal(t, testCommand(mkdirCmdWithP, t), expectedWithP, getTestPrompt("mkdir"))

	utils.CreateDir(testSubDir2, t)

	assert.Equal(t, testCommand(mkdirCmdWithP, t), expected, getTestPrompt("mkdir"))

	utils.RemoveDir(testDir1, t)
	utils.RemoveDir(testDir2, t)
}
