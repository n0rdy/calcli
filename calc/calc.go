package calc

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"n0rdy.foo/calcli/calc/parser"
	genparser "n0rdy.foo/calcli/calc/parser/gen"
)

type Calc struct {
	listener *parser.CalcListener
}

func NewCalc() Calc {
	return Calc{
		listener: parser.NewCalcListener(),
	}
}

func (c *Calc) Calculate(input string) (r float64, e error) {
	// recovery from panics of CalcListener
	defer func() {
		if r := recover(); r != nil {
			e = errors.New(fmt.Sprint(r))
		}
	}()
	defer c.listener.Reset()

	// set up the input stream
	is := antlr.NewInputStream(input)

	// custom error listener to panic on syntax errors and terminate the parsing immediately
	panicErrorListener := parser.NewPanicErrorListener()

	// init lexer
	lexer := genparser.NewCalcliLexer(is)
	// removing default error listeners to avoid printing errors to stdout
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(panicErrorListener)

	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// init parser
	p := genparser.NewCalcliParser(ts)
	// BailErrorStrategy to stop parsing after the first error
	p.SetErrorHandler(antlr.NewBailErrorStrategy())
	p.RemoveErrorListeners()
	p.AddErrorListener(panicErrorListener)

	// parse the input
	antlr.ParseTreeWalkerDefault.Walk(c.listener, p.Start_())

	if lexer.GetError() != nil {
		return 0, errors.New(lexer.GetError().GetMessage())
	}
	if p.GetError() != nil {
		return 0, errors.New(p.GetError().GetMessage())
	}

	return c.listener.Result(), nil
}
