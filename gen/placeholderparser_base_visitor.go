// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import "github.com/antlr4-go/antlr/v4"

type BasePlaceholderParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePlaceholderParserVisitor) VisitTemplate(ctx *TemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitText(ctx *TextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitPlaceholder(ctx *PlaceholderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitSimple_placeholder(ctx *Simple_placeholderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitBlock_placeholder(ctx *Block_placeholderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitAttribute_list(ctx *Attribute_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePlaceholderParserVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
