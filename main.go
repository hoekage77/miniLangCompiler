package main

import (
	"fmt"	
)

func main() {
    source := `
    let x = 5;
    if (x &gt; 4) {
        let y = x + 1;
        // This is a comment
        while (y &lt; 10) {
            y = y + 1;
        }
    }
    `
    lexer := NewLexer(source)
    lexer.tokenize()
    for _, token := range lexer.GetTokens() {
        fmt.Printf("%s: %s\n", token.Type, token.Value)
    }

	    
	
}