// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

	// EnterAttributeList is called when entering the attributeList production.
	EnterAttributeList(c *AttributeListContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterAttrValue is called when entering the attrValue production.
	EnterAttrValue(c *AttrValueContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// ExitTemplate is called when exiting the template production.
	ExitTemplate(c *TemplateContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitAttributeList is called when exiting the attributeList production.
	ExitAttributeList(c *AttributeListContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitAttrValue is called when exiting the attrValue production.
	ExitAttrValue(c *AttrValueContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)
}
