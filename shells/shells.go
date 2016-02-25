package shells

import (
	"github.com/nopecmd/nope/models"
	"github.com/rogpeppe/rog-go/reverse"
	"os"
)

var CurrentShell models.Shell

func init() {
	CurrentShell = models.Shell{Name: "fish", GetLastCmd: getLastLineFish}
}

func getLastLineBash() string {
	var fname = os.Getenv("HOME") + "/.bash_history"
	file, err := os.Open(fname)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var scanner = reverse.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func getLastLineFish() string {
	var fname = os.Getenv("HOME") + "/.config/fish/fish_history"
	file, err := os.Open(fname)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var scanner = reverse.NewScanner(file)

	// Easier way of parsing the fish yml
	var cmdCount = 0
	for cmdCount < 2 {
		scanner.Scan()
		if scanner.Text()[:6] == "- cmd:" {
			cmdCount++
		}
	}

	return scanner.Text()[7:]
}
