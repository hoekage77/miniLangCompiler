package main

import (
    "fmt"
)

// Parser represents a parser for the custom language.
type Parser struct {
    tokens       []Token
    pos          int
    currentToken Token
    symbolTable  *SymbolTable
}

// NewParser creates a new Parser instance with the given tokens and symbol table.
func NewParser(tokens []Token, symbolTable *SymbolTable) *Parser {
    return &Parser{
        tokens:       tokens,
        pos:          0,
        currentToken: tokens[0],
        symbolTable:  symbolTable,
    }
}

// advance moves to the next token in the list.
func (p *Parser) advance() {
    p.pos++
    if p.pos < len(p.tokens) {
        p.currentToken = p.tokens[p.pos]
    }
}

// expect checks if the current token matches the expected type and advances to the next token.
func (p *Parser) expect(tokenType TokenType) {
    if p.currentToken.Type == tokenType {
        p.advance()
    } else {
        panic(fmt.Sprintf("Syntax Error: Expected %s but found %s", tokenType, p.currentToken.Type))
    }
}

// parseProgram parses the entire program and returns the root AST node.
func (p *Parser) parseProgram() *ASTNode {
    program := &ASTNode{Type: NodeProgram, Body: []*ASTNode{}}
    for p.pos < len(p.tokens) {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Body = append(program.Body, stmt)
        }
    }
    return program
}

// parseStatement parses a single statement based on the current token.
func (p *Parser) parseStatement() *ASTNode {
    switch p.currentToken.Type {
    case LET:
        return p.parseDeclaration()
    case IF:
        return p.parseIfStatement()
    case WHILE:
        return p.parseWhileStatement()
    case LBRACE:
        return p.parseBlock()
    default:
        return p.parseExpression()
    }
}

// parseDeclaration parses a variable declaration statement.
func (p *Parser) parseDeclaration() *ASTNode {
    p.expect(LET)
    identifier := p.currentToken.Value
    p.expect(IDENTIFIER)
    p.expect(ASSIGN)
    value := p.parseExpression()

    // Add to symbol table
    p.symbolTable.Declare(identifier, VariableInfo{Identifier: identifier, Type: IntVarType})

    return &ASTNode{Type: NodeDeclaration, Value: identifier, Left: value}
}

// parseExpression parses an expression.
func (p *Parser) parseExpression() *ASTNode {
    return p.parseEquality()
}

// parseEquality parses an equality expression.
func (p *Parser) parseEquality() *ASTNode {
    node := p.parseComparison()
    for p.currentToken.Type == EQUAL {
        op := p.currentToken.Value
        p.advance()
        right := p.parseComparison()
        node = &ASTNode{Type: NodeBinaryOp, Value: op, Left: node, Right: right}
    }
    return node
}

// parseComparison parses a comparison expression.
func (p *Parser) parseComparison() *ASTNode {
    node := p.parseTerm()
    for p.currentToken.Type == GT || p.currentToken.Type == GE || p.currentToken.Type == LT || p.currentToken.Type == LE {
        op := p.currentToken.Value
        p.advance()
        right := p.parseTerm()
        node = &ASTNode{Type: NodeBinaryOp, Value: op, Left: node, Right: right}
    }
    return node
}

// parseTerm parses a term in an expression.
func (p *Parser) parseTerm() *ASTNode {
    node := p.parseFactor()
    for p.currentToken.Type == PLUS || p.currentToken.Type == MINUS {
        op := p.currentToken.Value
        p.advance()
        right := p.parseFactor()
        node = &ASTNode{Type: NodeBinaryOp, Value: op, Left: node, Right: right}
    }
    return node
}

// parseFactor parses a factor in an expression.
func (p *Parser) parseFactor() *ASTNode {
    node := p.parseUnary()
    for p.currentToken.Type == STAR || p.currentToken.Type == SLASH {
        op := p.currentToken.Value
        p.advance()
        right := p.parseUnary()
        node = &ASTNode{Type: NodeBinaryOp, Value: op, Left: node, Right: right}
    }
    return node
}

// parseUnary parses a unary expression.
func (p *Parser) parseUnary() *ASTNode {
    if p.currentToken.Type == BANG || p.currentToken.Type == MINUS {
        op := p.currentToken.Value
        p.advance()
        right := p.parsePrimary()
        return &ASTNode{Type: NodeUnaryOp, Value: op, Left: right}
    }
    return p.parsePrimary()
}

// parsePrimary parses a primary expression.
func (p *Parser) parsePrimary() *ASTNode {
    switch p.currentToken.Type {
    case NUMBER:
        value := p.currentToken.Value
        p.advance()
        return &ASTNode{Type: NodeLiteral, Value: value}
    case IDENTIFIER:
        value := p.currentToken.Value

        // Type check: Ensure the variable is declared
        if _, exists := p.symbolTable.Lookup(value); !exists {
            panic(fmt.Sprintf("Semantic Error: Use of undeclared identifier '%s'", value))
        }

        p.advance()
        return &ASTNode{Type: NodeIdentifier, Value: value}
    case LPAREN:
        p.advance()
        expr := p.parseExpression()
        p.expect(RPAREN)
        return expr
    default:
        panic(fmt.Sprintf("Syntax Error: Unexpected token %s", p.currentToken.Type))
    }
}

// parseIfStatement parses an if statement.
func (p *Parser) parseIfStatement() *ASTNode {
    p.expect(IF)
    p.expect(LPAREN)
    condition := p.parseExpression()
    p.expect(RPAREN)
    thenBlock := p.parseBlock()
    var elseBlock *ASTNode
    if p.currentToken.Type == ELSE {
        p.advance()
        elseBlock = p.parseBlock()
    }
    return &ASTNode{Type: NodeIfStatement, Left: condition, Right: thenBlock, Body: []*ASTNode{elseBlock}}
}

// parseWhileStatement parses a while statement.
func (p *Parser) parseWhileStatement() *ASTNode {
    p.expect(WHILE)
    p.expect(LPAREN)
    condition := p.parseExpression()
    p.expect(RPAREN)
    body := p.parseBlock()
    return &ASTNode{Type: NodeWhileStatement, Left: condition, Right: body}
}

// parseBlock parses a block of statements.
func (p *Parser) parseBlock() *ASTNode {
    p.expect(LBRACE)
    block := &ASTNode{Type: NodeBlock, Body: []*ASTNode{}}
    for p.currentToken.Type != RBRACE && p.pos < len(p.tokens) {
        stmt := p.parseStatement()
        if stmt != nil {
            block.Body = append(block.Body, stmt)
        }
    }
    p.expect(RBRACE)
    return block
}
