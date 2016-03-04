package match

import (
	"errors"
	"github.com/nopecmd/nope/models"
)

var cmdRules []models.Rule

func AddRule(isMatch func(models.Command) bool, getUndo func(models.Command) (string, error)) {
	var rule = models.Rule{IsMatch: isMatch, GetUndo: getUndo}
	cmdRules = append(cmdRules, rule)
}

func GetUndoCommand(cmd models.Command) (string, error) {
	for _, rule := range cmdRules {
		if rule.IsMatch(cmd) {
			return rule.GetUndo(cmd)
		}
	}

	return "", errors.New("Unable to match command")
}
