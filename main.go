// http://golangtutorials.blogspot.ie/2011/06/methods-on-structs.html
package main

import (
	"bufio"
	"fmt"
	"os"
)

type interpreter struct {
	parser Parser
}

// NewInterpreter creates a interpreter
func (i *interpreter) NewInterpreter() {}

// Interpret visit all nodes in the tree
func (i *interpreter) Interpret() {
	///tree := i.parser.expr() // TODO change to .parse
	//return i.visit()
}

func (i *interpreter) visit() {}

func main() {
	fmt.Println("Welcome a implementation of Pascal in GoLang")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')

		lexer := NewLexer(text)
		parser := NewParser(lexer)
		fmt.Println(parser.expr())
	}
}
