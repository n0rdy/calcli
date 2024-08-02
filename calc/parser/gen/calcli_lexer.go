// Code generated from Calcli.g4 by ANTLR 4.13.1. DO NOT EDIT.

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

type CalcliLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalcliLexerLexerStaticData struct {
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

func calclilexerLexerInit() {
	staticData := &CalcliLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'abs'", "'acos'", "'asin'", "'atan'", "'ceil'", "','", "'cos'",
		"'$'", "'.'", "'='", "'e'", "'!'", "'exp'", "'exp2'", "", "", "'floor'",
		"'ln'", "'log'", "'log2'", "'log10'", "'('", "'-'", "'mod'", "'nrt'",
		"'%'", "'percent'", "'pi'", "'+'", "'pmem'", "'^'", "'round'", "'$pr'",
		"')'", "'sin'", "'*'", "'/'", "'sqrt'", "'tan'",
	}
	staticData.SymbolicNames = []string{
		"", "ABS", "ACOS", "ASIN", "ATAN", "CEIL", "COMMA", "COS", "DOLLAR",
		"DOT", "EQUAL", "EULER", "EXCLAMATION", "EXP", "EXP2", "INT", "FLOAT",
		"FLOOR", "LN", "LOG", "LOG2", "LOG10", "LPAREN", "MINUS", "MOD", "NRT",
		"PERCENT", "PERCENTAGE", "PI", "PLUS", "PMEM", "POW", "ROUND", "PREVIOS",
		"RPAREN", "SIN", "STAR", "SLASH", "SQRT", "TAN", "VAR", "WS",
	}
	staticData.RuleNames = []string{
		"ABS", "ACOS", "ASIN", "ATAN", "CEIL", "COMMA", "COS", "DOLLAR", "DOT",
		"EQUAL", "EULER", "EXCLAMATION", "EXP", "EXP2", "INT", "FLOAT", "FLOOR",
		"LN", "LOG", "LOG2", "LOG10", "LPAREN", "MINUS", "MOD", "NRT", "PERCENT",
		"PERCENTAGE", "PI", "PLUS", "PMEM", "POW", "ROUND", "PREVIOS", "RPAREN",
		"SIN", "STAR", "SLASH", "SQRT", "TAN", "VAR", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 41, 248, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 4, 14, 134, 8, 14, 11, 14, 12, 14, 135, 1, 15, 4, 15, 139,
		8, 15, 11, 15, 12, 15, 140, 1, 15, 1, 15, 4, 15, 145, 8, 15, 11, 15, 12,
		15, 146, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 4, 39, 238, 8, 39, 11, 39,
		12, 39, 239, 1, 40, 4, 40, 243, 8, 40, 11, 40, 12, 40, 244, 1, 40, 1, 40,
		0, 0, 41, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19,
		10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37,
		19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55,
		28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73,
		37, 75, 38, 77, 39, 79, 40, 81, 41, 1, 0, 3, 1, 0, 48, 57, 3, 0, 48, 57,
		65, 90, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 252, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0,
		11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0,
		0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0,
		0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0,
		0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1,
		0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49,
		1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0,
		57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0,
		0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0,
		0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0,
		0, 0, 0, 81, 1, 0, 0, 0, 1, 83, 1, 0, 0, 0, 3, 87, 1, 0, 0, 0, 5, 92, 1,
		0, 0, 0, 7, 97, 1, 0, 0, 0, 9, 102, 1, 0, 0, 0, 11, 107, 1, 0, 0, 0, 13,
		109, 1, 0, 0, 0, 15, 113, 1, 0, 0, 0, 17, 115, 1, 0, 0, 0, 19, 117, 1,
		0, 0, 0, 21, 119, 1, 0, 0, 0, 23, 121, 1, 0, 0, 0, 25, 123, 1, 0, 0, 0,
		27, 127, 1, 0, 0, 0, 29, 133, 1, 0, 0, 0, 31, 138, 1, 0, 0, 0, 33, 148,
		1, 0, 0, 0, 35, 154, 1, 0, 0, 0, 37, 157, 1, 0, 0, 0, 39, 161, 1, 0, 0,
		0, 41, 166, 1, 0, 0, 0, 43, 172, 1, 0, 0, 0, 45, 174, 1, 0, 0, 0, 47, 176,
		1, 0, 0, 0, 49, 180, 1, 0, 0, 0, 51, 184, 1, 0, 0, 0, 53, 186, 1, 0, 0,
		0, 55, 194, 1, 0, 0, 0, 57, 197, 1, 0, 0, 0, 59, 199, 1, 0, 0, 0, 61, 204,
		1, 0, 0, 0, 63, 206, 1, 0, 0, 0, 65, 212, 1, 0, 0, 0, 67, 216, 1, 0, 0,
		0, 69, 218, 1, 0, 0, 0, 71, 222, 1, 0, 0, 0, 73, 224, 1, 0, 0, 0, 75, 226,
		1, 0, 0, 0, 77, 231, 1, 0, 0, 0, 79, 235, 1, 0, 0, 0, 81, 242, 1, 0, 0,
		0, 83, 84, 5, 97, 0, 0, 84, 85, 5, 98, 0, 0, 85, 86, 5, 115, 0, 0, 86,
		2, 1, 0, 0, 0, 87, 88, 5, 97, 0, 0, 88, 89, 5, 99, 0, 0, 89, 90, 5, 111,
		0, 0, 90, 91, 5, 115, 0, 0, 91, 4, 1, 0, 0, 0, 92, 93, 5, 97, 0, 0, 93,
		94, 5, 115, 0, 0, 94, 95, 5, 105, 0, 0, 95, 96, 5, 110, 0, 0, 96, 6, 1,
		0, 0, 0, 97, 98, 5, 97, 0, 0, 98, 99, 5, 116, 0, 0, 99, 100, 5, 97, 0,
		0, 100, 101, 5, 110, 0, 0, 101, 8, 1, 0, 0, 0, 102, 103, 5, 99, 0, 0, 103,
		104, 5, 101, 0, 0, 104, 105, 5, 105, 0, 0, 105, 106, 5, 108, 0, 0, 106,
		10, 1, 0, 0, 0, 107, 108, 5, 44, 0, 0, 108, 12, 1, 0, 0, 0, 109, 110, 5,
		99, 0, 0, 110, 111, 5, 111, 0, 0, 111, 112, 5, 115, 0, 0, 112, 14, 1, 0,
		0, 0, 113, 114, 5, 36, 0, 0, 114, 16, 1, 0, 0, 0, 115, 116, 5, 46, 0, 0,
		116, 18, 1, 0, 0, 0, 117, 118, 5, 61, 0, 0, 118, 20, 1, 0, 0, 0, 119, 120,
		5, 101, 0, 0, 120, 22, 1, 0, 0, 0, 121, 122, 5, 33, 0, 0, 122, 24, 1, 0,
		0, 0, 123, 124, 5, 101, 0, 0, 124, 125, 5, 120, 0, 0, 125, 126, 5, 112,
		0, 0, 126, 26, 1, 0, 0, 0, 127, 128, 5, 101, 0, 0, 128, 129, 5, 120, 0,
		0, 129, 130, 5, 112, 0, 0, 130, 131, 5, 50, 0, 0, 131, 28, 1, 0, 0, 0,
		132, 134, 7, 0, 0, 0, 133, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 30, 1, 0, 0, 0, 137, 139, 7,
		0, 0, 0, 138, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 138, 1, 0, 0,
		0, 140, 141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 3, 17, 8, 0, 143,
		145, 7, 0, 0, 0, 144, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 144,
		1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 32, 1, 0, 0, 0, 148, 149, 5, 102,
		0, 0, 149, 150, 5, 108, 0, 0, 150, 151, 5, 111, 0, 0, 151, 152, 5, 111,
		0, 0, 152, 153, 5, 114, 0, 0, 153, 34, 1, 0, 0, 0, 154, 155, 5, 108, 0,
		0, 155, 156, 5, 110, 0, 0, 156, 36, 1, 0, 0, 0, 157, 158, 5, 108, 0, 0,
		158, 159, 5, 111, 0, 0, 159, 160, 5, 103, 0, 0, 160, 38, 1, 0, 0, 0, 161,
		162, 5, 108, 0, 0, 162, 163, 5, 111, 0, 0, 163, 164, 5, 103, 0, 0, 164,
		165, 5, 50, 0, 0, 165, 40, 1, 0, 0, 0, 166, 167, 5, 108, 0, 0, 167, 168,
		5, 111, 0, 0, 168, 169, 5, 103, 0, 0, 169, 170, 5, 49, 0, 0, 170, 171,
		5, 48, 0, 0, 171, 42, 1, 0, 0, 0, 172, 173, 5, 40, 0, 0, 173, 44, 1, 0,
		0, 0, 174, 175, 5, 45, 0, 0, 175, 46, 1, 0, 0, 0, 176, 177, 5, 109, 0,
		0, 177, 178, 5, 111, 0, 0, 178, 179, 5, 100, 0, 0, 179, 48, 1, 0, 0, 0,
		180, 181, 5, 110, 0, 0, 181, 182, 5, 114, 0, 0, 182, 183, 5, 116, 0, 0,
		183, 50, 1, 0, 0, 0, 184, 185, 5, 37, 0, 0, 185, 52, 1, 0, 0, 0, 186, 187,
		5, 112, 0, 0, 187, 188, 5, 101, 0, 0, 188, 189, 5, 114, 0, 0, 189, 190,
		5, 99, 0, 0, 190, 191, 5, 101, 0, 0, 191, 192, 5, 110, 0, 0, 192, 193,
		5, 116, 0, 0, 193, 54, 1, 0, 0, 0, 194, 195, 5, 112, 0, 0, 195, 196, 5,
		105, 0, 0, 196, 56, 1, 0, 0, 0, 197, 198, 5, 43, 0, 0, 198, 58, 1, 0, 0,
		0, 199, 200, 5, 112, 0, 0, 200, 201, 5, 109, 0, 0, 201, 202, 5, 101, 0,
		0, 202, 203, 5, 109, 0, 0, 203, 60, 1, 0, 0, 0, 204, 205, 5, 94, 0, 0,
		205, 62, 1, 0, 0, 0, 206, 207, 5, 114, 0, 0, 207, 208, 5, 111, 0, 0, 208,
		209, 5, 117, 0, 0, 209, 210, 5, 110, 0, 0, 210, 211, 5, 100, 0, 0, 211,
		64, 1, 0, 0, 0, 212, 213, 5, 36, 0, 0, 213, 214, 5, 112, 0, 0, 214, 215,
		5, 114, 0, 0, 215, 66, 1, 0, 0, 0, 216, 217, 5, 41, 0, 0, 217, 68, 1, 0,
		0, 0, 218, 219, 5, 115, 0, 0, 219, 220, 5, 105, 0, 0, 220, 221, 5, 110,
		0, 0, 221, 70, 1, 0, 0, 0, 222, 223, 5, 42, 0, 0, 223, 72, 1, 0, 0, 0,
		224, 225, 5, 47, 0, 0, 225, 74, 1, 0, 0, 0, 226, 227, 5, 115, 0, 0, 227,
		228, 5, 113, 0, 0, 228, 229, 5, 114, 0, 0, 229, 230, 5, 116, 0, 0, 230,
		76, 1, 0, 0, 0, 231, 232, 5, 116, 0, 0, 232, 233, 5, 97, 0, 0, 233, 234,
		5, 110, 0, 0, 234, 78, 1, 0, 0, 0, 235, 237, 3, 15, 7, 0, 236, 238, 7,
		1, 0, 0, 237, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 237, 1, 0, 0,
		0, 239, 240, 1, 0, 0, 0, 240, 80, 1, 0, 0, 0, 241, 243, 7, 2, 0, 0, 242,
		241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 6, 40, 0, 0, 247, 82, 1, 0,
		0, 0, 6, 0, 135, 140, 146, 239, 244, 1, 6, 0, 0,
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

// CalcliLexerInit initializes any static state used to implement CalcliLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcliLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcliLexerInit() {
	staticData := &CalcliLexerLexerStaticData
	staticData.once.Do(calclilexerLexerInit)
}

// NewCalcliLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcliLexer(input antlr.CharStream) *CalcliLexer {
	CalcliLexerInit()
	l := new(CalcliLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalcliLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Calcli.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcliLexer tokens.
const (
	CalcliLexerABS         = 1
	CalcliLexerACOS        = 2
	CalcliLexerASIN        = 3
	CalcliLexerATAN        = 4
	CalcliLexerCEIL        = 5
	CalcliLexerCOMMA       = 6
	CalcliLexerCOS         = 7
	CalcliLexerDOLLAR      = 8
	CalcliLexerDOT         = 9
	CalcliLexerEQUAL       = 10
	CalcliLexerEULER       = 11
	CalcliLexerEXCLAMATION = 12
	CalcliLexerEXP         = 13
	CalcliLexerEXP2        = 14
	CalcliLexerINT         = 15
	CalcliLexerFLOAT       = 16
	CalcliLexerFLOOR       = 17
	CalcliLexerLN          = 18
	CalcliLexerLOG         = 19
	CalcliLexerLOG2        = 20
	CalcliLexerLOG10       = 21
	CalcliLexerLPAREN      = 22
	CalcliLexerMINUS       = 23
	CalcliLexerMOD         = 24
	CalcliLexerNRT         = 25
	CalcliLexerPERCENT     = 26
	CalcliLexerPERCENTAGE  = 27
	CalcliLexerPI          = 28
	CalcliLexerPLUS        = 29
	CalcliLexerPMEM        = 30
	CalcliLexerPOW         = 31
	CalcliLexerROUND       = 32
	CalcliLexerPREVIOS     = 33
	CalcliLexerRPAREN      = 34
	CalcliLexerSIN         = 35
	CalcliLexerSTAR        = 36
	CalcliLexerSLASH       = 37
	CalcliLexerSQRT        = 38
	CalcliLexerTAN         = 39
	CalcliLexerVAR         = 40
	CalcliLexerWS          = 41
)
