# Gophers 

Members and developers of the Go programming language community are popularly known as **"Gophers."**

This repository is a starting point for learning and running basic Go programs.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [1. Initialize a Go Module](#1-initialize-a-go-module)
  - [2. Build the Project](#2-build-the-project)
  - [3. Run the Project](#3-run-the-project)
- [Code Walkthrough: `hello.go`](#code-walkthrough-hellogo)
- [Go Language Concepts](#go-language-concepts)
  - [Variable Declaration Styles](#variable-declaration-styles)
  - [Static Typing and Type Inference](#static-typing-and-type-inference)
  - [Type Conversion](#type-conversion)
  - [Printing and Formatting Correctly](#printing-and-formatting-correctly)
- [Understanding `go.mod`](#understanding-go-mod)
  - [Key Elements](#key-elements)
  - [Example `go.mod` File](#example-go-mod-file)
  - [`go.mod` vs `go.sum`](#go-mod-vs-go-sum)
- [Command Reference](#command-reference)
  - [`go mod init main`](#go-mod-init-main)
  - [`go mod tidy`](#go-mod-tidy)
  - [`go get <dependency>`](#go-get-dependency)
  - [`go build`](#go-build)
  - [`go run .`](#go-run-)
  - [`go vet` / `go fmt`](#go-vet--go-fmt)
- [Typical Workflow](#typical-workflow)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## Overview

This project is a hands-on sandbox (`hello.go`) for learning core Go fundamentals: variable declaration styles, static typing, type inference, explicit type conversion, and the difference between the built-in `println` and the standard `fmt` package. It also documents the standard Go toolchain workflow — initializing a module, building a binary, and running code directly.

## Prerequisites

- [Go](https://go.dev/dl/) installed (this module targets Go **1.26.5**, see `go.mod`)
- Verify your installation:

  ```bash
  go version
  ```

## Project Structure

```
Gophers/
├── LICENSE
├── README.md
├── go.mod       # module definition (module "main", go 1.26.5)
├── hello.go     # sandbox program: variables, types, conversions, printing
└── main         # compiled binary produced by `go build` (safe to delete/regenerate)
```

> **Tip:** compiled binaries like `main` shouldn't normally be committed to version control. Consider adding a `.gitignore` with an entry for `/main` (or `/*.exe` on Windows) if this repo grows.

## Getting Started

Follow these three steps in order to go from a bare source file to a running program.

### 1. Initialize a Go Module

```bash
go mod init main
```

This creates a `go.mod` file, which turns this folder into a proper Go **module**. The module file tracks:

- The module's import path (here, `main`)
- The Go version used
- Any external dependencies (none yet, since this project is self-contained)

You only need to run this **once** per project — not before every build.

### 2. Build the Project

```bash
go build
```

This compiles all `.go` files in the current directory into a single executable binary (e.g., `main` or `Gophers` on macOS/Linux, `main.exe` on Windows). The binary is placed in the current directory and can be run independently, without needing Go installed.

```bash
./main
```

### 3. Run the Project

```bash
go run .
```

This compiles **and** runs the program in a single step, without leaving a binary behind. It's the fastest way to test changes while developing.

## Code Walkthrough: `hello.go`

```go
package main

func main() {
	println("Hi, this is the firt print in console using go")

	// variables
	var name string = "variable"
	println(name)

	name = "other vriable"
	println(name)

	var otherVariable2 = ""
	println(otherVariable2)

	var myIny int32 = 7
	myIny = myIny * 10
	println(myIny)

	myIny = 10

	println(name + string(myIny))

	println("%s", name, myIny, "he")
}
```

Line by line:

| Line(s) | What it demonstrates |
|---|---|
| `package main` | Every executable Go program lives in `package main` with a `func main()` entry point |
| `println("Hi, ...")` | The **built-in** `println` (lowercase, no import needed) writes to stderr — useful for quick debugging, but not meant for production output (see [Printing and Formatting Correctly](#printing-and-formatting-correctly)) |
| `var name string = "variable"` | Explicit variable declaration: keyword, identifier, type, value |
| `name = "other vriable"` | Reassignment — `name` was declared with `var`, so it's mutable and can be reassigned to any value **of the same type** (`string`) |
| `var otherVariable2 = ""` | Type is **inferred** from the value (`""` → `string`); no explicit type needed |
| `var myIny int32 = 7` | Explicit `int32` declaration. Go's default int type is platform-sized `int`, so `int32` here is deliberate/explicit |
| `myIny = myIny * 10` | Arithmetic on typed integers works as expected |
| `println(name + string(myIny))` | **Type conversion**: `string(myIny)` converts the `int32` to its Unicode code point as a string, not its decimal digits — see the pitfall below |
| `println("%s", name, myIny, "he")` | **Common pitfall**: `println` does **not** support `Printf`-style verbs like `%s`. It just prints each argument space-separated, so `%s` is printed literally, not substituted |

### Pitfalls Already in This File (and How to Fix Them)

1. **`string(int32)` doesn't produce digits.** `string(myIny)` where `myIny` is `10` converts the *rune* with code point 10 (a newline character), not the text `"10"`. To get `"10"`, use `strconv.Itoa(int(myIny))` (or `fmt.Sprintf("%d", myIny)`).
2. **`println` has no format verbs.** `println("%s", name, myIny, "he")` prints the literal string `%s`, then each remaining argument — it does **not** substitute `%s` with `name`. To format output, use `fmt.Printf`/`fmt.Sprintf` from the standard library instead:

   ```go
   import "fmt"

   fmt.Printf("%s %d\n", name, myIny)
   ```
3. **`println`/`print` are debugging-only builtins**, not part of the language spec's guaranteed API — they write to stderr, have no formatting, and their exact behavior can vary by Go implementation. Idiomatic Go programs use `fmt.Println`/`fmt.Printf` (which require `import "fmt"`) for real output.

## Go Language Concepts

### Variable Declaration Styles

Go offers several equivalent ways to declare a variable:

```go
var name string = "variable" // explicit type + value
var name = "variable"        // inferred type
name := "variable"           // short declaration (function bodies only, infers type)
var name string              // zero-value declaration (empty string "")
```

`:=` (short variable declaration) can only be used inside function bodies, not at package level, and it requires at least one new variable on the left-hand side.

### Static Typing and Type Inference

Go is **statically and strongly typed**: every variable has a fixed type at compile time, and the compiler rejects invalid operations between mismatched types instead of silently coercing them.

```go
var myIny int32 = 7
// myIny + name        // compile error: mismatched types int32 and string
```

When you omit the type (`var x = 5`), Go **infers** it from the right-hand side at compile time — this is still static typing, just with less typing (pun intended). It is *not* the same as dynamic typing: the inferred type is fixed for the variable's lifetime.

### Type Conversion

Go never implicitly converts between types — every conversion must be explicit, using the target type as a function:

```go
var i int32 = 65
s := string(i)        // "A" — converts the rune (code point) to its character
s2 := strconv.Itoa(int(i)) // "65" — converts the number to its decimal text
f := float64(i)        // 65 — numeric widening conversion
```

`string(someInt)` is a common beginner trap (see pitfall #1 above): it treats the integer as a **Unicode code point**, not as text to print. For number-to-text conversion, always reach for `strconv` or `fmt.Sprintf("%d", ...)`.

### Printing and Formatting Correctly

| Function | Import | Formatting? | Destination | Typical use |
|---|---|---|---|---|
| `println(...)` / `print(...)` | none (builtin) | No | stderr | Quick, throwaway debugging |
| `fmt.Println(...)` | `fmt` | No (adds spaces + newline) | stdout | Simple output |
| `fmt.Printf(format, ...)` | `fmt` | Yes (`%s`, `%d`, `%v`, `%T`, ...) | stdout | Formatted output |
| `fmt.Sprintf(format, ...)` | `fmt` | Yes | returns a `string` | Building formatted strings |

Rule of thumb: reach for `fmt` for anything beyond a one-off debug print — it's part of every idiomatic Go program.

## Understanding `go.mod`

The `go.mod` file is the **core of any Go project**. It sits at the root of the project and defines the **module** — a logical unit that groups your packages together. Its main purpose is to manage dependencies and versioning in a portable, reproducible way.

You don't edit `go.mod` by hand — it's managed through the `go` CLI (`go mod init`, `go mod tidy`, `go get`, etc.).

### Key Elements

A `go.mod` file always contains three key pieces of information:

| Element | Description |
|---|---|
| **Module path** | The unique identifier of your project — typically your repository URL, e.g. `github.com/user/repo` |
| **Go version** | The minimum Go language version required to compile the code |
| **Dependencies** | The list of third-party packages the project needs, each pinned to an exact version |

### Example `go.mod` File

This repository's actual `go.mod`:

```go
module main

go 1.26.5
```

A larger, real-world project would typically also declare a `require` block:

```go
module github.com/user/repo

go 1.22

require (
    github.com/some/package v1.8.1
    golang.org/x/text v0.14.0
)
```

### `go.mod` vs `go.sum`

`go.mod` always works together with a second file, **`go.sum`**, which is generated automatically alongside it:

| File | Purpose |
|---|---|
| `go.mod` | Declares the module path, Go version, and required dependencies |
| `go.sum` | Stores cryptographic checksums of every dependency, guaranteeing that downloads are secure and builds are reproducible |

Both files should be committed to version control — never edit `go.sum` by hand.

## Command Reference

### `go mod init main`

| | |
|---|---|
| **Purpose** | Initializes a new Go module named `main` in the current directory |
| **Creates** | `go.mod` |
| **When to run** | Once, when starting a new project |
| **Notes** | The name (`main`) is the module path. For real-world projects, this is typically a repository URL, e.g. `go mod init github.com/user/repo` |

### `go mod tidy`

| | |
|---|---|
| **Purpose** | The most commonly used module command. Scans your source code, automatically adds any dependencies you're using but haven't declared, and removes unused ones from `go.mod` |
| **Updates** | `go.mod` and `go.sum` |
| **When to run** | After adding, removing, or changing imports in your code |
| **Notes** | Keeps your module manifest accurate and minimal — always run it before committing |

### `go get <dependency>`

| | |
|---|---|
| **Purpose** | Downloads and adds a new library to your project, updating `go.mod` with the correct version |
| **Updates** | `go.mod` and `go.sum` |
| **When to run** | When you want to introduce a new third-party package |
| **Example** | `go get github.com/some/package@v1.8.1` |

### `go build`

| | |
|---|---|
| **Purpose** | Compiles the package and its dependencies into an executable binary |
| **Creates** | A compiled binary in the current directory |
| **When to run** | Before distributing or deploying the program |
| **Notes** | Does **not** run the program automatically — you must execute the resulting binary yourself |

### `go run .`

| | |
|---|---|
| **Purpose** | Compiles and immediately runs the package in the current directory |
| **Creates** | A temporary binary (discarded after execution) |
| **When to run** | During development, for quick iteration |
| **Notes** | The `.` tells Go to run the package in the current directory (equivalent to `go run hello.go` here) |

### `go vet` / `go fmt`

| | |
|---|---|
| **Purpose** | `go vet` statically analyzes code for suspicious constructs (e.g. wrong `Printf` verbs); `go fmt` rewrites files to Go's canonical formatting |
| **When to run** | `go vet ./...` before committing to catch bugs like the `%s` misuse in [Common Pitfalls](#pitfalls-already-in-this-file-and-how-to-fix-them); `gofmt -w .` (or `go fmt ./...`) to auto-format |
| **Notes** | Note that `go vet` only flags misused format verbs when using `fmt.Printf`-family functions — it does not catch the `println("%s", ...)` pattern in this file, since `println` is a builtin, not `fmt` |

## Typical Workflow

```bash
# 1. One-time setup
go mod init main

# 2. Iterate during development
go run .

# 3. Produce a final binary when ready
go build
./main
```

## Troubleshooting

| Problem | Likely Cause | Fix |
|---|---|---|
| `go: cannot find main module` | No `go.mod` file exists yet | Run `go mod init main` |
| `command not found: go` | Go isn't installed or isn't in your `PATH` | Install Go from [go.dev/dl](https://go.dev/dl/) |
| `go.mod already exists` | Module was already initialized | Skip step 1, or delete `go.mod` to start over |

## License

See [LICENSE](./LICENSE) for details.
