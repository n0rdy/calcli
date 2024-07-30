// Code generated from Calcli.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calcli

import "github.com/antlr4-go/antlr/v4"

// CalcliListener is a complete listener for a parse tree produced by CalcliParser.
type CalcliListener interface {
	antlr.ParseTreeListener

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterMathFunc is called when entering the mathFunc production.
	EnterMathFunc(c *MathFuncContext)

	// EnterAbsMathFunc is called when entering the absMathFunc production.
	EnterAbsMathFunc(c *AbsMathFuncContext)

	// EnterAcosMathFunc is called when entering the acosMathFunc production.
	EnterAcosMathFunc(c *AcosMathFuncContext)

	// EnterAsinMathFunc is called when entering the asinMathFunc production.
	EnterAsinMathFunc(c *AsinMathFuncContext)

	// EnterAtanMathFunc is called when entering the atanMathFunc production.
	EnterAtanMathFunc(c *AtanMathFuncContext)

	// EnterCeilMathFunc is called when entering the ceilMathFunc production.
	EnterCeilMathFunc(c *CeilMathFuncContext)

	// EnterCosMathFunc is called when entering the cosMathFunc production.
	EnterCosMathFunc(c *CosMathFuncContext)

	// EnterExpMathFunc is called when entering the expMathFunc production.
	EnterExpMathFunc(c *ExpMathFuncContext)

	// EnterExp2MathFunc is called when entering the exp2MathFunc production.
	EnterExp2MathFunc(c *Exp2MathFuncContext)

	// EnterFloorMathFunc is called when entering the floorMathFunc production.
	EnterFloorMathFunc(c *FloorMathFuncContext)

	// EnterLnMathFunc is called when entering the lnMathFunc production.
	EnterLnMathFunc(c *LnMathFuncContext)

	// EnterLogMathFunc is called when entering the logMathFunc production.
	EnterLogMathFunc(c *LogMathFuncContext)

	// EnterLog2MathFunc is called when entering the log2MathFunc production.
	EnterLog2MathFunc(c *Log2MathFuncContext)

	// EnterLog10MathFunc is called when entering the log10MathFunc production.
	EnterLog10MathFunc(c *Log10MathFuncContext)

	// EnterModMathFunc is called when entering the modMathFunc production.
	EnterModMathFunc(c *ModMathFuncContext)

	// EnterNrtMathFunc is called when entering the nrtMathFunc production.
	EnterNrtMathFunc(c *NrtMathFuncContext)

	// EnterPercentMathFunc is called when entering the percentMathFunc production.
	EnterPercentMathFunc(c *PercentMathFuncContext)

	// EnterRoundMathFunc is called when entering the roundMathFunc production.
	EnterRoundMathFunc(c *RoundMathFuncContext)

	// EnterSinMathFunc is called when entering the sinMathFunc production.
	EnterSinMathFunc(c *SinMathFuncContext)

	// EnterSqrtMathFunc is called when entering the sqrtMathFunc production.
	EnterSqrtMathFunc(c *SqrtMathFuncContext)

	// EnterTanMathFunc is called when entering the tanMathFunc production.
	EnterTanMathFunc(c *TanMathFuncContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterFactorial is called when entering the Factorial production.
	EnterFactorial(c *FactorialContext)

	// EnterMulDivMod is called when entering the MulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterNegation is called when entering the Negation production.
	EnterNegation(c *NegationContext)

	// EnterMathFunction is called when entering the MathFunction production.
	EnterMathFunction(c *MathFunctionContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterUnitValue is called when entering the UnitValue production.
	EnterUnitValue(c *UnitValueContext)

	// EnterPow is called when entering the Pow production.
	EnterPow(c *PowContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitMathFunc is called when exiting the mathFunc production.
	ExitMathFunc(c *MathFuncContext)

	// ExitAbsMathFunc is called when exiting the absMathFunc production.
	ExitAbsMathFunc(c *AbsMathFuncContext)

	// ExitAcosMathFunc is called when exiting the acosMathFunc production.
	ExitAcosMathFunc(c *AcosMathFuncContext)

	// ExitAsinMathFunc is called when exiting the asinMathFunc production.
	ExitAsinMathFunc(c *AsinMathFuncContext)

	// ExitAtanMathFunc is called when exiting the atanMathFunc production.
	ExitAtanMathFunc(c *AtanMathFuncContext)

	// ExitCeilMathFunc is called when exiting the ceilMathFunc production.
	ExitCeilMathFunc(c *CeilMathFuncContext)

	// ExitCosMathFunc is called when exiting the cosMathFunc production.
	ExitCosMathFunc(c *CosMathFuncContext)

	// ExitExpMathFunc is called when exiting the expMathFunc production.
	ExitExpMathFunc(c *ExpMathFuncContext)

	// ExitExp2MathFunc is called when exiting the exp2MathFunc production.
	ExitExp2MathFunc(c *Exp2MathFuncContext)

	// ExitFloorMathFunc is called when exiting the floorMathFunc production.
	ExitFloorMathFunc(c *FloorMathFuncContext)

	// ExitLnMathFunc is called when exiting the lnMathFunc production.
	ExitLnMathFunc(c *LnMathFuncContext)

	// ExitLogMathFunc is called when exiting the logMathFunc production.
	ExitLogMathFunc(c *LogMathFuncContext)

	// ExitLog2MathFunc is called when exiting the log2MathFunc production.
	ExitLog2MathFunc(c *Log2MathFuncContext)

	// ExitLog10MathFunc is called when exiting the log10MathFunc production.
	ExitLog10MathFunc(c *Log10MathFuncContext)

	// ExitModMathFunc is called when exiting the modMathFunc production.
	ExitModMathFunc(c *ModMathFuncContext)

	// ExitNrtMathFunc is called when exiting the nrtMathFunc production.
	ExitNrtMathFunc(c *NrtMathFuncContext)

	// ExitPercentMathFunc is called when exiting the percentMathFunc production.
	ExitPercentMathFunc(c *PercentMathFuncContext)

	// ExitRoundMathFunc is called when exiting the roundMathFunc production.
	ExitRoundMathFunc(c *RoundMathFuncContext)

	// ExitSinMathFunc is called when exiting the sinMathFunc production.
	ExitSinMathFunc(c *SinMathFuncContext)

	// ExitSqrtMathFunc is called when exiting the sqrtMathFunc production.
	ExitSqrtMathFunc(c *SqrtMathFuncContext)

	// ExitTanMathFunc is called when exiting the tanMathFunc production.
	ExitTanMathFunc(c *TanMathFuncContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitFactorial is called when exiting the Factorial production.
	ExitFactorial(c *FactorialContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitNegation is called when exiting the Negation production.
	ExitNegation(c *NegationContext)

	// ExitMathFunction is called when exiting the MathFunction production.
	ExitMathFunction(c *MathFunctionContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitUnitValue is called when exiting the UnitValue production.
	ExitUnitValue(c *UnitValueContext)

	// ExitPow is called when exiting the Pow production.
	ExitPow(c *PowContext)
}
