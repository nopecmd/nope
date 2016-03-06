package rules

import (
	"io/ioutil"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/nopecmd/nope/match"
	"github.com/nopecmd/nope/models"
	"github.com/nopecmd/nope/utils"
)

func isMatchMkdir(cmd models.Command) bool {
	return cmd.BaseCommand == mkdirBaseCommand
}

func findPathToDelete(dirs []string) (string, error) {
	if len(dirs) == 1 {
		return dirs[0], nil
	}

	var allButLast = dirs[:len(dirs)-1]
	var path = strings.Join(allButLast, "/")

	contents, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}

	if len(contents) > 1 {
		return strings.Join(dirs, "/"), nil
	}

	return findPathToDelete(allButLast)
}

func getCorrectPathsToDelete(rawPaths []string) ([]string, error) {
	var refactoredPaths []string
	for _, rawPath := range rawPaths {
		var dirs = strings.Split(rawPath, "/")
		pathToDelete, err := findPathToDelete(dirs)
		if err != nil {
			return nil, err
		}
		refactoredPaths = append(refactoredPaths, pathToDelete)
	}

	return refactoredPaths, nil
}

func getUndoMkdir(cmd models.Command) (string, error) {
	var mkdirFlags struct {
		Mode         string `short:"m"`
		CreateInterm bool   `short:"p"`
		Verbose      bool   `short:"v"`
	}
	filteredTokens, err := flags.ParseArgs(&mkdirFlags, cmd.TokensWithoutBase)
	if err != nil {
		return "", err
	}

	var paths = utils.GetFilePathsFromTokens(filteredTokens)

	if mkdirFlags.CreateInterm {
		paths, err = getCorrectPathsToDelete(paths)
		if err != nil {
			return "", nil
		}
	}

	return rmBaseCommand + " -rf " + strings.Join(paths, " "), nil
}

func init() {
	match.AddRule(isMatchMkdir, getUndoMkdir)
}
