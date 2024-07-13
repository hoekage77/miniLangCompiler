

package main

type TokenType string

const (
    // Keywords
    LET     TokenType = "LET"
    IF      TokenType = "IF"
    ELSE    TokenType = "ELSE"
    WHILE   TokenType = "WHILE"

    // Symbols
    LBRACE     TokenType = "LBRACE"
    RBRACE     TokenType = "RBRACE"
    LPAREN     TokenType = "LPAREN"
    RPAREN     TokenType = "RPAREN"
    SEMICOLON  TokenType = "SEMICOLON"
    ASSIGN     TokenType = "ASSIGN"
    EQUAL      TokenType = "EQUAL"
    GT         TokenType = "GT"
    GE         TokenType = "GE"
    LT         TokenType = "LT"
    LE         TokenType = "LE"
    PLUS       TokenType = "PLUS"
    MINUS      TokenType = "MINUS"
    STAR       TokenType = "STAR"
    SLASH      TokenType = "SLASH"
    BANG       TokenType = "BANG"

    // Literals
    NUMBER     TokenType = "NUMBER"
    IDENTIFIER TokenType = "IDENTIFIER"

    // Others
    WS         TokenType = "WS"
    COMMENT    TokenType = "COMMENT"
)

type Token struct {
    Type  TokenType
    Value string
}