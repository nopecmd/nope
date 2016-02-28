package main

import (
	"fmt"
	_ "github.com/nopecmd/nope/rules"
	"github.com/nopecmd/nope/shells"
)

func main() {
	// Print the previous command depending on 'CurrentShell'
	fmt.Println(shells.CurrentShell.GetLastCmd())
}
