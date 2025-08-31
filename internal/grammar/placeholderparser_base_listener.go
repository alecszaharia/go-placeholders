// Code generated from /Users/alecszaharia/GolandProjects/go-placeholders/internal/grammar/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

// EnterAttribute is called when production attribute is entered.
func (s *BasePlaceholderParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BasePlaceholderParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterSimplePlaceholder is called when production simplePlaceholder is entered.
func (s *BasePlaceholderParserListener) EnterSimplePlaceholder(ctx *SimplePlaceholderContext) {}

// ExitSimplePlaceholder is called when production simplePlaceholder is exited.
func (s *BasePlaceholderParserListener) ExitSimplePlaceholder(ctx *SimplePlaceholderContext) {}

// EnterBlockPlaceholder is called when production blockPlaceholder is entered.
func (s *BasePlaceholderParserListener) EnterBlockPlaceholder(ctx *BlockPlaceholderContext) {}

// ExitBlockPlaceholder is called when production blockPlaceholder is exited.
func (s *BasePlaceholderParserListener) ExitBlockPlaceholder(ctx *BlockPlaceholderContext) {}

// EnterBlockPlaceholderStart is called when production blockPlaceholderStart is entered.
func (s *BasePlaceholderParserListener) EnterBlockPlaceholderStart(ctx *BlockPlaceholderStartContext) {
}

// ExitBlockPlaceholderStart is called when production blockPlaceholderStart is exited.
func (s *BasePlaceholderParserListener) ExitBlockPlaceholderStart(ctx *BlockPlaceholderStartContext) {
}

// EnterBlockPlaceholderEnd is called when production blockPlaceholderEnd is entered.
func (s *BasePlaceholderParserListener) EnterBlockPlaceholderEnd(ctx *BlockPlaceholderEndContext) {}

// ExitBlockPlaceholderEnd is called when production blockPlaceholderEnd is exited.
func (s *BasePlaceholderParserListener) ExitBlockPlaceholderEnd(ctx *BlockPlaceholderEndContext) {}

// EnterBlockPlaceholderContent is called when production blockPlaceholderContent is entered.
func (s *BasePlaceholderParserListener) EnterBlockPlaceholderContent(ctx *BlockPlaceholderContentContext) {
}

// ExitBlockPlaceholderContent is called when production blockPlaceholderContent is exited.
func (s *BasePlaceholderParserListener) ExitBlockPlaceholderContent(ctx *BlockPlaceholderContentContext) {
}

// EnterPlaceholder is called when production placeholder is entered.
func (s *BasePlaceholderParserListener) EnterPlaceholder(ctx *PlaceholderContext) {}

// ExitPlaceholder is called when production placeholder is exited.
func (s *BasePlaceholderParserListener) ExitPlaceholder(ctx *PlaceholderContext) {}
