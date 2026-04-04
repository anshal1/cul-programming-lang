package let

import (
	"fmt"
	"slices"

	"github.com/anshal1/custom-language/parser"
	"github.com/anshal1/custom-language/utils"
)

type LetStatement struct {
	Type  string
	Name  string
	Value any
}

func ParseLetStatement(p *parser.Parser) {
	for {
		if p.CurrentToken().Symbol == "EOF" {
			break
		}
		token, err := p.Expect(utils.TT_LET)
		if err != nil {
			fmt.Printf("%+v, %v\n", token, err)
			break
		}

		typeToken, err := p.Expect(utils.TT_TYPE)
		if err != nil {
			fmt.Printf("%+v, %v\n", typeToken, err)
			break
		}

		nameToken, err := p.Expect(utils.TT_IDENT)
		if err != nil {
			fmt.Printf("%+v, %v\n", nameToken, err)
			break
		}

		assignToken, err := p.Expect(utils.TT_ASSIGN)
		if err != nil {
			fmt.Printf("%+v, %v\n", assignToken, err)
			break
		}
		// checking the type and parsing the value accordingly
		if typeToken.Symbol == "Type" {
			switch typeToken.Value {
			case "int":
				valueToken, err := p.Expect(utils.TT_INTEGER)
				if err != nil {
					fmt.Printf("%+v %+v\n", valueToken, err)
					break
				}
			case "str":
				valueToken, err := p.Expect(utils.TT_STRING)
				if err != nil {
					fmt.Printf("%+v %+v\n", valueToken, err)
					break
				}
			case "bool":
				valueToken, err := p.Expect(utils.TT_BOOLEAN)
				if err != nil {
					fmt.Printf("%+v %+v\n", valueToken, err)
					break
				}
			case "float":
				valueToken, err := p.Expect(utils.TT_FLOAT)
				if err != nil {
					fmt.Printf("%+v %+v\n", valueToken, err)
					break
				}
			case "null":
				valueToken, err := p.Expect(utils.TT_NULL)
				if err != nil {
					fmt.Printf("%+v %+v\n", valueToken, err)
					break
				}
			}

		} else if typeToken.Symbol == "Ident" {
			idx := slices.IndexFunc(p.Tokens, func(t utils.Token) bool {
				return t.Symbol == "Ident" && t.Value == typeToken.Value
			})
			if idx == -1 {
				fmt.Println("variable not found")
				break
			}
			fmt.Printf("%+v\n", p.Tokens[idx])
		}
		fmt.Printf("%+v\n", p.CurrentToken())
		semiColonToken, err := p.Expect(utils.TT_SEMICOLON)
		if err != nil {
			fmt.Printf("%+v, %v\n", semiColonToken, err)
			break
		}
	}

	// nameToken, err := p.Expect(utils.TT_)
	// if err != nil {
	// 	return nil, err
	// }

	// valueToken, err := p.Expect(utils.TT_COLON)
	// if err != nil {
	// 	return nil, err
	// }

	// value, err := p.Expect(utils.TT_VALUE)
	// if err != nil {
	// 	return nil, err
	// }

	// return &LetStatement{
	// 	Type:  token.Symbol,
	// 	Name:  nameToken.Value,
	// 	Value: value.Value,
	// }, nil
}
