// Code generated from /home/alex/GolandProjects/go-placeholders/internal/grammar/PlaceholderLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "'{{block_'", "'{{end_'", "'{{'", "", "", "", "'='", "'}}'", "",
		"'{'",
	}
	staticData.SymbolicNames = []string{
		"", "BLOCK_OPEN", "BLOCK_END", "OPEN", "TEXT", "ID", "INS_WS", "EQ",
		"CLOSE", "STRING", "LBRACE_AS_TEXT",
	}
	staticData.RuleNames = []string{
		"BLOCK_OPEN", "BLOCK_END", "OPEN", "TEXT", "LBRACE_AS_TEXT", "ID", "INS_WS",
		"EQ", "CLOSE", "INS_ANY", "STRING",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 107, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 53, 8, 3, 11, 3, 12,
		3, 54, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 63, 8, 5, 10, 5, 12, 5,
		66, 9, 5, 1, 6, 4, 6, 69, 8, 6, 11, 6, 12, 6, 70, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 5, 10, 90, 8, 10, 10, 10, 12, 10, 93, 9, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 5, 10, 100, 8, 10, 10, 10, 12, 10, 103, 9, 10, 1,
		10, 3, 10, 106, 8, 10, 0, 0, 11, 2, 1, 4, 2, 6, 3, 8, 4, 10, 10, 12, 5,
		14, 6, 16, 7, 18, 8, 20, 0, 22, 9, 2, 0, 1, 6, 1, 0, 123, 123, 2, 0, 65,
		90, 97, 122, 3, 0, 48, 57, 65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32,
		4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92,
		92, 114, 0, 2, 1, 0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1,
		0, 0, 0, 0, 10, 1, 0, 0, 0, 1, 12, 1, 0, 0, 0, 1, 14, 1, 0, 0, 0, 1, 16,
		1, 0, 0, 0, 1, 18, 1, 0, 0, 0, 1, 20, 1, 0, 0, 0, 1, 22, 1, 0, 0, 0, 2,
		24, 1, 0, 0, 0, 4, 35, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 52, 1, 0, 0, 0,
		10, 56, 1, 0, 0, 0, 12, 60, 1, 0, 0, 0, 14, 68, 1, 0, 0, 0, 16, 74, 1,
		0, 0, 0, 18, 76, 1, 0, 0, 0, 20, 81, 1, 0, 0, 0, 22, 105, 1, 0, 0, 0, 24,
		25, 5, 123, 0, 0, 25, 26, 5, 123, 0, 0, 26, 27, 5, 98, 0, 0, 27, 28, 5,
		108, 0, 0, 28, 29, 5, 111, 0, 0, 29, 30, 5, 99, 0, 0, 30, 31, 5, 107, 0,
		0, 31, 32, 5, 95, 0, 0, 32, 33, 1, 0, 0, 0, 33, 34, 6, 0, 0, 0, 34, 3,
		1, 0, 0, 0, 35, 36, 5, 123, 0, 0, 36, 37, 5, 123, 0, 0, 37, 38, 5, 101,
		0, 0, 38, 39, 5, 110, 0, 0, 39, 40, 5, 100, 0, 0, 40, 41, 5, 95, 0, 0,
		41, 42, 1, 0, 0, 0, 42, 43, 6, 1, 0, 0, 43, 5, 1, 0, 0, 0, 44, 45, 5, 123,
		0, 0, 45, 46, 5, 123, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 6, 2, 0, 0, 48,
		7, 1, 0, 0, 0, 49, 53, 8, 0, 0, 0, 50, 51, 5, 123, 0, 0, 51, 53, 8, 0,
		0, 0, 52, 49, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 52,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 57, 5, 123, 0, 0,
		57, 58, 1, 0, 0, 0, 58, 59, 6, 4, 1, 0, 59, 11, 1, 0, 0, 0, 60, 64, 7,
		1, 0, 0, 61, 63, 7, 2, 0, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64,
		62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 13, 1, 0, 0, 0, 66, 64, 1, 0, 0,
		0, 67, 69, 7, 3, 0, 0, 68, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68,
		1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 6, 6, 2, 0,
		73, 15, 1, 0, 0, 0, 74, 75, 5, 61, 0, 0, 75, 17, 1, 0, 0, 0, 76, 77, 5,
		125, 0, 0, 77, 78, 5, 125, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 6, 8, 3, 0,
		80, 19, 1, 0, 0, 0, 81, 82, 9, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 6,
		9, 1, 0, 84, 21, 1, 0, 0, 0, 85, 91, 5, 34, 0, 0, 86, 87, 5, 92, 0, 0,
		87, 90, 9, 0, 0, 0, 88, 90, 8, 4, 0, 0, 89, 86, 1, 0, 0, 0, 89, 88, 1,
		0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		94, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 106, 5, 34, 0, 0, 95, 101, 5, 39,
		0, 0, 96, 97, 5, 92, 0, 0, 97, 100, 9, 0, 0, 0, 98, 100, 8, 5, 0, 0, 99,
		96, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0,
		0, 0, 101, 102, 1, 0, 0, 0, 102, 104, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0,
		104, 106, 5, 39, 0, 0, 105, 85, 1, 0, 0, 0, 105, 95, 1, 0, 0, 0, 106, 23,
		1, 0, 0, 0, 11, 0, 1, 52, 54, 64, 70, 89, 91, 99, 101, 105, 4, 5, 1, 0,
		7, 4, 0, 6, 0, 0, 4, 0, 0,
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
	PlaceholderLexerBLOCK_OPEN     = 1
	PlaceholderLexerBLOCK_END      = 2
	PlaceholderLexerOPEN           = 3
	PlaceholderLexerTEXT           = 4
	PlaceholderLexerID             = 5
	PlaceholderLexerINS_WS         = 6
	PlaceholderLexerEQ             = 7
	PlaceholderLexerCLOSE          = 8
	PlaceholderLexerSTRING         = 9
	PlaceholderLexerLBRACE_AS_TEXT = 10
)

// PlaceholderLexerINSIDE is the PlaceholderLexer mode.
const PlaceholderLexerINSIDE = 1
