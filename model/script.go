package model

type Script struct {
	Name        string
	Description string
	Characters  map[string]*Character
	Scenes      map[string]*Scene
}

func NewScript(name, description string, characters map[string]*Character, scenes map[string]*Scene) *Script {
	return &Script{
		Name:        name,
		Description: description,
		Characters:  characters,
		Scenes:      scenes,
	}
}
