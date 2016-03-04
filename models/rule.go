package models

type Rule struct {
	IsMatch func(Command) bool
	GetUndo func(Command) (string, error)
}
