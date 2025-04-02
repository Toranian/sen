package main

import (
	"fmt"
	"os"
	"sen/repl"
)

func main() {
	fmt.Println("Sen Programming Language")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
