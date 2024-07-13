// parser_test.go

package main

import (
    "fmt"
    "reflect"
    "testing"
)

func TestParser(t *testing.T) {
    symbolTable := NewSymbolTable()
    tests := []struct {
        source   string
        expected *ASTNode
    }{
        {
            source:   "let a = 5 + 3;",
            expected: &ASTNode{
                Type: NodeProgram,
                Body: []*ASTNode{
                    {
                        Type:  NodeDeclaration,
                        Value: "a",
                        Left: &ASTNode{
                            Type:  NodeBinaryOp,
                            Value: "+",
                            Left: &ASTNode{
                                Type:  NodeLiteral,
                                Value: "5",
                            },
                            Right: &ASTNode{
                                Type:  NodeLiteral,
                                Value: "3",
                            },
                        },
                }},
            },
        },
        {
            source: `if (x &gt; 4) {
                       let y = x + 1;
                     }`,
            expected: &ASTNode{
                Type: NodeProgram,
                Body: []*ASTNode{
                    {
                        Type: NodeIfStatement,
                        Left: &ASTNode{
                            Type:  NodeBinaryOp,
                            Value: ">",
                            Left: &ASTNode{
                                Type:  NodeIdentifier,
                                Value: "x",
                            },
                            Right: &ASTNode{
                                Type:  NodeLiteral,
                                Value: "4",
                            },
                        },
                        Right: &ASTNode{
                            Type: NodeBlock,
                            Body: []*ASTNode{
                                {
                                    Type:  NodeDeclaration,
                                    Value: "y",
                                    Left: &ASTNode{
                                        Type:  NodeBinaryOp,
                                        Value: "+",
                                        Left: &ASTNode{
                                            Type:  NodeIdentifier,
                                            Value: "x",
                                        },
                                        Right: &ASTNode{
                                            Type:  NodeLiteral,
                                            Value: "1",
                                        },
                                    },
                }},
                        },
                        Body: []*ASTNode{
                            nil,
                        },
                    },
                },
            },
        },
    }

    for _, test := range tests {
        lexer := NewLexer(test.source)
        lexer.tokenize()
        parser := NewParser(lexer.GetTokens(), symbolTable)
        ast := parser.parseProgram()

        if !reflect.DeepEqual(ast, test.expected) {
            t.Errorf("For source '%s'\nExpected %+v\nbut got  %+v\n", test.source, test.expected, ast)
        } else {
            fmt.Printf("Test passed for source: %s\n", test.source)
        }
    }
}

