// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "", "'{{'", "'}}'", "'end_'", "','", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "OPEN", "CLOSE", "END", "COMMA", "EQUALS", "IDENT", "INT",
		"DQSTRING", "SQSTRING", "WSIN",
	}
	staticData.RuleNames = []string{
		"template", "text", "placeholder", "simple_placeholder", "block_placeholder",
		"attribute_list", "attribute", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 11, 66, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 5, 0, 19, 8, 0, 10,
		0, 12, 0, 22, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2, 30, 8, 2,
		1, 3, 1, 3, 1, 3, 3, 3, 35, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 42,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		5, 5, 55, 8, 5, 10, 5, 12, 5, 58, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 1, 0, 8, 10, 63, 0,
		20, 1, 0, 0, 0, 2, 25, 1, 0, 0, 0, 4, 29, 1, 0, 0, 0, 6, 31, 1, 0, 0, 0,
		8, 38, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12, 59, 1, 0, 0, 0, 14, 63, 1, 0,
		0, 0, 16, 19, 3, 2, 1, 0, 17, 19, 3, 4, 2, 0, 18, 16, 1, 0, 0, 0, 18, 17,
		1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0, 0,
		21, 23, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 24, 5, 0, 0, 1, 24, 1, 1, 0,
		0, 0, 25, 26, 5, 1, 0, 0, 26, 3, 1, 0, 0, 0, 27, 30, 3, 6, 3, 0, 28, 30,
		3, 8, 4, 0, 29, 27, 1, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 5, 1, 0, 0, 0,
		31, 32, 5, 2, 0, 0, 32, 34, 5, 7, 0, 0, 33, 35, 3, 10, 5, 0, 34, 33, 1,
		0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 5, 3, 0, 0, 37,
		7, 1, 0, 0, 0, 38, 39, 5, 2, 0, 0, 39, 41, 5, 7, 0, 0, 40, 42, 3, 10, 5,
		0, 41, 40, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44,
		5, 3, 0, 0, 44, 45, 3, 0, 0, 0, 45, 46, 5, 2, 0, 0, 46, 47, 5, 4, 0, 0,
		47, 48, 5, 7, 0, 0, 48, 49, 5, 3, 0, 0, 49, 50, 4, 4, 0, 1, 50, 9, 1, 0,
		0, 0, 51, 56, 3, 12, 6, 0, 52, 53, 5, 5, 0, 0, 53, 55, 3, 12, 6, 0, 54,
		52, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0,
		0, 57, 11, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 60, 5, 7, 0, 0, 60, 61,
		5, 6, 0, 0, 61, 62, 3, 14, 7, 0, 62, 13, 1, 0, 0, 0, 63, 64, 7, 0, 0, 0,
		64, 15, 1, 0, 0, 0, 6, 18, 20, 29, 34, 41, 56,
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
	PlaceholderParserEOF      = antlr.TokenEOF
	PlaceholderParserTEXT     = 1
	PlaceholderParserOPEN     = 2
	PlaceholderParserCLOSE    = 3
	PlaceholderParserEND      = 4
	PlaceholderParserCOMMA    = 5
	PlaceholderParserEQUALS   = 6
	PlaceholderParserIDENT    = 7
	PlaceholderParserINT      = 8
	PlaceholderParserDQSTRING = 9
	PlaceholderParserSQSTRING = 10
	PlaceholderParserWSIN     = 11
)

// PlaceholderParser rules.
const (
	PlaceholderParserRULE_template           = 0
	PlaceholderParserRULE_text               = 1
	PlaceholderParserRULE_placeholder        = 2
	PlaceholderParserRULE_simple_placeholder = 3
	PlaceholderParserRULE_block_placeholder  = 4
	PlaceholderParserRULE_attribute_list     = 5
	PlaceholderParserRULE_attribute          = 6
	PlaceholderParserRULE_value              = 7
)

// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllText() []ITextContext
	Text(i int) ITextContext
	AllPlaceholder() []IPlaceholderContext
	Placeholder(i int) IPlaceholderContext

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

func (s *TemplateContext) EOF() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserEOF, 0)
}

func (s *TemplateContext) AllText() []ITextContext {
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

func (s *TemplateContext) Text(i int) ITextContext {
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

func (s *TemplateContext) AllPlaceholder() []IPlaceholderContext {
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

func (s *TemplateContext) Placeholder(i int) IPlaceholderContext {
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
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PlaceholderParserTEXT || _la == PlaceholderParserOPEN {
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PlaceholderParserTEXT:
			{
				p.SetState(16)
				p.Text()
			}

		case PlaceholderParserOPEN:
			{
				p.SetState(17)
				p.Placeholder()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
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
	p.EnterRule(localctx, 2, PlaceholderParserRULE_text)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
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

// IPlaceholderContext is an interface to support dynamic dispatch.
type IPlaceholderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Simple_placeholder() ISimple_placeholderContext
	Block_placeholder() IBlock_placeholderContext

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

func (s *PlaceholderContext) Simple_placeholder() ISimple_placeholderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_placeholderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_placeholderContext)
}

func (s *PlaceholderContext) Block_placeholder() IBlock_placeholderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlock_placeholderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlock_placeholderContext)
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
	p.EnterRule(localctx, 4, PlaceholderParserRULE_placeholder)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Simple_placeholder()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Block_placeholder()
		}

	case antlr.ATNInvalidAltNumber:
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

// ISimple_placeholderContext is an interface to support dynamic dispatch.
type ISimple_placeholderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	CLOSE() antlr.TerminalNode
	Attribute_list() IAttribute_listContext

	// IsSimple_placeholderContext differentiates from other interfaces.
	IsSimple_placeholderContext()
}

type Simple_placeholderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_placeholderContext() *Simple_placeholderContext {
	var p = new(Simple_placeholderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_simple_placeholder
	return p
}

func InitEmptySimple_placeholderContext(p *Simple_placeholderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_simple_placeholder
}

func (*Simple_placeholderContext) IsSimple_placeholderContext() {}

func NewSimple_placeholderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_placeholderContext {
	var p = new(Simple_placeholderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_simple_placeholder

	return p
}

func (s *Simple_placeholderContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_placeholderContext) OPEN() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserOPEN, 0)
}

func (s *Simple_placeholderContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserIDENT, 0)
}

func (s *Simple_placeholderContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCLOSE, 0)
}

func (s *Simple_placeholderContext) Attribute_list() IAttribute_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttribute_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttribute_listContext)
}

func (s *Simple_placeholderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_placeholderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_placeholderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterSimple_placeholder(s)
	}
}

func (s *Simple_placeholderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitSimple_placeholder(s)
	}
}

func (s *Simple_placeholderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitSimple_placeholder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Simple_placeholder() (localctx ISimple_placeholderContext) {
	localctx = NewSimple_placeholderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PlaceholderParserRULE_simple_placeholder)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(PlaceholderParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(PlaceholderParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PlaceholderParserIDENT {
		{
			p.SetState(33)
			p.Attribute_list()
		}

	}
	{
		p.SetState(36)
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

// IBlock_placeholderContext is an interface to support dynamic dispatch.
type IBlock_placeholderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHead returns the head token.
	GetHead() antlr.Token

	// GetTail returns the tail token.
	GetTail() antlr.Token

	// SetHead sets the head token.
	SetHead(antlr.Token)

	// SetTail sets the tail token.
	SetTail(antlr.Token)

	// Getter signatures
	AllOPEN() []antlr.TerminalNode
	OPEN(i int) antlr.TerminalNode
	AllCLOSE() []antlr.TerminalNode
	CLOSE(i int) antlr.TerminalNode
	Template() ITemplateContext
	END() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	Attribute_list() IAttribute_listContext

	// IsBlock_placeholderContext differentiates from other interfaces.
	IsBlock_placeholderContext()
}

type Block_placeholderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	head   antlr.Token
	tail   antlr.Token
}

func NewEmptyBlock_placeholderContext() *Block_placeholderContext {
	var p = new(Block_placeholderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_block_placeholder
	return p
}

func InitEmptyBlock_placeholderContext(p *Block_placeholderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_block_placeholder
}

func (*Block_placeholderContext) IsBlock_placeholderContext() {}

func NewBlock_placeholderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Block_placeholderContext {
	var p = new(Block_placeholderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_block_placeholder

	return p
}

func (s *Block_placeholderContext) GetParser() antlr.Parser { return s.parser }

func (s *Block_placeholderContext) GetHead() antlr.Token { return s.head }

func (s *Block_placeholderContext) GetTail() antlr.Token { return s.tail }

func (s *Block_placeholderContext) SetHead(v antlr.Token) { s.head = v }

func (s *Block_placeholderContext) SetTail(v antlr.Token) { s.tail = v }

func (s *Block_placeholderContext) AllOPEN() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserOPEN)
}

func (s *Block_placeholderContext) OPEN(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserOPEN, i)
}

func (s *Block_placeholderContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserCLOSE)
}

func (s *Block_placeholderContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCLOSE, i)
}

func (s *Block_placeholderContext) Template() ITemplateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemplateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITemplateContext)
}

func (s *Block_placeholderContext) END() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserEND, 0)
}

func (s *Block_placeholderContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserIDENT)
}

func (s *Block_placeholderContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserIDENT, i)
}

func (s *Block_placeholderContext) Attribute_list() IAttribute_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttribute_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttribute_listContext)
}

func (s *Block_placeholderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Block_placeholderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Block_placeholderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterBlock_placeholder(s)
	}
}

func (s *Block_placeholderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitBlock_placeholder(s)
	}
}

func (s *Block_placeholderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitBlock_placeholder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Block_placeholder() (localctx IBlock_placeholderContext) {
	localctx = NewBlock_placeholderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PlaceholderParserRULE_block_placeholder)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(PlaceholderParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)

		var _m = p.Match(PlaceholderParserIDENT)

		localctx.(*Block_placeholderContext).head = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PlaceholderParserIDENT {
		{
			p.SetState(40)
			p.Attribute_list()
		}

	}
	{
		p.SetState(43)
		p.Match(PlaceholderParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Template()
	}
	{
		p.SetState(45)
		p.Match(PlaceholderParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(PlaceholderParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)

		var _m = p.Match(PlaceholderParserIDENT)

		localctx.(*Block_placeholderContext).tail = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(PlaceholderParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(49)

	if !((func() string {
		if localctx.(*Block_placeholderContext).GetHead() == nil {
			return ""
		} else {
			return localctx.(*Block_placeholderContext).GetHead().GetText()
		}
	}()) == (func() string {
		if localctx.(*Block_placeholderContext).GetTail() == nil {
			return ""
		} else {
			return localctx.(*Block_placeholderContext).GetTail().GetText()
		}
	}())) {
		p.SetError(antlr.NewFailedPredicateException(p, " $head.text == $tail.text ", ""))
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

// IAttribute_listContext is an interface to support dynamic dispatch.
type IAttribute_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAttribute_listContext differentiates from other interfaces.
	IsAttribute_listContext()
}

type Attribute_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribute_listContext() *Attribute_listContext {
	var p = new(Attribute_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_attribute_list
	return p
}

func InitEmptyAttribute_listContext(p *Attribute_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_attribute_list
}

func (*Attribute_listContext) IsAttribute_listContext() {}

func NewAttribute_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attribute_listContext {
	var p = new(Attribute_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_attribute_list

	return p
}

func (s *Attribute_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Attribute_listContext) AllAttribute() []IAttributeContext {
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

func (s *Attribute_listContext) Attribute(i int) IAttributeContext {
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

func (s *Attribute_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserCOMMA)
}

func (s *Attribute_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCOMMA, i)
}

func (s *Attribute_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attribute_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attribute_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterAttribute_list(s)
	}
}

func (s *Attribute_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitAttribute_list(s)
	}
}

func (s *Attribute_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitAttribute_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Attribute_list() (localctx IAttribute_listContext) {
	localctx = NewAttribute_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PlaceholderParserRULE_attribute_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Attribute()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PlaceholderParserCOMMA {
		{
			p.SetState(52)
			p.Match(PlaceholderParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Attribute()
		}

		p.SetState(58)
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

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Value() IValueContext

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *AttributeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserIDENT, 0)
}

func (s *AttributeContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserEQUALS, 0)
}

func (s *AttributeContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
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
	p.EnterRule(localctx, 12, PlaceholderParserRULE_attribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(PlaceholderParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(PlaceholderParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Value()
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	DQSTRING() antlr.TerminalNode
	SQSTRING() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) INT() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserINT, 0)
}

func (s *ValueContext) DQSTRING() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserDQSTRING, 0)
}

func (s *ValueContext) SQSTRING() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserSQSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PlaceholderParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1792) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

func (p *PlaceholderParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *Block_placeholderContext = nil
		if localctx != nil {
			t = localctx.(*Block_placeholderContext)
		}
		return p.Block_placeholder_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PlaceholderParser) Block_placeholder_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return (func() string {
			if localctx.(*Block_placeholderContext).GetHead() == nil {
				return ""
			} else {
				return localctx.(*Block_placeholderContext).GetHead().GetText()
			}
		}()) == (func() string {
			if localctx.(*Block_placeholderContext).GetTail() == nil {
				return ""
			} else {
				return localctx.(*Block_placeholderContext).GetTail().GetText()
			}
		}())

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
