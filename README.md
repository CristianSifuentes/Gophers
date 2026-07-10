# Gophers 

![Go](https://img.shields.io/badge/Go-1.26.5-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![Status](https://img.shields.io/badge/status-learning%20sandbox-yellow)

Members and developers of the Go programming language community are popularly known as **"Gophers."**

This repository is a hands-on, commit-by-commit journey through core Go fundamentals — variables, types, control flow, data structures, and the standard toolchain — captured in a single evolving file, `hello.go`.

## Table of Contents

- [Overview](#overview)
- [Project Evolution (Commit History)](#project-evolution-commit-history)
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
  - [Inspecting Types with `reflect`](#inspecting-types-with-reflect)
  - [Constants](#constants)
  - [Control Flow: `if` / `else if` / `else`](#control-flow-if--else-if--else)
  - [No Ternary Operator — and How to Fake One](#no-ternary-operator--and-how-to-fake-one)
  - [Arrays](#arrays)
  - [Maps](#maps)
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

This project is a hands-on sandbox (`hello.go`) for learning core Go fundamentals, built incrementally one concept at a time:

- Variable declaration styles (`var`, inferred `var`, short `:=`) and reassignment
- Static typing, type inference, and explicit type conversion
- Numeric types (`int32`, `float64`) and arithmetic
- Booleans and boolean logic (`&&`, `==`)
- Runtime type inspection with the `reflect` package
- The difference between the built-in `println`/`print` and the standard `fmt` package
- Constants (`const`)
- Control flow (`if` / `else if` / `else`)
- Emulating a ternary operator, since Go deliberately has none
- Fixed-size arrays and bounds checking
- Maps (Go's built-in hash table) and `make`

It also documents the standard Go toolchain workflow — initializing a module, building a binary, and running code directly — and calls out the beginner pitfalls this file deliberately walks into (and how to fix each one).

## Project Evolution (Commit History)

This project grew incrementally, commit by commit, mirroring the natural order in which a newcomer learns Go. Each commit below introduced (or reinforced) a specific concept — reading them in order is itself a short Go tutorial.

| Commit | Message | Concept Introduced |
|---|---|---|
| `0b5fd9b` | **Initial commit** | Repository scaffolding: `.gitignore`, `LICENSE` (Apache-2.0), and an empty `README.md`. No Go code yet — just the project's legal and hygiene baseline. |
| `5e73f2a` | **Add initial code** | The first `hello.go` is born — a minimal 5-line file, the seed the rest of the sandbox grows from. |
| `3ee9fdb` | **Initial create, build and run go project** | The file becomes a real **Go module**: `go mod init` produces `go.mod` (`module main`, `go 1.26.5`), the project is built (`go build`, producing the `main` binary), and run for the first time. The README also receives its first full pass — prerequisites, structure, and the build/run workflow. |
| `aed46ac` | **Delete ugly icon** | Small housekeeping — a cosmetic README cleanup, a reminder that documentation quality matters even in a learning repo. |
| `b6f54e0` | **Working with variables, string, int32 and format combination** | The heart of the first language lesson: `hello.go` expands to demonstrate **variable declaration** (`var name string = "variable"`), **reassignment**, **type inference** (`var otherVariable2 = ""`), an explicit **`int32`** declaration and arithmetic (`myIny * 10`), **type conversion** (`string(myIny)`), and the classic **`println` has no format verbs** pitfall (`println("%s", name, myIny, "he")`). |
| `16e8946` | **Update documentation** | The README is rewritten to fully explain every concept above in depth: variable styles, static typing vs. inference, type conversion pitfalls, `println` vs. `fmt`, the anatomy of `go.mod`, and a full Go CLI command reference. |
| `9eee060`, `60134df`, `8b9b20a`, `f77e782` | **Branch merges** (`feature/one` ↔ `develop` ↔ `main`) | No new code — these commits synchronize the three long-lived branches so that documentation and code fixes made on one branch propagate to the others, keeping the branch graph consistent. |
| `0693e57` | **Add new commits-details section** | The README gains this very commit-history table, turning `git log` into a guided tour. |
| `ec2bf4e` | **Working with float64** | Introduces `var myFloat = 6.5` and mixed-type arithmetic via explicit conversion (`float64(myIny) + myFloat`) — reinforcing that Go never implicitly widens numeric types. |
| `7e1cbff` | **Geeting context for import fmt and reflect and working with boolean in Go** | Adds the `reflect` import and `reflect.TypeOf(...)` calls to print a variable's concrete type at runtime, plus the first `bool` variable (`myBool`). This is also where a real bug slipped in: `reflect.Type` is an interface, and the builtin `print`/`println` cannot accept it — see [Inspecting Types with `reflect`](#inspecting-types-with-reflect). |
| `b384d25` | **Working with := Go** | Introduces the short variable declaration form, `myString := "Variable"` — the idiomatic, terse alternative to `var x = ...` inside function bodies. |
| `01aa16e` | **working with ternary mimics** | Go has no `?:` ternary operator. This commit adds a `ternary(cond, whenTrue, whenFalse string) string` helper function so a conditional value can still be produced as a single expression and passed straight into `fmt.Println(...)`. |
| `213fe84` | **Workint with control flow** | Adds a full `if / else if / else` chain (`if true && myBool { ... } else if myString == "X" { ... } else { ... }`), demonstrating boolean operators and Go's C-like (but parenthesis-free) conditional syntax. |
| `4934440` | **.** | A tiny one-line tweak — no functional change worth documenting on its own, folded into the surrounding lesson. |
| `ba3f283` | **working with []int** | Introduces Go's first **data structure**: a fixed-size array, `var myArray [3]int`, plus `len()` and indexed access/assignment. |
| `01b0e9d` | **Go is a smart lenguaje - Invalid argument: index 4 out of bounds [0:3]compilerInvalidIndex** | Deliberately documents (as a comment) what happens when you index past an array's bounds — the Go **compiler**, not just the runtime, catches constant out-of-range indices at compile time. |
| `6d32190` | **Working with - func make(t Type, size ...int) Type** | Introduces Go's second core data structure: a `map[string]int` built with `make`, populated with key/value pairs, and printed with `fmt.Println`. |

> **Reading tip:** if you want to relive the learning path, check out each commit in order (`git checkout <hash>`) and diff `hello.go` against the previous one — you'll see the file grow one Go concept at a time.

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
├── hello.go     # sandbox program: variables, types, control flow, arrays, maps
└── main         # compiled binary produced by `go build` (safe to delete/regenerate)
```

> **Tip:** compiled binaries like `main` shouldn't normally be committed to version control. Consider adding a `.gitignore` entry for `/main` (or `/*.exe` on Windows) if this repo grows.

## Getting Started

Follow these three steps in order to go from a bare source file to a running program.

### 1. Initialize a Go Module

```bash
go mod init main
```

This creates a `go.mod` file, which turns this folder into a proper Go **module**. The module file tracks:

- The module's import path (here, `main`)
- The Go version used
- Any external dependencies (none yet — `fmt` and `reflect` are both part of the standard library, so no `require` block is needed)

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

import (
	"fmt"
	"reflect"
)

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

	fmt.Println("reflect int32", reflect.TypeOf(myIny))

	var myFloat = 6.5
	println(myFloat)
	fmt.Println("reflect float64", reflect.TypeOf(myFloat))
	println(float64(myIny) + myFloat)

	var myBool bool = true
	println(myBool)
	fmt.Println(ternary(myBool, "The variable is true", "The variable is false"))

	// variable declared and initialized abbreviated way
	myString := "Variable"
	fmt.Println(myString)

	// Constants
	const myConstant = "Constant"
	fmt.Println(myConstant)

	// Control Flow
	if true && myBool {
		println("True")
	} else if myString == "X" {
		println("myString = X")
	} else {
		println("else")
	}

	// Data Structure: Array
	var myArray [3]int
	println(len(myArray))
	println("Array[2]", myArray[2])

	myArray[2] = 3
	println("Before Array[2]", myArray[2])

	// Data Structure: Map
	myMap := make(map[string]int)
	myMap["Cris"] = 35
	myMap["Artias"] = 20
	myMap["John"] = 29
	fmt.Println(myMap)
}

// ternary mimics the ?: operator Go doesn't have: a function call is an
// expression, so its result can be passed directly into another call.
func ternary(cond bool, whenTrue, whenFalse string) string {
	if cond {
		return whenTrue
	}
	return whenFalse
}
```

Line by line:

| Line(s) | What it demonstrates |
|---|---|
| `package main` / `import (...)` | Every executable Go program lives in `package main` with a `func main()` entry point. Multiple standard-library imports are grouped in one parenthesized block — the idiomatic style (`gofmt` will rewrite two separate `import` lines into this form). |
| `println("Hi, ...")` | The **built-in** `println` (lowercase, no import needed) writes to stderr — useful for quick debugging, but not meant for production output (see [Printing and Formatting Correctly](#printing-and-formatting-correctly)). |
| `var name string = "variable"` | Explicit variable declaration: keyword, identifier, type, value. |
| `name = "other vriable"` | Reassignment — `name` was declared with `var`, so it's mutable and can be reassigned to any value **of the same type** (`string`). |
| `var otherVariable2 = ""` | Type is **inferred** from the value (`""` → `string`); no explicit type needed. |
| `var myIny int32 = 7` | Explicit `int32` declaration. Go's default int type is platform-sized `int`, so `int32` here is deliberate/explicit. |
| `myIny = myIny * 10` | Arithmetic on typed integers works as expected. |
| `println(name + string(myIny))` | **Type conversion**: `string(myIny)` converts the `int32` to its Unicode code point as a string, not its decimal digits — see [Pitfalls](#pitfalls-already-in-this-file-and-how-to-fix-them). |
| `println("%s", name, myIny, "he")` | **Common pitfall**: `println` does **not** support `Printf`-style verbs like `%s`. It just prints each argument space-separated, so `%s` is printed literally, not substituted. |
| `fmt.Println("reflect int32", reflect.TypeOf(myIny))` | **Runtime type inspection**: `reflect.TypeOf` returns the concrete `reflect.Type` of any value (here, `int32`). Note it's printed with `fmt.Println`, not the builtin `println` — see [Inspecting Types with `reflect`](#inspecting-types-with-reflect) for why that distinction matters. |
| `var myFloat = 6.5` / `float64(myIny) + myFloat` | Introduces `float64` (Go's default floating-point type) and shows that mixed-type arithmetic (`int32 + float64`) requires an **explicit conversion** — Go never widens numeric types implicitly. |
| `var myBool bool = true` | Explicit `bool` declaration; Go's boolean type has exactly two values, `true` and `false` — no truthy/falsy coercion from other types. |
| `ternary(myBool, "...", "...")` | Calls the hand-rolled `ternary` helper (defined below `main`) to get a conditional string in one expression — see [No Ternary Operator](#no-ternary-operator--and-how-to-fake-one). |
| `myString := "Variable"` | **Short variable declaration** (`:=`) — declares and initializes in one step, inferring the type. Only legal inside function bodies. |
| `const myConstant = "Constant"` | Declares a compile-time **constant** — see [Constants](#constants). |
| `if true && myBool { ... } else if ... else { ... }` | A full **control-flow** chain using the `&&` logical AND and `==` comparison operators — see [Control Flow](#control-flow-if--else-if--else). |
| `var myArray [3]int` / `myArray[2] = 3` | A fixed-size **array** of 3 integers, zero-valued (`[0 0 0]`) until assigned — see [Arrays](#arrays). |
| `myMap := make(map[string]int)` | Creates a **map** (hash table) using the built-in `make` function, then populates it with three key/value pairs — see [Maps](#maps). |
| `func ternary(...) string { ... }` | A package-level function declared *after* `main` — Go doesn't require forward declarations; the compiler resolves all top-level names in a single pass. |

### Pitfalls Already in This File (and How to Fix Them)

1. **`string(int32)` doesn't produce digits.** `string(myIny)` where `myIny` is `10` converts the *rune* with code point 10 (a newline character), not the text `"10"`. To get `"10"`, use `strconv.Itoa(int(myIny))` (or `fmt.Sprintf("%d", myIny)`). Modern Go toolchains (`go vet`) flag this exact pattern with *"conversion from int32 to string yields a string of one rune"*.
2. **`println` has no format verbs.** `println("%s", name, myIny, "he")` prints the literal string `%s`, then each remaining argument — it does **not** substitute `%s` with `name`. To format output, use `fmt.Printf`/`fmt.Sprintf` from the standard library instead:

   ```go
   fmt.Printf("%s %d\n", name, myIny)
   ```
3. **`println`/`print` are debugging-only builtins**, not part of the language spec's guaranteed API — they write to stderr, have no formatting, and their exact behavior can vary by Go implementation. Idiomatic Go programs use `fmt.Println`/`fmt.Printf` for real output.
4. **The builtin `print`/`println` can't take arbitrary types.** An earlier version of this file tried `print("reflect int32", reflect.TypeOf(myIny))`, which fails to compile: the builtin printers only accept boolean, numeric, string, or pointer arguments — not the `reflect.Type` interface value returned by `reflect.TypeOf`. The fix is to use `fmt.Println` instead, which accepts any type.
5. **`if` is a statement, not an expression.** An earlier attempt wrote `println(if myBool { "true" })` — a syntax error, since Go requires every function argument to be an expression, and `if` never produces a value. The fix (now in the file) is the `ternary` helper: a function *call* is an expression, so `ternary(cond, a, b)` can be dropped straight into `fmt.Println(...)`.
6. **Indexing an array out of bounds is a compile-time error when the index is a constant.** `myArray[4]` on a `[3]int` array fails with *"invalid argument: index 4 out of bounds [0:3]"* — caught by the compiler before the program ever runs, because array length is part of the type (`[3]int` and `[4]int` are different types). This is different from a `slice`, where an out-of-range constant index can't always be caught statically and instead panics at runtime.

## Go Language Concepts

### Variable Declaration Styles

Go offers several equivalent ways to declare a variable:

```go
var name string = "variable" // explicit type + value
var name = "variable"        // inferred type
name := "variable"           // short declaration (function bodies only, infers type)
var name string              // zero-value declaration (empty string "")
```

`:=` (short variable declaration) can only be used inside function bodies, not at package level, and it requires at least one new variable on the left-hand side. It's the form most idiomatic Go code reaches for by default — `var` is reserved for cases where you want the zero value, need an explicit type that differs from the inferred one, or are declaring at package scope.

### Static Typing and Type Inference

Go is **statically and strongly typed**: every variable has a fixed type at compile time, and the compiler rejects invalid operations between mismatched types instead of silently coercing them.

```go
var myIny int32 = 7
// myIny + name        // compile error: mismatched types int32 and string
```

When you omit the type (`var x = 5` or `x := 5`), Go **infers** it from the right-hand side at compile time — this is still static typing, just with less typing (pun intended). It is *not* the same as dynamic typing: the inferred type is fixed for the variable's lifetime.

### Type Conversion

Go never implicitly converts between types — every conversion must be explicit, using the target type as a function:

```go
var i int32 = 65
s := string(i)             // "A" — converts the rune (code point) to its character
s2 := strconv.Itoa(int(i)) // "65" — converts the number to its decimal text
f := float64(i)            // 65 — numeric widening conversion
```

`string(someInt)` is a common beginner trap (see [Pitfall #1](#pitfalls-already-in-this-file-and-how-to-fix-them)): it treats the integer as a **Unicode code point**, not as text to print. For number-to-text conversion, always reach for `strconv` or `fmt.Sprintf("%d", ...)`.

This same rule is why `float64(myIny) + myFloat` needs the explicit `float64(...)` — Go will not automatically promote an `int32` to a `float64` for you, even though the conversion is "safe."

### Printing and Formatting Correctly

| Function | Import | Formatting? | Destination | Typical use |
|---|---|---|---|---|
| `println(...)` / `print(...)` | none (builtin) | No | stderr | Quick, throwaway debugging; only boolean/numeric/string/pointer args |
| `fmt.Println(...)` | `fmt` | No (adds spaces + newline) | stdout | Simple output of arbitrary types |
| `fmt.Printf(format, ...)` | `fmt` | Yes (`%s`, `%d`, `%v`, `%T`, ...) | stdout | Formatted output |
| `fmt.Sprintf(format, ...)` | `fmt` | Yes | returns a `string` | Building formatted strings |

Rule of thumb: reach for `fmt` for anything beyond a one-off debug print of a primitive value — it's part of every idiomatic Go program, and it's the only option once interfaces, structs, or `reflect.Type` values enter the picture (see [Pitfall #4](#pitfalls-already-in-this-file-and-how-to-fix-them)).

### Inspecting Types with `reflect`

The standard library's [`reflect`](https://pkg.go.dev/reflect) package lets a program examine its own values and types at **runtime**:

```go
import "reflect"

var myIny int32 = 7
fmt.Println(reflect.TypeOf(myIny)) // prints: int32
```

`reflect.TypeOf(v)` returns a `reflect.Type` — an **interface**, not a primitive value. That's why it must be printed with `fmt.Println`/`fmt.Printf` rather than the builtin `println`, which only understands a fixed set of primitive argument kinds (see [Pitfall #4](#pitfalls-already-in-this-file-and-how-to-fix-them)).

`reflect` is powerful but "expert-level" Go: it bypasses static typing and compile-time guarantees, trading them for runtime flexibility. It's invaluable for writing generic serializers, ORMs, and testing helpers — but idiomatic Go code favors interfaces and generics over `reflect` wherever possible, reserving it for the cases where the type truly isn't known until runtime.

### Constants

```go
const myConstant = "Constant"
```

`const` declares a value that's fixed at **compile time** and can never be reassigned — attempting `myConstant = "other"` is a compile error. Unlike `var`:

- Constants can only hold basic types (numbers, strings, booleans, runes) — never slices, maps, or structs.
- Untyped constants (like `const Pi = 3.14159`) adopt whatever type context they're used in, which lets the same constant be used as an `int32` in one expression and a `float64` in another without an explicit conversion.
- Multiple related constants are often grouped with `iota` for auto-incrementing enums — a pattern worth exploring once this file's basics feel comfortable.

### Control Flow: `if` / `else if` / `else`

```go
if true && myBool {
	println("True")
} else if myString == "X" {
	println("myString = X")
} else {
	println("else")
}
```

Key differences from C-family languages:

- **No parentheses** around the condition (`if (true && myBool)` would still compile, but `gofmt` strips the parens — Go style omits them).
- **Mandatory braces**, even for a single-statement body — there's no bare `if x println(x)`.
- The condition must be a `bool` expression; there's no implicit conversion from `0`/`""`/`nil` to `false` like in some other languages.
- `&&` (logical AND) and `||` (logical OR) short-circuit, just as you'd expect; `==` compares values of comparable types.

### No Ternary Operator — and How to Fake One

Go deliberately omits the `?:` ternary operator found in C, JavaScript, and many other languages — the language spec favors explicit `if` statements for readability over compact conditional expressions. Since `if` is a **statement**, it cannot be used as a function argument:

```go
println(if myBool { "true" })  // ✗ syntax error: unexpected keyword if
```

This file's fix is a small helper that turns the conditional into a **function call**, which *is* a valid expression:

```go
func ternary(cond bool, whenTrue, whenFalse string) string {
	if cond {
		return whenTrue
	}
	return whenFalse
}

fmt.Println(ternary(myBool, "The variable is true", "The variable is false"))
```

Other common ways to get the same effect for a single value:

```go
// A local variable set inside an if/else
msg := "false"
if myBool {
	msg = "true"
}

// A map lookup keyed by the boolean (works well for a fixed set of outcomes)
labels := map[bool]string{true: "true", false: "false"}
fmt.Println(labels[myBool])
```

### Arrays

```go
var myArray [3]int      // [0 0 0] — a fixed-size, zero-valued array of length 3
println(len(myArray))   // 3
myArray[2] = 3           // assign by index (0-based)
println(myArray[2])      // 3
```

Key properties:

- **Length is part of the type.** `[3]int` and `[4]int` are different, incompatible types — you can't assign one to the other or pass a `[3]int` where a `[4]int` is expected.
- **Value semantics.** Arrays are copied by value on assignment or when passed to a function — mutating a copy never affects the original. (This is why Go code overwhelmingly prefers **slices**, `[]T`, which wrap a reference to an underlying array and support growth via `append`.)
- **Bounds are checked.** Indexing with a constant outside `[0, len)` — e.g. `myArray[4]` on a length-3 array — is a **compile-time** error (`invalid argument: index 4 out of bounds [0:3]`) because the compiler knows the array's length statically. Indexing with a *variable* index that's out of range instead panics at **runtime**.

### Maps

```go
myMap := make(map[string]int) // create an empty map[string]int
myMap["Cris"] = 35            // insert/update by key
myMap["Artias"] = 20
myMap["John"] = 29
fmt.Println(myMap)            // map[Artias:20 Cris:35 John:29]
```

Key properties:

- **`make` is required** (or a map literal, `map[string]int{"Cris": 35}`) — the zero value of a map type is `nil`, and writing to a `nil` map panics. Reading from a `nil` map, however, is safe and returns the zero value.
- **Unordered.** `fmt.Println` on a map sorts keys for deterministic, readable output, but iteration order with `for k, v := range myMap` is intentionally randomized by the runtime — never rely on map order.
- **Lookups return two values**: `v, ok := myMap["Cris"]` — `ok` is `false` if the key isn't present, letting you distinguish "missing key" from "key present with the zero value."
- **Deleting** a key uses the builtin `delete(myMap, "Cris")`.

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

Note there's no `require` block: `fmt` and `reflect` are both part of the Go **standard library**, which ships with every Go installation and never needs to be declared as a dependency.

A larger, real-world project would typically also declare a `require` block for third-party packages:

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

Both files should be committed to version control — never edit `go.sum` by hand. (This repository has no `go.sum` yet, since it has zero external dependencies.)

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
| **Purpose** | `go vet` statically analyzes code for suspicious constructs (e.g. wrong `Printf` verbs, the `string(int32)` conversion pattern); `go fmt` rewrites files to Go's canonical formatting |
| **When to run** | `go vet ./...` before committing to catch bugs like the pitfalls in [this file](#pitfalls-already-in-this-file-and-how-to-fix-them); `gofmt -w .` (or `go fmt ./...`) to auto-format |
| **Notes** | `go vet` only flags misused format verbs and suspicious conversions in `fmt`-family and builtin-conversion contexts — it will not, for example, catch a semantic ternary-as-statement mistake, since that's a parse error the compiler itself already rejects |

## Typical Workflow

```bash
# 1. One-time setup
go mod init main

# 2. Iterate during development
go run .

# 3. Check for suspicious code and enforce formatting
go vet ./...
gofmt -l .

# 4. Produce a final binary when ready
go build
./main
```

## Troubleshooting

| Problem | Likely Cause | Fix |
|---|---|---|
| `go: cannot find main module` | No `go.mod` file exists yet | Run `go mod init main` |
| `command not found: go` | Go isn't installed or isn't in your `PATH` | Install Go from [go.dev/dl](https://go.dev/dl/) |
| `go.mod already exists` | Module was already initialized | Skip step 1, or delete `go.mod` to start over |
| `invalid argument: index N out of bounds [0:len]` | A constant index is used past a fixed-size array's declared length | Use an index within `[0, len(array))`, or switch to a slice with `append` if the size needs to grow |
| `unexpected keyword if, expected expression` | Trying to use `if` as if it were an expression (e.g. inside a function call) | Move the conditional to its own statement, or wrap it in a helper function — see [No Ternary Operator](#no-ternary-operator--and-how-to-fake-one) |
| `<value> (variable of type reflect.Type) is not a valid type for a builtin argument` | Passing a `reflect.Type` (or any non-primitive/interface value) to builtin `print`/`println` | Use `fmt.Println`/`fmt.Printf` instead — they accept any type |

## License

See [LICENSE](./LICENSE) for details.
