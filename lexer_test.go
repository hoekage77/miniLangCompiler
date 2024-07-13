package main

import (
    "fmt"
    "reflect"
    "testing"
)

func TestLexer(t *testing.T) {
    tests := []struct {
        source   string
        expected []Token
    }{
        {
            source:   "let a = 5 + 3;",
            expected: []Token{{LET, "let"}, {IDENTIFIER, "a"}, {ASSIGN, "="}, {NUMBER, "5"}, {PLUS, "+"}, {NUMBER, "3"}, {SEMICOLON, ";"}},
        },
        {
            source:   "// This is a comment\nlet a = 5;  // Inline comment",
            expected: []Token{{LET, "let"}, {IDENTIFIER, "a"}, {ASSIGN, "="}, {NUMBER, "5"}, {SEMICOLON, ";"}},
        },
        {
            source: `if (x &gt; 4) {
                       let y = x + 1;
                     }`,
            expected: []Token{
                {IF, "if"}, {LPAREN, "("}, {IDENTIFIER, "x"}, {GT, ">"}, {NUMBER, "4"}, {RPAREN, ")"}, {LBRACE, "{"},
                {LET, "let"}, {IDENTIFIER, "y"}, {ASSIGN, "="}, {IDENTIFIER, "x"}, {PLUS, "+"}, {NUMBER, "1"}, {SEMICOLON, ";"},
                {RBRACE, "}"},
            },
        },
    }

    for _, test := range tests {
        lexer := NewLexer(test.source)
        lexer.tokenize()
        tokens := lexer.GetTokens()

        if !reflect.DeepEqual(tokens, test.expected) {
            t.Errorf("For source '%s'\nExpected %+v\nbut got  %+v\n", test.source, test.expected, tokens)
        } else {
            fmt.Printf("Test passed for source: %s\n", test.source)
        }
    }
}

