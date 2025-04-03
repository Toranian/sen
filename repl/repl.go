package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sen/evaluator"
	"sen/lexer"
	"sen/object"
	"sen/parser"
	// "sen/token"
	"time"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// Run reads a file, tokenizes it, parses it, evaluates it, and prints the result with execution time
func Run(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	env := object.NewEnvironment()

	// Read entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Initialize lexer, parser, and measure execution time
	l := lexer.New(string(content))
	p := parser.New(l)
	start := time.Now() // Start timing execution

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(os.Stdout, p.Errors())
		os.Exit(1)
	}

	// Evaluate the parsed program
	evaluator.Eval(program, env)
	elapsed := time.Since(start) // Calculate execution time

	// Print execution time
	fmt.Printf("\nExecution Time: %s\n", elapsed)
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Encountered errors when parsing: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
