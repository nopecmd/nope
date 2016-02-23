package match

type rule struct {
	IsMatch func(string) bool
	GetUndo func(string) string
}

var Cmds []rule

func AddCommand(isMatch func(string) bool, getUndo func(string) string) {
	var cmdRule = rule{IsMatch: isMatch, GetUndo: getUndo}
	Cmds = append(Cmds, cmdRule)
}
