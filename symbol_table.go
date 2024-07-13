// symbol_table.go

package main

import (
    "fmt"
)

type VarType string

const (
    IntVarType    VarType = "INT"
    UnknownVarType VarType = "UNKNOWN"
)

type VariableInfo struct {
    Identifier string
    Type       VarType
}

type SymbolTable struct {
    table map[string]interface{}
}

func NewSymbolTable() *SymbolTable {
    return &SymbolTable{
        table: make(map[string]interface{}),
    }
}

func (st *SymbolTable) Declare(identifier string, value interface{}) {
    if _, exists := st.table[identifier]; exists {
        panic(fmt.Sprintf("Identifier '%s' already declared", identifier))
    }
    st.table[identifier] = value
}

func (st *SymbolTable) Lookup(identifier string) (interface{}, bool) {
    value, exists := st.table[identifier]
    return value, exists
}