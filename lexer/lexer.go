package lexer

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

// import (
// 	"bufio"
// 	"log"
// 	"strings"
// )

// type Variables struct {
// 	SYMBOL   string
// 	TYPE     string
// 	NAME     string
// 	OPERATOR string
// 	VALUE    string
// 	LINE_NO  int
// }

// func getStatementType(line string) string {
// 	if strings.HasPrefix(line, "let") {
// 		return "let-variable"
// 	}
// 	if strings.HasPrefix(line, "const") {
// 		return "const-variable"
// 	}
// 	return "unknown"
// }

// func getVariableType(line string) string {
// 	trimed := strings.TrimSpace(line)
// 	fields := strings.Fields(trimed)
// 	if len(fields) < 3 {
// 		return ""
// 	}
// 	return fields[1]
// }

// func getVariableName(line string) string {
// 	trimed := strings.TrimSpace(line)
// 	fields := strings.Fields(trimed)
// 	if len(fields) < 3 {
// 		return ""
// 	}
// 	return fields[2]
// }

// func getVariableValue(line string) string {
// 	trimed := strings.TrimSpace(line)
// 	fields := strings.Fields(trimed)
// 	if len(fields) < 3 {
// 		return ""
// 	}
// 	return fields[3]
// }
// func getOperator(line string) string {
// 	trimed := strings.TrimSpace(line)
// 	fields := strings.Fields(trimed)
// 	if len(fields) < 3 {
// 		return ""
// 	}
// 	return fields[3]
// }

// func Lexer(reader *bufio.Reader) []Variables {
// 	var variables []Variables = make([]Variables, 0)
// 	var lineNo int = 1
// 	for {
// 		l, err := reader.ReadString(';')
// 		if err != nil {
// 			break
// 		}
// 		line := strings.ReplaceAll(strings.TrimSpace(l), ";", "")
// 		statementType := getStatementType(line)
// 		switch statementType {
// 		case "let-variable":
// 			vType := getVariableType(line)
// 			vName := getVariableName(line)
// 			vValue := getVariableValue(line)
// 			vOperator := getOperator(line)
// 			variable := Variables{
// 				SYMBOL:   "let",
// 				TYPE:     vType,
// 				NAME:     vName,
// 				VALUE:    vValue,
// 				OPERATOR: vOperator,
// 				LINE_NO:  lineNo,
// 			}
// 			variables = append(variables, variable)

// 		case "const-variable":
// 			vType := getVariableType(line)
// 			vName := getVariableName(line)
// 			vValue := getVariableValue(line)
// 			vOperator := getOperator(line)
// 			variable := Variables{
// 				SYMBOL:   "const",
// 				TYPE:     vType,
// 				NAME:     vName,
// 				VALUE:    vValue,
// 				OPERATOR: vOperator,
// 				LINE_NO:  lineNo,
// 			}
// 			variables = append(variables, variable)

// 		default:
// 			log.Fatal("Unkown code: ", line)
// 		}
// 		lineNo++
// 	}
// 	return variables
// }

// v2 lexer
var (
	TT_LET        = "Let"
	TT_CONST      = "Const"
	TT_INTEGER    = "Int_Type"
	TT_FLOAT      = "Float_Type"
	TT_STRING     = "String_Type"
	TT_BOOLEAN    = "Boolean_Type"
	TT_NULL       = "Null_Type"
	TT_SEMICOLON  = "Semicolon"
	TT_ASSIGN     = "Assign"
	TT_PLUS       = "Plus"
	TT_MINUS      = "Minus"
	TT_MULTIPLY   = "Multiply"
	TT_DIVIDE     = "Divide"
	TT_EOF        = "EOF"
	TT_FUNCTION   = "Func"
	TT_IDENT      = "Ident"
	ILLEGAL       = "Illegal"
	TT_PRINT_FUNC = "PrintFunc"
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

var keywords = map[string]string{
	"let":   TT_LET,
	"const": TT_CONST,
	"float": TT_FLOAT,
	"str":   TT_STRING,
	"bool":  TT_BOOLEAN,
	"null":  TT_NULL,
	"int":   TT_INTEGER,
	"print": TT_PRINT_FUNC,
}

type Token struct {
	Symbol string
	Value  string
	LineNo int
}

type Tokenizer struct {
	input  string
	pos    int
	lineNo int
}

func NewTokenizer(input string, lineNo int) *Tokenizer {
	return &Tokenizer{
		input:  input,
		pos:    0,
		lineNo: lineNo,
	}
}

func (t *Tokenizer) readString() Token {
	// ignoring the double quotes
	t.pos++
	start := t.pos

	for t.pos < len(t.input) && t.input[t.pos] != '"' {
		t.pos++
	}
	return Token{Symbol: TT_STRING, Value: t.input[start:t.pos], LineNo: t.lineNo}
}

func (t *Tokenizer) readIdentifier() Token {
	start := t.pos
	for isLetter(t.input[t.pos]) && t.pos <= len(t.input) {
		t.pos++
	}
	word := t.input[start:t.pos]
	if symbol, ok := keywords[word]; ok {
		return Token{Symbol: symbol, Value: word, LineNo: t.lineNo}
	}
	return Token{Symbol: TT_IDENT, Value: word, LineNo: t.lineNo}
}

func (l *Tokenizer) skipWhitespace() {
	for l.pos < len(l.input) {
		ch := l.input[l.pos]
		switch ch {
		case '\n':
			l.lineNo++
			l.pos++
		case ' ', '\t', '\r':
			l.pos++
		default:
			return
		}
	}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func (l *Tokenizer) readNumber() Token {
	start := l.pos
	for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
		l.pos++
	}
	return Token{Symbol: TT_INTEGER, Value: l.input[start:l.pos], LineNo: l.lineNo}
}

func (t *Tokenizer) Next() (Token, error) {
	t.skipWhitespace()
	if t.pos >= len(t.input) {
		return Token{Symbol: TT_EOF, Value: "", LineNo: t.lineNo}, io.EOF
	}
	switch rune(t.input[t.pos]) {
	case PLUS:
		token := Token{Symbol: TT_PLUS, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case MINUS:
		token := Token{Symbol: TT_MINUS, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case MULTIPLY:
		token := Token{Symbol: TT_MULTIPLY, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case DIVIDE:
		token := Token{Symbol: TT_DIVIDE, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case ASSIGN:
		token := Token{Symbol: TT_ASSIGN, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case SEMICOLON:
		token := Token{Symbol: TT_SEMICOLON, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case STRING:
		return t.readString(), nil
	}
	if isDigit(t.input[t.pos]) {
		return t.readNumber(), nil
	}
	if isLetter(t.input[t.pos]) {
		return t.readIdentifier(), nil
	}
	t.pos++
	return Token{Symbol: ILLEGAL, Value: string(t.input[t.pos-1]), LineNo: t.lineNo}, errors.New("unknown token at position " + strconv.Itoa(t.pos-1) + ": " + string(t.input[t.pos-1]))
}

func Lexer(reader *bufio.Reader) []Token {
	var tokens []Token = make([]Token, 0)
	var lineNo int = 1
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line := strings.TrimSpace(l)
		if line == "" {
			continue
		}
		lexer := NewTokenizer(line, lineNo)
		for {
			token, err := lexer.Next()
			if err != nil {
				break
			}
			if token.Symbol == TT_EOF {
				break
			}
			tokens = append(tokens, token)
		}
		lineNo++
	}
	return tokens
}
