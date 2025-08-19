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

	// EnterAttribute_list is called when entering the attribute_list production.
	EnterAttribute_list(c *Attribute_listContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitAttribute_list is called when exiting the attribute_list production.
	ExitAttribute_list(c *Attribute_listContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)
}
