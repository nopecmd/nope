package main

import (
	"fmt"
	"github.com/rogpeppe/rog-go/reverse"
	"os"

	_ "github.com/nopecmd/nope/commands"
	_ "github.com/nopecmd/nope/match"
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
	// Do nothing for now
}
