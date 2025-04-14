# Sen Programming Language

A simple programming language interpreter written in Go. It features a Go and Rust-like syntax, with the absence of semicolons. Although the language is quite simple, it's also powerful. The initial implementation is based off of the "Monkey" programming language, which you can learn more about [here.](https://monkeylang.org/)

## Features

The programming language features:

- UTF-8 input
- Variable assignments
- Ints, strings, and booleans
- While & break statements
- Control flow via if and else statements
- Arrays
- Prefix & infix expressions
- Function statements and literals
- Closures
- Block statements and scope
- Errors that state where the error occured

## Future Goals

- Classes
- Modules
- A standard library
- File input and output
- Sen as an executable file
- For loops

## Running Code Files

To run Sen code, write the code in a `.sen` file. To execute: `go run main.go <path to .sen file>`

## Running via REPL

To try Sen out through the Read-eval-print-loop, just run `go run main.go` and you'll be able to interact with it.

## Examples

Map example

```rust
fn map(array, func) {
  length = len(array)
  mapped = []
  count = 0

  while (count < length) {
    mapped = push(mapped, func(array[count]))
    count = count + 1
  }

  return mapped
}


fn square(x) {
  x * x
}

arr = [1, 2, 3, 4, 5, 6]

// "out" outputs values to the console
out(map(arr, square))
//>>> [1, 4, 9, 16, 25, 36]


```

Closures Example

```rust
fn makeMult(a, b) {
  fn(c) {a * b * c}
}

// This will create a new function that will multiply our number by 10
mult = makeMult(2, 5)

// 2 * 5 * 5 = 50
out(mult(5))

//>>> 50
```

For a more detailed writeup, view the full blog post at [www.isaacmorrow.me/posts/sen](https://isaacmorrow.me/posts/sen)
