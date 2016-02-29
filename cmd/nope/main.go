package main

import (
	"fmt"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/parse"
	_ "github.com/nopecmd/nope/rules"
	"os"
	"strings"
)

func main() {
	// Runs the command supplied by the argument passed to nope
	var rawCmd = "cd"
	if len(os.Args) > 1 {
		rawCmd = strings.Join(os.Args[1:], " ")
	}
	cmd, err := parse.ParseCommand(rawCmd)
	if err != nil {
		panic(err)
	}

	undo, err := match.GetUndoCommand(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(undo)
}
