package model

type Character struct {
	Name   string
	Script string
	Task   string
}

func NewCharacter(name, script, target string) *Character {
	return &Character{
		Name:   name,
		Script: script,
		Task:   target,
	}
}
