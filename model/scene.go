package model

type Scene struct {
	Name        string
	Description string
}

func NewScene(name, description string) *Scene {
	return &Scene{
		Name:        name,
		Description: description,
	}
}
