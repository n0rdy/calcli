// Code generated from Calcli.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calcli

import "github.com/antlr4-go/antlr/v4"

// BaseCalcliListener is a complete listener for a parse tree produced by CalcliParser.
type BaseCalcliListener struct{}

var _ CalcliListener = &BaseCalcliListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcliListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcliListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcliListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcliListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseCalcliListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseCalcliListener) ExitUnit(ctx *UnitContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCalcliListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCalcliListener) ExitConstant(ctx *ConstantContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseCalcliListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseCalcliListener) ExitNumber(ctx *NumberContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseCalcliListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseCalcliListener) ExitVariable(ctx *VariableContext) {}

// EnterMathFunc is called when production mathFunc is entered.
func (s *BaseCalcliListener) EnterMathFunc(ctx *MathFuncContext) {}

// ExitMathFunc is called when production mathFunc is exited.
func (s *BaseCalcliListener) ExitMathFunc(ctx *MathFuncContext) {}

// EnterAbsMathFunc is called when production absMathFunc is entered.
func (s *BaseCalcliListener) EnterAbsMathFunc(ctx *AbsMathFuncContext) {}

// ExitAbsMathFunc is called when production absMathFunc is exited.
func (s *BaseCalcliListener) ExitAbsMathFunc(ctx *AbsMathFuncContext) {}

// EnterAcosMathFunc is called when production acosMathFunc is entered.
func (s *BaseCalcliListener) EnterAcosMathFunc(ctx *AcosMathFuncContext) {}

// ExitAcosMathFunc is called when production acosMathFunc is exited.
func (s *BaseCalcliListener) ExitAcosMathFunc(ctx *AcosMathFuncContext) {}

// EnterAsinMathFunc is called when production asinMathFunc is entered.
func (s *BaseCalcliListener) EnterAsinMathFunc(ctx *AsinMathFuncContext) {}

// ExitAsinMathFunc is called when production asinMathFunc is exited.
func (s *BaseCalcliListener) ExitAsinMathFunc(ctx *AsinMathFuncContext) {}

// EnterAtanMathFunc is called when production atanMathFunc is entered.
func (s *BaseCalcliListener) EnterAtanMathFunc(ctx *AtanMathFuncContext) {}

// ExitAtanMathFunc is called when production atanMathFunc is exited.
func (s *BaseCalcliListener) ExitAtanMathFunc(ctx *AtanMathFuncContext) {}

// EnterCeilMathFunc is called when production ceilMathFunc is entered.
func (s *BaseCalcliListener) EnterCeilMathFunc(ctx *CeilMathFuncContext) {}

// ExitCeilMathFunc is called when production ceilMathFunc is exited.
func (s *BaseCalcliListener) ExitCeilMathFunc(ctx *CeilMathFuncContext) {}

// EnterCosMathFunc is called when production cosMathFunc is entered.
func (s *BaseCalcliListener) EnterCosMathFunc(ctx *CosMathFuncContext) {}

// ExitCosMathFunc is called when production cosMathFunc is exited.
func (s *BaseCalcliListener) ExitCosMathFunc(ctx *CosMathFuncContext) {}

// EnterExpMathFunc is called when production expMathFunc is entered.
func (s *BaseCalcliListener) EnterExpMathFunc(ctx *ExpMathFuncContext) {}

// ExitExpMathFunc is called when production expMathFunc is exited.
func (s *BaseCalcliListener) ExitExpMathFunc(ctx *ExpMathFuncContext) {}

// EnterExp2MathFunc is called when production exp2MathFunc is entered.
func (s *BaseCalcliListener) EnterExp2MathFunc(ctx *Exp2MathFuncContext) {}

// ExitExp2MathFunc is called when production exp2MathFunc is exited.
func (s *BaseCalcliListener) ExitExp2MathFunc(ctx *Exp2MathFuncContext) {}

// EnterFloorMathFunc is called when production floorMathFunc is entered.
func (s *BaseCalcliListener) EnterFloorMathFunc(ctx *FloorMathFuncContext) {}

// ExitFloorMathFunc is called when production floorMathFunc is exited.
func (s *BaseCalcliListener) ExitFloorMathFunc(ctx *FloorMathFuncContext) {}

// EnterLnMathFunc is called when production lnMathFunc is entered.
func (s *BaseCalcliListener) EnterLnMathFunc(ctx *LnMathFuncContext) {}

// ExitLnMathFunc is called when production lnMathFunc is exited.
func (s *BaseCalcliListener) ExitLnMathFunc(ctx *LnMathFuncContext) {}

// EnterLogMathFunc is called when production logMathFunc is entered.
func (s *BaseCalcliListener) EnterLogMathFunc(ctx *LogMathFuncContext) {}

// ExitLogMathFunc is called when production logMathFunc is exited.
func (s *BaseCalcliListener) ExitLogMathFunc(ctx *LogMathFuncContext) {}

// EnterLog2MathFunc is called when production log2MathFunc is entered.
func (s *BaseCalcliListener) EnterLog2MathFunc(ctx *Log2MathFuncContext) {}

// ExitLog2MathFunc is called when production log2MathFunc is exited.
func (s *BaseCalcliListener) ExitLog2MathFunc(ctx *Log2MathFuncContext) {}

// EnterLog10MathFunc is called when production log10MathFunc is entered.
func (s *BaseCalcliListener) EnterLog10MathFunc(ctx *Log10MathFuncContext) {}

// ExitLog10MathFunc is called when production log10MathFunc is exited.
func (s *BaseCalcliListener) ExitLog10MathFunc(ctx *Log10MathFuncContext) {}

// EnterModMathFunc is called when production modMathFunc is entered.
func (s *BaseCalcliListener) EnterModMathFunc(ctx *ModMathFuncContext) {}

// ExitModMathFunc is called when production modMathFunc is exited.
func (s *BaseCalcliListener) ExitModMathFunc(ctx *ModMathFuncContext) {}

// EnterNrtMathFunc is called when production nrtMathFunc is entered.
func (s *BaseCalcliListener) EnterNrtMathFunc(ctx *NrtMathFuncContext) {}

// ExitNrtMathFunc is called when production nrtMathFunc is exited.
func (s *BaseCalcliListener) ExitNrtMathFunc(ctx *NrtMathFuncContext) {}

// EnterPercentMathFunc is called when production percentMathFunc is entered.
func (s *BaseCalcliListener) EnterPercentMathFunc(ctx *PercentMathFuncContext) {}

// ExitPercentMathFunc is called when production percentMathFunc is exited.
func (s *BaseCalcliListener) ExitPercentMathFunc(ctx *PercentMathFuncContext) {}

// EnterRoundMathFunc is called when production roundMathFunc is entered.
func (s *BaseCalcliListener) EnterRoundMathFunc(ctx *RoundMathFuncContext) {}

// ExitRoundMathFunc is called when production roundMathFunc is exited.
func (s *BaseCalcliListener) ExitRoundMathFunc(ctx *RoundMathFuncContext) {}

// EnterSinMathFunc is called when production sinMathFunc is entered.
func (s *BaseCalcliListener) EnterSinMathFunc(ctx *SinMathFuncContext) {}

// ExitSinMathFunc is called when production sinMathFunc is exited.
func (s *BaseCalcliListener) ExitSinMathFunc(ctx *SinMathFuncContext) {}

// EnterSqrtMathFunc is called when production sqrtMathFunc is entered.
func (s *BaseCalcliListener) EnterSqrtMathFunc(ctx *SqrtMathFuncContext) {}

// ExitSqrtMathFunc is called when production sqrtMathFunc is exited.
func (s *BaseCalcliListener) ExitSqrtMathFunc(ctx *SqrtMathFuncContext) {}

// EnterTanMathFunc is called when production tanMathFunc is entered.
func (s *BaseCalcliListener) EnterTanMathFunc(ctx *TanMathFuncContext) {}

// ExitTanMathFunc is called when production tanMathFunc is exited.
func (s *BaseCalcliListener) ExitTanMathFunc(ctx *TanMathFuncContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalcliListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalcliListener) ExitStart(ctx *StartContext) {}

// EnterFactorial is called when production Factorial is entered.
func (s *BaseCalcliListener) EnterFactorial(ctx *FactorialContext) {}

// ExitFactorial is called when production Factorial is exited.
func (s *BaseCalcliListener) ExitFactorial(ctx *FactorialContext) {}

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseCalcliListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseCalcliListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterNegation is called when production Negation is entered.
func (s *BaseCalcliListener) EnterNegation(ctx *NegationContext) {}

// ExitNegation is called when production Negation is exited.
func (s *BaseCalcliListener) ExitNegation(ctx *NegationContext) {}

// EnterMathFunction is called when production MathFunction is entered.
func (s *BaseCalcliListener) EnterMathFunction(ctx *MathFunctionContext) {}

// ExitMathFunction is called when production MathFunction is exited.
func (s *BaseCalcliListener) ExitMathFunction(ctx *MathFunctionContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalcliListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalcliListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseCalcliListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseCalcliListener) ExitParens(ctx *ParensContext) {}

// EnterUnitValue is called when production UnitValue is entered.
func (s *BaseCalcliListener) EnterUnitValue(ctx *UnitValueContext) {}

// ExitUnitValue is called when production UnitValue is exited.
func (s *BaseCalcliListener) ExitUnitValue(ctx *UnitValueContext) {}

// EnterPow is called when production Pow is entered.
func (s *BaseCalcliListener) EnterPow(ctx *PowContext) {}

// ExitPow is called when production Pow is exited.
func (s *BaseCalcliListener) ExitPow(ctx *PowContext) {}
