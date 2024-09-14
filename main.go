package main

import (
	"fmt"
	"os"
)

// Structure to represent function calls
type FunctionCall struct {
	From string
	To   string
}

// Sample data
var functionCalls = []FunctionCall{
	{"Main", "FunctionA"},
	{"FunctionA", "FunctionB"},
	{"FunctionB", "FunctionC"},
	{"FunctionC", "FunctionD"},
	{"Main", "FunctionE"},
}

func main() {
	// Open file to write the Mermaid diagram
	file, err := os.Create("diagram.mmd")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write Mermaid graph syntax
	file.WriteString("graph TD;\n")

	// Write each function call as a Mermaid relationship
	for _, call := range functionCalls {
		line := fmt.Sprintf("    %s --> %s;\n", call.From, call.To)
		file.WriteString(line)
	}

	fmt.Println("Mermaid diagram generated: diagram.mmd")
}
