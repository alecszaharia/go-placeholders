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

// EnterText is called when production text is entered.
func (s *BasePlaceholderParserListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BasePlaceholderParserListener) ExitText(ctx *TextContext) {}

// EnterPlaceholder is called when production placeholder is entered.
func (s *BasePlaceholderParserListener) EnterPlaceholder(ctx *PlaceholderContext) {}

// ExitPlaceholder is called when production placeholder is exited.
func (s *BasePlaceholderParserListener) ExitPlaceholder(ctx *PlaceholderContext) {}

// EnterSimple_placeholder is called when production simple_placeholder is entered.
func (s *BasePlaceholderParserListener) EnterSimple_placeholder(ctx *Simple_placeholderContext) {}

// ExitSimple_placeholder is called when production simple_placeholder is exited.
func (s *BasePlaceholderParserListener) ExitSimple_placeholder(ctx *Simple_placeholderContext) {}

// EnterBlock_placeholder is called when production block_placeholder is entered.
func (s *BasePlaceholderParserListener) EnterBlock_placeholder(ctx *Block_placeholderContext) {}

// ExitBlock_placeholder is called when production block_placeholder is exited.
func (s *BasePlaceholderParserListener) ExitBlock_placeholder(ctx *Block_placeholderContext) {}

// EnterAttribute_list is called when production attribute_list is entered.
func (s *BasePlaceholderParserListener) EnterAttribute_list(ctx *Attribute_listContext) {}

// ExitAttribute_list is called when production attribute_list is exited.
func (s *BasePlaceholderParserListener) ExitAttribute_list(ctx *Attribute_listContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BasePlaceholderParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BasePlaceholderParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterValue is called when production value is entered.
func (s *BasePlaceholderParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePlaceholderParserListener) ExitValue(ctx *ValueContext) {}
