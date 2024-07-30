// Code generated from Calcli.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calcli

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

type CalcliParser struct {
	*antlr.BaseParser
}

var CalcliParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcliParserInit() {
	staticData := &CalcliParserStaticData
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
		"unit", "constant", "number", "variable", "mathFunc", "absMathFunc",
		"acosMathFunc", "asinMathFunc", "atanMathFunc", "ceilMathFunc", "cosMathFunc",
		"expMathFunc", "exp2MathFunc", "floorMathFunc", "lnMathFunc", "logMathFunc",
		"log2MathFunc", "log10MathFunc", "modMathFunc", "nrtMathFunc", "percentMathFunc",
		"roundMathFunc", "sinMathFunc", "sqrtMathFunc", "tanMathFunc", "start",
		"expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 226, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 1, 0, 1, 0, 3, 0, 58, 8, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 86, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 3, 26, 208, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 221, 8, 26, 10, 26, 12, 26, 224,
		9, 26, 1, 26, 0, 1, 52, 27, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 0, 4, 2, 0,
		9, 9, 26, 26, 1, 0, 13, 14, 2, 0, 24, 24, 33, 34, 2, 0, 21, 21, 27, 27,
		226, 0, 57, 1, 0, 0, 0, 2, 59, 1, 0, 0, 0, 4, 61, 1, 0, 0, 0, 6, 63, 1,
		0, 0, 0, 8, 85, 1, 0, 0, 0, 10, 87, 1, 0, 0, 0, 12, 92, 1, 0, 0, 0, 14,
		97, 1, 0, 0, 0, 16, 102, 1, 0, 0, 0, 18, 107, 1, 0, 0, 0, 20, 112, 1, 0,
		0, 0, 22, 117, 1, 0, 0, 0, 24, 122, 1, 0, 0, 0, 26, 127, 1, 0, 0, 0, 28,
		132, 1, 0, 0, 0, 30, 137, 1, 0, 0, 0, 32, 144, 1, 0, 0, 0, 34, 149, 1,
		0, 0, 0, 36, 154, 1, 0, 0, 0, 38, 161, 1, 0, 0, 0, 40, 168, 1, 0, 0, 0,
		42, 175, 1, 0, 0, 0, 44, 180, 1, 0, 0, 0, 46, 185, 1, 0, 0, 0, 48, 190,
		1, 0, 0, 0, 50, 195, 1, 0, 0, 0, 52, 207, 1, 0, 0, 0, 54, 58, 3, 4, 2,
		0, 55, 58, 3, 2, 1, 0, 56, 58, 3, 6, 3, 0, 57, 54, 1, 0, 0, 0, 57, 55,
		1, 0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 1, 1, 0, 0, 0, 59, 60, 7, 0, 0, 0,
		60, 3, 1, 0, 0, 0, 61, 62, 7, 1, 0, 0, 62, 5, 1, 0, 0, 0, 63, 64, 5, 30,
		0, 0, 64, 7, 1, 0, 0, 0, 65, 86, 3, 10, 5, 0, 66, 86, 3, 12, 6, 0, 67,
		86, 3, 14, 7, 0, 68, 86, 3, 16, 8, 0, 69, 86, 3, 18, 9, 0, 70, 86, 3, 20,
		10, 0, 71, 86, 3, 22, 11, 0, 72, 86, 3, 24, 12, 0, 73, 86, 3, 26, 13, 0,
		74, 86, 3, 28, 14, 0, 75, 86, 3, 30, 15, 0, 76, 86, 3, 32, 16, 0, 77, 86,
		3, 34, 17, 0, 78, 86, 3, 36, 18, 0, 79, 86, 3, 38, 19, 0, 80, 86, 3, 40,
		20, 0, 81, 86, 3, 42, 21, 0, 82, 86, 3, 44, 22, 0, 83, 86, 3, 46, 23, 0,
		84, 86, 3, 48, 24, 0, 85, 65, 1, 0, 0, 0, 85, 66, 1, 0, 0, 0, 85, 67, 1,
		0, 0, 0, 85, 68, 1, 0, 0, 0, 85, 69, 1, 0, 0, 0, 85, 70, 1, 0, 0, 0, 85,
		71, 1, 0, 0, 0, 85, 72, 1, 0, 0, 0, 85, 73, 1, 0, 0, 0, 85, 74, 1, 0, 0,
		0, 85, 75, 1, 0, 0, 0, 85, 76, 1, 0, 0, 0, 85, 77, 1, 0, 0, 0, 85, 78,
		1, 0, 0, 0, 85, 79, 1, 0, 0, 0, 85, 80, 1, 0, 0, 0, 85, 81, 1, 0, 0, 0,
		85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 9, 1, 0,
		0, 0, 87, 88, 5, 1, 0, 0, 88, 89, 5, 20, 0, 0, 89, 90, 3, 52, 26, 0, 90,
		91, 5, 31, 0, 0, 91, 11, 1, 0, 0, 0, 92, 93, 5, 2, 0, 0, 93, 94, 5, 20,
		0, 0, 94, 95, 3, 52, 26, 0, 95, 96, 5, 31, 0, 0, 96, 13, 1, 0, 0, 0, 97,
		98, 5, 3, 0, 0, 98, 99, 5, 20, 0, 0, 99, 100, 3, 52, 26, 0, 100, 101, 5,
		31, 0, 0, 101, 15, 1, 0, 0, 0, 102, 103, 5, 4, 0, 0, 103, 104, 5, 20, 0,
		0, 104, 105, 3, 52, 26, 0, 105, 106, 5, 31, 0, 0, 106, 17, 1, 0, 0, 0,
		107, 108, 5, 5, 0, 0, 108, 109, 5, 20, 0, 0, 109, 110, 3, 52, 26, 0, 110,
		111, 5, 31, 0, 0, 111, 19, 1, 0, 0, 0, 112, 113, 5, 7, 0, 0, 113, 114,
		5, 20, 0, 0, 114, 115, 3, 52, 26, 0, 115, 116, 5, 31, 0, 0, 116, 21, 1,
		0, 0, 0, 117, 118, 5, 11, 0, 0, 118, 119, 5, 20, 0, 0, 119, 120, 3, 52,
		26, 0, 120, 121, 5, 31, 0, 0, 121, 23, 1, 0, 0, 0, 122, 123, 5, 12, 0,
		0, 123, 124, 5, 20, 0, 0, 124, 125, 3, 52, 26, 0, 125, 126, 5, 31, 0, 0,
		126, 25, 1, 0, 0, 0, 127, 128, 5, 15, 0, 0, 128, 129, 5, 20, 0, 0, 129,
		130, 3, 52, 26, 0, 130, 131, 5, 31, 0, 0, 131, 27, 1, 0, 0, 0, 132, 133,
		5, 16, 0, 0, 133, 134, 5, 20, 0, 0, 134, 135, 3, 52, 26, 0, 135, 136, 5,
		31, 0, 0, 136, 29, 1, 0, 0, 0, 137, 138, 5, 17, 0, 0, 138, 139, 5, 20,
		0, 0, 139, 140, 3, 52, 26, 0, 140, 141, 5, 6, 0, 0, 141, 142, 3, 52, 26,
		0, 142, 143, 5, 31, 0, 0, 143, 31, 1, 0, 0, 0, 144, 145, 5, 18, 0, 0, 145,
		146, 5, 20, 0, 0, 146, 147, 3, 52, 26, 0, 147, 148, 5, 31, 0, 0, 148, 33,
		1, 0, 0, 0, 149, 150, 5, 19, 0, 0, 150, 151, 5, 20, 0, 0, 151, 152, 3,
		52, 26, 0, 152, 153, 5, 31, 0, 0, 153, 35, 1, 0, 0, 0, 154, 155, 5, 22,
		0, 0, 155, 156, 5, 20, 0, 0, 156, 157, 3, 52, 26, 0, 157, 158, 5, 6, 0,
		0, 158, 159, 3, 52, 26, 0, 159, 160, 5, 31, 0, 0, 160, 37, 1, 0, 0, 0,
		161, 162, 5, 23, 0, 0, 162, 163, 5, 20, 0, 0, 163, 164, 3, 52, 26, 0, 164,
		165, 5, 6, 0, 0, 165, 166, 3, 52, 26, 0, 166, 167, 5, 31, 0, 0, 167, 39,
		1, 0, 0, 0, 168, 169, 5, 25, 0, 0, 169, 170, 5, 20, 0, 0, 170, 171, 3,
		52, 26, 0, 171, 172, 5, 6, 0, 0, 172, 173, 3, 52, 26, 0, 173, 174, 5, 31,
		0, 0, 174, 41, 1, 0, 0, 0, 175, 176, 5, 29, 0, 0, 176, 177, 5, 20, 0, 0,
		177, 178, 3, 52, 26, 0, 178, 179, 5, 31, 0, 0, 179, 43, 1, 0, 0, 0, 180,
		181, 5, 32, 0, 0, 181, 182, 5, 20, 0, 0, 182, 183, 3, 52, 26, 0, 183, 184,
		5, 31, 0, 0, 184, 45, 1, 0, 0, 0, 185, 186, 5, 35, 0, 0, 186, 187, 5, 20,
		0, 0, 187, 188, 3, 52, 26, 0, 188, 189, 5, 31, 0, 0, 189, 47, 1, 0, 0,
		0, 190, 191, 5, 36, 0, 0, 191, 192, 5, 20, 0, 0, 192, 193, 3, 52, 26, 0,
		193, 194, 5, 31, 0, 0, 194, 49, 1, 0, 0, 0, 195, 196, 3, 52, 26, 0, 196,
		197, 5, 0, 0, 1, 197, 51, 1, 0, 0, 0, 198, 199, 6, 26, -1, 0, 199, 208,
		3, 0, 0, 0, 200, 201, 5, 21, 0, 0, 201, 208, 3, 52, 26, 7, 202, 203, 5,
		20, 0, 0, 203, 204, 3, 52, 26, 0, 204, 205, 5, 31, 0, 0, 205, 208, 1, 0,
		0, 0, 206, 208, 3, 8, 4, 0, 207, 198, 1, 0, 0, 0, 207, 200, 1, 0, 0, 0,
		207, 202, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 222, 1, 0, 0, 0, 209,
		210, 10, 3, 0, 0, 210, 211, 5, 28, 0, 0, 211, 221, 3, 52, 26, 4, 212, 213,
		10, 2, 0, 0, 213, 214, 7, 2, 0, 0, 214, 221, 3, 52, 26, 3, 215, 216, 10,
		1, 0, 0, 216, 217, 7, 3, 0, 0, 217, 221, 3, 52, 26, 2, 218, 219, 10, 4,
		0, 0, 219, 221, 5, 10, 0, 0, 220, 209, 1, 0, 0, 0, 220, 212, 1, 0, 0, 0,
		220, 215, 1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222,
		220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 53, 1, 0, 0, 0, 224, 222, 1,
		0, 0, 0, 5, 57, 85, 207, 220, 222,
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

// CalcliParserInit initializes any static state used to implement CalcliParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalcliParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcliParserInit() {
	staticData := &CalcliParserStaticData
	staticData.once.Do(calcliParserInit)
}

// NewCalcliParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalcliParser(input antlr.TokenStream) *CalcliParser {
	CalcliParserInit()
	this := new(CalcliParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalcliParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Calcli.g4"

	return this
}

// CalcliParser tokens.
const (
	CalcliParserEOF         = antlr.TokenEOF
	CalcliParserABS         = 1
	CalcliParserACOS        = 2
	CalcliParserASIN        = 3
	CalcliParserATAN        = 4
	CalcliParserCEIL        = 5
	CalcliParserCOMMA       = 6
	CalcliParserCOS         = 7
	CalcliParserDOT         = 8
	CalcliParserEULER       = 9
	CalcliParserEXCLAMATION = 10
	CalcliParserEXP         = 11
	CalcliParserEXP2        = 12
	CalcliParserINT         = 13
	CalcliParserFLOAT       = 14
	CalcliParserFLOOR       = 15
	CalcliParserLN          = 16
	CalcliParserLOG         = 17
	CalcliParserLOG2        = 18
	CalcliParserLOG10       = 19
	CalcliParserLPAREN      = 20
	CalcliParserMINUS       = 21
	CalcliParserMOD         = 22
	CalcliParserNRT         = 23
	CalcliParserPERCENT     = 24
	CalcliParserPERCENTAGE  = 25
	CalcliParserPI          = 26
	CalcliParserPLUS        = 27
	CalcliParserPOW         = 28
	CalcliParserROUND       = 29
	CalcliParserPREVIOS     = 30
	CalcliParserRPAREN      = 31
	CalcliParserSIN         = 32
	CalcliParserSTAR        = 33
	CalcliParserSLASH       = 34
	CalcliParserSQRT        = 35
	CalcliParserTAN         = 36
	CalcliParserWS          = 37
)

// CalcliParser rules.
const (
	CalcliParserRULE_unit            = 0
	CalcliParserRULE_constant        = 1
	CalcliParserRULE_number          = 2
	CalcliParserRULE_variable        = 3
	CalcliParserRULE_mathFunc        = 4
	CalcliParserRULE_absMathFunc     = 5
	CalcliParserRULE_acosMathFunc    = 6
	CalcliParserRULE_asinMathFunc    = 7
	CalcliParserRULE_atanMathFunc    = 8
	CalcliParserRULE_ceilMathFunc    = 9
	CalcliParserRULE_cosMathFunc     = 10
	CalcliParserRULE_expMathFunc     = 11
	CalcliParserRULE_exp2MathFunc    = 12
	CalcliParserRULE_floorMathFunc   = 13
	CalcliParserRULE_lnMathFunc      = 14
	CalcliParserRULE_logMathFunc     = 15
	CalcliParserRULE_log2MathFunc    = 16
	CalcliParserRULE_log10MathFunc   = 17
	CalcliParserRULE_modMathFunc     = 18
	CalcliParserRULE_nrtMathFunc     = 19
	CalcliParserRULE_percentMathFunc = 20
	CalcliParserRULE_roundMathFunc   = 21
	CalcliParserRULE_sinMathFunc     = 22
	CalcliParserRULE_sqrtMathFunc    = 23
	CalcliParserRULE_tanMathFunc     = 24
	CalcliParserRULE_start           = 25
	CalcliParserRULE_expr            = 26
)

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	Constant() IConstantContext
	Variable() IVariableContext

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *UnitContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *UnitContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *CalcliParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcliParserRULE_unit)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcliParserINT, CalcliParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Number()
		}

	case CalcliParserEULER, CalcliParserPI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Constant()
		}

	case CalcliParserPREVIOS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Variable()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EULER() antlr.TerminalNode
	PI() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) EULER() antlr.TerminalNode {
	return s.GetToken(CalcliParserEULER, 0)
}

func (s *ConstantContext) PI() antlr.TerminalNode {
	return s.GetToken(CalcliParserPI, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *CalcliParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalcliParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcliParserEULER || _la == CalcliParserPI) {
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) INT() antlr.TerminalNode {
	return s.GetToken(CalcliParserINT, 0)
}

func (s *NumberContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(CalcliParserFLOAT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *CalcliParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcliParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CalcliParserINT || _la == CalcliParserFLOAT) {
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

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PREVIOS() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) PREVIOS() antlr.TerminalNode {
	return s.GetToken(CalcliParserPREVIOS, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *CalcliParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcliParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(CalcliParserPREVIOS)
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

// IMathFuncContext is an interface to support dynamic dispatch.
type IMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AbsMathFunc() IAbsMathFuncContext
	AcosMathFunc() IAcosMathFuncContext
	AsinMathFunc() IAsinMathFuncContext
	AtanMathFunc() IAtanMathFuncContext
	CeilMathFunc() ICeilMathFuncContext
	CosMathFunc() ICosMathFuncContext
	ExpMathFunc() IExpMathFuncContext
	Exp2MathFunc() IExp2MathFuncContext
	FloorMathFunc() IFloorMathFuncContext
	LnMathFunc() ILnMathFuncContext
	LogMathFunc() ILogMathFuncContext
	Log2MathFunc() ILog2MathFuncContext
	Log10MathFunc() ILog10MathFuncContext
	ModMathFunc() IModMathFuncContext
	NrtMathFunc() INrtMathFuncContext
	PercentMathFunc() IPercentMathFuncContext
	RoundMathFunc() IRoundMathFuncContext
	SinMathFunc() ISinMathFuncContext
	SqrtMathFunc() ISqrtMathFuncContext
	TanMathFunc() ITanMathFuncContext

	// IsMathFuncContext differentiates from other interfaces.
	IsMathFuncContext()
}

type MathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathFuncContext() *MathFuncContext {
	var p = new(MathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_mathFunc
	return p
}

func InitEmptyMathFuncContext(p *MathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_mathFunc
}

func (*MathFuncContext) IsMathFuncContext() {}

func NewMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathFuncContext {
	var p = new(MathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_mathFunc

	return p
}

func (s *MathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *MathFuncContext) AbsMathFunc() IAbsMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsMathFuncContext)
}

func (s *MathFuncContext) AcosMathFunc() IAcosMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAcosMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAcosMathFuncContext)
}

func (s *MathFuncContext) AsinMathFunc() IAsinMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsinMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsinMathFuncContext)
}

func (s *MathFuncContext) AtanMathFunc() IAtanMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtanMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtanMathFuncContext)
}

func (s *MathFuncContext) CeilMathFunc() ICeilMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICeilMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICeilMathFuncContext)
}

func (s *MathFuncContext) CosMathFunc() ICosMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICosMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICosMathFuncContext)
}

func (s *MathFuncContext) ExpMathFunc() IExpMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpMathFuncContext)
}

func (s *MathFuncContext) Exp2MathFunc() IExp2MathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExp2MathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExp2MathFuncContext)
}

func (s *MathFuncContext) FloorMathFunc() IFloorMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloorMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloorMathFuncContext)
}

func (s *MathFuncContext) LnMathFunc() ILnMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILnMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILnMathFuncContext)
}

func (s *MathFuncContext) LogMathFunc() ILogMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogMathFuncContext)
}

func (s *MathFuncContext) Log2MathFunc() ILog2MathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILog2MathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILog2MathFuncContext)
}

func (s *MathFuncContext) Log10MathFunc() ILog10MathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILog10MathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILog10MathFuncContext)
}

func (s *MathFuncContext) ModMathFunc() IModMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModMathFuncContext)
}

func (s *MathFuncContext) NrtMathFunc() INrtMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INrtMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INrtMathFuncContext)
}

func (s *MathFuncContext) PercentMathFunc() IPercentMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPercentMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPercentMathFuncContext)
}

func (s *MathFuncContext) RoundMathFunc() IRoundMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoundMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoundMathFuncContext)
}

func (s *MathFuncContext) SinMathFunc() ISinMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISinMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISinMathFuncContext)
}

func (s *MathFuncContext) SqrtMathFunc() ISqrtMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISqrtMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISqrtMathFuncContext)
}

func (s *MathFuncContext) TanMathFunc() ITanMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITanMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITanMathFuncContext)
}

func (s *MathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterMathFunc(s)
	}
}

func (s *MathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitMathFunc(s)
	}
}

func (p *CalcliParser) MathFunc() (localctx IMathFuncContext) {
	localctx = NewMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcliParserRULE_mathFunc)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcliParserABS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.AbsMathFunc()
		}

	case CalcliParserACOS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.AcosMathFunc()
		}

	case CalcliParserASIN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.AsinMathFunc()
		}

	case CalcliParserATAN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.AtanMathFunc()
		}

	case CalcliParserCEIL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.CeilMathFunc()
		}

	case CalcliParserCOS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.CosMathFunc()
		}

	case CalcliParserEXP:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(71)
			p.ExpMathFunc()
		}

	case CalcliParserEXP2:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(72)
			p.Exp2MathFunc()
		}

	case CalcliParserFLOOR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(73)
			p.FloorMathFunc()
		}

	case CalcliParserLN:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(74)
			p.LnMathFunc()
		}

	case CalcliParserLOG:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(75)
			p.LogMathFunc()
		}

	case CalcliParserLOG2:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(76)
			p.Log2MathFunc()
		}

	case CalcliParserLOG10:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(77)
			p.Log10MathFunc()
		}

	case CalcliParserMOD:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(78)
			p.ModMathFunc()
		}

	case CalcliParserNRT:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(79)
			p.NrtMathFunc()
		}

	case CalcliParserPERCENTAGE:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(80)
			p.PercentMathFunc()
		}

	case CalcliParserROUND:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(81)
			p.RoundMathFunc()
		}

	case CalcliParserSIN:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(82)
			p.SinMathFunc()
		}

	case CalcliParserSQRT:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(83)
			p.SqrtMathFunc()
		}

	case CalcliParserTAN:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(84)
			p.TanMathFunc()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IAbsMathFuncContext is an interface to support dynamic dispatch.
type IAbsMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	ABS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsAbsMathFuncContext differentiates from other interfaces.
	IsAbsMathFuncContext()
}

type AbsMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyAbsMathFuncContext() *AbsMathFuncContext {
	var p = new(AbsMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_absMathFunc
	return p
}

func InitEmptyAbsMathFuncContext(p *AbsMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_absMathFunc
}

func (*AbsMathFuncContext) IsAbsMathFuncContext() {}

func NewAbsMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsMathFuncContext {
	var p = new(AbsMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_absMathFunc

	return p
}

func (s *AbsMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsMathFuncContext) GetNum() IExprContext { return s.num }

func (s *AbsMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *AbsMathFuncContext) ABS() antlr.TerminalNode {
	return s.GetToken(CalcliParserABS, 0)
}

func (s *AbsMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *AbsMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *AbsMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AbsMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterAbsMathFunc(s)
	}
}

func (s *AbsMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitAbsMathFunc(s)
	}
}

func (p *CalcliParser) AbsMathFunc() (localctx IAbsMathFuncContext) {
	localctx = NewAbsMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalcliParserRULE_absMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(CalcliParserABS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)

		var _x = p.expr(0)

		localctx.(*AbsMathFuncContext).num = _x
	}
	{
		p.SetState(90)
		p.Match(CalcliParserRPAREN)
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

// IAcosMathFuncContext is an interface to support dynamic dispatch.
type IAcosMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	ACOS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsAcosMathFuncContext differentiates from other interfaces.
	IsAcosMathFuncContext()
}

type AcosMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyAcosMathFuncContext() *AcosMathFuncContext {
	var p = new(AcosMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_acosMathFunc
	return p
}

func InitEmptyAcosMathFuncContext(p *AcosMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_acosMathFunc
}

func (*AcosMathFuncContext) IsAcosMathFuncContext() {}

func NewAcosMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AcosMathFuncContext {
	var p = new(AcosMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_acosMathFunc

	return p
}

func (s *AcosMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *AcosMathFuncContext) GetNum() IExprContext { return s.num }

func (s *AcosMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *AcosMathFuncContext) ACOS() antlr.TerminalNode {
	return s.GetToken(CalcliParserACOS, 0)
}

func (s *AcosMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *AcosMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *AcosMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AcosMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcosMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AcosMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterAcosMathFunc(s)
	}
}

func (s *AcosMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitAcosMathFunc(s)
	}
}

func (p *CalcliParser) AcosMathFunc() (localctx IAcosMathFuncContext) {
	localctx = NewAcosMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalcliParserRULE_acosMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(CalcliParserACOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)

		var _x = p.expr(0)

		localctx.(*AcosMathFuncContext).num = _x
	}
	{
		p.SetState(95)
		p.Match(CalcliParserRPAREN)
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

// IAsinMathFuncContext is an interface to support dynamic dispatch.
type IAsinMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	ASIN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsAsinMathFuncContext differentiates from other interfaces.
	IsAsinMathFuncContext()
}

type AsinMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyAsinMathFuncContext() *AsinMathFuncContext {
	var p = new(AsinMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_asinMathFunc
	return p
}

func InitEmptyAsinMathFuncContext(p *AsinMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_asinMathFunc
}

func (*AsinMathFuncContext) IsAsinMathFuncContext() {}

func NewAsinMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsinMathFuncContext {
	var p = new(AsinMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_asinMathFunc

	return p
}

func (s *AsinMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *AsinMathFuncContext) GetNum() IExprContext { return s.num }

func (s *AsinMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *AsinMathFuncContext) ASIN() antlr.TerminalNode {
	return s.GetToken(CalcliParserASIN, 0)
}

func (s *AsinMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *AsinMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *AsinMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsinMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsinMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsinMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterAsinMathFunc(s)
	}
}

func (s *AsinMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitAsinMathFunc(s)
	}
}

func (p *CalcliParser) AsinMathFunc() (localctx IAsinMathFuncContext) {
	localctx = NewAsinMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalcliParserRULE_asinMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(CalcliParserASIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)

		var _x = p.expr(0)

		localctx.(*AsinMathFuncContext).num = _x
	}
	{
		p.SetState(100)
		p.Match(CalcliParserRPAREN)
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

// IAtanMathFuncContext is an interface to support dynamic dispatch.
type IAtanMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	ATAN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsAtanMathFuncContext differentiates from other interfaces.
	IsAtanMathFuncContext()
}

type AtanMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyAtanMathFuncContext() *AtanMathFuncContext {
	var p = new(AtanMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_atanMathFunc
	return p
}

func InitEmptyAtanMathFuncContext(p *AtanMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_atanMathFunc
}

func (*AtanMathFuncContext) IsAtanMathFuncContext() {}

func NewAtanMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtanMathFuncContext {
	var p = new(AtanMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_atanMathFunc

	return p
}

func (s *AtanMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *AtanMathFuncContext) GetNum() IExprContext { return s.num }

func (s *AtanMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *AtanMathFuncContext) ATAN() antlr.TerminalNode {
	return s.GetToken(CalcliParserATAN, 0)
}

func (s *AtanMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *AtanMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *AtanMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AtanMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtanMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtanMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterAtanMathFunc(s)
	}
}

func (s *AtanMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitAtanMathFunc(s)
	}
}

func (p *CalcliParser) AtanMathFunc() (localctx IAtanMathFuncContext) {
	localctx = NewAtanMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CalcliParserRULE_atanMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(CalcliParserATAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)

		var _x = p.expr(0)

		localctx.(*AtanMathFuncContext).num = _x
	}
	{
		p.SetState(105)
		p.Match(CalcliParserRPAREN)
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

// ICeilMathFuncContext is an interface to support dynamic dispatch.
type ICeilMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	CEIL() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsCeilMathFuncContext differentiates from other interfaces.
	IsCeilMathFuncContext()
}

type CeilMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyCeilMathFuncContext() *CeilMathFuncContext {
	var p = new(CeilMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_ceilMathFunc
	return p
}

func InitEmptyCeilMathFuncContext(p *CeilMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_ceilMathFunc
}

func (*CeilMathFuncContext) IsCeilMathFuncContext() {}

func NewCeilMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CeilMathFuncContext {
	var p = new(CeilMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_ceilMathFunc

	return p
}

func (s *CeilMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *CeilMathFuncContext) GetNum() IExprContext { return s.num }

func (s *CeilMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *CeilMathFuncContext) CEIL() antlr.TerminalNode {
	return s.GetToken(CalcliParserCEIL, 0)
}

func (s *CeilMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *CeilMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *CeilMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CeilMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CeilMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterCeilMathFunc(s)
	}
}

func (s *CeilMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitCeilMathFunc(s)
	}
}

func (p *CalcliParser) CeilMathFunc() (localctx ICeilMathFuncContext) {
	localctx = NewCeilMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CalcliParserRULE_ceilMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(CalcliParserCEIL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)

		var _x = p.expr(0)

		localctx.(*CeilMathFuncContext).num = _x
	}
	{
		p.SetState(110)
		p.Match(CalcliParserRPAREN)
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

// ICosMathFuncContext is an interface to support dynamic dispatch.
type ICosMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	COS() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsCosMathFuncContext differentiates from other interfaces.
	IsCosMathFuncContext()
}

type CosMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyCosMathFuncContext() *CosMathFuncContext {
	var p = new(CosMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_cosMathFunc
	return p
}

func InitEmptyCosMathFuncContext(p *CosMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_cosMathFunc
}

func (*CosMathFuncContext) IsCosMathFuncContext() {}

func NewCosMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CosMathFuncContext {
	var p = new(CosMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_cosMathFunc

	return p
}

func (s *CosMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *CosMathFuncContext) GetNum() IExprContext { return s.num }

func (s *CosMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *CosMathFuncContext) COS() antlr.TerminalNode {
	return s.GetToken(CalcliParserCOS, 0)
}

func (s *CosMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *CosMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *CosMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CosMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CosMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterCosMathFunc(s)
	}
}

func (s *CosMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitCosMathFunc(s)
	}
}

func (p *CalcliParser) CosMathFunc() (localctx ICosMathFuncContext) {
	localctx = NewCosMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CalcliParserRULE_cosMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(CalcliParserCOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)

		var _x = p.expr(0)

		localctx.(*CosMathFuncContext).num = _x
	}
	{
		p.SetState(115)
		p.Match(CalcliParserRPAREN)
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

// IExpMathFuncContext is an interface to support dynamic dispatch.
type IExpMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	EXP() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsExpMathFuncContext differentiates from other interfaces.
	IsExpMathFuncContext()
}

type ExpMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyExpMathFuncContext() *ExpMathFuncContext {
	var p = new(ExpMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_expMathFunc
	return p
}

func InitEmptyExpMathFuncContext(p *ExpMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_expMathFunc
}

func (*ExpMathFuncContext) IsExpMathFuncContext() {}

func NewExpMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpMathFuncContext {
	var p = new(ExpMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_expMathFunc

	return p
}

func (s *ExpMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpMathFuncContext) GetNum() IExprContext { return s.num }

func (s *ExpMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *ExpMathFuncContext) EXP() antlr.TerminalNode {
	return s.GetToken(CalcliParserEXP, 0)
}

func (s *ExpMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *ExpMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *ExpMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterExpMathFunc(s)
	}
}

func (s *ExpMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitExpMathFunc(s)
	}
}

func (p *CalcliParser) ExpMathFunc() (localctx IExpMathFuncContext) {
	localctx = NewExpMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CalcliParserRULE_expMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(CalcliParserEXP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(119)

		var _x = p.expr(0)

		localctx.(*ExpMathFuncContext).num = _x
	}
	{
		p.SetState(120)
		p.Match(CalcliParserRPAREN)
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

// IExp2MathFuncContext is an interface to support dynamic dispatch.
type IExp2MathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	EXP2() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsExp2MathFuncContext differentiates from other interfaces.
	IsExp2MathFuncContext()
}

type Exp2MathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyExp2MathFuncContext() *Exp2MathFuncContext {
	var p = new(Exp2MathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_exp2MathFunc
	return p
}

func InitEmptyExp2MathFuncContext(p *Exp2MathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_exp2MathFunc
}

func (*Exp2MathFuncContext) IsExp2MathFuncContext() {}

func NewExp2MathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp2MathFuncContext {
	var p = new(Exp2MathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_exp2MathFunc

	return p
}

func (s *Exp2MathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp2MathFuncContext) GetNum() IExprContext { return s.num }

func (s *Exp2MathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *Exp2MathFuncContext) EXP2() antlr.TerminalNode {
	return s.GetToken(CalcliParserEXP2, 0)
}

func (s *Exp2MathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *Exp2MathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *Exp2MathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Exp2MathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp2MathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp2MathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterExp2MathFunc(s)
	}
}

func (s *Exp2MathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitExp2MathFunc(s)
	}
}

func (p *CalcliParser) Exp2MathFunc() (localctx IExp2MathFuncContext) {
	localctx = NewExp2MathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CalcliParserRULE_exp2MathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(CalcliParserEXP2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)

		var _x = p.expr(0)

		localctx.(*Exp2MathFuncContext).num = _x
	}
	{
		p.SetState(125)
		p.Match(CalcliParserRPAREN)
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

// IFloorMathFuncContext is an interface to support dynamic dispatch.
type IFloorMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	FLOOR() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsFloorMathFuncContext differentiates from other interfaces.
	IsFloorMathFuncContext()
}

type FloorMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyFloorMathFuncContext() *FloorMathFuncContext {
	var p = new(FloorMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_floorMathFunc
	return p
}

func InitEmptyFloorMathFuncContext(p *FloorMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_floorMathFunc
}

func (*FloorMathFuncContext) IsFloorMathFuncContext() {}

func NewFloorMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloorMathFuncContext {
	var p = new(FloorMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_floorMathFunc

	return p
}

func (s *FloorMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *FloorMathFuncContext) GetNum() IExprContext { return s.num }

func (s *FloorMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *FloorMathFuncContext) FLOOR() antlr.TerminalNode {
	return s.GetToken(CalcliParserFLOOR, 0)
}

func (s *FloorMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *FloorMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *FloorMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FloorMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloorMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterFloorMathFunc(s)
	}
}

func (s *FloorMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitFloorMathFunc(s)
	}
}

func (p *CalcliParser) FloorMathFunc() (localctx IFloorMathFuncContext) {
	localctx = NewFloorMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CalcliParserRULE_floorMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(CalcliParserFLOOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)

		var _x = p.expr(0)

		localctx.(*FloorMathFuncContext).num = _x
	}
	{
		p.SetState(130)
		p.Match(CalcliParserRPAREN)
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

// ILnMathFuncContext is an interface to support dynamic dispatch.
type ILnMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	LN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsLnMathFuncContext differentiates from other interfaces.
	IsLnMathFuncContext()
}

type LnMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyLnMathFuncContext() *LnMathFuncContext {
	var p = new(LnMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_lnMathFunc
	return p
}

func InitEmptyLnMathFuncContext(p *LnMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_lnMathFunc
}

func (*LnMathFuncContext) IsLnMathFuncContext() {}

func NewLnMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LnMathFuncContext {
	var p = new(LnMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_lnMathFunc

	return p
}

func (s *LnMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *LnMathFuncContext) GetNum() IExprContext { return s.num }

func (s *LnMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *LnMathFuncContext) LN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLN, 0)
}

func (s *LnMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *LnMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *LnMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LnMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LnMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LnMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterLnMathFunc(s)
	}
}

func (s *LnMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitLnMathFunc(s)
	}
}

func (p *CalcliParser) LnMathFunc() (localctx ILnMathFuncContext) {
	localctx = NewLnMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CalcliParserRULE_lnMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(CalcliParserLN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)

		var _x = p.expr(0)

		localctx.(*LnMathFuncContext).num = _x
	}
	{
		p.SetState(135)
		p.Match(CalcliParserRPAREN)
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

// ILogMathFuncContext is an interface to support dynamic dispatch.
type ILogMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// GetBase returns the base rule contexts.
	GetBase() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// SetBase sets the base rule contexts.
	SetBase(IExprContext)

	// Getter signatures
	LOG() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsLogMathFuncContext differentiates from other interfaces.
	IsLogMathFuncContext()
}

type LogMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
	base   IExprContext
}

func NewEmptyLogMathFuncContext() *LogMathFuncContext {
	var p = new(LogMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_logMathFunc
	return p
}

func InitEmptyLogMathFuncContext(p *LogMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_logMathFunc
}

func (*LogMathFuncContext) IsLogMathFuncContext() {}

func NewLogMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogMathFuncContext {
	var p = new(LogMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_logMathFunc

	return p
}

func (s *LogMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *LogMathFuncContext) GetNum() IExprContext { return s.num }

func (s *LogMathFuncContext) GetBase() IExprContext { return s.base }

func (s *LogMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *LogMathFuncContext) SetBase(v IExprContext) { s.base = v }

func (s *LogMathFuncContext) LOG() antlr.TerminalNode {
	return s.GetToken(CalcliParserLOG, 0)
}

func (s *LogMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *LogMathFuncContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CalcliParserCOMMA, 0)
}

func (s *LogMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *LogMathFuncContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LogMathFuncContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *LogMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterLogMathFunc(s)
	}
}

func (s *LogMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitLogMathFunc(s)
	}
}

func (p *CalcliParser) LogMathFunc() (localctx ILogMathFuncContext) {
	localctx = NewLogMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CalcliParserRULE_logMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Match(CalcliParserLOG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)

		var _x = p.expr(0)

		localctx.(*LogMathFuncContext).num = _x
	}
	{
		p.SetState(140)
		p.Match(CalcliParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)

		var _x = p.expr(0)

		localctx.(*LogMathFuncContext).base = _x
	}
	{
		p.SetState(142)
		p.Match(CalcliParserRPAREN)
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

// ILog2MathFuncContext is an interface to support dynamic dispatch.
type ILog2MathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	LOG2() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsLog2MathFuncContext differentiates from other interfaces.
	IsLog2MathFuncContext()
}

type Log2MathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyLog2MathFuncContext() *Log2MathFuncContext {
	var p = new(Log2MathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_log2MathFunc
	return p
}

func InitEmptyLog2MathFuncContext(p *Log2MathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_log2MathFunc
}

func (*Log2MathFuncContext) IsLog2MathFuncContext() {}

func NewLog2MathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Log2MathFuncContext {
	var p = new(Log2MathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_log2MathFunc

	return p
}

func (s *Log2MathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *Log2MathFuncContext) GetNum() IExprContext { return s.num }

func (s *Log2MathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *Log2MathFuncContext) LOG2() antlr.TerminalNode {
	return s.GetToken(CalcliParserLOG2, 0)
}

func (s *Log2MathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *Log2MathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *Log2MathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Log2MathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log2MathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Log2MathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterLog2MathFunc(s)
	}
}

func (s *Log2MathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitLog2MathFunc(s)
	}
}

func (p *CalcliParser) Log2MathFunc() (localctx ILog2MathFuncContext) {
	localctx = NewLog2MathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CalcliParserRULE_log2MathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(CalcliParserLOG2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)

		var _x = p.expr(0)

		localctx.(*Log2MathFuncContext).num = _x
	}
	{
		p.SetState(147)
		p.Match(CalcliParserRPAREN)
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

// ILog10MathFuncContext is an interface to support dynamic dispatch.
type ILog10MathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	LOG10() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsLog10MathFuncContext differentiates from other interfaces.
	IsLog10MathFuncContext()
}

type Log10MathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyLog10MathFuncContext() *Log10MathFuncContext {
	var p = new(Log10MathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_log10MathFunc
	return p
}

func InitEmptyLog10MathFuncContext(p *Log10MathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_log10MathFunc
}

func (*Log10MathFuncContext) IsLog10MathFuncContext() {}

func NewLog10MathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Log10MathFuncContext {
	var p = new(Log10MathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_log10MathFunc

	return p
}

func (s *Log10MathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *Log10MathFuncContext) GetNum() IExprContext { return s.num }

func (s *Log10MathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *Log10MathFuncContext) LOG10() antlr.TerminalNode {
	return s.GetToken(CalcliParserLOG10, 0)
}

func (s *Log10MathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *Log10MathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *Log10MathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Log10MathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log10MathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Log10MathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterLog10MathFunc(s)
	}
}

func (s *Log10MathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitLog10MathFunc(s)
	}
}

func (p *CalcliParser) Log10MathFunc() (localctx ILog10MathFuncContext) {
	localctx = NewLog10MathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CalcliParserRULE_log10MathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(CalcliParserLOG10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)

		var _x = p.expr(0)

		localctx.(*Log10MathFuncContext).num = _x
	}
	{
		p.SetState(152)
		p.Match(CalcliParserRPAREN)
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

// IModMathFuncContext is an interface to support dynamic dispatch.
type IModMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Getter signatures
	MOD() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsModMathFuncContext differentiates from other interfaces.
	IsModMathFuncContext()
}

type ModMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IExprContext
	right  IExprContext
}

func NewEmptyModMathFuncContext() *ModMathFuncContext {
	var p = new(ModMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_modMathFunc
	return p
}

func InitEmptyModMathFuncContext(p *ModMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_modMathFunc
}

func (*ModMathFuncContext) IsModMathFuncContext() {}

func NewModMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModMathFuncContext {
	var p = new(ModMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_modMathFunc

	return p
}

func (s *ModMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ModMathFuncContext) GetLeft() IExprContext { return s.left }

func (s *ModMathFuncContext) GetRight() IExprContext { return s.right }

func (s *ModMathFuncContext) SetLeft(v IExprContext) { s.left = v }

func (s *ModMathFuncContext) SetRight(v IExprContext) { s.right = v }

func (s *ModMathFuncContext) MOD() antlr.TerminalNode {
	return s.GetToken(CalcliParserMOD, 0)
}

func (s *ModMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *ModMathFuncContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CalcliParserCOMMA, 0)
}

func (s *ModMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *ModMathFuncContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ModMathFuncContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ModMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterModMathFunc(s)
	}
}

func (s *ModMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitModMathFunc(s)
	}
}

func (p *CalcliParser) ModMathFunc() (localctx IModMathFuncContext) {
	localctx = NewModMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CalcliParserRULE_modMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(CalcliParserMOD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)

		var _x = p.expr(0)

		localctx.(*ModMathFuncContext).left = _x
	}
	{
		p.SetState(157)
		p.Match(CalcliParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)

		var _x = p.expr(0)

		localctx.(*ModMathFuncContext).right = _x
	}
	{
		p.SetState(159)
		p.Match(CalcliParserRPAREN)
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

// INrtMathFuncContext is an interface to support dynamic dispatch.
type INrtMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// GetDegree returns the degree rule contexts.
	GetDegree() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// SetDegree sets the degree rule contexts.
	SetDegree(IExprContext)

	// Getter signatures
	NRT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsNrtMathFuncContext differentiates from other interfaces.
	IsNrtMathFuncContext()
}

type NrtMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
	degree IExprContext
}

func NewEmptyNrtMathFuncContext() *NrtMathFuncContext {
	var p = new(NrtMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_nrtMathFunc
	return p
}

func InitEmptyNrtMathFuncContext(p *NrtMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_nrtMathFunc
}

func (*NrtMathFuncContext) IsNrtMathFuncContext() {}

func NewNrtMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NrtMathFuncContext {
	var p = new(NrtMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_nrtMathFunc

	return p
}

func (s *NrtMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *NrtMathFuncContext) GetNum() IExprContext { return s.num }

func (s *NrtMathFuncContext) GetDegree() IExprContext { return s.degree }

func (s *NrtMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *NrtMathFuncContext) SetDegree(v IExprContext) { s.degree = v }

func (s *NrtMathFuncContext) NRT() antlr.TerminalNode {
	return s.GetToken(CalcliParserNRT, 0)
}

func (s *NrtMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *NrtMathFuncContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CalcliParserCOMMA, 0)
}

func (s *NrtMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *NrtMathFuncContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *NrtMathFuncContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *NrtMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NrtMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NrtMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterNrtMathFunc(s)
	}
}

func (s *NrtMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitNrtMathFunc(s)
	}
}

func (p *CalcliParser) NrtMathFunc() (localctx INrtMathFuncContext) {
	localctx = NewNrtMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CalcliParserRULE_nrtMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(CalcliParserNRT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)

		var _x = p.expr(0)

		localctx.(*NrtMathFuncContext).num = _x
	}
	{
		p.SetState(164)
		p.Match(CalcliParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)

		var _x = p.expr(0)

		localctx.(*NrtMathFuncContext).degree = _x
	}
	{
		p.SetState(166)
		p.Match(CalcliParserRPAREN)
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

// IPercentMathFuncContext is an interface to support dynamic dispatch.
type IPercentMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// GetFrom returns the from rule contexts.
	GetFrom() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// SetFrom sets the from rule contexts.
	SetFrom(IExprContext)

	// Getter signatures
	PERCENTAGE() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsPercentMathFuncContext differentiates from other interfaces.
	IsPercentMathFuncContext()
}

type PercentMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
	from   IExprContext
}

func NewEmptyPercentMathFuncContext() *PercentMathFuncContext {
	var p = new(PercentMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_percentMathFunc
	return p
}

func InitEmptyPercentMathFuncContext(p *PercentMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_percentMathFunc
}

func (*PercentMathFuncContext) IsPercentMathFuncContext() {}

func NewPercentMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PercentMathFuncContext {
	var p = new(PercentMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_percentMathFunc

	return p
}

func (s *PercentMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *PercentMathFuncContext) GetNum() IExprContext { return s.num }

func (s *PercentMathFuncContext) GetFrom() IExprContext { return s.from }

func (s *PercentMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *PercentMathFuncContext) SetFrom(v IExprContext) { s.from = v }

func (s *PercentMathFuncContext) PERCENTAGE() antlr.TerminalNode {
	return s.GetToken(CalcliParserPERCENTAGE, 0)
}

func (s *PercentMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *PercentMathFuncContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CalcliParserCOMMA, 0)
}

func (s *PercentMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *PercentMathFuncContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PercentMathFuncContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *PercentMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PercentMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterPercentMathFunc(s)
	}
}

func (s *PercentMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitPercentMathFunc(s)
	}
}

func (p *CalcliParser) PercentMathFunc() (localctx IPercentMathFuncContext) {
	localctx = NewPercentMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CalcliParserRULE_percentMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(CalcliParserPERCENTAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)

		var _x = p.expr(0)

		localctx.(*PercentMathFuncContext).num = _x
	}
	{
		p.SetState(171)
		p.Match(CalcliParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)

		var _x = p.expr(0)

		localctx.(*PercentMathFuncContext).from = _x
	}
	{
		p.SetState(173)
		p.Match(CalcliParserRPAREN)
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

// IRoundMathFuncContext is an interface to support dynamic dispatch.
type IRoundMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	ROUND() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsRoundMathFuncContext differentiates from other interfaces.
	IsRoundMathFuncContext()
}

type RoundMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyRoundMathFuncContext() *RoundMathFuncContext {
	var p = new(RoundMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_roundMathFunc
	return p
}

func InitEmptyRoundMathFuncContext(p *RoundMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_roundMathFunc
}

func (*RoundMathFuncContext) IsRoundMathFuncContext() {}

func NewRoundMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoundMathFuncContext {
	var p = new(RoundMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_roundMathFunc

	return p
}

func (s *RoundMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *RoundMathFuncContext) GetNum() IExprContext { return s.num }

func (s *RoundMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *RoundMathFuncContext) ROUND() antlr.TerminalNode {
	return s.GetToken(CalcliParserROUND, 0)
}

func (s *RoundMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *RoundMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *RoundMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RoundMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoundMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterRoundMathFunc(s)
	}
}

func (s *RoundMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitRoundMathFunc(s)
	}
}

func (p *CalcliParser) RoundMathFunc() (localctx IRoundMathFuncContext) {
	localctx = NewRoundMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CalcliParserRULE_roundMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(CalcliParserROUND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)

		var _x = p.expr(0)

		localctx.(*RoundMathFuncContext).num = _x
	}
	{
		p.SetState(178)
		p.Match(CalcliParserRPAREN)
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

// ISinMathFuncContext is an interface to support dynamic dispatch.
type ISinMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	SIN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsSinMathFuncContext differentiates from other interfaces.
	IsSinMathFuncContext()
}

type SinMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptySinMathFuncContext() *SinMathFuncContext {
	var p = new(SinMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_sinMathFunc
	return p
}

func InitEmptySinMathFuncContext(p *SinMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_sinMathFunc
}

func (*SinMathFuncContext) IsSinMathFuncContext() {}

func NewSinMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinMathFuncContext {
	var p = new(SinMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_sinMathFunc

	return p
}

func (s *SinMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *SinMathFuncContext) GetNum() IExprContext { return s.num }

func (s *SinMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *SinMathFuncContext) SIN() antlr.TerminalNode {
	return s.GetToken(CalcliParserSIN, 0)
}

func (s *SinMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *SinMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *SinMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SinMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SinMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterSinMathFunc(s)
	}
}

func (s *SinMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitSinMathFunc(s)
	}
}

func (p *CalcliParser) SinMathFunc() (localctx ISinMathFuncContext) {
	localctx = NewSinMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CalcliParserRULE_sinMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(CalcliParserSIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)

		var _x = p.expr(0)

		localctx.(*SinMathFuncContext).num = _x
	}
	{
		p.SetState(183)
		p.Match(CalcliParserRPAREN)
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

// ISqrtMathFuncContext is an interface to support dynamic dispatch.
type ISqrtMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	SQRT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsSqrtMathFuncContext differentiates from other interfaces.
	IsSqrtMathFuncContext()
}

type SqrtMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptySqrtMathFuncContext() *SqrtMathFuncContext {
	var p = new(SqrtMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_sqrtMathFunc
	return p
}

func InitEmptySqrtMathFuncContext(p *SqrtMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_sqrtMathFunc
}

func (*SqrtMathFuncContext) IsSqrtMathFuncContext() {}

func NewSqrtMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SqrtMathFuncContext {
	var p = new(SqrtMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_sqrtMathFunc

	return p
}

func (s *SqrtMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *SqrtMathFuncContext) GetNum() IExprContext { return s.num }

func (s *SqrtMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *SqrtMathFuncContext) SQRT() antlr.TerminalNode {
	return s.GetToken(CalcliParserSQRT, 0)
}

func (s *SqrtMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *SqrtMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *SqrtMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SqrtMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrtMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SqrtMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterSqrtMathFunc(s)
	}
}

func (s *SqrtMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitSqrtMathFunc(s)
	}
}

func (p *CalcliParser) SqrtMathFunc() (localctx ISqrtMathFuncContext) {
	localctx = NewSqrtMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CalcliParserRULE_sqrtMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(CalcliParserSQRT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)

		var _x = p.expr(0)

		localctx.(*SqrtMathFuncContext).num = _x
	}
	{
		p.SetState(188)
		p.Match(CalcliParserRPAREN)
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

// ITanMathFuncContext is an interface to support dynamic dispatch.
type ITanMathFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num rule contexts.
	GetNum() IExprContext

	// SetNum sets the num rule contexts.
	SetNum(IExprContext)

	// Getter signatures
	TAN() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext

	// IsTanMathFuncContext differentiates from other interfaces.
	IsTanMathFuncContext()
}

type TanMathFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    IExprContext
}

func NewEmptyTanMathFuncContext() *TanMathFuncContext {
	var p = new(TanMathFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_tanMathFunc
	return p
}

func InitEmptyTanMathFuncContext(p *TanMathFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_tanMathFunc
}

func (*TanMathFuncContext) IsTanMathFuncContext() {}

func NewTanMathFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TanMathFuncContext {
	var p = new(TanMathFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_tanMathFunc

	return p
}

func (s *TanMathFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *TanMathFuncContext) GetNum() IExprContext { return s.num }

func (s *TanMathFuncContext) SetNum(v IExprContext) { s.num = v }

func (s *TanMathFuncContext) TAN() antlr.TerminalNode {
	return s.GetToken(CalcliParserTAN, 0)
}

func (s *TanMathFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *TanMathFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *TanMathFuncContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TanMathFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanMathFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TanMathFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterTanMathFunc(s)
	}
}

func (s *TanMathFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitTanMathFunc(s)
	}
}

func (p *CalcliParser) TanMathFunc() (localctx ITanMathFuncContext) {
	localctx = NewTanMathFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CalcliParserRULE_tanMathFunc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(CalcliParserTAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Match(CalcliParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)

		var _x = p.expr(0)

		localctx.(*TanMathFuncContext).num = _x
	}
	{
		p.SetState(193)
		p.Match(CalcliParserRPAREN)
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

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcliParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *CalcliParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CalcliParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.expr(0)
	}
	{
		p.SetState(196)
		p.Match(CalcliParserEOF)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcliParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcliParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FactorialContext struct {
	ExprContext
	num IExprContext
	op  antlr.Token
}

func NewFactorialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FactorialContext {
	var p = new(FactorialContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FactorialContext) GetOp() antlr.Token { return s.op }

func (s *FactorialContext) SetOp(v antlr.Token) { s.op = v }

func (s *FactorialContext) GetNum() IExprContext { return s.num }

func (s *FactorialContext) SetNum(v IExprContext) { s.num = v }

func (s *FactorialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorialContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FactorialContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(CalcliParserEXCLAMATION, 0)
}

func (s *FactorialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterFactorial(s)
	}
}

func (s *FactorialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitFactorial(s)
	}
}

type MulDivModContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModContext {
	var p = new(MulDivModContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivModContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModContext) GetLeft() IExprContext { return s.left }

func (s *MulDivModContext) GetRight() IExprContext { return s.right }

func (s *MulDivModContext) SetLeft(v IExprContext) { s.left = v }

func (s *MulDivModContext) SetRight(v IExprContext) { s.right = v }

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivModContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *MulDivModContext) STAR() antlr.TerminalNode {
	return s.GetToken(CalcliParserSTAR, 0)
}

func (s *MulDivModContext) SLASH() antlr.TerminalNode {
	return s.GetToken(CalcliParserSLASH, 0)
}

func (s *MulDivModContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(CalcliParserPERCENT, 0)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

type NegationContext struct {
	ExprContext
	right IExprContext
}

func NewNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationContext {
	var p = new(NegationContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NegationContext) GetRight() IExprContext { return s.right }

func (s *NegationContext) SetRight(v IExprContext) { s.right = v }

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CalcliParserMINUS, 0)
}

func (s *NegationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitNegation(s)
	}
}

type MathFunctionContext struct {
	ExprContext
}

func NewMathFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MathFunctionContext {
	var p = new(MathFunctionContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MathFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathFunctionContext) MathFunc() IMathFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathFuncContext)
}

func (s *MathFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterMathFunction(s)
	}
}

func (s *MathFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitMathFunction(s)
	}
}

type AddSubContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetLeft() IExprContext { return s.left }

func (s *AddSubContext) GetRight() IExprContext { return s.right }

func (s *AddSubContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddSubContext) SetRight(v IExprContext) { s.right = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AddSubContext) PLUS() antlr.TerminalNode {
	return s.GetToken(CalcliParserPLUS, 0)
}

func (s *AddSubContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CalcliParserMINUS, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type ParensContext struct {
	ExprContext
	inner IExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetInner() IExprContext { return s.inner }

func (s *ParensContext) SetInner(v IExprContext) { s.inner = v }

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserLPAREN, 0)
}

func (s *ParensContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CalcliParserRPAREN, 0)
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitParens(s)
	}
}

type UnitValueContext struct {
	ExprContext
}

func NewUnitValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnitValueContext {
	var p = new(UnitValueContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnitValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitValueContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *UnitValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterUnitValue(s)
	}
}

func (s *UnitValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitUnitValue(s)
	}
}

type PowContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowContext {
	var p = new(PowContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PowContext) GetOp() antlr.Token { return s.op }

func (s *PowContext) SetOp(v antlr.Token) { s.op = v }

func (s *PowContext) GetLeft() IExprContext { return s.left }

func (s *PowContext) GetRight() IExprContext { return s.right }

func (s *PowContext) SetLeft(v IExprContext) { s.left = v }

func (s *PowContext) SetRight(v IExprContext) { s.right = v }

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PowContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *PowContext) POW() antlr.TerminalNode {
	return s.GetToken(CalcliParserPOW, 0)
}

func (s *PowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.EnterPow(s)
	}
}

func (s *PowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcliListener); ok {
		listenerT.ExitPow(s)
	}
}

func (p *CalcliParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CalcliParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, CalcliParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcliParserEULER, CalcliParserINT, CalcliParserFLOAT, CalcliParserPI, CalcliParserPREVIOS:
		localctx = NewUnitValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(199)
			p.Unit()
		}

	case CalcliParserMINUS:
		localctx = NewNegationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(200)
			p.Match(CalcliParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)

			var _x = p.expr(7)

			localctx.(*NegationContext).right = _x
		}

	case CalcliParserLPAREN:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(202)
			p.Match(CalcliParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)

			var _x = p.expr(0)

			localctx.(*ParensContext).inner = _x
		}
		{
			p.SetState(204)
			p.Match(CalcliParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalcliParserABS, CalcliParserACOS, CalcliParserASIN, CalcliParserATAN, CalcliParserCEIL, CalcliParserCOS, CalcliParserEXP, CalcliParserEXP2, CalcliParserFLOOR, CalcliParserLN, CalcliParserLOG, CalcliParserLOG2, CalcliParserLOG10, CalcliParserMOD, CalcliParserNRT, CalcliParserPERCENTAGE, CalcliParserROUND, CalcliParserSIN, CalcliParserSQRT, CalcliParserTAN:
		localctx = NewMathFunctionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)
			p.MathFunc()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*PowContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CalcliParserRULE_expr)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(210)

					var _m = p.Match(CalcliParserPOW)

					localctx.(*PowContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(211)

					var _x = p.expr(4)

					localctx.(*PowContext).right = _x
				}

			case 2:
				localctx = NewMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulDivModContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CalcliParserRULE_expr)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(213)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&25786580992) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(214)

					var _x = p.expr(3)

					localctx.(*MulDivModContext).right = _x
				}

			case 3:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddSubContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CalcliParserRULE_expr)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(216)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcliParserMINUS || _la == CalcliParserPLUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(217)

					var _x = p.expr(2)

					localctx.(*AddSubContext).right = _x
				}

			case 4:
				localctx = NewFactorialContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*FactorialContext).num = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CalcliParserRULE_expr)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(219)

					var _m = p.Match(CalcliParserEXCLAMATION)

					localctx.(*FactorialContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *CalcliParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 26:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcliParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
