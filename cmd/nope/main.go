package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)

	var x = ""
	for scanner.Scan() {
		x = scanner.Text()
	}
	return x
}

func main() {
	fmt.Println(getLastLine(HISTFILE))
}
