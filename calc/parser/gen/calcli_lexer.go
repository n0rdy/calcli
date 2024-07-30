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
		"'.'", "'e'", "'!'", "'exp'", "'exp2'", "", "", "'floor'", "'ln'", "'log'",
		"'log2'", "'log10'", "'('", "'-'", "'mod'", "'nrt'", "'%'", "'percent'",
		"'pi'", "'+'", "'^'", "'round'", "'$pr'", "')'", "'sin'", "'*'", "'/'",
		"'sqrt'", "'tan'",
	}
	staticData.SymbolicNames = []string{
		"", "ABS", "ACOS", "ASIN", "ATAN", "CEIL", "COMMA", "COS", "DOT", "EULER",
		"EXCLAMATION", "EXP", "EXP2", "INT", "FLOAT", "FLOOR", "LN", "LOG",
		"LOG2", "LOG10", "LPAREN", "MINUS", "MOD", "NRT", "PERCENT", "PERCENTAGE",
		"PI", "PLUS", "POW", "ROUND", "PREVIOS", "RPAREN", "SIN", "STAR", "SLASH",
		"SQRT", "TAN", "WS",
	}
	staticData.RuleNames = []string{
		"ABS", "ACOS", "ASIN", "ATAN", "CEIL", "COMMA", "COS", "DOT", "EULER",
		"EXCLAMATION", "EXP", "EXP2", "INT", "FLOAT", "FLOOR", "LN", "LOG",
		"LOG2", "LOG10", "LPAREN", "MINUS", "MOD", "NRT", "PERCENT", "PERCENTAGE",
		"PI", "PLUS", "POW", "ROUND", "PREVIOS", "RPAREN", "SIN", "STAR", "SLASH",
		"SQRT", "TAN", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 225, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 4, 12, 122, 8, 12, 11, 12, 12, 12, 123, 1, 13, 4, 13, 127, 8, 13, 11,
		13, 12, 13, 128, 1, 13, 1, 13, 4, 13, 133, 8, 13, 11, 13, 12, 13, 134,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 4, 36,
		220, 8, 36, 11, 36, 12, 36, 221, 1, 36, 1, 36, 0, 0, 37, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 1, 0, 2, 1, 0, 48,
		57, 3, 0, 9, 10, 13, 13, 32, 32, 228, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 79, 1, 0, 0, 0, 5, 84, 1, 0, 0, 0, 7, 89,
		1, 0, 0, 0, 9, 94, 1, 0, 0, 0, 11, 99, 1, 0, 0, 0, 13, 101, 1, 0, 0, 0,
		15, 105, 1, 0, 0, 0, 17, 107, 1, 0, 0, 0, 19, 109, 1, 0, 0, 0, 21, 111,
		1, 0, 0, 0, 23, 115, 1, 0, 0, 0, 25, 121, 1, 0, 0, 0, 27, 126, 1, 0, 0,
		0, 29, 136, 1, 0, 0, 0, 31, 142, 1, 0, 0, 0, 33, 145, 1, 0, 0, 0, 35, 149,
		1, 0, 0, 0, 37, 154, 1, 0, 0, 0, 39, 160, 1, 0, 0, 0, 41, 162, 1, 0, 0,
		0, 43, 164, 1, 0, 0, 0, 45, 168, 1, 0, 0, 0, 47, 172, 1, 0, 0, 0, 49, 174,
		1, 0, 0, 0, 51, 182, 1, 0, 0, 0, 53, 185, 1, 0, 0, 0, 55, 187, 1, 0, 0,
		0, 57, 189, 1, 0, 0, 0, 59, 195, 1, 0, 0, 0, 61, 199, 1, 0, 0, 0, 63, 201,
		1, 0, 0, 0, 65, 205, 1, 0, 0, 0, 67, 207, 1, 0, 0, 0, 69, 209, 1, 0, 0,
		0, 71, 214, 1, 0, 0, 0, 73, 219, 1, 0, 0, 0, 75, 76, 5, 97, 0, 0, 76, 77,
		5, 98, 0, 0, 77, 78, 5, 115, 0, 0, 78, 2, 1, 0, 0, 0, 79, 80, 5, 97, 0,
		0, 80, 81, 5, 99, 0, 0, 81, 82, 5, 111, 0, 0, 82, 83, 5, 115, 0, 0, 83,
		4, 1, 0, 0, 0, 84, 85, 5, 97, 0, 0, 85, 86, 5, 115, 0, 0, 86, 87, 5, 105,
		0, 0, 87, 88, 5, 110, 0, 0, 88, 6, 1, 0, 0, 0, 89, 90, 5, 97, 0, 0, 90,
		91, 5, 116, 0, 0, 91, 92, 5, 97, 0, 0, 92, 93, 5, 110, 0, 0, 93, 8, 1,
		0, 0, 0, 94, 95, 5, 99, 0, 0, 95, 96, 5, 101, 0, 0, 96, 97, 5, 105, 0,
		0, 97, 98, 5, 108, 0, 0, 98, 10, 1, 0, 0, 0, 99, 100, 5, 44, 0, 0, 100,
		12, 1, 0, 0, 0, 101, 102, 5, 99, 0, 0, 102, 103, 5, 111, 0, 0, 103, 104,
		5, 115, 0, 0, 104, 14, 1, 0, 0, 0, 105, 106, 5, 46, 0, 0, 106, 16, 1, 0,
		0, 0, 107, 108, 5, 101, 0, 0, 108, 18, 1, 0, 0, 0, 109, 110, 5, 33, 0,
		0, 110, 20, 1, 0, 0, 0, 111, 112, 5, 101, 0, 0, 112, 113, 5, 120, 0, 0,
		113, 114, 5, 112, 0, 0, 114, 22, 1, 0, 0, 0, 115, 116, 5, 101, 0, 0, 116,
		117, 5, 120, 0, 0, 117, 118, 5, 112, 0, 0, 118, 119, 5, 50, 0, 0, 119,
		24, 1, 0, 0, 0, 120, 122, 7, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122, 123, 1,
		0, 0, 0, 123, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 26, 1, 0, 0,
		0, 125, 127, 7, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132,
		3, 15, 7, 0, 131, 133, 7, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 134, 1, 0,
		0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 28, 1, 0, 0, 0,
		136, 137, 5, 102, 0, 0, 137, 138, 5, 108, 0, 0, 138, 139, 5, 111, 0, 0,
		139, 140, 5, 111, 0, 0, 140, 141, 5, 114, 0, 0, 141, 30, 1, 0, 0, 0, 142,
		143, 5, 108, 0, 0, 143, 144, 5, 110, 0, 0, 144, 32, 1, 0, 0, 0, 145, 146,
		5, 108, 0, 0, 146, 147, 5, 111, 0, 0, 147, 148, 5, 103, 0, 0, 148, 34,
		1, 0, 0, 0, 149, 150, 5, 108, 0, 0, 150, 151, 5, 111, 0, 0, 151, 152, 5,
		103, 0, 0, 152, 153, 5, 50, 0, 0, 153, 36, 1, 0, 0, 0, 154, 155, 5, 108,
		0, 0, 155, 156, 5, 111, 0, 0, 156, 157, 5, 103, 0, 0, 157, 158, 5, 49,
		0, 0, 158, 159, 5, 48, 0, 0, 159, 38, 1, 0, 0, 0, 160, 161, 5, 40, 0, 0,
		161, 40, 1, 0, 0, 0, 162, 163, 5, 45, 0, 0, 163, 42, 1, 0, 0, 0, 164, 165,
		5, 109, 0, 0, 165, 166, 5, 111, 0, 0, 166, 167, 5, 100, 0, 0, 167, 44,
		1, 0, 0, 0, 168, 169, 5, 110, 0, 0, 169, 170, 5, 114, 0, 0, 170, 171, 5,
		116, 0, 0, 171, 46, 1, 0, 0, 0, 172, 173, 5, 37, 0, 0, 173, 48, 1, 0, 0,
		0, 174, 175, 5, 112, 0, 0, 175, 176, 5, 101, 0, 0, 176, 177, 5, 114, 0,
		0, 177, 178, 5, 99, 0, 0, 178, 179, 5, 101, 0, 0, 179, 180, 5, 110, 0,
		0, 180, 181, 5, 116, 0, 0, 181, 50, 1, 0, 0, 0, 182, 183, 5, 112, 0, 0,
		183, 184, 5, 105, 0, 0, 184, 52, 1, 0, 0, 0, 185, 186, 5, 43, 0, 0, 186,
		54, 1, 0, 0, 0, 187, 188, 5, 94, 0, 0, 188, 56, 1, 0, 0, 0, 189, 190, 5,
		114, 0, 0, 190, 191, 5, 111, 0, 0, 191, 192, 5, 117, 0, 0, 192, 193, 5,
		110, 0, 0, 193, 194, 5, 100, 0, 0, 194, 58, 1, 0, 0, 0, 195, 196, 5, 36,
		0, 0, 196, 197, 5, 112, 0, 0, 197, 198, 5, 114, 0, 0, 198, 60, 1, 0, 0,
		0, 199, 200, 5, 41, 0, 0, 200, 62, 1, 0, 0, 0, 201, 202, 5, 115, 0, 0,
		202, 203, 5, 105, 0, 0, 203, 204, 5, 110, 0, 0, 204, 64, 1, 0, 0, 0, 205,
		206, 5, 42, 0, 0, 206, 66, 1, 0, 0, 0, 207, 208, 5, 47, 0, 0, 208, 68,
		1, 0, 0, 0, 209, 210, 5, 115, 0, 0, 210, 211, 5, 113, 0, 0, 211, 212, 5,
		114, 0, 0, 212, 213, 5, 116, 0, 0, 213, 70, 1, 0, 0, 0, 214, 215, 5, 116,
		0, 0, 215, 216, 5, 97, 0, 0, 216, 217, 5, 110, 0, 0, 217, 72, 1, 0, 0,
		0, 218, 220, 7, 1, 0, 0, 219, 218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221,
		219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224,
		6, 36, 0, 0, 224, 74, 1, 0, 0, 0, 5, 0, 123, 128, 134, 221, 1, 6, 0, 0,
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
	CalcliLexerDOT         = 8
	CalcliLexerEULER       = 9
	CalcliLexerEXCLAMATION = 10
	CalcliLexerEXP         = 11
	CalcliLexerEXP2        = 12
	CalcliLexerINT         = 13
	CalcliLexerFLOAT       = 14
	CalcliLexerFLOOR       = 15
	CalcliLexerLN          = 16
	CalcliLexerLOG         = 17
	CalcliLexerLOG2        = 18
	CalcliLexerLOG10       = 19
	CalcliLexerLPAREN      = 20
	CalcliLexerMINUS       = 21
	CalcliLexerMOD         = 22
	CalcliLexerNRT         = 23
	CalcliLexerPERCENT     = 24
	CalcliLexerPERCENTAGE  = 25
	CalcliLexerPI          = 26
	CalcliLexerPLUS        = 27
	CalcliLexerPOW         = 28
	CalcliLexerROUND       = 29
	CalcliLexerPREVIOS     = 30
	CalcliLexerRPAREN      = 31
	CalcliLexerSIN         = 32
	CalcliLexerSTAR        = 33
	CalcliLexerSLASH       = 34
	CalcliLexerSQRT        = 35
	CalcliLexerTAN         = 36
	CalcliLexerWS          = 37
)
