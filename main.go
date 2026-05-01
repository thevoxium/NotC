package main

import (
	"notc/repl"
	"os"
	"strconv"
)

func main() {
	mode, _ := strconv.Atoi(os.Args[1])
	repl.Start(os.Stdin, mode)
}
