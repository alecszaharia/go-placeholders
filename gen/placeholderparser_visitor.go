// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PlaceholderParser.
type PlaceholderParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PlaceholderParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#placeholder.
	VisitPlaceholder(ctx *PlaceholderContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#simple_placeholder.
	VisitSimple_placeholder(ctx *Simple_placeholderContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#block_placeholder.
	VisitBlock_placeholder(ctx *Block_placeholderContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#attribute_list.
	VisitAttribute_list(ctx *Attribute_listContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
