// Code generated from /home/alex/GolandProjects/go-placeholders/internal/grammar/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PlaceholderParser struct {
	*antlr.BaseParser
}

var PlaceholderParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func placeholderparserParserInit() {
	staticData := &PlaceholderParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "'='", "", "", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "BLOCK_OPEN", "BLOCK_END", "OPEN", "TEXT", "ID", "INS_WS", "EQ",
		"CLOSE", "STRING", "LBRACE_AS_TEXT",
	}
	staticData.RuleNames = []string{
		"template", "content", "text", "attribute", "simplePlaceholder", "blockPlaceholder",
		"blockPlaceholderStart", "blockPlaceholderEnd", "blockPlaceholderContent",
		"placeholder",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1, 29, 9, 1, 1, 2, 1, 2,
		1, 3, 4, 3, 34, 8, 3, 11, 3, 12, 3, 35, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 5, 4, 45, 8, 4, 10, 4, 12, 4, 48, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 59, 8, 6, 10, 6, 12, 6, 62, 9, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 3, 9, 74, 8, 9,
		1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 0, 71, 0, 20, 1,
		0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 30, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0, 8, 41,
		1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 55, 1, 0, 0, 0, 14, 65, 1, 0, 0, 0,
		16, 69, 1, 0, 0, 0, 18, 73, 1, 0, 0, 0, 20, 21, 3, 2, 1, 0, 21, 22, 5,
		0, 0, 1, 22, 1, 1, 0, 0, 0, 23, 26, 3, 4, 2, 0, 24, 26, 3, 18, 9, 0, 25,
		23, 1, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0,
		0, 27, 28, 1, 0, 0, 0, 28, 3, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 5,
		4, 0, 0, 31, 5, 1, 0, 0, 0, 32, 34, 5, 6, 0, 0, 33, 32, 1, 0, 0, 0, 34,
		35, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 1, 0, 0,
		0, 37, 38, 5, 5, 0, 0, 38, 39, 5, 7, 0, 0, 39, 40, 5, 9, 0, 0, 40, 7, 1,
		0, 0, 0, 41, 42, 5, 3, 0, 0, 42, 46, 5, 5, 0, 0, 43, 45, 3, 6, 3, 0, 44,
		43, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0,
		0, 47, 49, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 50, 5, 8, 0, 0, 50, 9, 1,
		0, 0, 0, 51, 52, 3, 12, 6, 0, 52, 53, 3, 16, 8, 0, 53, 54, 3, 14, 7, 0,
		54, 11, 1, 0, 0, 0, 55, 56, 5, 1, 0, 0, 56, 60, 5, 5, 0, 0, 57, 59, 3,
		6, 3, 0, 58, 57, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60,
		61, 1, 0, 0, 0, 61, 63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 64, 5, 8, 0,
		0, 64, 13, 1, 0, 0, 0, 65, 66, 5, 2, 0, 0, 66, 67, 5, 5, 0, 0, 67, 68,
		5, 8, 0, 0, 68, 15, 1, 0, 0, 0, 69, 70, 3, 2, 1, 0, 70, 17, 1, 0, 0, 0,
		71, 74, 3, 8, 4, 0, 72, 74, 3, 10, 5, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1,
		0, 0, 0, 74, 19, 1, 0, 0, 0, 6, 25, 27, 35, 46, 60, 73,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PlaceholderParserInit initializes any static state used to implement PlaceholderParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPlaceholderParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PlaceholderParserInit() {
	staticData := &PlaceholderParserParserStaticData
	staticData.once.Do(placeholderparserParserInit)
}

// NewPlaceholderParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPlaceholderParser(input antlr.TokenStream) *PlaceholderParser {
	PlaceholderParserInit()
	this := new(PlaceholderParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PlaceholderParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PlaceholderParser.g4"

	return this
}

// PlaceholderParser tokens.
const (
	PlaceholderParserEOF            = antlr.TokenEOF
	PlaceholderParserBLOCK_OPEN     = 1
	PlaceholderParserBLOCK_END      = 2
	PlaceholderParserOPEN           = 3
	PlaceholderParserTEXT           = 4
	PlaceholderParserID             = 5
	PlaceholderParserINS_WS         = 6
	PlaceholderParserEQ             = 7
	PlaceholderParserCLOSE          = 8
	PlaceholderParserSTRING         = 9
	PlaceholderParserLBRACE_AS_TEXT = 10
)

// PlaceholderParser rules.
const (
	PlaceholderParserRULE_template                = 0
	PlaceholderParserRULE_content                 = 1
	PlaceholderParserRULE_text                    = 2
	PlaceholderParserRULE_attribute               = 3
	PlaceholderParserRULE_simplePlaceholder       = 4
	PlaceholderParserRULE_blockPlaceholder        = 5
	PlaceholderParserRULE_blockPlaceholderStart   = 6
	PlaceholderParserRULE_blockPlaceholderEnd     = 7
	PlaceholderParserRULE_blockPlaceholderContent = 8
	PlaceholderParserRULE_placeholder             = 9
)

// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Content() IContentContext
	EOF() antlr.TerminalNode

	// IsTemplateContext differentiates from other interfaces.
	IsTemplateContext()
}

type TemplateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateContext() *TemplateContext {
	var p = new(TemplateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_template
	return p
}

func InitEmptyTemplateContext(p *TemplateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_template
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) Content() IContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *TemplateContext) EOF() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserEOF, 0)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Template() (localctx ITemplateContext) {
	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PlaceholderParserRULE_template)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Content()
	}
	{
		p.SetState(21)
		p.Match(PlaceholderParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllText() []ITextContext
	Text(i int) ITextContext
	AllPlaceholder() []IPlaceholderContext
	Placeholder(i int) IPlaceholderContext

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_content
	return p
}

func InitEmptyContentContext(p *ContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_content
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) AllText() []ITextContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITextContext); ok {
			len++
		}
	}

	tst := make([]ITextContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITextContext); ok {
			tst[i] = t.(ITextContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Text(i int) ITextContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITextContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ContentContext) AllPlaceholder() []IPlaceholderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPlaceholderContext); ok {
			len++
		}
	}

	tst := make([]IPlaceholderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPlaceholderContext); ok {
			tst[i] = t.(IPlaceholderContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Placeholder(i int) IPlaceholderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlaceholderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlaceholderContext)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitContent(s)
	}
}

func (s *ContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PlaceholderParserRULE_content)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26) != 0 {
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PlaceholderParserTEXT:
			{
				p.SetState(23)
				p.Text()
			}

		case PlaceholderParserBLOCK_OPEN, PlaceholderParserOPEN:
			{
				p.SetState(24)
				p.Placeholder()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT() antlr.TerminalNode

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_text
	return p
}

func InitEmptyTextContext(p *TextContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_text
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserTEXT, 0)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitText(s)
	}
}

func (s *TextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitText(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PlaceholderParserRULE_text)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(PlaceholderParserTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// Getter signatures
	EQ() antlr.TerminalNode
	ID() antlr.TerminalNode
	STRING() antlr.TerminalNode
	AllINS_WS() []antlr.TerminalNode
	INS_WS(i int) antlr.TerminalNode

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	value  antlr.Token
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_attribute
	return p
}

func InitEmptyAttributeContext(p *AttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_attribute
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) GetName() antlr.Token { return s.name }

func (s *AttributeContext) GetValue() antlr.Token { return s.value }

func (s *AttributeContext) SetName(v antlr.Token) { s.name = v }

func (s *AttributeContext) SetValue(v antlr.Token) { s.value = v }

func (s *AttributeContext) EQ() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserEQ, 0)
}

func (s *AttributeContext) ID() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserID, 0)
}

func (s *AttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserSTRING, 0)
}

func (s *AttributeContext) AllINS_WS() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserINS_WS)
}

func (s *AttributeContext) INS_WS(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserINS_WS, i)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (s *AttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PlaceholderParserRULE_attribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PlaceholderParserINS_WS {
		{
			p.SetState(32)
			p.Match(PlaceholderParserINS_WS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)

		var _m = p.Match(PlaceholderParserID)

		localctx.(*AttributeContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Match(PlaceholderParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)

		var _m = p.Match(PlaceholderParserSTRING)

		localctx.(*AttributeContext).value = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimplePlaceholderContext is an interface to support dynamic dispatch.
type ISimplePlaceholderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPlaceholderName returns the placeholderName token.
	GetPlaceholderName() antlr.Token

	// SetPlaceholderName sets the placeholderName token.
	SetPlaceholderName(antlr.Token)

	// Getter signatures
	OPEN() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsSimplePlaceholderContext differentiates from other interfaces.
	IsSimplePlaceholderContext()
}

type SimplePlaceholderContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	placeholderName antlr.Token
}

func NewEmptySimplePlaceholderContext() *SimplePlaceholderContext {
	var p = new(SimplePlaceholderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_simplePlaceholder
	return p
}

func InitEmptySimplePlaceholderContext(p *SimplePlaceholderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_simplePlaceholder
}

func (*SimplePlaceholderContext) IsSimplePlaceholderContext() {}

func NewSimplePlaceholderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimplePlaceholderContext {
	var p = new(SimplePlaceholderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_simplePlaceholder

	return p
}

func (s *SimplePlaceholderContext) GetParser() antlr.Parser { return s.parser }

func (s *SimplePlaceholderContext) GetPlaceholderName() antlr.Token { return s.placeholderName }

func (s *SimplePlaceholderContext) SetPlaceholderName(v antlr.Token) { s.placeholderName = v }

func (s *SimplePlaceholderContext) OPEN() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserOPEN, 0)
}

func (s *SimplePlaceholderContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCLOSE, 0)
}

func (s *SimplePlaceholderContext) ID() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserID, 0)
}

func (s *SimplePlaceholderContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *SimplePlaceholderContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *SimplePlaceholderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimplePlaceholderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimplePlaceholderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterSimplePlaceholder(s)
	}
}

func (s *SimplePlaceholderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitSimplePlaceholder(s)
	}
}

func (s *SimplePlaceholderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitSimplePlaceholder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) SimplePlaceholder() (localctx ISimplePlaceholderContext) {
	localctx = NewSimplePlaceholderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PlaceholderParserRULE_simplePlaceholder)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(PlaceholderParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)

		var _m = p.Match(PlaceholderParserID)

		localctx.(*SimplePlaceholderContext).placeholderName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PlaceholderParserINS_WS {
		{
			p.SetState(43)
			p.Attribute()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(PlaceholderParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockPlaceholderContext is an interface to support dynamic dispatch.
type IBlockPlaceholderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BlockPlaceholderStart() IBlockPlaceholderStartContext
	BlockPlaceholderContent() IBlockPlaceholderContentContext
	BlockPlaceholderEnd() IBlockPlaceholderEndContext

	// IsBlockPlaceholderContext differentiates from other interfaces.
	IsBlockPlaceholderContext()
}

type BlockPlaceholderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockPlaceholderContext() *BlockPlaceholderContext {
	var p = new(BlockPlaceholderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholder
	return p
}

func InitEmptyBlockPlaceholderContext(p *BlockPlaceholderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholder
}

func (*BlockPlaceholderContext) IsBlockPlaceholderContext() {}

func NewBlockPlaceholderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockPlaceholderContext {
	var p = new(BlockPlaceholderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholder

	return p
}

func (s *BlockPlaceholderContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockPlaceholderContext) BlockPlaceholderStart() IBlockPlaceholderStartContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockPlaceholderStartContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockPlaceholderStartContext)
}

func (s *BlockPlaceholderContext) BlockPlaceholderContent() IBlockPlaceholderContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockPlaceholderContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockPlaceholderContentContext)
}

func (s *BlockPlaceholderContext) BlockPlaceholderEnd() IBlockPlaceholderEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockPlaceholderEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockPlaceholderEndContext)
}

func (s *BlockPlaceholderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockPlaceholderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockPlaceholderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterBlockPlaceholder(s)
	}
}

func (s *BlockPlaceholderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitBlockPlaceholder(s)
	}
}

func (s *BlockPlaceholderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitBlockPlaceholder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) BlockPlaceholder() (localctx IBlockPlaceholderContext) {
	localctx = NewBlockPlaceholderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PlaceholderParserRULE_blockPlaceholder)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.BlockPlaceholderStart()
	}
	{
		p.SetState(52)
		p.BlockPlaceholderContent()
	}
	{
		p.SetState(53)
		p.BlockPlaceholderEnd()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockPlaceholderStartContext is an interface to support dynamic dispatch.
type IBlockPlaceholderStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPlaceholderName returns the placeholderName token.
	GetPlaceholderName() antlr.Token

	// SetPlaceholderName sets the placeholderName token.
	SetPlaceholderName(antlr.Token)

	// Getter signatures
	BLOCK_OPEN() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

	// IsBlockPlaceholderStartContext differentiates from other interfaces.
	IsBlockPlaceholderStartContext()
}

type BlockPlaceholderStartContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	placeholderName antlr.Token
}

func NewEmptyBlockPlaceholderStartContext() *BlockPlaceholderStartContext {
	var p = new(BlockPlaceholderStartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderStart
	return p
}

func InitEmptyBlockPlaceholderStartContext(p *BlockPlaceholderStartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderStart
}

func (*BlockPlaceholderStartContext) IsBlockPlaceholderStartContext() {}

func NewBlockPlaceholderStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockPlaceholderStartContext {
	var p = new(BlockPlaceholderStartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderStart

	return p
}

func (s *BlockPlaceholderStartContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockPlaceholderStartContext) GetPlaceholderName() antlr.Token { return s.placeholderName }

func (s *BlockPlaceholderStartContext) SetPlaceholderName(v antlr.Token) { s.placeholderName = v }

func (s *BlockPlaceholderStartContext) BLOCK_OPEN() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserBLOCK_OPEN, 0)
}

func (s *BlockPlaceholderStartContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCLOSE, 0)
}

func (s *BlockPlaceholderStartContext) ID() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserID, 0)
}

func (s *BlockPlaceholderStartContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *BlockPlaceholderStartContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *BlockPlaceholderStartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockPlaceholderStartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockPlaceholderStartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterBlockPlaceholderStart(s)
	}
}

func (s *BlockPlaceholderStartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitBlockPlaceholderStart(s)
	}
}

func (s *BlockPlaceholderStartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitBlockPlaceholderStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) BlockPlaceholderStart() (localctx IBlockPlaceholderStartContext) {
	localctx = NewBlockPlaceholderStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PlaceholderParserRULE_blockPlaceholderStart)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.Match(PlaceholderParserBLOCK_OPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)

		var _m = p.Match(PlaceholderParserID)

		localctx.(*BlockPlaceholderStartContext).placeholderName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PlaceholderParserINS_WS {
		{
			p.SetState(57)
			p.Attribute()
		}

		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(PlaceholderParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockPlaceholderEndContext is an interface to support dynamic dispatch.
type IBlockPlaceholderEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPlaceholderName returns the placeholderName token.
	GetPlaceholderName() antlr.Token

	// SetPlaceholderName sets the placeholderName token.
	SetPlaceholderName(antlr.Token)

	// Getter signatures
	BLOCK_END() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsBlockPlaceholderEndContext differentiates from other interfaces.
	IsBlockPlaceholderEndContext()
}

type BlockPlaceholderEndContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	placeholderName antlr.Token
}

func NewEmptyBlockPlaceholderEndContext() *BlockPlaceholderEndContext {
	var p = new(BlockPlaceholderEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderEnd
	return p
}

func InitEmptyBlockPlaceholderEndContext(p *BlockPlaceholderEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderEnd
}

func (*BlockPlaceholderEndContext) IsBlockPlaceholderEndContext() {}

func NewBlockPlaceholderEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockPlaceholderEndContext {
	var p = new(BlockPlaceholderEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderEnd

	return p
}

func (s *BlockPlaceholderEndContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockPlaceholderEndContext) GetPlaceholderName() antlr.Token { return s.placeholderName }

func (s *BlockPlaceholderEndContext) SetPlaceholderName(v antlr.Token) { s.placeholderName = v }

func (s *BlockPlaceholderEndContext) BLOCK_END() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserBLOCK_END, 0)
}

func (s *BlockPlaceholderEndContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCLOSE, 0)
}

func (s *BlockPlaceholderEndContext) ID() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserID, 0)
}

func (s *BlockPlaceholderEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockPlaceholderEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockPlaceholderEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterBlockPlaceholderEnd(s)
	}
}

func (s *BlockPlaceholderEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitBlockPlaceholderEnd(s)
	}
}

func (s *BlockPlaceholderEndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitBlockPlaceholderEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) BlockPlaceholderEnd() (localctx IBlockPlaceholderEndContext) {
	localctx = NewBlockPlaceholderEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PlaceholderParserRULE_blockPlaceholderEnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(PlaceholderParserBLOCK_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)

		var _m = p.Match(PlaceholderParserID)

		localctx.(*BlockPlaceholderEndContext).placeholderName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(PlaceholderParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockPlaceholderContentContext is an interface to support dynamic dispatch.
type IBlockPlaceholderContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Content() IContentContext

	// IsBlockPlaceholderContentContext differentiates from other interfaces.
	IsBlockPlaceholderContentContext()
}

type BlockPlaceholderContentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockPlaceholderContentContext() *BlockPlaceholderContentContext {
	var p = new(BlockPlaceholderContentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderContent
	return p
}

func InitEmptyBlockPlaceholderContentContext(p *BlockPlaceholderContentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderContent
}

func (*BlockPlaceholderContentContext) IsBlockPlaceholderContentContext() {}

func NewBlockPlaceholderContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockPlaceholderContentContext {
	var p = new(BlockPlaceholderContentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_blockPlaceholderContent

	return p
}

func (s *BlockPlaceholderContentContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockPlaceholderContentContext) Content() IContentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *BlockPlaceholderContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockPlaceholderContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockPlaceholderContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterBlockPlaceholderContent(s)
	}
}

func (s *BlockPlaceholderContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitBlockPlaceholderContent(s)
	}
}

func (s *BlockPlaceholderContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitBlockPlaceholderContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) BlockPlaceholderContent() (localctx IBlockPlaceholderContentContext) {
	localctx = NewBlockPlaceholderContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PlaceholderParserRULE_blockPlaceholderContent)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Content()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPlaceholderContext is an interface to support dynamic dispatch.
type IPlaceholderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimplePlaceholder() ISimplePlaceholderContext
	BlockPlaceholder() IBlockPlaceholderContext

	// IsPlaceholderContext differentiates from other interfaces.
	IsPlaceholderContext()
}

type PlaceholderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlaceholderContext() *PlaceholderContext {
	var p = new(PlaceholderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_placeholder
	return p
}

func InitEmptyPlaceholderContext(p *PlaceholderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_placeholder
}

func (*PlaceholderContext) IsPlaceholderContext() {}

func NewPlaceholderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlaceholderContext {
	var p = new(PlaceholderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_placeholder

	return p
}

func (s *PlaceholderContext) GetParser() antlr.Parser { return s.parser }

func (s *PlaceholderContext) SimplePlaceholder() ISimplePlaceholderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplePlaceholderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimplePlaceholderContext)
}

func (s *PlaceholderContext) BlockPlaceholder() IBlockPlaceholderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockPlaceholderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockPlaceholderContext)
}

func (s *PlaceholderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlaceholderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlaceholderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterPlaceholder(s)
	}
}

func (s *PlaceholderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitPlaceholder(s)
	}
}

func (s *PlaceholderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitPlaceholder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Placeholder() (localctx IPlaceholderContext) {
	localctx = NewPlaceholderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PlaceholderParserRULE_placeholder)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PlaceholderParserOPEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.SimplePlaceholder()
		}

	case PlaceholderParserBLOCK_OPEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(72)
			p.BlockPlaceholder()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
