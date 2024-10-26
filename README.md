# Monkey Language Interpreter

A Go implementation of the Monkey programming language, based on the book ["Writing An Interpreter In Go"](https://interpreterbook.com/) by Thorsten Ball.

## Overview

Monkey is a programming language that supports:

- Variable bindings
- Integer arithmetic
- Built-in functions
- First-class and higher-order functions
- Closures
- String data structure
- Array data structure
- Hash data structure

## Installation

### Prerequisites

- Go 1.11 or higher

### Steps

1. Clone the repository:

```bash
git clone https://github.com/yourusername/monkey-interpreter
cd monkey-interpreter
```

2. Build the interpreter:

```bash
go build -o monkey
```

## Usage

### Interactive REPL

To start the REPL (Read-Eval-Print Loop):

```bash
./monkey
```

You'll see a prompt where you can type Monkey code:

```
Hello! This is the Monkey programming language!
Feel free to type in commands
>>
```

### Running Monkey Files

To execute a Monkey source file:

```bash
./monkey ./examples/fibonacci.monkey
```

## Language Features

### Variables

```javascript
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

### Functions

```javascript
// Define functions
let add = fn(a, b) { 
    return a + b; 
};

// Higher-order functions
let twice = fn(f, x) {
    return f(f(x));
};
```

### Arrays

```javascript
let myArray = [1, 2, 3, 4, 5];
let first = myArray[0];
```

### Hashes

```javascript
let myHash = {"name": "Monkey", "age": 1};
let name = myHash["name"];
```

### Built-in Functions

```javascript
// Length of arrays or strings
len([1, 2, 3]); // => 3
len("hello");   // => 5

// First element of array
first([1, 2, 3]); // => 1

// Last element of array
last([1, 2, 3]); // => 3

// Rest of array (all elements except first)
rest([1, 2, 3]); // => [2, 3]

// Push element to array
push([1, 2], 3); // => [1, 2, 3]

// Print values
puts("Hello, World!");
```

## Example Programs

### Fibonacci Sequence

```javascript
let fibonacci = fn(x) {
    if (x == 0) {
        return 0;
    } else {
        if (x == 1) {
            return 1;
        } else {
            return fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

fibonacci(10);
```

### Map Function

```javascript
let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            return accumulated;
        } else {
            return iter(rest(arr), push(accumulated, f(first(arr))));
        }
    };
    return iter(arr, []);
};

let double = fn(x) { return x * 2 };
map([1, 2, 3, 4], double); // => [2, 4, 6, 8]
```

## Development

### Project Structure

```
monkey/
├── ast/        - Abstract Syntax Tree definitions
├── evaluator/  - Expression evaluator
├── lexer/      - Lexical analyzer
├── object/     - Object system
├── parser/     - Parser implementation
├── repl/       - REPL implementation
├── token/      - Token definitions
└── main.go     - Entry point
```

### Running Tests

To run the test suite:

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thorsten Ball for the book ["Writing An Interpreter In Go"](https://interpreterbook.com/)
- The Go programming language team
