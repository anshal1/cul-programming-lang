package parser

import (
	"fmt"

	"github.com/anshal1/custom-language/utils"
)

type Parser struct {
	tokens  []utils.Token
	current int
	Tokens  []utils.Token
}

func NewParser(tokens []utils.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
		Tokens:  tokens,
	}
}

func (p *Parser) Next() utils.Token {
	if p.current >= len(p.tokens) {
		return utils.Token{}
	}
	token := p.tokens[p.current]
	p.current++
	return token
}

func (p *Parser) CurrentToken() utils.Token {
	if p.current >= len(p.tokens) {
		return utils.Token{}
	}
	return p.tokens[p.current]
}

func (p *Parser) Expect(tokenType string) (utils.Token, error) {
	token := p.CurrentToken()
	if token.Symbol != tokenType {
		return utils.Token{}, fmt.Errorf("expected %s, got %s", tokenType, token.Symbol)
	}
	p.Next()
	return token, nil
}
