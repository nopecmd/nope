package models

type Rule struct {
	IsMatch func(string) bool
	GetUndo func(string) string
}
