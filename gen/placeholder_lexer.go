// Code generated from /Users/alex/GolandProjects/antlr/gramair/PlaceholderLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

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
		"", "", "'{{'", "'}}'", "'end_'", "','", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "OPEN", "CLOSE", "END", "COMMA", "EQUALS", "IDENT", "INT",
		"DQSTRING", "SQSTRING", "WSIN",
	}
	staticData.RuleNames = []string{
		"TEXT", "OPEN", "CLOSE", "END", "COMMA", "EQUALS", "IDENT", "INT", "DQSTRING",
		"SQSTRING", "WSIN",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 11, 91, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7,
		3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7,
		9, 2, 10, 7, 10, 1, 0, 1, 0, 1, 0, 4, 0, 28, 8, 0, 11, 0, 12, 0, 29, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 53, 8, 6, 10,
		6, 12, 6, 56, 9, 6, 1, 7, 4, 7, 59, 8, 7, 11, 7, 12, 7, 60, 1, 8, 1, 8,
		1, 8, 1, 8, 5, 8, 67, 8, 8, 10, 8, 12, 8, 70, 9, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 5, 9, 78, 8, 9, 10, 9, 12, 9, 81, 9, 9, 1, 9, 1, 9, 1, 10,
		4, 10, 86, 8, 10, 11, 10, 12, 10, 87, 1, 10, 1, 10, 0, 0, 11, 2, 1, 4,
		2, 6, 3, 8, 4, 10, 5, 12, 6, 14, 7, 16, 8, 18, 9, 20, 10, 22, 11, 2, 0,
		1, 7, 1, 0, 123, 123, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39,
		92, 92, 3, 0, 9, 10, 13, 13, 32, 32, 98, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0,
		0, 0, 1, 6, 1, 0, 0, 0, 1, 8, 1, 0, 0, 0, 1, 10, 1, 0, 0, 0, 1, 12, 1,
		0, 0, 0, 1, 14, 1, 0, 0, 0, 1, 16, 1, 0, 0, 0, 1, 18, 1, 0, 0, 0, 1, 20,
		1, 0, 0, 0, 1, 22, 1, 0, 0, 0, 2, 27, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6,
		36, 1, 0, 0, 0, 8, 41, 1, 0, 0, 0, 10, 46, 1, 0, 0, 0, 12, 48, 1, 0, 0,
		0, 14, 50, 1, 0, 0, 0, 16, 58, 1, 0, 0, 0, 18, 62, 1, 0, 0, 0, 20, 73,
		1, 0, 0, 0, 22, 85, 1, 0, 0, 0, 24, 28, 8, 0, 0, 0, 25, 26, 5, 123, 0,
		0, 26, 28, 8, 0, 0, 0, 27, 24, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29,
		1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 3, 1, 0, 0, 0,
		31, 32, 5, 123, 0, 0, 32, 33, 5, 123, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35,
		6, 1, 0, 0, 35, 5, 1, 0, 0, 0, 36, 37, 5, 125, 0, 0, 37, 38, 5, 125, 0,
		0, 38, 39, 1, 0, 0, 0, 39, 40, 6, 2, 1, 0, 40, 7, 1, 0, 0, 0, 41, 42, 5,
		101, 0, 0, 42, 43, 5, 110, 0, 0, 43, 44, 5, 100, 0, 0, 44, 45, 5, 95, 0,
		0, 45, 9, 1, 0, 0, 0, 46, 47, 5, 44, 0, 0, 47, 11, 1, 0, 0, 0, 48, 49,
		5, 61, 0, 0, 49, 13, 1, 0, 0, 0, 50, 54, 7, 1, 0, 0, 51, 53, 7, 2, 0, 0,
		52, 51, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1,
		0, 0, 0, 55, 15, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 59, 7, 3, 0, 0, 58,
		57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0,
		0, 61, 17, 1, 0, 0, 0, 62, 68, 5, 34, 0, 0, 63, 67, 8, 4, 0, 0, 64, 65,
		5, 92, 0, 0, 65, 67, 9, 0, 0, 0, 66, 63, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0,
		67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 71, 1,
		0, 0, 0, 70, 68, 1, 0, 0, 0, 71, 72, 5, 34, 0, 0, 72, 19, 1, 0, 0, 0, 73,
		79, 5, 39, 0, 0, 74, 78, 8, 5, 0, 0, 75, 76, 5, 92, 0, 0, 76, 78, 9, 0,
		0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		82, 83, 5, 39, 0, 0, 83, 21, 1, 0, 0, 0, 84, 86, 7, 6, 0, 0, 85, 84, 1,
		0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 90, 6, 10, 2, 0, 90, 23, 1, 0, 0, 0, 11, 0, 1, 27,
		29, 54, 60, 66, 68, 77, 79, 87, 3, 5, 1, 0, 4, 0, 0, 6, 0, 0,
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
	PlaceholderLexerTEXT     = 1
	PlaceholderLexerOPEN     = 2
	PlaceholderLexerCLOSE    = 3
	PlaceholderLexerEND      = 4
	PlaceholderLexerCOMMA    = 5
	PlaceholderLexerEQUALS   = 6
	PlaceholderLexerIDENT    = 7
	PlaceholderLexerINT      = 8
	PlaceholderLexerDQSTRING = 9
	PlaceholderLexerSQSTRING = 10
	PlaceholderLexerWSIN     = 11
)

// PlaceholderLexerINSIDE is the PlaceholderLexer mode.
const PlaceholderLexerINSIDE = 1
