miniLang Compiler

Overview:

This Go Lang Compiler is a simple, educational compiler developed using Go language. Its primary purpose is to translate a source code written in a custom language into an intermediate representation and eventually to a target machine code.

Features
- Lexical Analysis: Tokenizes the input source code.
- Syntax Analysis: Parses tokens into an Abstract Syntax Tree (AST).
- Semantic Analysis: Ensures semantic correctness by type checking and building a symbol table.
- Intermediate Code Generation: Translates the AST to an intermediate representation (IR).
- Optimization: Applies optimization techniques to the IR.
- Target Code Generation: Generates the final machine code from the optimized IR.
- Error Handling: Includes syntax, semantic, and runtime error handling.
- Testing: Unit tests for various compiler components.

## Setup Instructions

To set up the project, follow these steps:

1. **Clone the repository:**
   ```sh
   git clone https://github.com/hoekage77/MiniLangCompiler.git
   cd MiniLangCompiler
   ```

2. **Install dependencies:**
   Ensure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).

3. **Run the project:**
   ```sh
   go run main.go
   ```

4. **Run tests:**
   ```sh
   go test ./...
   ```

## Usage Examples

Here are some examples of how to use the MiniLang Compiler:

### Example 1: Basic Variable Declaration and Assignment

Source Code:
```plaintext
let a = 5 + 3;
```

Expected Output:
```plaintext
LET: let
IDENTIFIER: a
ASSIGN: =
NUMBER: 5
PLUS: +
NUMBER: 3
SEMICOLON: ;
```

### Example 2: If Statement

Source Code:
```plaintext
if (x > 4) {
   let y = x + 1;
}
```

Expected Output:
```plaintext
IF: if
LPAREN: (
IDENTIFIER: x
GT: >
NUMBER: 4
RPAREN: )
LBRACE: {
LET: let
IDENTIFIER: y
ASSIGN: =
IDENTIFIER: x
PLUS: +
NUMBER: 1
SEMICOLON: ;
RBRACE: }
```

### Example 3: While Loop

Source Code:
```plaintext
while (a < 10) {
   a = a + 1;
}
```

Expected Output:
```plaintext
WHILE: while
LPAREN: (
IDENTIFIER: a
LT: <
NUMBER: 10
RPAREN: )
LBRACE: {
IDENTIFIER: a
ASSIGN: =
IDENTIFIER: a
PLUS: +
NUMBER: 1
SEMICOLON: ;
RBRACE: }
```
