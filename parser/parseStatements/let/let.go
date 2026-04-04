package let

import (
	"fmt"
	"strconv"

	"github.com/anshal1/custom-language/parser"
	"github.com/anshal1/custom-language/utils"
)

type LetStatement struct {
	Type  string
	Name  string
	Value any
}

func parseTypeAndValue(valueToken utils.Token, typeToken utils.Token, p *parser.Parser) (utils.Token, error) {
	if valueToken.Symbol != utils.TT_IDENT {
		switch typeToken.Value {
		case utils.Int:
			token, err := p.Expect(valueToken.Symbol)
			_, err = strconv.Atoi(valueToken.Value)
			if err != nil {
				return utils.Token{}, err
			}
			return token, err
		case utils.Str:
			token, err := p.Expect(valueToken.Symbol)
			if err != nil {
				return utils.Token{}, err
			}
			return token, err
		}
	}
	p.Next()
	return utils.Token{}, nil
}

func parse(token utils.Token, p *parser.Parser) (LetStatement, error) {
	statement := LetStatement{}
	if p.CurrentToken().Symbol == "EOF" {
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
	value, err := parseTypeAndValue(p.CurrentToken(), typeToken, p)
	if err != nil {
		return statement, err
	}
	statement.Value = value.Value

	_, err = p.Expect(utils.TT_SEMICOLON)
	if err != nil {
		return statement, err
	}
	statement.Name = nameToken.Value
	statement.Type = typeToken.Value
	return statement, nil
}

func ParseLetStatement(p *parser.Parser) (*[]LetStatement, error) {
	letStatement := make([]LetStatement, 0)
	for {
		statement, err := parse(p.CurrentToken(), p)
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
