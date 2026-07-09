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
- [Typical Workflow](#typical-workflow)
- [Troubleshooting](#troubleshooting)
- [License](#license)

## Overview

This project contains a minimal Go program (`hello.go`) meant to demonstrate the basic Go toolchain workflow: initializing a module, building a binary, and running the code directly.

## Prerequisites

- [Go](https://go.dev/dl/) installed (version 1.21+ recommended)
- Verify your installation:

  ```bash
  go version
  ```

## Project Structure

```
Gophers/
├── LICENSE
├── README.md
└── hello.go
```

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
