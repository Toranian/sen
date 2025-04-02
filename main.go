package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sen/repl"
)

func main() {
	fmt.Println("Sen Programming Language")

	// Check if a file is passed as an argument
	if len(os.Args) > 1 {
		filename := os.Args[1]

		// Ensure the file has a .sen extension
		if filepath.Ext(filename) == ".sen" {
			repl.Run(filename)
			return
		} else {
			fmt.Println("Error: Only .sen files are supported")
			os.Exit(1)
		}
	}

	// Default to REPL mode if no file is provided
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
