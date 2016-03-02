package nope_test

import (
	"testing"

	"github.com/nopecmd/nope/parse"
	"github.com/stretchr/testify/assert"
)

func TestParseCommandSplit(t *testing.T) {
	var commandWithRandomSpaces = "git  commit  -m \"there   are random spaces in here\""
	cmd, err := parse.ParseCommand(commandWithRandomSpaces)
	if err != nil {
		t.Errorf(formatError(commandWithRandomSpaces, "could not parse command"))
	}
	if len(cmd.Tokens) != 4 {
		t.Errorf(formatError(commandWithRandomSpaces, "incorrectly tokenized command"))
	}
	assert.Equal(t, cmd.RawCommandString, "git commit -m there   are random spaces in here", "Parsing removes unnecessary spaces in the raw command")
}
