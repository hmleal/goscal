package main

import (
    "fmt"
    "bufio"
    "os"
)

const INTEGER string = "INTEGER"
const PLUS string = "PLUS"
const MINUS string = "MINUS"
const EOF string = "EOF"

type Token struct {
    kind  string
    value string
}

type Parser struct {
}

type Interpreter struct {
}

func main() {
    fmt.Println("Welcome a implementation of Pascal in GoLang")

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">>> ")
        text, _ := reader.ReadString('\n')

        l := NewLexer(text)
        fmt.Println(l.get_next_token())
        fmt.Println(l.get_next_token())
        fmt.Println(l.get_next_token())
        fmt.Println(l.get_next_token())
    }
}
