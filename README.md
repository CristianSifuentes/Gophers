# Gophers 

![Go](https://img.shields.io/badge/Go-1.26.5-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![Status](https://img.shields.io/badge/status-learning%20sandbox-yellow)

Members and developers of the Go programming language community are popularly known as **"Gophers."**

This repository is a hands-on, commit-by-commit journey through core Go fundamentals — variables, types, control flow, data structures, custom types, and the standard toolchain — captured in a single evolving file, `hello.go`.

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
  - [`container/list`: a Doubly Linked List](#containerlist-a-doubly-linked-list)
  - [Loops: Go's Only Looping Keyword, `for`](#loops-gos-only-looping-keyword-for)
  - [Structs: Custom Types](#structs-custom-types)
  - [Methods and Receivers: Value vs. Pointer](#methods-and-receivers-value-vs-pointer)
- [Advanced Go Concepts](#advanced-go-concepts)
  - [Advanced Concurrency Patterns](#advanced-concurrency-patterns)
  - [High-Performance Memory Management](#high-performance-memory-management)
  - [Type System Mastery & Generics](#type-system-mastery--generics)
  - [Runtime Metaprogramming & Optimization](#runtime-metaprogramming--optimization)
  - [Advanced Error Handling](#advanced-error-handling)
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
- Maps (Go's built-in hash table), both via `make` and map literals
- `container/list`, the standard library's doubly linked list, and pointer semantics
- Go's single looping keyword, `for`, in its C-style, condition-only, and `range` forms
- Custom types with `struct`, and methods with value vs. pointer receivers

It also documents the standard Go toolchain workflow — initializing a module, building a binary, and running code directly — and calls out the beginner pitfalls this file deliberately walks into (and how to fix each one).

## Project Evolution (Commit History)

This project grew incrementally, commit by commit, mirroring the natural order in which a newcomer learns Go. Each commit below introduced (or reinforced) a specific concept — reading them in order is itself a short Go tutorial.

| Commit | Message | Concept Introduced |
|---|---|---|
| `0b5fd9b` | **Initial commit** | Repository scaffolding: `.gitignore`, `LICENSE` (Apache-2.0), and an empty `README.md`. No Go code yet — just the project's legal and hygiene baseline. |
| `5e73f2a` | **Add initial code** | The first `hello.go` is born — a minimal 5-line file, the seed the rest of the sandbox grows from. |
| `3ee9fdb` | **Initial create, build and run go project** | The file becomes a real **Go module**: `go mod init` produces `go.mod` (`module main`, `go 1.26.5`), the project is built (`go build`), and run for the first time. The README also receives its first full pass. |
| `aed46ac` | **Delete ugly icon** | Small housekeeping — a cosmetic README cleanup. |
| `b6f54e0` | **Working with variables, string, int32 and format combination** | The heart of the first lesson: **variable declaration** (`var name string = "variable"`), **reassignment**, **type inference** (`var otherVariable2 = ""`), an explicit **`int32`** declaration and arithmetic, **type conversion** (`string(myIny)`), and the classic **`println` has no format verbs** pitfall. |
| `16e8946` | **Update documentation** | The README is rewritten to explain every concept above in depth: variable styles, static typing vs. inference, type conversion pitfalls, `println` vs. `fmt`, `go.mod`, and a Go CLI command reference. |
| `9eee060`, `60134df`, `8b9b20a`, `f77e782` | **Branch merges** (`feature/one` ↔ `develop` ↔ `main`) | No new code — these commits synchronize the three long-lived branches. |
| `0693e57` | **Add new commits-details section** | The README gains a commit-history table, turning `git log` into a guided tour. |
| `ec2bf4e` | **Working with float64** | Introduces `var myFloat = 6.5` and mixed-type arithmetic via explicit conversion (`float64(myIny) + myFloat`) — reinforcing that Go never implicitly widens numeric types. |
| `7e1cbff` | **Geeting context for import fmt and reflect and working with boolean in Go** | Adds the `reflect` import and `reflect.TypeOf(...)` to print a variable's concrete type at runtime, plus the first `bool` variable. Also where a real bug slipped in: `reflect.Type` is an interface the builtin `print`/`println` can't accept — fixed by switching to `fmt.Println`. |
| `b384d25` | **Working with := Go** | Introduces the short variable declaration form, `myString := "Variable"`. |
| `01aa16e` | **working with ternary mimics** | Go has no `?:` operator. Adds a `ternary(cond, whenTrue, whenFalse string) string` helper so a conditional value can be produced as one expression. |
| `213fe84` | **Workint with control flow** | Adds a full `if / else if / else` chain using `&&` and `==`. |
| `4934440` | **.** | A tiny one-line tweak, folded into the surrounding lesson. |
| `ba3f283` | **working with []int** | Introduces Go's first fixed-size data structure: `var myArray [3]int`, plus `len()` and indexed access/assignment. |
| `01b0e9d` | **Go is a smart lenguaje - Invalid argument: index 4 out of bounds [0:3]compilerInvalidIndex** | Documents (as a comment) that indexing an array past its bounds with a constant index is a **compile-time** error, not a runtime panic. |
| `6d32190` | **Working with - func make(t Type, size ...int) Type** | Introduces maps: `make(map[string]int)`, populated with key/value pairs and printed. |
| `229b068` | **Add documentation** | README sync pass to catch up with the map work above. |
| `ec34821` | **Working with map[string]int{"string": int, "string": int}** | Adds the **map literal** form, `map[string]int{"Cris": 36, "Artias": 30}` — an alternative to `make` when the initial contents are known upfront. |
| `f615525` | **Working with ontainer/list and getting pointer context** | Brings in the standard library's `container/list` (a doubly linked list) via `list.New()` and `PushBack`, and introduces **pointers**: `myList.Back().Value` is accessed through a `*list.Element`, the first explicit exposure to Go's address-of-a-value model. |
| `bdffddf` | **Working with for using List** | Iterates the list with a classic C-style `for i := 0; i < myList.Len(); i++ { ... }` loop — Go's *only* looping construct, reused here in its most explicit form. |
| `4519226` | **working with for and list and map** | Extends the same `for` pattern to a map's length and adds a `for index, value := range myArray` loop, showing the **`range`** form that Go favors for iterating collections idiomatically. |
| `6a67a4a` | **working with type - structure** | Introduces the **`struct`** keyword — Go's mechanism for custom, composite types — with a `Person{Name, Age}` type, a value-receiver method (`Greet`), a pointer-receiver method (`Birthday`), a zero-value struct, and a `[]Person` slice iterated with `range`. This is the file's first real step from *procedural scripting* toward *object-oriented-flavored* Go. |

> **Reading tip:** if you want to relive the learning path, check out each commit in order (`git checkout <hash>`) and diff `hello.go` against the previous one — you'll see the file grow one Go concept at a time, from `println("Hi")` all the way to custom types with methods.

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
├── hello.go     # sandbox program: variables, control flow, data structures, structs & methods
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
- Any external dependencies (none yet — `fmt`, `reflect`, and `container/list` are all part of the standard library, so no `require` block is needed)

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
	"container/list"
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

	// short declaration
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

	// Data Structure: Map (via make, then via literal)
	myMap := make(map[string]int)
	myMap["Cris"] = 35
	myMap["Artias"] = 20
	myMap["John"] = 29
	println(myMap)

	myMap2 := map[string]int{"Cris": 36, "Artias": 30}
	myMap2["Cris"] = 36
	println(myMap2)

	// Data Structure: container/list (doubly linked list) + pointers
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	myList.PushBack(4)
	println(myList.Back().Value)

	// Loops (Go's only looping keyword: for)
	for i := 0; i < myList.Len(); i++ {
		println(i)
	}
	for i := 0; i < len(myMap2); i++ {
		println(i)
	}
	for index, value := range myArray {
		fmt.Println(index, value)
	}

	// Data Structure: Struct (custom type) + methods
	cris := Person{Name: "Cris", Age: 35}
	fmt.Println(cris)
	fmt.Println(cris.Greet())

	cris.Birthday()
	fmt.Println("After birthday:", cris.Age)

	var empty Person
	fmt.Println(empty)

	people := []Person{
		{Name: "Artias", Age: 20},
		{Name: "John", Age: 29},
	}
	for _, person := range people {
		fmt.Println(person.Greet())
	}
}

// ternary mimics the ?: operator Go doesn't have: a function call is an
// expression, so its result can be passed directly into another call.
func ternary(cond bool, whenTrue, whenFalse string) string {
	if cond {
		return whenTrue
	}
	return whenFalse
}

// Person is a custom type: a struct groups named fields under one type.
type Person struct {
	Name string
	Age  int
}

// Greet uses a value receiver: p is a copy, so this method can never
// modify the caller's original value.
func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

// Birthday uses a pointer receiver: p points at the caller's original
// Person, so p.Age++ is visible to the caller after the call returns.
func (p *Person) Birthday() {
	p.Age++
}
```

Line by line:

| Line(s) | What it demonstrates |
|---|---|
| `package main` / `import (...)` | Every executable Go program lives in `package main` with a `func main()` entry point. Multiple standard-library imports are grouped in one parenthesized block — the idiomatic style. |
| `println("Hi, ...")` | The **built-in** `println` (lowercase, no import needed) writes to stderr — useful for quick debugging, but not meant for production output (see [Printing and Formatting Correctly](#printing-and-formatting-correctly)). |
| `var name string = "variable"` | Explicit variable declaration: keyword, identifier, type, value. |
| `name = "other vriable"` | Reassignment — `name` was declared with `var`, so it's mutable and can be reassigned to any value **of the same type**. |
| `var otherVariable2 = ""` | Type is **inferred** from the value (`""` → `string`). |
| `var myIny int32 = 7` | Explicit `int32` declaration; Go's default int type is platform-sized `int`, so `int32` here is deliberate. |
| `println(name + string(myIny))` | **Type conversion**: `string(myIny)` converts the `int32` to its Unicode code point as a string, not its decimal digits — see [Pitfalls](#pitfalls-already-in-this-file-and-how-to-fix-them). |
| `println("%s", name, myIny, "he")` | **Common pitfall**: `println` has no `Printf`-style verbs — `%s` prints literally. |
| `reflect.TypeOf(myIny)` | **Runtime type inspection** — returns a `reflect.Type` interface, which is why it's printed with `fmt.Println`, not builtin `println` (see [Inspecting Types with `reflect`](#inspecting-types-with-reflect)). |
| `var myFloat = 6.5` / `float64(myIny) + myFloat` | Introduces `float64` and shows mixed-type arithmetic requires an **explicit conversion**. |
| `var myBool bool = true` | Explicit `bool` declaration; exactly two values, no truthy/falsy coercion. |
| `ternary(myBool, "...", "...")` | Calls the hand-rolled `ternary` helper to get a conditional value in one expression — see [No Ternary Operator](#no-ternary-operator--and-how-to-fake-one). |
| `myString := "Variable"` | **Short variable declaration** — declares and initializes in one step. Function bodies only. |
| `const myConstant = "Constant"` | Declares a compile-time **constant** — see [Constants](#constants). |
| `if true && myBool { ... } else if ... else { ... }` | A full **control-flow** chain using `&&` and `==` — see [Control Flow](#control-flow-if--else-if--else). |
| `var myArray [3]int` / `myArray[2] = 3` | A fixed-size **array**, zero-valued until assigned — see [Arrays](#arrays). |
| `make(map[string]int)` / `map[string]int{...}` | Two ways to create a **map**: via `make` (empty, populate later) or a **map literal** (populated at creation) — see [Maps](#maps). |
| `list.New()` / `PushBack` / `.Back().Value` | The standard library's `container/list`, a doubly linked list, and Go's first explicit **pointer** exposure (`*list.Element`) — see [`container/list`](#containerlist-a-doubly-linked-list). |
| `for i := 0; i < N; i++` / `for index, value := range x` | Go's **only** looping keyword, `for`, in its C-style and `range` forms — see [Loops](#loops-gos-only-looping-keyword-for). |
| `type Person struct { ... }` | Declares a custom composite type — see [Structs](#structs-custom-types). |
| `func (p Person) Greet() string` / `func (p *Person) Birthday()` | Methods with value vs. pointer **receivers** — see [Methods and Receivers](#methods-and-receivers-value-vs-pointer). |
| `func ternary(...)` / `type Person struct{...}` declared after `main` | Go doesn't require forward declarations; the compiler resolves all package-level names in a single pass, so helper types/functions can live below the code that uses them. |

### Pitfalls Already in This File (and How to Fix Them)

1. **`string(int32)` doesn't produce digits.** `string(myIny)` where `myIny` is `10` converts the *rune* with code point 10 (a newline character), not the text `"10"`. Use `strconv.Itoa(int(myIny))` or `fmt.Sprintf("%d", myIny)` instead. Modern `go vet` flags this exact pattern.
2. **`println` has no format verbs.** `println("%s", name, myIny, "he")` prints the literal string `%s`, then each remaining argument. Use `fmt.Printf("%s %d\n", name, myIny)` for real formatting.
3. **`println`/`print` are debugging-only builtins**, not part of the language spec's guaranteed API. Idiomatic Go uses `fmt.Println`/`fmt.Printf` for real output.
4. **The builtin `print`/`println` can't reliably take arbitrary types.** Per the language spec, they only guarantee support for boolean, numeric, string, and pointer arguments. `reflect.TypeOf(myIny)` returns an interface value (`reflect.Type`), so it's printed with `fmt.Println` instead. Likewise, `println(myMap)` and `println(myList.Back().Value)` in this file happen to compile under the `gc` toolchain (which is more permissive than the spec requires), but they aren't portable — prefer `fmt.Println` for maps, structs, and interface values everywhere.
5. **`if` is a statement, not an expression.** `println(if myBool { "true" })` is a syntax error, since every function argument must be an expression and `if` never produces a value. The fix is the `ternary` helper: a function *call* is an expression, so `ternary(cond, a, b)` drops straight into `fmt.Println(...)`.
6. **Indexing an array out of bounds is a compile-time error when the index is a constant.** `myArray[4]` on a `[3]int` fails with *"invalid argument: index 4 out of bounds [0:3]"* — caught before the program runs, because array length is part of the type. A *variable* out-of-range index instead panics at runtime.
7. **`list.Element.Value` is `any` (`interface{}`).** `container/list` stores values as `any`, so retrieving one back (`myList.Back().Value`) loses static type information — using it as a specific type again requires a type assertion (`v := myList.Back().Value.(int)`), which panics if the stored value isn't actually an `int`. This is the classic trade-off of pre-generics-style containers versus a typed slice.

## Go Language Concepts

### Variable Declaration Styles

Go offers several equivalent ways to declare a variable:

```go
var name string = "variable" // explicit type + value
var name = "variable"        // inferred type
name := "variable"           // short declaration (function bodies only, infers type)
var name string              // zero-value declaration (empty string "")
```

`:=` can only be used inside function bodies and requires at least one new variable on the left-hand side. It's the form most idiomatic Go code reaches for by default — `var` is reserved for the zero value, an explicit differing type, or package-level scope.

### Static Typing and Type Inference

Go is **statically and strongly typed**: every variable has a fixed type at compile time, and the compiler rejects invalid operations between mismatched types instead of silently coercing them.

```go
var myIny int32 = 7
// myIny + name        // compile error: mismatched types int32 and string
```

When you omit the type (`var x = 5` or `x := 5`), Go **infers** it from the right-hand side at compile time — still static typing, just with less typing.

### Type Conversion

Go never implicitly converts between types — every conversion must be explicit, using the target type as a function:

```go
var i int32 = 65
s := string(i)             // "A" — converts the rune (code point) to its character
s2 := strconv.Itoa(int(i)) // "65" — converts the number to its decimal text
f := float64(i)            // 65 — numeric widening conversion
```

`string(someInt)` is a common beginner trap (see [Pitfall #1](#pitfalls-already-in-this-file-and-how-to-fix-them)): it treats the integer as a **Unicode code point**, not text. This same rule is why `float64(myIny) + myFloat` needs the explicit `float64(...)` — Go will not auto-promote an `int32` to a `float64`.

### Printing and Formatting Correctly

| Function | Import | Formatting? | Destination | Typical use |
|---|---|---|---|---|
| `println(...)` / `print(...)` | none (builtin) | No | stderr | Quick debugging; spec-guaranteed only for bool/numeric/string/pointer args |
| `fmt.Println(...)` | `fmt` | No (adds spaces + newline) | stdout | Simple output of arbitrary types |
| `fmt.Printf(format, ...)` | `fmt` | Yes (`%s`, `%d`, `%v`, `%T`, ...) | stdout | Formatted output |
| `fmt.Sprintf(format, ...)` | `fmt` | Yes | returns a `string` | Building formatted strings |

Rule of thumb: reach for `fmt` for anything beyond a one-off debug print of a primitive value — it's the only portable option once maps, structs, or interface values (like `reflect.Type` or `list.Element.Value`) enter the picture.

### Inspecting Types with `reflect`

```go
import "reflect"

var myIny int32 = 7
fmt.Println(reflect.TypeOf(myIny)) // prints: int32
```

`reflect.TypeOf(v)` returns a `reflect.Type` — an **interface**, not a primitive — which is why it's printed with `fmt.Println` rather than builtin `println` (see [Pitfall #4](#pitfalls-already-in-this-file-and-how-to-fix-them)).

`reflect` bypasses static typing for runtime flexibility — invaluable for serializers, ORMs, and test helpers, but idiomatic Go favors interfaces and generics wherever the type can be known at compile time instead.

### Constants

```go
const myConstant = "Constant"
```

`const` declares a value fixed at **compile time**; reassignment is a compile error. Unlike `var`:

- Constants can only hold basic types (numbers, strings, booleans, runes) — never slices, maps, or structs.
- Untyped constants (`const Pi = 3.14159`) adopt whatever type context they're used in.
- Related constants are often grouped with `iota` for auto-incrementing enums.

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

Key differences from C-family languages: **no parentheses** around the condition, **mandatory braces** even for one statement, and the condition must be a real `bool` — no implicit `0`/`""`/`nil` → `false` coercion. `&&`/`||` short-circuit as expected.

### No Ternary Operator — and How to Fake One

Go deliberately omits `?:`. Since `if` is a **statement**, it can't be a function argument:

```go
println(if myBool { "true" })  // ✗ syntax error: unexpected keyword if
```

This file's fix wraps the conditional in a **function call**, which *is* an expression:

```go
func ternary(cond bool, whenTrue, whenFalse string) string {
	if cond {
		return whenTrue
	}
	return whenFalse
}

fmt.Println(ternary(myBool, "The variable is true", "The variable is false"))
```

Other common alternatives for a single value: a local variable set inside `if`/`else`, or a `map[bool]string{true: "...", false: "..."}` lookup.

### Arrays

```go
var myArray [3]int      // [0 0 0] — fixed-size, zero-valued
println(len(myArray))   // 3
myArray[2] = 3
```

- **Length is part of the type** — `[3]int` and `[4]int` are incompatible.
- **Value semantics** — arrays copy on assignment or function call, which is why idiomatic Go usually prefers **slices** (`[]T`) for anything that needs to grow.
- **Bounds checked at compile time** for constant indices out of range (see [Pitfall #6](#pitfalls-already-in-this-file-and-how-to-fix-them)); variable indices panic at runtime instead.

### Maps

```go
myMap := make(map[string]int)  // create empty, populate later
myMap["Cris"] = 35

myMap2 := map[string]int{"Cris": 36, "Artias": 30} // map literal: create + populate at once
myMap2["Cris"] = 36                                 // still mutable afterward
```

- **`make` or a literal is required** — the zero value of a map is `nil`, and writing to a `nil` map panics (reading is safe and returns the zero value).
- **Unordered** — iteration order via `range` is intentionally randomized; `fmt.Println` sorts keys only for display purposes.
- **Two-value lookups**: `v, ok := myMap["Cris"]` — `ok` distinguishes "missing key" from "key present with zero value."
- **`delete(myMap, "Cris")`** removes a key.

### `container/list`: a Doubly Linked List

```go
import "container/list"

myList := list.New()
myList.PushBack(1)
myList.PushBack(2)
back := myList.Back() // *list.Element — a pointer
fmt.Println(back.Value)
```

- `list.New()` returns a `*list.List` — a **pointer** to the list header, so every method (`PushBack`, `PushFront`, `Remove`, ...) mutates the same underlying list without needing to reassign the variable.
- Each element is a `*list.Element`; `.Value` is typed `any`, so retrieving a specific type back requires a type assertion — see [Pitfall #7](#pitfalls-already-in-this-file-and-how-to-fix-them).
- **When to reach for it:** `container/list` shines when you need O(1) insertion/removal from both ends *and* from the middle given an element pointer. For most everyday Go code, a slice (with `append`) is simpler, more cache-friendly, and idiomatic — reserve linked lists for cases with genuinely frequent middle-of-sequence mutation.

### Loops: Go's Only Looping Keyword, `for`

Go has exactly one looping keyword — no `while`, no `do...while` — and `for` covers every case:

```go
// 1. C-style: init; condition; post
for i := 0; i < myList.Len(); i++ {
	println(i)
}

// 2. Condition-only ("while" equivalent)
for i < 10 {
	i++
}

// 3. Infinite loop
for {
	break // needs an explicit exit
}

// 4. range — the idiomatic way to iterate arrays, slices, maps, strings, and channels
for index, value := range myArray {
	fmt.Println(index, value)
}
```

Prefer `range` whenever you're walking a whole collection — it's shorter, avoids off-by-one bugs, and communicates *"iterate everything"* more clearly than a manually indexed C-style loop, which is better reserved for cases needing a custom stride, an early skip, or an index derived from something other than the collection itself.

### Structs: Custom Types

```go
type Person struct {
	Name string
	Age  int
}

cris := Person{Name: "Cris", Age: 35}  // named-field literal: order-independent, self-documenting
var empty Person                       // zero-value struct: Person{Name: "", Age: 0}
people := []Person{                     // a slice of structs composes cleanly with everything above
	{Name: "Artias", Age: 20},
	{Name: "John", Age: 29},
}
```

`struct` is Go's answer to "group related fields under one named type" — there's no `class` keyword. A struct is a value type: assigning one `Person` to another, or passing it to a function, copies all its fields (contrast with the reference-like `*list.List` above). Fields starting with an uppercase letter (`Name`, `Age`) are **exported** — visible outside the package they're defined in; lowercase fields are package-private.

### Methods and Receivers: Value vs. Pointer

Go doesn't have classes or inheritance, but it lets you attach functions to a type via a **receiver**, turning them into methods:

```go
// Value receiver: p is a copy. Safe, but can never mutate the original.
func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

// Pointer receiver: p points at the original. Required to mutate state.
func (p *Person) Birthday() {
	p.Age++
}
```

```go
cris := Person{Name: "Cris", Age: 35}
cris.Birthday()               // Go automatically rewrites this to (&cris).Birthday()
fmt.Println(cris.Age)         // 36 — the mutation is visible, because Birthday had a pointer receiver
```

**Rule of thumb:**

| Use a **value** receiver when... | Use a **pointer** receiver when... |
|---|---|
| The method doesn't need to modify the receiver | The method needs to modify the receiver's fields |
| The struct is small (cheap to copy) | The struct is large (avoid copying overhead) |
| You want the method to work on both values and pointers uniformly | Consistency: if *any* method on the type needs a pointer receiver, idiomatic Go uses pointer receivers for *all* of that type's methods |

Go automatically takes the address of an addressable value (like a local variable `cris`) to call a pointer-receiver method, so you rarely need to write `(&cris).Birthday()` yourself — but this auto-addressing does **not** happen for values stored in a map or returned from a function, which is a common source of "cannot call pointer method" compile errors.

## Advanced Go Concepts

Beyond basic syntax, building high-performance, scalable systems in Go means writing code that aligns with **mechanical sympathy** — how the runtime's scheduler, memory model, and concurrency primitives actually behave. Each concept below lives in its own runnable package under [`advanced/`](./advanced), so it can be built, run, and studied independently: `go run ./advanced/<package>`.

### Advanced Concurrency Patterns

Code: [`advanced/concurrency/main.go`](./advanced/concurrency/main.go)

Go's concurrency shines when goroutines are orchestrated systematically rather than launched arbitrarily.

- **Context Propagation** — `context.Context` manages request scopes, carries metadata, and broadcasts cancellation down deep call trees. Blocking operations should always honor `ctx.Done()` to avoid goroutine leaks.
- **Worker Pools** — a fixed set of worker goroutines pulling from a shared, buffered channel bounds concurrency and prevents unbounded resource consumption, regardless of how many jobs are queued.
- **The `select` Multiplexer** — adding a `default` case makes a channel operation non-blocking; `select` is also how timeout timers, cancellation channels, and data channels get multiplexed cleanly in one statement.
- **Pipeline Patterns** — independent processing stages connected by channels (`generate -> square -> print` in the example) let data stream through concurrently, each stage on its own goroutine.

```bash
go run ./advanced/concurrency
```

### High-Performance Memory Management

Code: [`advanced/memory/main.go`](./advanced/memory/main.go)

Optimizing allocations lowers Garbage Collection (GC) overhead and improves throughput.

- **Heap vs. Stack Analysis** — run `go build -gcflags="-m" ./advanced/memory` to see which variables "escape to the heap." Minimizing heap allocations reduces both GC pressure and cache misses.
- **Memory Reuse** — `sync.Pool` caches and reuses heavily allocated structures (buffers, decoders, etc.), keeping the memory footprint stable under traffic spikes instead of allocating fresh on every request.
- **Data Semantics** — value receivers copy the receiver (cheap for small structs, and enforce immutability by default); pointer receivers mutate shared state in place and avoid copying large structures.

```bash
go run ./advanced/memory
go build -gcflags="-m" ./advanced/memory
```

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

Note there's no `require` block: `fmt`, `reflect`, and `container/list` are all part of the Go **standard library**, which ships with every Go installation and never needs to be declared as a dependency.

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
| **Purpose** | `go vet` statically analyzes code for suspicious constructs (wrong `Printf` verbs, the `string(int32)` conversion pattern); `go fmt` rewrites files to Go's canonical formatting |
| **When to run** | `go vet ./...` before committing to catch bugs like the pitfalls in [this file](#pitfalls-already-in-this-file-and-how-to-fix-them); `gofmt -w .` (or `go fmt ./...`) to auto-format |
| **Notes** | `go vet` only flags misused format verbs and suspicious conversions in `fmt`-family and builtin-conversion contexts — a genuine parse error (like `if` used as an expression) is instead rejected by the compiler itself, before `vet` ever runs |

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
| `interface conversion: interface {} is X, not Y` | Type-asserting a `list.Element.Value` (or any `any`) to the wrong concrete type | Use the two-value form, `v, ok := element.Value.(int)`, and check `ok` before trusting `v` |
| `cannot call pointer method on X` / `cannot take the address of X` | Calling a pointer-receiver method on a value that isn't addressable (e.g. a map value, or a literal) | Store the value in a local variable first, or make the collection hold pointers (`[]*Person`) instead of values |

## License

See [LICENSE](./LICENSE) for details.
