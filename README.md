# Sen Programming Language

A simple programming language written in Go. Based off of the "Monkey" programming language.

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

## Future Goals

- Classes
- Modules
- A standard library
- File input and output
- Sen as an executable file
- For loops

## Running Code

To run Sen code, write the code in a `.sen` file. To execute: `go run main.go <path to .sen file>`

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

For a more detailed writeup, view the full blog post at (www.isaacmorrow.me/posts/sen)[https://isaacmorrow.me/posts/sen]
