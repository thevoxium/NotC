package repl

import (
	"bufio"
	"fmt"
	"io"
	"notc/ast"
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
			for i := 0; i < len(program.Statements); i++ {
				ts := program.Statements[i].(*ast.TypeStatement)
				if ts != nil {
					fmt.Printf("Type: %T\n", ts)
					fmt.Printf("%+v\n", ts)
					fmt.Printf("%+v\n", ts.TypeName)
				}
				// fmt.Printf("Type: %T\n",program.Statements[i])
				// fmt.Printf("%+v\n", program.Statements[i])
			}
		} else {
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		}

	}
}
