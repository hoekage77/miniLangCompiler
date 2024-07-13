// lexer.go

package main

import (
    "fmt"
    "regexp"
)

type Lexer struct {
    source string
    tokens []Token
}

func NewLexer(source string) *Lexer {
    return &Lexer{
        source: source,
        tokens: []Token{},
    }
}

func (lexer *Lexer) tokenize() {
    patterns := map[TokenType]*regexp.Regexp{
        LET:       regexp.MustCompile(`\blet\b`),
        IF:        regexp.MustCompile(`\bif\b`),
        ELSE:      regexp.MustCompile(`\belse\b`),
        WHILE:     regexp.MustCompile(`\bwhile\b`),
        LBRACE:    regexp.MustCompile(`\{`),
        RBRACE:    regexp.MustCompile(`\}`),
        LPAREN:    regexp.MustCompile(`\(`),
        RPAREN:    regexp.MustCompile(`\)`),
        SEMICOLON: regexp.MustCompile(`;`),
        ASSIGN:    regexp.MustCompile(`=`),
        EQUAL:     regexp.MustCompile(`==`),
        GT:        regexp.MustCompile(`&gt;`),
        GE:        regexp.MustCompile(`&gt;=`),
        LT:        regexp.MustCompile(`&lt;`),
        LE:        regexp.MustCompile(`&lt;=`),
        PLUS:      regexp.MustCompile(`\+`),
        MINUS:     regexp.MustCompile(`-`),
        STAR:      regexp.MustCompile(`\*`),
        SLASH:     regexp.MustCompile(`/`),
        BANG:      regexp.MustCompile(`!`),
        NUMBER:    regexp.MustCompile(`[0-9]+`),
        IDENTIFIER: regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9]*`),
        WS:        regexp.MustCompile(`[ \r\t\n]+`),
        COMMENT:   regexp.MustCompile(`//.*`),
    }

    position := 0
    for position < len(lexer.source) {
        match := false
        for tokenType, pattern := range patterns {
            loc := pattern.FindStringIndex(lexer.source[position:])
            if loc != nil && loc[0] == 0 {
                value := lexer.source[position : position+loc[1]]
                if tokenType != WS && tokenType != COMMENT {
                    lexer.tokens = append(lexer.tokens, Token{Type: tokenType, Value: value})
                }
                position += loc[1]
                match = true
                break
            }
        }
        if !match {
            fmt.Printf("Unexpected character: %c\n", lexer.source[position])
            position++
        }
    }
}

func (lexer *Lexer) GetTokens() []Token {
    return lexer.tokens
}