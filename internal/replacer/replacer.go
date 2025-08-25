package replacer

import (
	"antrl/internal/model"

	"github.com/antlr4-go/antlr/v4"
)

func RebuildWithPlaceholderReplace(inputTokens *antlr.CommonTokenStream, nodes []*model.Node) string {
	rewriter := antlr.NewTokenStreamRewriter(inputTokens)

	for _, n := range nodes {
		replaceNode(rewriter, n)
	}

	return ""

}

func replaceNode(rewriter *antlr.TokenStreamRewriter, node *model.Node) {
	switch node.Type {
	case model.NodeText:
		// No replacement needed for text nodes

	case model.NodePlaceholder:

		break
	}
}
