package main

import (
	"notc/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin)
}
