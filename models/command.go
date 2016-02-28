package models

type Command struct {
	RawCommandString string
	BaseCommand      string
	Tokens           []string
	Flags            map[string][]string
}
