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

func TestParseCommandFlags(t *testing.T) {
	var commandWithLotsOfFlags = "curl -O URL1 -O URL2 -m \"there   are spaces here\" --data 'also here'"
	cmd, err := parse.ParseCommand(commandWithLotsOfFlags)
	if err != nil {
		t.Errorf(formatError(commandWithLotsOfFlags, "could not parse command"))
	}

	var result = map[string][]string{
		"O":    []string{"URL1", "URL2"},
		"m":    []string{"there   are spaces here"},
		"data": []string{"also here"},
	}
	assert.Equal(t, cmd.Flags, result, "Parsing should create the appropriate map")
}
