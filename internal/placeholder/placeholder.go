package placeholder

import "antrl/internal/model"

type DisplayType int

type HandlerFunc func(ctx *model.Context) (string, error)

var DisplayMap = map[DisplayType]string{DisplayInline: "inline", DisplayBlock: "block"}

const (
	DisplayInline DisplayType = 1
	DisplayBlock  DisplayType = 2
)

type Placeholder struct {
	Name            string            // this is the unique name of the placeholder. This will be used in the template like {{ Name }}
	Label           string            // this is the label shown in the editor.
	Group           string            // this is the group name used in editor.
	Display         DisplayType       // this is the display type of the placeholder. It can be inline or block.
	Attrs           map[string]string // this is the default attributes of the placeholder.
	VaryAttrs       []string          // this is the attributes that can be varied by the user.
	Handler         HandlerFunc       // this is the handler function that will be called when the placeholder is rendered.
	FallbackHandler HandlerFunc       // this is the fallback handler function that will be called when the main handler fails.
}
