package lexer

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/anshal1/custom-language/utils"
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

func (t *Tokenizer) readString() utils.Token {
	// ignoring the double quotes
	t.pos++
	start := t.pos

	for t.pos < len(t.input) {
		if t.input[t.pos] == '"' {
			t.pos++
			break
		}
		t.pos++
	}
	var builder strings.Builder
	builder.WriteRune('"')
	builder.WriteString(t.input[start : t.pos-1])
	builder.WriteRune('"')
	return utils.Token{Symbol: utils.TT_STRING, Value: builder.String(), LineNo: t.lineNo}
}

func (t *Tokenizer) readIdentifier() utils.Token {
	start := t.pos
	for isLetter(t.input[t.pos]) && t.pos <= len(t.input) {
		t.pos++
	}
	word := t.input[start:t.pos]
	if symbol, ok := utils.Keywords[word]; ok {
		return utils.Token{Symbol: symbol, Value: word, LineNo: t.lineNo}
	}
	return utils.Token{Symbol: utils.TT_IDENT, Value: word, LineNo: t.lineNo}
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

func (l *Tokenizer) readNumber() utils.Token {
	start := l.pos
	for l.pos < len(l.input) && isDigit(l.input[l.pos]) {
		l.pos++
	}
	return utils.Token{Symbol: utils.TT_INTEGER, Value: l.input[start:l.pos], LineNo: l.lineNo}
}

func (t *Tokenizer) Next() (utils.Token, error) {
	t.skipWhitespace()
	if t.pos >= len(t.input) {
		return utils.Token{Symbol: utils.TT_END_OF_STATEMENT, Value: utils.TT_END_OF_STATEMENT, LineNo: t.lineNo}, nil
	}
	switch rune(t.input[t.pos]) {
	case utils.PLUS:
		token := utils.Token{Symbol: utils.TT_PLUS, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case utils.MINUS:
		token := utils.Token{Symbol: utils.TT_MINUS, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case utils.MULTIPLY:
		token := utils.Token{Symbol: utils.TT_MULTIPLY, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case utils.DIVIDE:
		token := utils.Token{Symbol: utils.TT_DIVIDE, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case utils.ASSIGN:
		token := utils.Token{Symbol: utils.TT_ASSIGN, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case utils.SEMICOLON:
		token := utils.Token{Symbol: utils.TT_SEMICOLON, Value: string(t.input[t.pos]), LineNo: t.lineNo}
		t.pos++
		return token, nil
	case utils.STRING:
		return t.readString(), nil
	}
	if isDigit(t.input[t.pos]) {
		return t.readNumber(), nil
	}
	if isLetter(t.input[t.pos]) {
		return t.readIdentifier(), nil
	}
	t.pos++
	return utils.Token{Symbol: utils.ILLEGAL, Value: string(t.input[t.pos-1]), LineNo: t.lineNo}, errors.New("unknown token at position " + strconv.Itoa(t.pos-1) + ": " + string(t.input[t.pos-1]))
}

func Lexer(reader *bufio.Reader) []utils.Token {
	var tokens []utils.Token = make([]utils.Token, 0)
	var lineNo int = 1
	for {
		l, err := reader.ReadString(';')
		if err != nil {
			if errors.Is(err, io.EOF) {
				token := utils.Token{Symbol: utils.TT_EOF, Value: utils.TT_EOF, LineNo: lineNo}
				tokens = append(tokens, token)
			}
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
			if token.Symbol == utils.TT_END_OF_STATEMENT {
				break
			}
			tokens = append(tokens, token)
		}
		lineNo++
	}
	return tokens
}
