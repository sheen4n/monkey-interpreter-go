package main

import (
	"fmt"
	"io/ioutil"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Check if a file path is provided as an argument
	if len(os.Args) > 1 {
		// Get the file path from command line argument
		filePath := os.Args[1]
		executeFile(filePath)
	} else {
		// No file provided, start REPL
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}

func executeFile(filePath string) {
	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	env := object.NewEnvironment()
	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
		os.Exit(1)
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func printParserErrors(errors []string) {
	fmt.Printf("parser errors:\n")
	for _, msg := range errors {
		fmt.Printf("\t%s\n", msg)
	}
}
