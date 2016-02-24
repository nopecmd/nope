package match

import (
	"errors"
	"github.com/nopecmd/nope/models"
)

var cmdRules []models.Rule

func AddCommand(isMatch func(string) bool, getUndo func(string) string) {
	var cmdRule = models.Rule{IsMatch: isMatch, GetUndo: getUndo}
	cmdRules = append(cmdRules, cmdRule)
}

func GetUndoCommand(prevCmd string) (string, error) {
	for _, rule := range cmdRules {
		if rule.IsMatch(prevCmd) {
			return rule.GetUndo(prevCmd), nil
		}
	}

	return "", errors.New("Unable to match command")
}
