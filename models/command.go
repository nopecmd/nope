package models

type Command struct {
	RawCommandString string
	BaseCommand      string
	Tokens           []string
}
