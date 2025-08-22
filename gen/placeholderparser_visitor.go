// Code generated from /home/alex/GolandProjects/go-placeholders/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PlaceholderParser.
type PlaceholderParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PlaceholderParser#template.
	VisitTemplate(ctx *TemplateContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#content.
	VisitContent(ctx *ContentContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#simplePlaceholder.
	VisitSimplePlaceholder(ctx *SimplePlaceholderContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#blockPlaceholder.
	VisitBlockPlaceholder(ctx *BlockPlaceholderContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#blockPlaceholderStart.
	VisitBlockPlaceholderStart(ctx *BlockPlaceholderStartContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#blockPlaceholderEnd.
	VisitBlockPlaceholderEnd(ctx *BlockPlaceholderEndContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#blockPlaceholderContent.
	VisitBlockPlaceholderContent(ctx *BlockPlaceholderContentContext) interface{}

	// Visit a parse tree produced by PlaceholderParser#placeholder.
	VisitPlaceholder(ctx *PlaceholderContext) interface{}
}
