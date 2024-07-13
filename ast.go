// ast.go

package main

type NodeType string

const (
    NodeProgram        NodeType = "Program"
    NodeDeclaration    NodeType = "Declaration"
    NodeAssignment     NodeType = "Assignment"
    NodeIfStatement    NodeType = "IfStatement"
    NodeWhileStatement NodeType = "WhileStatement"
    NodeBlock          NodeType = "Block"
    NodeExpression     NodeType = "Expression"
    NodeBinaryOp       NodeType = "BinaryOp"
    NodeUnaryOp        NodeType = "UnaryOp"
    NodeLiteral        NodeType = "Literal"
    NodeIdentifier     NodeType = "Identifier"
)

type ASTNode struct {
    Type  NodeType
    Value string
    Left  *ASTNode
    Right *ASTNode
    Body  []*ASTNode
}