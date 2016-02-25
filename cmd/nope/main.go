package main

import (
	"fmt"
	"github.com/rogpeppe/rog-go/reverse"
	"os"

	_ "github.com/nopecmd/nope/commands"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
	"github.com/nopecmd/nope/shells"
)

func main() {
	// Print the previous command depending on 'CurrentShell'
	fmt.Println(shells.CurrentShell.GetLastCmd())

	// test cd
	var cdRaw = "cd .."
	var cdCmd = parse.ParseCommand(cdRaw)

	undoCd, err := match.GetUndoCommand(cdCmd)
	if err != nil {
		fmt.Println("donezo")
	}
	fmt.Println(undoCd)

	// test git add -A with extra spaces
	var gitAddAllRaw = "git    add    -A"
	var gitAddAllCmd = parse.ParseCommand(gitAddAllRaw)

	undoGitAddAll, err := match.GetUndoCommand(gitAddAllCmd)
	if err != nil {
		fmt.Println("donezo")
	}
	fmt.Println(undoGitAddAll)
}
