package models

type Shell struct {
	Name       string
	GetLastCmd func() string
}
