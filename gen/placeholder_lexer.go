// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

var blocksTagStack []string

func pushTag(name string) { blocksTagStack = append(blocksTagStack, name) }
func popTag() string {
	if len(blocksTagStack) == 0 {
		return ""
	}
	t := blocksTagStack[len(blocksTagStack)-1]
	blocksTagStack = blocksTagStack[:len(blocksTagStack)-1]
	return t
}
func peekTag() string {
	if len(blocksTagStack) == 0 {
		return ""
	}
	return blocksTagStack[len(blocksTagStack)-1]
}

// "{{A}}" -> "A"
func openNameFromText(s string) string {
	// assumes s starts with "{{" and ends with "}}"
	if len(s) >= 4 {
		return s[2 : len(s)-2]
	}
	return s
}

// "{{end_A}}" -> "A"
func endNameFromText(s string) string {
	// "{{" (2) + "end_" (4) = 6
	if len(s) >= 8 {
		return s[6 : len(s)-2]
	}
	return s
}

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PlaceholderLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PlaceholderLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func placeholderlexerLexerInit() {
	staticData := &PlaceholderLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE", "INSIDE",
	}
	staticData.LiteralNames = []string{
		"", "'{{end_'", "'{{'", "", "", "'}}'",
	}
	staticData.SymbolicNames = []string{
		"", "OPEN_END", "OPEN", "TEXT", "ID", "CLOSE", "STRING",
	}
	staticData.RuleNames = []string{
		"OPEN_END", "OPEN", "TEXT", "ID", "CLOSE", "INS_ANY", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 73, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7,
		3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 4,
		2, 34, 8, 2, 11, 2, 12, 2, 35, 1, 3, 4, 3, 39, 8, 3, 11, 3, 12, 3, 40,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 5, 6, 56, 8, 6, 10, 6, 12, 6, 59, 9, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 5, 6, 66, 8, 6, 10, 6, 12, 6, 69, 9, 6, 1, 6, 3, 6, 72, 8, 6, 0, 0,
		7, 2, 1, 4, 2, 6, 3, 8, 4, 10, 5, 12, 0, 14, 6, 2, 0, 1, 4, 1, 0, 123,
		123, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 10, 10, 13, 13, 34, 34, 92, 92,
		4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 79, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0,
		0, 0, 0, 6, 1, 0, 0, 0, 1, 8, 1, 0, 0, 0, 1, 10, 1, 0, 0, 0, 1, 12, 1,
		0, 0, 0, 1, 14, 1, 0, 0, 0, 2, 16, 1, 0, 0, 0, 4, 25, 1, 0, 0, 0, 6, 33,
		1, 0, 0, 0, 8, 38, 1, 0, 0, 0, 10, 42, 1, 0, 0, 0, 12, 47, 1, 0, 0, 0,
		14, 71, 1, 0, 0, 0, 16, 17, 5, 123, 0, 0, 17, 18, 5, 123, 0, 0, 18, 19,
		5, 101, 0, 0, 19, 20, 5, 110, 0, 0, 20, 21, 5, 100, 0, 0, 21, 22, 5, 95,
		0, 0, 22, 23, 1, 0, 0, 0, 23, 24, 6, 0, 0, 0, 24, 3, 1, 0, 0, 0, 25, 26,
		5, 123, 0, 0, 26, 27, 5, 123, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 6, 1, 0,
		0, 29, 5, 1, 0, 0, 0, 30, 34, 8, 0, 0, 0, 31, 32, 5, 123, 0, 0, 32, 34,
		8, 0, 0, 0, 33, 30, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 7, 1, 0, 0, 0, 37, 39, 7, 1,
		0, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41,
		1, 0, 0, 0, 41, 9, 1, 0, 0, 0, 42, 43, 5, 125, 0, 0, 43, 44, 5, 125, 0,
		0, 44, 45, 1, 0, 0, 0, 45, 46, 6, 4, 1, 0, 46, 11, 1, 0, 0, 0, 47, 48,
		9, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 6, 5, 2, 0, 50, 13, 1, 0, 0, 0,
		51, 57, 5, 34, 0, 0, 52, 53, 5, 92, 0, 0, 53, 56, 9, 0, 0, 0, 54, 56, 8,
		2, 0, 0, 55, 52, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0, 57,
		55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1, 0, 0,
		0, 60, 72, 5, 34, 0, 0, 61, 67, 5, 39, 0, 0, 62, 63, 5, 92, 0, 0, 63, 66,
		9, 0, 0, 0, 64, 66, 8, 3, 0, 0, 65, 62, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0,
		66, 69, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 70, 1,
		0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 72, 5, 39, 0, 0, 71, 51, 1, 0, 0, 0, 71,
		61, 1, 0, 0, 0, 72, 15, 1, 0, 0, 0, 10, 0, 1, 33, 35, 40, 55, 57, 65, 67,
		71, 3, 5, 1, 0, 4, 0, 0, 7, 3, 0,
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

// PlaceholderLexerInit initializes any static state used to implement PlaceholderLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPlaceholderLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PlaceholderLexerInit() {
	staticData := &PlaceholderLexerLexerStaticData
	staticData.once.Do(placeholderlexerLexerInit)
}

// NewPlaceholderLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPlaceholderLexer(input antlr.CharStream) *PlaceholderLexer {
	PlaceholderLexerInit()
	l := new(PlaceholderLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PlaceholderLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PlaceholderLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PlaceholderLexer tokens.
const (
	PlaceholderLexerOPEN_END = 1
	PlaceholderLexerOPEN     = 2
	PlaceholderLexerTEXT     = 3
	PlaceholderLexerID       = 4
	PlaceholderLexerCLOSE    = 5
	PlaceholderLexerSTRING   = 6
)

// PlaceholderLexerINSIDE is the PlaceholderLexer mode.
const PlaceholderLexerINSIDE = 1
