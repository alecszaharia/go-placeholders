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
		"", "'{{end_'", "'{{'", "", "", "'}}'",
	}
	staticData.SymbolicNames = []string{
		"", "OPEN_END", "OPEN", "TEXT", "ID", "CLOSE", "STRING",
	}
	staticData.RuleNames = []string{
		"template", "content", "text", "placeholder",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 6, 33, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 5, 1, 14, 8, 1, 10, 1, 12, 1, 17, 9, 1, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 31, 8, 3,
		1, 3, 0, 0, 4, 0, 2, 4, 6, 0, 0, 31, 0, 8, 1, 0, 0, 0, 2, 15, 1, 0, 0,
		0, 4, 18, 1, 0, 0, 0, 6, 20, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 0,
		0, 1, 10, 1, 1, 0, 0, 0, 11, 14, 3, 4, 2, 0, 12, 14, 3, 6, 3, 0, 13, 11,
		1, 0, 0, 0, 13, 12, 1, 0, 0, 0, 14, 17, 1, 0, 0, 0, 15, 13, 1, 0, 0, 0,
		15, 16, 1, 0, 0, 0, 16, 3, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 18, 19, 5, 3,
		0, 0, 19, 5, 1, 0, 0, 0, 20, 21, 5, 2, 0, 0, 21, 22, 5, 4, 0, 0, 22, 23,
		6, 3, -1, 0, 23, 30, 5, 5, 0, 0, 24, 25, 3, 2, 1, 0, 25, 26, 5, 1, 0, 0,
		26, 27, 5, 4, 0, 0, 27, 28, 6, 3, -1, 0, 28, 29, 5, 5, 0, 0, 29, 31, 1,
		0, 0, 0, 30, 24, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 7, 1, 0, 0, 0, 3,
		13, 15, 30,
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
	PlaceholderParserOPEN_END = 1
	PlaceholderParserOPEN     = 2
	PlaceholderParserTEXT     = 3
	PlaceholderParserID       = 4
	PlaceholderParserCLOSE    = 5
	PlaceholderParserSTRING   = 6
)

// PlaceholderParser rules.
const (
	PlaceholderParserRULE_template    = 0
	PlaceholderParserRULE_content     = 1
	PlaceholderParserRULE_text        = 2
	PlaceholderParserRULE_placeholder = 3
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
		p.SetState(8)
		p.Content()
	}
	{
		p.SetState(9)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PlaceholderParserOPEN || _la == PlaceholderParserTEXT {
		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PlaceholderParserTEXT:
			{
				p.SetState(11)
				p.Text()
			}

		case PlaceholderParserOPEN:
			{
				p.SetState(12)
				p.Placeholder()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(17)
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
		p.SetState(18)
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

	// GetPlaceholderName returns the placeholderName token.
	GetPlaceholderName() antlr.Token

	// GetPlaceholderEndName returns the placeholderEndName token.
	GetPlaceholderEndName() antlr.Token

	// SetPlaceholderName sets the placeholderName token.
	SetPlaceholderName(antlr.Token)

	// SetPlaceholderEndName sets the placeholderEndName token.
	SetPlaceholderEndName(antlr.Token)

	// Getter signatures
	OPEN() antlr.TerminalNode
	AllCLOSE() []antlr.TerminalNode
	CLOSE(i int) antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	Content() IContentContext
	OPEN_END() antlr.TerminalNode

	// IsPlaceholderContext differentiates from other interfaces.
	IsPlaceholderContext()
}

type PlaceholderContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	placeholderName    antlr.Token
	placeholderEndName antlr.Token
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

func (s *PlaceholderContext) GetPlaceholderName() antlr.Token { return s.placeholderName }

func (s *PlaceholderContext) GetPlaceholderEndName() antlr.Token { return s.placeholderEndName }

func (s *PlaceholderContext) SetPlaceholderName(v antlr.Token) { s.placeholderName = v }

func (s *PlaceholderContext) SetPlaceholderEndName(v antlr.Token) { s.placeholderEndName = v }

func (s *PlaceholderContext) OPEN() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserOPEN, 0)
}

func (s *PlaceholderContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserCLOSE)
}

func (s *PlaceholderContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCLOSE, i)
}

func (s *PlaceholderContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserID)
}

func (s *PlaceholderContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserID, i)
}

func (s *PlaceholderContext) Content() IContentContext {
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

func (s *PlaceholderContext) OPEN_END() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserOPEN_END, 0)
}

func (s *PlaceholderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlaceholderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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
	p.EnterRule(localctx, 6, PlaceholderParserRULE_placeholder)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Match(PlaceholderParserOPEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)

		var _m = p.Match(PlaceholderParserID)

		localctx.(*PlaceholderContext).placeholderName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	pushTag(localctx.(*PlaceholderContext).GetPlaceholderName().GetText())
	{
		p.SetState(23)
		p.Match(PlaceholderParserCLOSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(24)
			p.Content()
		}
		{
			p.SetState(25)
			p.Match(PlaceholderParserOPEN_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(26)

			var _m = p.Match(PlaceholderParserID)

			localctx.(*PlaceholderContext).placeholderEndName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		name := localctx.(*PlaceholderContext).GetPlaceholderEndName().GetText()
		tag := peekTag()
		if tag == name {
			popTag()
			// keep as END_TAG
		} else {
			goto errorExit
		}

		{
			p.SetState(28)
			p.Match(PlaceholderParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
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
