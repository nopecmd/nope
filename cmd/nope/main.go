package main

import (
	"fmt"
	"github.com/rogpeppe/rog-go/reverse"
	"os"

	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
)

// Don't know how to make this a constant
var HISTFILE = os.Getenv("HOME") + "/.bash_history"

func getLastLine(fname string) string {
	file, err := os.Open(fname)

	if err != nil {
		fmt.Println("Can't open bash history file!")
		panic(err)
	}
	defer file.Close()

	var scanner = reverse.NewScanner(file)
	scanner.Scan()

	return scanner.Text()
}

func main() {
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
