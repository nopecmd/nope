package match

import "errors"

type rule struct {
	IsMatch func(string) bool
	GetUndo func(string) string
}

var cmdRules []rule

func AddCommand(isMatch func(string) bool, getUndo func(string) string) {
	var cmdRule = rule{IsMatch: isMatch, GetUndo: getUndo}
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
