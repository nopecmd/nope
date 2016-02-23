package main

import (
	"fmt"

	"github.com/nopecmd/nope/match"

	_ "github.com/nopecmd/nope/commands"
)

func main() {
	fmt.Println(match.Cmds[0].GetUndo("cd weqwe"))
	fmt.Println(match.Cmds[1].GetUndo("cd weqwe"))
}
