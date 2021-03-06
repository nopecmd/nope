package shells

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/nopecmd/nope/models"
	"github.com/rogpeppe/rog-go/reverse"
)

var CurrentShell models.Shell

const (
	bashHistoryPath = "/.bash_history"
	fishHistoryPath = "/.config/fish/fish_history"
)

var shellNames = map[string]models.Shell{
	"bash": models.Shell{Name: "bash", GetLastCmd: getLastLineBash, Delimiter: " && "},
	"fish": models.Shell{Name: "fish", GetLastCmd: getLastLineFish, Delimiter: ";"},
}

func init() {
	CurrentShell = getShellFromProcess()
}

func getShellFromProcess() models.Shell {
	proc, err := ps.FindProcess(os.Getppid())
	if err != nil {
		panic(err)
	}

	return shellNames[proc.Executable()]
}

// used for testing purposes
func SetShell(name string) {
	CurrentShell = shellNames[name]
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

func ConcatCommands(commands []string) string {
	return strings.Join(commands, fmt.Sprintf(" %s ", CurrentShell.Delimiter))
}
