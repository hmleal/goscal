package main

import "testing"


func TestToString(t *testing.T) {
    var v Token
    v = Token{kind: INTEGER, value: "1"}
    if v.toInteger() != 1 {
        t.Error("Expected 1, got ", v.value)
    }
}
