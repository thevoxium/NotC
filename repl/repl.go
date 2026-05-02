package repl

import (
	"bufio"
	"fmt"
	"io"
	"notc/lexer"
	"notc/parser"
	"notc/token"
)

const PROMPT = "NotC >> "

func Start(in io.Reader, mode int) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "quit" {
			return
		}

		l := lexer.NewLexer(line)

		if mode == 1 {
			p := parser.NewParser(l)
			program := p.ParseProgram()
			fmt.Println(program.String())
		} else {
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		}

	}
}
