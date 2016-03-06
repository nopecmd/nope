package parse

import (
	"testing"

	"github.com/nopecmd/nope/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseCommandSplit(t *testing.T) {
	var commandWithRandomSpaces = "git  commit  -m \"there   are random spaces in here\""
	cmd, err := ParseCommand(commandWithRandomSpaces)
	if err != nil {
		t.Errorf(utils.FormatError(commandWithRandomSpaces, "could not parse command"))
	}
	if len(cmd.Tokens) != 4 {
		t.Errorf(utils.FormatError(commandWithRandomSpaces, "could not match command"))
	}
	assert.Equal(t, cmd.RawCommandString, "git commit -m there   are random spaces in here", "Parsing removes unnecessary spaces in the raw command")
}
