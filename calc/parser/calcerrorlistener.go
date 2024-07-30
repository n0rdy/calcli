package parser

import (
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

type PanicErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewPanicErrorListener() *PanicErrorListener {
	return &PanicErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (pel *PanicErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}
