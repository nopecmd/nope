package utils

import (
	"os"
	"testing"
)

func FormatError(rawCmd string, msg string) string {
	return rawCmd + " command failed: " + msg
}

func CreateFile(fileName string, t *testing.T) {
	if _, err := os.Create(fileName); err != nil {
		t.Errorf("could not create file: " + fileName)
	}
}

func RemoveFile(fileName string, t *testing.T) {
	if err := os.Remove(fileName); err != nil {
		t.Errorf("could not remove file: " + fileName)
	}
}

func CreateDir(dirName string, t *testing.T) {
	if err := os.MkdirAll(dirName, 0777); err != nil {
		t.Errorf("could not create dir: " + dirName)
	}
}

func RemoveDir(dirName string, t *testing.T) {
	if err := os.RemoveAll(dirName); err != nil {
		t.Errorf("could not remove direcotry " + dirName)
	}
}
