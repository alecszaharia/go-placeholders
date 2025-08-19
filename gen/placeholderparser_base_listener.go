// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import "github.com/antlr4-go/antlr/v4"

// BasePlaceholderParserListener is a complete listener for a parse tree produced by PlaceholderParser.
type BasePlaceholderParserListener struct{}

var _ PlaceholderParserListener = &BasePlaceholderParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePlaceholderParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePlaceholderParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePlaceholderParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePlaceholderParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTemplate is called when production template is entered.
func (s *BasePlaceholderParserListener) EnterTemplate(ctx *TemplateContext) {}

// ExitTemplate is called when production template is exited.
func (s *BasePlaceholderParserListener) ExitTemplate(ctx *TemplateContext) {}

// EnterContent is called when production content is entered.
func (s *BasePlaceholderParserListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BasePlaceholderParserListener) ExitContent(ctx *ContentContext) {}

// EnterText is called when production text is entered.
func (s *BasePlaceholderParserListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BasePlaceholderParserListener) ExitText(ctx *TextContext) {}

// EnterAttributeList is called when production attributeList is entered.
func (s *BasePlaceholderParserListener) EnterAttributeList(ctx *AttributeListContext) {}

// ExitAttributeList is called when production attributeList is exited.
func (s *BasePlaceholderParserListener) ExitAttributeList(ctx *AttributeListContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BasePlaceholderParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BasePlaceholderParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterAttrValue is called when production attrValue is entered.
func (s *BasePlaceholderParserListener) EnterAttrValue(ctx *AttrValueContext) {}

// ExitAttrValue is called when production attrValue is exited.
func (s *BasePlaceholderParserListener) ExitAttrValue(ctx *AttrValueContext) {}

// EnterBlock is called when production block is entered.
func (s *BasePlaceholderParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasePlaceholderParserListener) ExitBlock(ctx *BlockContext) {}
