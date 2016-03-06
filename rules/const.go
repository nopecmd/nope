package rules

// base commands
const (
	mvBaseCommand    = "mv"
	mkdirBaseCommand = "mkdir"
	touchBaseCommand = "touch"
	rmBaseCommand    = "rm"
	cdBaseCommand    = "cd"
)

// full commands
const gitAddAllFullCommand = "git add -A"

// undo commands
const (
	gitAddAllUndoCommand = "git reset"
	cdUndoCommand        = "cd -"
)
