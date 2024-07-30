package parser

import (
	"math"
	genparser "n0rdy.foo/calcli/calc/parser/gen"
	"n0rdy.foo/calcli/calc/utils"
	"strconv"
)

const (
	pi = 3.141592653589793
	e  = 2.718281828459045
)

var (
	factorials = map[int]int{
		0:  1,
		1:  1,
		2:  2,
		3:  6,
		4:  24,
		5:  120,
		6:  720,
		7:  5040,
		8:  40320,
		9:  362880,
		10: 3628800,
		11: 39916800,
		12: 479001600,
	}
)

type CalcListener struct {
	*genparser.BaseCalcliListener

	stack utils.Stack
}

func NewCalcListener() *CalcListener {
	return &CalcListener{
		stack: utils.Stack{},
	}
}

func (cl *CalcListener) Result() float64 {
	return cl.stack.Pop()
}

func (cl *CalcListener) Reset() {
	cl.stack = utils.Stack{}
}

// ExitConstant is called when production constant is exited.
func (cl *CalcListener) ExitConstant(ctx *genparser.ConstantContext) {
	if ctx.PI() != nil {
		cl.stack.Push(pi)
	} else if ctx.EULER() != nil {
		cl.stack.Push(e)
	} else {
		panic("unknown constant: " + ctx.GetText())
	}
}

// ExitVariable is called when production variable is exited.
func (cl *CalcListener) ExitVariable(ctx *genparser.VariableContext) {
	if ctx.PREVIOS() != nil {
		cl.stack.Push(utils.PreviousResult)
	} else {
		panic("unknown variable: " + ctx.GetText())
	}
}

// ExitNumber is called when production number is exited.
func (cl *CalcListener) ExitNumber(ctx *genparser.NumberContext) {
	if ctx.INT() != nil {
		i, err := strconv.Atoi(ctx.GetText())
		if err != nil {
			panic("failed to parse int: " + err.Error())
		}
		// we convert int to float64 to be able to perform operations like multiplication and division on it
		// and get the float-like result rather than integer:
		// e.g.: 5 / 2 = 2.5, not 2
		cl.stack.Push(float64(i))
	} else if ctx.FLOAT() != nil {
		f, err := strconv.ParseFloat(ctx.GetText(), 64)
		if err != nil {
			panic("failed to parse float: " + err.Error())
		}
		cl.stack.Push(f)
	} else {
		panic("unknown number: " + ctx.GetText())
	}
}

// ExitNegation is called when production Negation is exited.
func (cl *CalcListener) ExitNegation(ctx *genparser.NegationContext) {
	val := cl.stack.Pop()
	cl.stack.Push(-val)
}

// ExitAbsMathFunc is called when production absMathFunc is exited.
func (cl *CalcListener) ExitAbsMathFunc(ctx *genparser.AbsMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Abs(num))
}

// ExitAcosMathFunc is called when production acosMathFunc is exited.
func (cl *CalcListener) ExitAcosMathFunc(ctx *genparser.AcosMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Acos(num))
}

// ExitAsinMathFunc is called when production asinMathFunc is exited.
func (cl *CalcListener) ExitAsinMathFunc(ctx *genparser.AsinMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Asin(num))
}

// ExitAtanMathFunc is called when production atanMathFunc is exited.
func (cl *CalcListener) ExitAtanMathFunc(ctx *genparser.AtanMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Atan(num))
}

// ExitCeilMathFunc is called when production ceilMathFunc is exited.
func (cl *CalcListener) ExitCeilMathFunc(ctx *genparser.CeilMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Ceil(num))
}

// ExitCosMathFunc is called when production cosMathFunc is exited.
func (cl *CalcListener) ExitCosMathFunc(ctx *genparser.CosMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Cos(num))
}

// ExitExpMathFunc is called when production expMathFunc is exited.
func (cl *CalcListener) ExitExpMathFunc(ctx *genparser.ExpMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Exp(num))
}

// ExitExp2MathFunc is called when production exp2MathFunc is exited.
func (cl *CalcListener) ExitExp2MathFunc(ctx *genparser.Exp2MathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Exp2(num))
}

// ExitFloorMathFunc is called when production floorMathFunc is exited.
func (cl *CalcListener) ExitFloorMathFunc(ctx *genparser.FloorMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Floor(num))
}

// ExitLnMathFunc is called when production lnMathFunc is exited.
func (cl *CalcListener) ExitLnMathFunc(ctx *genparser.LnMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Log(num))
}

// ExitLogMathFunc is called when production logMathFunc is exited.
func (cl *CalcListener) ExitLogMathFunc(ctx *genparser.LogMathFuncContext) {
	base := cl.stack.Pop()
	num := cl.stack.Pop()
	cl.stack.Push(math.Log(num) / math.Log(base))
}

// ExitLog2MathFunc is called when production log2MathFunc is exited.
func (cl *CalcListener) ExitLog2MathFunc(ctx *genparser.Log2MathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Log2(num))
}

// ExitLog10MathFunc is called when production log10MathFunc is exited.
func (cl *CalcListener) ExitLog10MathFunc(ctx *genparser.Log10MathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Log10(num))
}

// ExitModMathFunc is called when production modMathFunc is exited.
func (cl *CalcListener) ExitModMathFunc(ctx *genparser.ModMathFuncContext) {
	right := cl.stack.Pop()
	left := cl.stack.Pop()
	cl.stack.Push(math.Mod(left, right))
}

// ExitNrtMathFunc is called when production nrtMathFunc is exited.
func (cl *CalcListener) ExitNrtMathFunc(ctx *genparser.NrtMathFuncContext) {
	degree := cl.stack.Pop()
	num := cl.stack.Pop()
	cl.stack.Push(math.Pow(num, 1/degree))
}

// ExitPercentMathFunc is called when production percentMathFunc is exited.
func (cl *CalcListener) ExitPercentMathFunc(ctx *genparser.PercentMathFuncContext) {
	from := cl.stack.Pop()
	num := cl.stack.Pop()
	cl.stack.Push(num / from * 100)
}

// ExitRoundMathFunc is called when production roundMathFunc is exited.
func (cl *CalcListener) ExitRoundMathFunc(ctx *genparser.RoundMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Round(num))
}

// ExitSinMathFunc is called when production sinMathFunc is exited.
func (cl *CalcListener) ExitSinMathFunc(ctx *genparser.SinMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Sin(num))
}

// ExitSqrtMathFunc is called when production sqrtMathFunc is exited.
func (cl *CalcListener) ExitSqrtMathFunc(ctx *genparser.SqrtMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Sqrt(num))
}

// ExitTanMathFunc is called when production tanMathFunc is exited.
func (cl *CalcListener) ExitTanMathFunc(ctx *genparser.TanMathFuncContext) {
	num := cl.stack.Pop()
	cl.stack.Push(math.Tan(num))
}

// ExitFactorial is called when production Factorial is exited.
func (cl *CalcListener) ExitFactorial(ctx *genparser.FactorialContext) {
	num := cl.stack.Pop()
	numInt, floatValue := math.Modf(num)
	if floatValue != 0 {
		panic("factorial is only defined for integers")
	}
	if numInt < 0 {
		panic("factorial is only defined for non-negative integers")
	}

	res := cl.factorial(int(numInt))
	cl.stack.Push(float64(res))
}

// ExitPow is called when production Pow is exited.
func (cl *CalcListener) ExitPow(ctx *genparser.PowContext) {
	right := cl.stack.Pop()
	left := cl.stack.Pop()
	pow := math.Pow(left, right)
	cl.stack.Push(pow)
}

// ExitMulDivMod is called when production MulDivMod is exited.
func (cl *CalcListener) ExitMulDivMod(ctx *genparser.MulDivModContext) {
	right := cl.stack.Pop()
	left := cl.stack.Pop()

	switch ctx.GetOp().GetTokenType() {
	case genparser.CalcliLexerSTAR:
		cl.stack.Push(left * right)
	case genparser.CalcliLexerSLASH:
		cl.stack.Push(left / right)
	case genparser.CalcliLexerPERCENT:
		cl.stack.Push(math.Mod(left, right))
	default:
		panic("unknown operator: " + ctx.GetOp().GetText())
	}
}

// ExitAddSub is called when production AddSub is exited.
func (cl *CalcListener) ExitAddSub(ctx *genparser.AddSubContext) {
	right := cl.stack.Pop()
	left := cl.stack.Pop()

	switch ctx.GetOp().GetTokenType() {
	case genparser.CalcliLexerPLUS:
		cl.stack.Push(left + right)
	case genparser.CalcliLexerMINUS:
		cl.stack.Push(left - right)
	default:
		panic("unknown operator: " + ctx.GetOp().GetText())
	}
}

func (cl *CalcListener) factorial(n int) int {
	if factorials[n] != 0 {
		return factorials[n]
	}
	return n * cl.factorial(n-1)
}
