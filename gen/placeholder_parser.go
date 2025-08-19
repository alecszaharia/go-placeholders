// Code generated from /home/alex/GolandProjects/go-placeholders/gramair/PlaceholderParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PlaceholderParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "strings"

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
		"", "'{{end_'", "'{{'", "", "", "", "'='", "", "'}}'", "", "'{'",
	}
	staticData.SymbolicNames = []string{
		"", "OPENEND", "OPEN", "TEXT", "ID", "INT", "EQUALS", "INS_WS", "CLOSE",
		"STRING", "LBRACE_AS_TEXT",
	}
	staticData.RuleNames = []string{
		"template", "content", "text", "attribute_list", "attribute", "value",
		"block",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 74, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 20, 8, 1,
		10, 1, 12, 1, 23, 9, 1, 1, 2, 1, 2, 1, 3, 4, 3, 28, 8, 3, 11, 3, 12, 3,
		29, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 40, 8, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 45, 8, 6, 1, 6, 3, 6, 48, 8, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 55, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 62, 8,
		6, 1, 6, 1, 6, 3, 6, 66, 8, 6, 1, 6, 3, 6, 69, 8, 6, 1, 6, 3, 6, 72, 8,
		6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 1, 2, 0, 5, 5, 9, 9, 77, 0,
		14, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 24, 1, 0, 0, 0, 6, 27, 1, 0, 0, 0,
		8, 31, 1, 0, 0, 0, 10, 35, 1, 0, 0, 0, 12, 71, 1, 0, 0, 0, 14, 15, 3, 2,
		1, 0, 15, 16, 5, 0, 0, 1, 16, 1, 1, 0, 0, 0, 17, 20, 3, 4, 2, 0, 18, 20,
		3, 12, 6, 0, 19, 17, 1, 0, 0, 0, 19, 18, 1, 0, 0, 0, 20, 23, 1, 0, 0, 0,
		21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 3, 1, 0, 0, 0, 23, 21, 1, 0,
		0, 0, 24, 25, 5, 3, 0, 0, 25, 5, 1, 0, 0, 0, 26, 28, 3, 8, 4, 0, 27, 26,
		1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0,
		30, 7, 1, 0, 0, 0, 31, 32, 5, 4, 0, 0, 32, 33, 5, 6, 0, 0, 33, 34, 3, 10,
		5, 0, 34, 9, 1, 0, 0, 0, 35, 36, 7, 0, 0, 0, 36, 11, 1, 0, 0, 0, 37, 39,
		5, 2, 0, 0, 38, 40, 5, 7, 0, 0, 39, 38, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0,
		40, 41, 1, 0, 0, 0, 41, 42, 5, 4, 0, 0, 42, 44, 4, 6, 0, 1, 43, 45, 3,
		6, 3, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 47, 1, 0, 0, 0, 46,
		48, 5, 7, 0, 0, 47, 46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0,
		0, 49, 50, 5, 8, 0, 0, 50, 51, 3, 2, 1, 0, 51, 52, 5, 1, 0, 0, 52, 54,
		5, 4, 0, 0, 53, 55, 5, 7, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 57, 5, 8, 0, 0, 57, 58, 4, 6, 1, 1, 58, 72, 1,
		0, 0, 0, 59, 61, 5, 2, 0, 0, 60, 62, 5, 7, 0, 0, 61, 60, 1, 0, 0, 0, 61,
		62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 65, 5, 4, 0, 0, 64, 66, 3, 6, 3,
		0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 69,
		5, 7, 0, 0, 68, 67, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 72, 5, 8, 0, 0, 71, 37, 1, 0, 0, 0, 71, 59, 1, 0, 0, 0, 72, 13, 1,
		0, 0, 0, 11, 19, 21, 29, 39, 44, 47, 54, 61, 65, 68, 71,
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
	PlaceholderParserOPENEND        = 1
	PlaceholderParserOPEN           = 2
	PlaceholderParserTEXT           = 3
	PlaceholderParserID             = 4
	PlaceholderParserINT            = 5
	PlaceholderParserEQUALS         = 6
	PlaceholderParserINS_WS         = 7
	PlaceholderParserCLOSE          = 8
	PlaceholderParserSTRING         = 9
	PlaceholderParserLBRACE_AS_TEXT = 10
)

// PlaceholderParser rules.
const (
	PlaceholderParserRULE_template       = 0
	PlaceholderParserRULE_content        = 1
	PlaceholderParserRULE_text           = 2
	PlaceholderParserRULE_attribute_list = 3
	PlaceholderParserRULE_attribute      = 4
	PlaceholderParserRULE_value          = 5
	PlaceholderParserRULE_block          = 6
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
		p.SetState(14)
		p.Content()
	}
	{
		p.SetState(15)
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
	AllBlock() []IBlockContext
	Block(i int) IBlockContext

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

func (s *ContentContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *ContentContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PlaceholderParserOPEN || _la == PlaceholderParserTEXT {
		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PlaceholderParserTEXT:
			{
				p.SetState(17)
				p.Text()
			}

		case PlaceholderParserOPEN:
			{
				p.SetState(18)
				p.Block()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(23)
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
		p.SetState(24)
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

// IAttribute_listContext is an interface to support dynamic dispatch.
type IAttribute_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttribute() []IAttributeContext
	Attribute(i int) IAttributeContext

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
	p.EnterRule(localctx, 6, PlaceholderParserRULE_attribute_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PlaceholderParserID {
		{
			p.SetState(26)
			p.Attribute()
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

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
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

func (s *AttributeContext) ID() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserID, 0)
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
	p.EnterRule(localctx, 8, PlaceholderParserRULE_attribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(PlaceholderParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Match(PlaceholderParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
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
	STRING() antlr.TerminalNode

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

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserSTRING, 0)
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
	p.EnterRule(localctx, 10, PlaceholderParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PlaceholderParserINT || _la == PlaceholderParserSTRING) {
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBlockName returns the blockName token.
	GetBlockName() antlr.Token

	// GetBlockEndName returns the blockEndName token.
	GetBlockEndName() antlr.Token

	// SetBlockName sets the blockName token.
	SetBlockName(antlr.Token)

	// SetBlockEndName sets the blockEndName token.
	SetBlockEndName(antlr.Token)

	// Getter signatures
	OPEN() antlr.TerminalNode
	AllCLOSE() []antlr.TerminalNode
	CLOSE(i int) antlr.TerminalNode
	Content() IContentContext
	OPENEND() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllINS_WS() []antlr.TerminalNode
	INS_WS(i int) antlr.TerminalNode
	Attribute_list() IAttribute_listContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blockName    antlr.Token
	blockEndName antlr.Token
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PlaceholderParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlaceholderParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) GetBlockName() antlr.Token { return s.blockName }

func (s *BlockContext) GetBlockEndName() antlr.Token { return s.blockEndName }

func (s *BlockContext) SetBlockName(v antlr.Token) { s.blockName = v }

func (s *BlockContext) SetBlockEndName(v antlr.Token) { s.blockEndName = v }

func (s *BlockContext) OPEN() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserOPEN, 0)
}

func (s *BlockContext) AllCLOSE() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserCLOSE)
}

func (s *BlockContext) CLOSE(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserCLOSE, i)
}

func (s *BlockContext) Content() IContentContext {
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

func (s *BlockContext) OPENEND() antlr.TerminalNode {
	return s.GetToken(PlaceholderParserOPENEND, 0)
}

func (s *BlockContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserID)
}

func (s *BlockContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserID, i)
}

func (s *BlockContext) AllINS_WS() []antlr.TerminalNode {
	return s.GetTokens(PlaceholderParserINS_WS)
}

func (s *BlockContext) INS_WS(i int) antlr.TerminalNode {
	return s.GetToken(PlaceholderParserINS_WS, i)
}

func (s *BlockContext) Attribute_list() IAttribute_listContext {
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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlaceholderParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlaceholderParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlaceholderParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PlaceholderParserRULE_block)
	var _la int

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(PlaceholderParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PlaceholderParserINS_WS {
			{
				p.SetState(38)
				p.Match(PlaceholderParserINS_WS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(41)

			var _m = p.Match(PlaceholderParserID)

			localctx.(*BlockContext).blockName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(42)

		if !(!strings.HasPrefix(localctx.(*BlockContext).GetBlockName().GetText(), "end_")) {
			p.SetError(antlr.NewFailedPredicateException(p, " !strings.HasPrefix($blockName.GetText(), \"end_\") ", ""))
			goto errorExit
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PlaceholderParserID {
			{
				p.SetState(43)
				p.Attribute_list()
			}

		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PlaceholderParserINS_WS {
			{
				p.SetState(46)
				p.Match(PlaceholderParserINS_WS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(49)
			p.Match(PlaceholderParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(50)
			p.Content()
		}
		{
			p.SetState(51)
			p.Match(PlaceholderParserOPENEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)

			var _m = p.Match(PlaceholderParserID)

			localctx.(*BlockContext).blockEndName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PlaceholderParserINS_WS {
			{
				p.SetState(53)
				p.Match(PlaceholderParserINS_WS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(56)
			p.Match(PlaceholderParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(57)

		if !(localctx.(*BlockContext).GetBlockName().GetText() == localctx.(*BlockContext).GetBlockEndName().GetText()) {
			p.SetError(antlr.NewFailedPredicateException(p, " $blockName.GetText() == $blockEndName.GetText() ", ""))
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(PlaceholderParserOPEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PlaceholderParserINS_WS {
			{
				p.SetState(60)
				p.Match(PlaceholderParserINS_WS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(63)

			var _m = p.Match(PlaceholderParserID)

			localctx.(*BlockContext).blockName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PlaceholderParserID {
			{
				p.SetState(64)
				p.Attribute_list()
			}

		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PlaceholderParserINS_WS {
			{
				p.SetState(67)
				p.Match(PlaceholderParserINS_WS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(70)
			p.Match(PlaceholderParserCLOSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (p *PlaceholderParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *BlockContext = nil
		if localctx != nil {
			t = localctx.(*BlockContext)
		}
		return p.Block_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PlaceholderParser) Block_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return !strings.HasPrefix(localctx.(*BlockContext).GetBlockName().GetText(), "end_")

	case 1:
		return localctx.(*BlockContext).GetBlockName().GetText() == localctx.(*BlockContext).GetBlockEndName().GetText()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
