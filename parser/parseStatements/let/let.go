package let

import (
	"fmt"

	"github.com/anshal1/custom-language/environment"
	"github.com/anshal1/custom-language/parser"
	"github.com/anshal1/custom-language/utils"
)

type LetStatement struct {
	Type  string
	Name  string
	Value any
}

func parseValue(valueToken utils.Token, typeToken utils.Token, p *parser.Parser, env *environment.Environment) (utils.Token, error) {
	if valueToken.Symbol != utils.TT_IDENT {

		switch typeToken.Value {
		case utils.Int:
			if valueToken.Symbol != utils.TT_INTEGER {
				return valueToken, fmt.Errorf("line %d: type mismatch — declared %s but got %s", valueToken.LineNo, typeToken.Value, valueToken.Symbol)
			}
			p.Next()
			return valueToken, nil
		case utils.Float:
			if valueToken.Symbol != utils.TT_FLOAT {
				return valueToken, fmt.Errorf("line %d: type mismatch — declared %s but got %s", valueToken.LineNo, typeToken.Value, valueToken.Symbol)
			}
			p.Next()
			return valueToken, nil
		case utils.Str:
			if valueToken.Symbol != utils.TT_STRING {
				return valueToken, fmt.Errorf("line %d: type mismatch — declared %s but got %s", valueToken.LineNo, typeToken.Value, valueToken.Symbol)
			}
			p.Next()
			return valueToken, nil
		case utils.Bool:
			if valueToken.Symbol != utils.TT_BOOLEAN {
				return valueToken, fmt.Errorf("line %d: type mismatch — declared %s but got %s", valueToken.LineNo, typeToken.Value, valueToken.Symbol)
			}
			p.Next()
			return valueToken, nil
		case utils.Null:
			if valueToken.Symbol != utils.TT_NULL {
				return valueToken, fmt.Errorf("line %d: type mismatch — declared %s but got %s", valueToken.LineNo, typeToken.Value, valueToken.Symbol)
			}
			p.Next()
			return valueToken, nil
		default:
			return utils.Token{}, fmt.Errorf("line %d: unexpected value %q", valueToken.LineNo, valueToken.Value)
		}
	}
	val, ok := env.Get(valueToken.Value)
	if !ok {
		return utils.Token{}, fmt.Errorf("line %d: undefined variable %q", valueToken.LineNo, valueToken.Value)
	}
	if val.Type != typeToken.Value {
		return utils.Token{}, fmt.Errorf("line %d: type mismatch — cannot assign %s to %s", valueToken.LineNo, val.Type, typeToken.Value)
	}
	p.Next()
	return val.Token, nil
}

func parse(p *parser.Parser, env *environment.Environment) (LetStatement, error) {
	statement := LetStatement{}
	if p.CurrentToken().Symbol == "EOF" {
		return statement, nil
	}
	if p.CurrentToken().Symbol != utils.TT_LET {
		return statement, nil
	}

	token, err := p.Expect(utils.TT_LET)
	if err != nil {
		fmt.Printf("%+v, %v\n", token, err)

		return statement, err
	}

	typeToken, err := p.Expect(utils.TT_TYPE)
	if err != nil {
		fmt.Printf("%+v, %v\n", typeToken, err)

		return statement, err
	}

	nameToken, err := p.Expect(utils.TT_IDENT)
	if err != nil {
		fmt.Printf("%+v, %v\n", nameToken, err)
		return statement, err
	}

	assignToken, err := p.Expect(utils.TT_ASSIGN)
	if err != nil {
		fmt.Printf("%+v, %v\n", assignToken, err)
		return statement, err
	}

	// value comes before semiColon
	value, err := parseValue(p.CurrentToken(), typeToken, p, env)
	if err != nil {
		return statement, err
	}
	statement.Value = value.Value

	_, err = p.Expect(utils.TT_SEMICOLON)
	if err != nil {
		return statement, err
	}
	statement.Name = nameToken.Value
	statement.Type = "VARIABLE"
	env.Set(nameToken.Value, environment.Value{
		Type:  typeToken.Value,
		Value: value.Value,
		Token: value,
	})
	return statement, nil
}

func ParseLetStatement(p *parser.Parser) (*[]LetStatement, error) {
	letStatement := make([]LetStatement, 0)
	env := environment.NewEnv()
	for {
		statement, err := parse(p, env)
		if err != nil {
			return &letStatement, err
		}
		if statement.Name == "" {
			break
		}
		letStatement = append(letStatement, statement)
	}
	return &letStatement, nil
}
