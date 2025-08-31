// Code generated from /Users/alecszaharia/GolandProjects/go-placeholders/internal/grammar/PlaceholderLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "", "", "", "", "", "", "'='", "", "", "'{'",
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
		4, 0, 10, 139, 6, -1, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3,
		7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9,
		7, 9, 2, 10, 7, 10, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 29, 8, 0, 10, 0, 12,
		0, 32, 9, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 63, 8, 2, 10,
		2, 12, 2, 66, 9, 2, 1, 2, 1, 2, 1, 3, 4, 3, 71, 8, 3, 11, 3, 12, 3, 72,
		1, 3, 1, 3, 1, 3, 5, 3, 78, 8, 3, 10, 3, 12, 3, 81, 9, 3, 5, 3, 83, 8,
		3, 10, 3, 12, 3, 86, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 94,
		8, 5, 10, 5, 12, 5, 97, 9, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 5, 8, 104,
		8, 8, 10, 8, 12, 8, 107, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 122, 8, 10, 10, 10, 12,
		10, 125, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 132, 8, 10, 10,
		10, 12, 10, 135, 9, 10, 1, 10, 3, 10, 138, 8, 10, 0, 0, 11, 2, 1, 4, 2,
		6, 3, 8, 4, 10, 10, 12, 5, 14, 6, 16, 7, 18, 8, 20, 0, 22, 9, 2, 0, 1,
		7, 1, 0, 123, 123, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 125, 125, 4, 0, 10, 10, 13,
		13, 34, 34, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 150, 0, 2, 1,
		0, 0, 0, 0, 4, 1, 0, 0, 0, 0, 6, 1, 0, 0, 0, 0, 8, 1, 0, 0, 0, 0, 10, 1,
		0, 0, 0, 1, 12, 1, 0, 0, 0, 1, 14, 1, 0, 0, 0, 1, 16, 1, 0, 0, 0, 1, 18,
		1, 0, 0, 0, 1, 20, 1, 0, 0, 0, 1, 22, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4,
		42, 1, 0, 0, 0, 6, 58, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 87, 1, 0, 0,
		0, 12, 91, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 100, 1, 0, 0, 0, 18, 105,
		1, 0, 0, 0, 20, 113, 1, 0, 0, 0, 22, 137, 1, 0, 0, 0, 24, 25, 5, 123, 0,
		0, 25, 26, 5, 123, 0, 0, 26, 30, 1, 0, 0, 0, 27, 29, 5, 32, 0, 0, 28, 27,
		1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0,
		31, 33, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33, 34, 5, 98, 0, 0, 34, 35, 5,
		108, 0, 0, 35, 36, 5, 111, 0, 0, 36, 37, 5, 99, 0, 0, 37, 38, 5, 107, 0,
		0, 38, 39, 5, 95, 0, 0, 39, 40, 1, 0, 0, 0, 40, 41, 6, 0, 0, 0, 41, 3,
		1, 0, 0, 0, 42, 43, 5, 123, 0, 0, 43, 44, 5, 123, 0, 0, 44, 48, 1, 0, 0,
		0, 45, 47, 5, 32, 0, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46,
		1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0,
		51, 52, 5, 101, 0, 0, 52, 53, 5, 110, 0, 0, 53, 54, 5, 100, 0, 0, 54, 55,
		5, 95, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 6, 1, 0, 0, 57, 5, 1, 0, 0, 0,
		58, 59, 5, 123, 0, 0, 59, 60, 5, 123, 0, 0, 60, 64, 1, 0, 0, 0, 61, 63,
		5, 32, 0, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0,
		64, 65, 1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68, 6,
		2, 0, 0, 68, 7, 1, 0, 0, 0, 69, 71, 8, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71,
		72, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 84, 1, 0, 0,
		0, 74, 75, 5, 123, 0, 0, 75, 79, 8, 0, 0, 0, 76, 78, 8, 0, 0, 0, 77, 76,
		1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0,
		80, 83, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 74, 1, 0, 0, 0, 83, 86, 1,
		0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 9, 1, 0, 0, 0, 86,
		84, 1, 0, 0, 0, 87, 88, 5, 123, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 6, 4,
		1, 0, 90, 11, 1, 0, 0, 0, 91, 95, 7, 1, 0, 0, 92, 94, 7, 2, 0, 0, 93, 92,
		1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0,
		96, 13, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 99, 7, 3, 0, 0, 99, 15, 1,
		0, 0, 0, 100, 101, 5, 61, 0, 0, 101, 17, 1, 0, 0, 0, 102, 104, 5, 32, 0,
		0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 105,
		106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 109,
		5, 125, 0, 0, 109, 110, 5, 125, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 6,
		8, 2, 0, 112, 19, 1, 0, 0, 0, 113, 114, 8, 4, 0, 0, 114, 115, 1, 0, 0,
		0, 115, 116, 6, 9, 1, 0, 116, 21, 1, 0, 0, 0, 117, 123, 5, 34, 0, 0, 118,
		119, 5, 92, 0, 0, 119, 122, 9, 0, 0, 0, 120, 122, 8, 5, 0, 0, 121, 118,
		1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0,
		0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0,
		126, 138, 5, 34, 0, 0, 127, 133, 5, 39, 0, 0, 128, 129, 5, 92, 0, 0, 129,
		132, 9, 0, 0, 0, 130, 132, 8, 6, 0, 0, 131, 128, 1, 0, 0, 0, 131, 130,
		1, 0, 0, 0, 132, 135, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 136, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 138, 5, 39, 0, 0,
		137, 117, 1, 0, 0, 0, 137, 127, 1, 0, 0, 0, 138, 23, 1, 0, 0, 0, 15, 0,
		1, 30, 48, 64, 72, 79, 84, 95, 105, 121, 123, 131, 133, 137, 3, 5, 1, 0,
		7, 4, 0, 4, 0, 0,
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
