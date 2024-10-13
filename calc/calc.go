package calc

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"n0rdy.foo/calcli/calc/parser"
	genparser "n0rdy.foo/calcli/calc/parser/gen"
	"strings"
)

type CalcProcessor struct {
	listener *parser.CalcListener
}

func NewCalcProcessor() CalcProcessor {
	return CalcProcessor{
		listener: parser.NewCalcListener(),
	}
}

func (cp *CalcProcessor) Process(input string) (r *parser.CalcResult, e error) {
	// recovery from panics of CalcListener
	defer func() {
		if r := recover(); r != nil {
			e = errors.New(fmt.Sprint(r))
		}
	}()
	defer cp.listener.Reset()

	// set up the input stream
	is := antlr.NewInputStream(cp.removeSpaces(input))

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
	antlr.ParseTreeWalkerDefault.Walk(cp.listener, p.Start_())

	if lexer.GetError() != nil {
		return nil, errors.New(lexer.GetError().GetMessage())
	}
	if p.GetError() != nil {
		return nil, errors.New(p.GetError().GetMessage())
	}

	return cp.listener.Result(), nil
}

func (cp *CalcProcessor) removeSpaces(text string) string {
	return strings.ReplaceAll(text, " ", "")
}
