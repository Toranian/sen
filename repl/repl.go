package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sen/lexer"
	"sen/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

// Run reads a file, tokenizes it using the lexer, and prints the tokens
func Run(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the entire file content while preserving line endings
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Pass the full content (with newlines) to the lexer
	l := lexer.New(string(content))

	// Tokenize and print tokens
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
