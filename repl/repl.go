package repl

import (
	"bufio"
	"fmt"
	"io"
	"notc/lexer"
	"notc/token"
)

const PROMPT = "NotC >> "

func Start(in io.Reader) {
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
