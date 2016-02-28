package shells

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"github.com/nopecmd/nope/models"
	"github.com/rogpeppe/rog-go/reverse"
	"os"
)

var CurrentShell models.Shell

const (
	bashHistoryPath = "/.bash_history"
	fishHistoryPath = "/.config/fish/fish_history"
)

func init() {
	CurrentShell = getShellFromProcess()
}

func getShellFromProcess() models.Shell {
	var getLastLine = map[string]func() string{
		"bash": getLastLineBash,
		"fish": getLastLineFish,
	}

	proc, err := ps.FindProcess(os.Getppid())

	if err != nil {
		panic(err)
	}

	var procName string = proc.Executable()

	return models.Shell{Name: procName, GetLastCmd: getLastLine[procName]}
}

func getFileScannerFromPath(path string) (*reverse.Scanner, *os.File) {
	file, err := os.Open(os.Getenv("HOME") + path)

	if err != nil {
		fmt.Println("If you could turn on your shell history, that'd be awesome! Thanks! :)")
		panic(err)
	}
	return reverse.NewScanner(file), file
}

func getLastLineBash() string {
	scanner, file := getFileScannerFromPath(bashHistoryPath)

	scanner.Scan()
	file.Close()
	return scanner.Text()
}

func getLastLineFish() string {
	scanner, file := getFileScannerFromPath(fishHistoryPath)

	// Easier way of parsing the fish yml
	var cmdCount = 0
	for cmdCount < 2 {
		scanner.Scan()
		if scanner.Text()[:6] == "- cmd:" {
			cmdCount++
		}
	}

	file.Close()
	return scanner.Text()[7:]
}
