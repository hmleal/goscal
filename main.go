package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome a implementation of Pascal in GoLang v2")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')

		lexer := NewLexer(text)
		parser := NewParser(lexer)
		interpreter := Interpreter{parser}

		fmt.Println(interpreter.Interpret())
	}
}
