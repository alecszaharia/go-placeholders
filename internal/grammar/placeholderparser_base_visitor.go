// Code generated from /Users/alecszaharia/GolandProjects/go-placeholders/internal/grammar/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import "github.com/antlr4-go/antlr/v4"

type BasePlaceholderParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePlaceholderParserVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitContent(ctx *ContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitSimplePlaceholder(ctx *SimplePlaceholderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitBlockPlaceholder(ctx *BlockPlaceholderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitBlockPlaceholderStart(ctx *BlockPlaceholderStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitBlockPlaceholderEnd(ctx *BlockPlaceholderEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitBlockPlaceholderContent(ctx *BlockPlaceholderContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitPlaceholder(ctx *PlaceholderContext) interface{} {
	return v.VisitChildren(ctx)
}
