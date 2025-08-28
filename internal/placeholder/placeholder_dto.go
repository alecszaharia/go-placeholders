package placeholder

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type PlaceholderJSON interface {
	ToDTO() *PlaceholderDTO
}

type PlaceholderDTO struct {
	Id          string            `json:"id,omitempty"`
	Label       string            `json:"label"`
	Group       string            `json:"group"` // this is the group name used in editor.
	Display     string            `json:"display"`
	Attr        map[string]string `json:"attr"`
	VaryAttr    []string          `json:"varyAttr"`
	Placeholder string            `json:"placeholder"`
}

func (p *Placeholder) attrToParams() string {
	params := ""
	for k, v := range p.Attrs {
		v = strings.ReplaceAll(v, `"`, `\"`)
		params += fmt.Sprintf(` %s="%s"`, k, v)
	}
	return params
}

func (p *Placeholder) ToDTO() *PlaceholderDTO {
	return &PlaceholderDTO{
		Id:          uuid.New().String(),
		Label:       p.Label,
		Group:       p.Group,
		Display:     DisplayMap[p.Display],
		Attr:        p.Attrs,
		VaryAttr:    p.VaryAttrs,
		Placeholder: fmt.Sprintf("{{ %s%s }}", p.Name, p.attrToParams()),
	}
}
