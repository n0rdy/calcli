grammar Calcli;

// Tokens:
ABS
    : 'abs'
    ;

ACOS
    : 'acos'
    ;

ASIN
    : 'asin'
    ;

ATAN
    : 'atan'
    ;

CEIL
    : 'ceil'
    ;

COMMA
    : ','
    ;

COS
    : 'cos'
    ;

DOLLAR
    : '$'
    ;

DOT
    : '.'
    ;

EQUAL
    : '='
    ;

EULER
    : 'e'
    ;

EXCLAMATION
    : '!'
    ;

EXP
    : 'exp'
    ;

EXP2
    : 'exp2'
    ;

INT
    : [0-9 ]+
    ;

FLOAT
    : INT DOT INT
    ;

FLOOR
    : 'floor'
    ;

LN
    : 'ln'
    ;

LOG
    : 'log'
    ;

LOG2
    : 'log2'
    ;

LOG10
    : 'log10'
    ;

LPAREN
    : '('
    ;

MINUS
    : '-'
    ;

MOD
    : 'mod'
    ;

NRT
    : 'nrt'
    ;

PERCENT
    : '%'
    ;

PERCENTAGE
    : 'percent'
    ;

PI
    : 'pi'
    ;

PLUS
    : '+'
    ;

PMEM
    : 'pmem'
    ;

POW
    : '^'
    ;

ROUND
    : 'round'
    ;

PREVIOS
    : '$pr'
    ;

RPAREN
    : ')'
    ;

SIN
    : 'sin'
    ;

STAR
    : '*'
    ;

SLASH
    : '/'
    ;

SQRT
    : 'sqrt'
    ;

TAN
    : 'tan'
    ;

VAR
    : DOLLAR [0-9a-zA-Z]+
    ;

WS
    : [ \t\r\n]+ -> skip
    ;

// Rules:

// units:
unit
    : number
    | constant
    | variable
    ;

constant
    : EULER
    | PI
    ;

number
    : INT
    | FLOAT
    ;

variable
    : PREVIOS
    | VAR
    ;

// math functions:
mathFunc
    : absMathFunc
    | acosMathFunc
    | asinMathFunc
    | atanMathFunc
    | ceilMathFunc
    | cosMathFunc
    | expMathFunc
    | exp2MathFunc
    | floorMathFunc
    | lnMathFunc
    | logMathFunc
    | log2MathFunc
    | log10MathFunc
    | modMathFunc
    | nrtMathFunc
    | percentMathFunc
    | roundMathFunc
    | sinMathFunc
    | sqrtMathFunc
    | tanMathFunc
    ;

absMathFunc
    : ABS LPAREN num=expr RPAREN
    ;

acosMathFunc
    : ACOS LPAREN num=expr RPAREN
    ;

asinMathFunc
    : ASIN LPAREN num=expr RPAREN
    ;

atanMathFunc
    : ATAN LPAREN num=expr RPAREN
    ;

ceilMathFunc
    : CEIL LPAREN num=expr RPAREN
    ;

cosMathFunc
    : COS LPAREN num=expr RPAREN
    ;

expMathFunc
    : EXP LPAREN num=expr RPAREN
    ;

exp2MathFunc
    : EXP2 LPAREN num=expr RPAREN
    ;

floorMathFunc
    : FLOOR LPAREN num=expr RPAREN
    ;

lnMathFunc
    : LN LPAREN num=expr RPAREN
    ;

logMathFunc
    : LOG LPAREN num=expr COMMA base=expr RPAREN
    ;

log2MathFunc
    : LOG2 LPAREN num=expr RPAREN
    ;

log10MathFunc
    : LOG10 LPAREN num=expr RPAREN
    ;

modMathFunc
    : MOD LPAREN left=expr COMMA right=expr RPAREN
    ;

nrtMathFunc
    : NRT LPAREN num=expr COMMA degree=expr RPAREN
    ;

percentMathFunc
    : PERCENTAGE LPAREN num=expr COMMA from=expr RPAREN
    ;

roundMathFunc
    : ROUND LPAREN num=expr RPAREN
    ;

sinMathFunc
    : SIN LPAREN num=expr RPAREN
    ;

sqrtMathFunc
    : SQRT LPAREN num=expr RPAREN
    ;

tanMathFunc
    : TAN LPAREN num=expr RPAREN
    ;

// system functions:
sysFunc
    : printMemory
    ;

printMemory
    : PMEM LPAREN RPAREN
    ;

// Expressions
start
    : expr EOF
    | assign EOF
    | sysFunc EOF
    ;

expr
     : unit                                                 # UnitValue
     | MINUS right=expr                                     # Negation
     | LPAREN inner=expr RPAREN                             # Parens
     | mathFunc                                             # MathFunction
     | num=expr op=EXCLAMATION                              # Factorial
     | left=expr op=POW right=expr                          # Pow
     | left=expr op=(STAR | SLASH | PERCENT) right=expr     # MulDivMod
     | left=expr op=(PLUS | MINUS) right=expr               # AddSub
     ;

assign
    : VAR EQUAL val=expr
    ;