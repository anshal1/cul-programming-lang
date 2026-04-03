package main

import (
	"bufio"
	"fmt"
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

	variables := lexer.Lexer(reader)
	for _, v := range variables {
		fmt.Printf("%+v\n", v)
	}
}
