package utils

var (
	TT_LET              = "Let"
	TT_CONST            = "Const"
	TT_INTEGER          = "Int_Value"
	TT_FLOAT            = "Float_Value"
	TT_STRING           = "String_Value"
	TT_BOOLEAN          = "Boolean_Value"
	TT_NULL             = "Null_Value"
	TT_SEMICOLON        = "Semicolon"
	TT_ASSIGN           = "Assign"
	TT_PLUS             = "Plus"
	TT_MINUS            = "Minus"
	TT_MULTIPLY         = "Multiply"
	TT_DIVIDE           = "Divide"
	TT_EOF              = "EOF"
	TT_FUNCTION         = "Func"
	TT_IDENT            = "Ident"
	ILLEGAL             = "Illegal"
	TT_PRINT_FUNC       = "PrintFunc"
	TT_TYPE             = "Type"
	TT_END_OF_STATEMENT = "EndOfStatement"
)

var (
	PLUS      = '+'
	MINUS     = '-'
	MULTIPLY  = '*'
	DIVIDE    = '/'
	ASSIGN    = '='
	SEMICOLON = ';'
	STRING    = '"'
)

var Keywords = map[string]string{
	"let":   TT_LET,
	"const": TT_CONST,
	"float": TT_TYPE,
	"str":   TT_TYPE,
	"bool":  TT_TYPE,
	"null":  TT_TYPE,
	"int":   TT_TYPE,
	"print": TT_PRINT_FUNC,
}

var (
	Str   = "str"
	Int   = "int"
	Float = "float"
	Bool  = "bool"
	Null  = "null"
)

type Token struct {
	Symbol string
	Value  string
	LineNo int
}
