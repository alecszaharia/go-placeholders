package placeholder

import "errors"

type Group struct {
	placeholders map[string]*Placeholder
}

func New() *Group {
	return &Group{
		placeholders: make(map[string]*Placeholder, 5),
	}
}

func (g *Group) AddPlaceholder(placeholder *Placeholder) error {
	if _, exists := g.placeholders[placeholder.Name]; exists {
		return errors.New("placeholder with name " + placeholder.Name + " already exists in group " + placeholder.Group)
	}

	g.placeholders[placeholder.Name] = placeholder

	return nil
}

func (g *Group) GetPlaceholderByName(name string) *Placeholder {
	return g.placeholders[name]
}
