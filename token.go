package main

const INTEGER string = "INTEGER"
const PLUS string = "PLUS"
const MINUS string = "MINUS"
const EOF string = "EOF"

type Token struct {
    kind  string
    value string
}
