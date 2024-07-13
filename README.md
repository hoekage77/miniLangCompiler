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
