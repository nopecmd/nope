package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func IsValidFilePath(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func GetFilePathsFromTokens(tokens []string) []string {
	var filePaths []string

	for _, token := range tokens {
		if IsValidFilePath(token) {
			filePaths = append(filePaths, token)
		}
	}
	return filePaths
}

func CleanPath(path string) string {
	// TODO: add more stuff to clean, like replacing `~` with `/Users/someone/`
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	return path
}

func WasRenamed(fromTokens []string, toTokens []string) (bool, error) {
	if len(fromTokens) == 0 || len(toTokens) == 0 {
		return false, errors.New("Incorrect paths")
	}

	var target = fromTokens[len(fromTokens)-1]
	var destPath = strings.Join(toTokens, "/")

	// first check if destPath is a directory
	dir, err := os.Stat(destPath)
	if err != nil {
		return false, err
	}

	// if destPath is not a directory, then file was renamed
	if !dir.IsDir() {
		return true, nil
	}

	// otherwise, check if destPath contains target
	contents, err := ioutil.ReadDir(destPath)
	if err != nil {
		return false, err
	}

	var contains = false
	for _, fileOrDir := range contents {
		if fileOrDir.Name() == target {
			contains = true
			break
		}
	}
	if !contains {
		return true, nil
	}

	return false, nil
}
