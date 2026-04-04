package main

import (
	"bufio"
	"log"
	"os"

	"github.com/anshal1/custom-language/lexer"
)

func main() {
	file, err := os.Open("./index.cul")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)

	tokens := lexer.Lexer(reader)
	for _, token := range tokens {
		log.Printf("%+v\n", token)
	}
	// parser := parser.NewParser(tokens)
	// let.ParseLetStatement(parser)
}
