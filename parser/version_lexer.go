// Code generated from Version.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 41, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 35, 10, 8, 12, 8, 14,
	8, 38, 11, 8, 5, 8, 40, 10, 8, 2, 2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 3, 2, 4, 3, 2, 51, 59, 3, 2, 50, 59, 2, 42, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 19, 3, 2,
	2, 2, 7, 21, 3, 2, 2, 2, 9, 24, 3, 2, 2, 2, 11, 27, 3, 2, 2, 2, 13, 29,
	3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 18, 7, 62, 2, 2, 18, 4, 3, 2, 2, 2,
	19, 20, 7, 64, 2, 2, 20, 6, 3, 2, 2, 2, 21, 22, 7, 64, 2, 2, 22, 23, 7,
	63, 2, 2, 23, 8, 3, 2, 2, 2, 24, 25, 7, 62, 2, 2, 25, 26, 7, 63, 2, 2,
	26, 10, 3, 2, 2, 2, 27, 28, 7, 128, 2, 2, 28, 12, 3, 2, 2, 2, 29, 30, 7,
	48, 2, 2, 30, 14, 3, 2, 2, 2, 31, 40, 7, 50, 2, 2, 32, 36, 9, 2, 2, 2,
	33, 35, 9, 3, 2, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3,
	2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39,
	31, 3, 2, 2, 2, 39, 32, 3, 2, 2, 2, 40, 16, 3, 2, 2, 2, 5, 2, 36, 39, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'<'", "'>'", "'>='", "'<='", "'~'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "Number",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "Number",
}

type VersionLexer struct {
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

func NewVersionLexer(input antlr.CharStream) *VersionLexer {

	l := new(VersionLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Version.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// VersionLexer tokens.
const (
	VersionLexerT__0   = 1
	VersionLexerT__1   = 2
	VersionLexerT__2   = 3
	VersionLexerT__3   = 4
	VersionLexerT__4   = 5
	VersionLexerT__5   = 6
	VersionLexerNumber = 7
)
