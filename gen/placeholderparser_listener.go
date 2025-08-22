// Code generated from /home/alex/GolandProjects/go-placeholders/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import "github.com/antlr4-go/antlr/v4"

// PlaceholderParserListener is a complete listener for a parse tree produced by PlaceholderParser.
type PlaceholderParserListener interface {
	antlr.ParseTreeListener

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterSimplePlaceholder is called when entering the simplePlaceholder production.
	EnterSimplePlaceholder(c *SimplePlaceholderContext)

	// EnterBlockPlaceholder is called when entering the blockPlaceholder production.
	EnterBlockPlaceholder(c *BlockPlaceholderContext)

	// EnterBlockPlaceholderStart is called when entering the blockPlaceholderStart production.
	EnterBlockPlaceholderStart(c *BlockPlaceholderStartContext)

	// EnterBlockPlaceholderEnd is called when entering the blockPlaceholderEnd production.
	EnterBlockPlaceholderEnd(c *BlockPlaceholderEndContext)

	// EnterBlockPlaceholderContent is called when entering the blockPlaceholderContent production.
	EnterBlockPlaceholderContent(c *BlockPlaceholderContentContext)

	// EnterPlaceholder is called when entering the placeholder production.
	EnterPlaceholder(c *PlaceholderContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitSimplePlaceholder is called when exiting the simplePlaceholder production.
	ExitSimplePlaceholder(c *SimplePlaceholderContext)

	// ExitBlockPlaceholder is called when exiting the blockPlaceholder production.
	ExitBlockPlaceholder(c *BlockPlaceholderContext)

	// ExitBlockPlaceholderStart is called when exiting the blockPlaceholderStart production.
	ExitBlockPlaceholderStart(c *BlockPlaceholderStartContext)

	// ExitBlockPlaceholderEnd is called when exiting the blockPlaceholderEnd production.
	ExitBlockPlaceholderEnd(c *BlockPlaceholderEndContext)

	// ExitBlockPlaceholderContent is called when exiting the blockPlaceholderContent production.
	ExitBlockPlaceholderContent(c *BlockPlaceholderContentContext)

	// ExitPlaceholder is called when exiting the placeholder production.
	ExitPlaceholder(c *PlaceholderContext)
}
