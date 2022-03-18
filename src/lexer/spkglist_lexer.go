// Code generated from ./Spkglist.g4 by ANTLR 4.8. DO NOT EDIT.

package lexer

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 80, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 6, 2, 25,
	10, 2, 13, 2, 14, 2, 26, 3, 2, 3, 2, 3, 3, 5, 3, 32, 10, 3, 3, 3, 6, 3,
	35, 10, 3, 13, 3, 14, 3, 36, 3, 4, 3, 4, 7, 4, 41, 10, 4, 12, 4, 14, 4,
	44, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 50, 10, 5, 12, 5, 14, 5, 53, 11,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 59, 10, 6, 12, 6, 14, 6, 62, 11, 6, 6,
	6, 64, 10, 6, 13, 6, 14, 6, 65, 3, 7, 3, 7, 3, 7, 5, 7, 71, 10, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 3, 2, 8, 4, 2, 11, 11,
	34, 34, 4, 2, 12, 12, 15, 15, 3, 2, 98, 98, 3, 2, 50, 59, 4, 2, 67, 92,
	99, 124, 8, 2, 35, 36, 38, 49, 60, 66, 93, 93, 95, 97, 125, 128, 2, 83,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 3, 24, 3, 2, 2, 2, 5, 34, 3, 2, 2, 2, 7, 38, 3, 2, 2,
	2, 9, 47, 3, 2, 2, 2, 11, 63, 3, 2, 2, 2, 13, 70, 3, 2, 2, 2, 15, 72, 3,
	2, 2, 2, 17, 74, 3, 2, 2, 2, 19, 76, 3, 2, 2, 2, 21, 78, 3, 2, 2, 2, 23,
	25, 9, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 24, 3, 2, 2,
	2, 26, 27, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 29, 8, 2, 2, 2, 29, 4, 3,
	2, 2, 2, 30, 32, 7, 15, 2, 2, 31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32,
	33, 3, 2, 2, 2, 33, 35, 7, 12, 2, 2, 34, 31, 3, 2, 2, 2, 35, 36, 3, 2,
	2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 6, 3, 2, 2, 2, 38, 42,
	7, 37, 2, 2, 39, 41, 10, 3, 2, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3, 2, 2,
	2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42,
	3, 2, 2, 2, 45, 46, 8, 4, 2, 2, 46, 8, 3, 2, 2, 2, 47, 51, 5, 15, 8, 2,
	48, 50, 10, 4, 2, 2, 49, 48, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3,
	2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54,
	55, 5, 15, 8, 2, 55, 10, 3, 2, 2, 2, 56, 60, 5, 13, 7, 2, 57, 59, 7, 34,
	2, 2, 58, 57, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 64, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 56, 3, 2, 2, 2,
	64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 12, 3,
	2, 2, 2, 67, 71, 5, 19, 10, 2, 68, 71, 5, 17, 9, 2, 69, 71, 5, 21, 11,
	2, 70, 67, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 69, 3, 2, 2, 2, 71, 14,
	3, 2, 2, 2, 72, 73, 7, 98, 2, 2, 73, 16, 3, 2, 2, 2, 74, 75, 9, 5, 2, 2,
	75, 18, 3, 2, 2, 2, 76, 77, 9, 6, 2, 2, 77, 20, 3, 2, 2, 2, 78, 79, 9,
	7, 2, 2, 79, 22, 3, 2, 2, 2, 11, 2, 26, 31, 36, 42, 51, 60, 65, 70, 3,
	8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "WS", "NL", "COMMENT", "QUOTED_PACKAGE", "BARE_PACKAGE",
}

var lexerRuleNames = []string{
	"WS", "NL", "COMMENT", "QUOTED_PACKAGE", "BARE_PACKAGE", "PKG_BARE", "BACKTICK",
	"DIGIT", "ALPHA", "SYMBOL",
}

type Spkglist struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSpkglist(input antlr.CharStream) *Spkglist {

	l := new(Spkglist)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Spkglist.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Spkglist tokens.
const (
	SpkglistWS             = 1
	SpkglistNL             = 2
	SpkglistCOMMENT        = 3
	SpkglistQUOTED_PACKAGE = 4
	SpkglistBARE_PACKAGE   = 5
)
