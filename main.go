package main

import (
    "fmt"
    "bufio"
    "os"
)

type Interpreter struct {
}

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
