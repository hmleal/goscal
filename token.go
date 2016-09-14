package main

import "strconv"

const (
    INTEGER string = "INTEGER"
    PLUS string = "PLUS"
    MINUS string = "MINUS"
    EOF string = "EOF"
)

type Token struct {
    kind  string
    value string
}

func (t *Token) toInteger() int {
    i, _ := strconv.Atoi(t.value)
    return i
}
