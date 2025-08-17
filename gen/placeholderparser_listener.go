// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import "github.com/antlr4-go/antlr/v4"

// PlaceholderParserListener is a complete listener for a parse tree produced by PlaceholderParser.
type PlaceholderParserListener interface {
	antlr.ParseTreeListener

	// EnterTemplate is called when entering the template production.
	EnterTemplate(c *TemplateContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterPlaceholder is called when entering the placeholder production.
	EnterPlaceholder(c *PlaceholderContext)

	// EnterSimple_placeholder is called when entering the simple_placeholder production.
	EnterSimple_placeholder(c *Simple_placeholderContext)

	// EnterBlock_placeholder is called when entering the block_placeholder production.
	EnterBlock_placeholder(c *Block_placeholderContext)

	// EnterAttribute_list is called when entering the attribute_list production.
	EnterAttribute_list(c *Attribute_listContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitPlaceholder is called when exiting the placeholder production.
	ExitPlaceholder(c *PlaceholderContext)

	// ExitSimple_placeholder is called when exiting the simple_placeholder production.
	ExitSimple_placeholder(c *Simple_placeholderContext)

	// ExitBlock_placeholder is called when exiting the block_placeholder production.
	ExitBlock_placeholder(c *Block_placeholderContext)

	// ExitAttribute_list is called when exiting the attribute_list production.
	ExitAttribute_list(c *Attribute_listContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
