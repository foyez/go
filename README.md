# Golang

**[Go Playground](https://play.golang.org/)**

Go is a **compiled, statically typed programming language** designed at Google in **2007** and **open-sourced in 2009**.

**Core Design Goals:**
- **Simplicity**: Easy to learn, minimal keywords
- **Performance**: Compiled to native machine code
- **Concurrency**: Built-in support via goroutines and channels
- **Productivity**: Fast compilation, strong tooling

**Key Characteristics:**

| Feature | Description |
|---------|-------------|
| **Type System** | Statically typed with type inference |
| **Compilation** | Compiles to single binary (no dependencies) |
| **Memory Management** | Automatic garbage collection |
| **Concurrency Model** | Goroutines (lightweight threads) and channels |
| **Package Management** | Go modules (go.mod) |

**Real-World Use Cases:**
- Backend services (Docker, Kubernetes)
- Cloud infrastructure (Terraform, Consul)
- DevOps tools (GitHub CLI, Hugo)
- Distributed systems (Prometheus, etcd)

---

## Why Go?

<details>
<summary>View contents</summary>

### 1. Strongly Typed

In Go, **every variable has a type**, and that type must be known at **compile time**.

```go
var age int = 25
name := "Alice" // inferred as string
```

* You cannot accidentally mix incompatible types
* Many bugs are caught **before the program runs**

This leads to safer and more predictable code.

---

### 2. Statically Typed

Once a variable has a type, **it cannot change** during runtime.

```go
var x int = 10
x = "hello" // ❌ compile-time error
```

Benefits:

* Better performance
* Easier reasoning about code
* Strong IDE support (autocompletion, refactoring)

---

### 3. High Performance (Near C / C++)

Go is a **compiled language** that produces native machine code.

* Much faster than interpreted languages
* Slightly slower than C/C++ due to garbage collection
* Excellent performance for servers and network applications

```
Go vs Interpreted Languages:
- 10-100x faster than Python/Ruby/Node.js
- Near C/C++ performance with safety guarantees

Go vs JVM Languages:
- Faster startup times (no JVM warmup)
- Lower memory footprint
- Single binary deployment
```

---

### 4. Simple and Readable (Python-like)

Go intentionally avoids complex language features.

* No inheritance
* No operator overloading
* Minimal keywords

The result is:

* Code that looks similar across projects
* Easier onboarding for new developers
* Less "clever" but more maintainable code

---

### 5. Fast Compilation & Execution

Go is known for **very fast compile times**.

* Designed for large codebases
* Ideal for microservices
* Quick feedback loop during development

---

### 6. Automatic Garbage Collection

Go automatically manages memory:

* Allocates memory when needed
* Frees unused memory

This reduces:

* Memory leaks
* Manual memory management bugs

Unlike Java, Go’s garbage collector is designed to have **low latency**, making it suitable for high-performance systems.

---

### 7. Concurrency Made Simple

Concurrency is a **core feature** of Go, not an add-on.

Key concepts:

* **Goroutines**: Lightweight threads managed by Go
* **Channels**: Safe communication between goroutines
* **sync package**: Mutexes, WaitGroups, etc.

```go
// Traditional threading (complex)
Thread t = new Thread(new Runnable() {
    public void run() { /* work */ }
});
t.start();

// Go goroutines (simple)
go doWork()
```

Go can efficiently run **thousands or millions of concurrent tasks**.

---

### Key Reasons to Choose Go (In short)

* Easy to learn and read
* Built-in concurrency (goroutines & channels)
* Compiles to a single binary
* Fast execution (close to C/C++)
* Strong standard library
* Memory-safe (garbage collected)

---

</details>

## Parallelism vs Multithreading vs Concurrency

<details>
<summary>View contents</summary>

### Parallelism

Running multiple tasks **at the same time** on multiple CPU cores.

Example:

* Downloading
* Uploading
* Rendering

All happening simultaneously.

---

### Multithreading

Using multiple OS threads to do work concurrently.

Example:

* Watching a YouTube video
* Commenting
* Loading recommendations

Go handles threads internally so developers don’t manage them directly.

---

### Concurrency (Go’s Core Strength)

Concurrency means **handling many tasks at once**, but not necessarily executing them simultaneously.

Example:

* Multiple users booking tickets
* Multiple users editing the same document

Go excels at structuring concurrent programs safely and clearly.

---

</details>

## Go vs Other Languages

<details>
<summary>View contents</summary>

### Go vs C

| Feature           | Go             | C                   |
| ----------------- | -------------- | ------------------- |
| Level             | High-level     | Low-level           |
| Memory management | Automatic (GC) | Manual              |
| Performance       | Very fast      | Extremely fast      |
| Safety            | High           | Low                 |
| Concurrency       | Built-in       | Manual (OS threads) |
| Ease of learning  | Easy           | Hard                |

**Choose Go when**

* Building servers, APIs, microservices
* You want speed + safety

**Choose C when**

* Writing OS, drivers, embedded systems
* You need full hardware control

**Interview Tip:**
*Go trades low-level control for developer productivity and safety.*

---

### Go vs C++

| Feature      | Go             | C++                     |
| ------------ | -------------- | ----------------------- |
| Syntax       | Simple         | Complex                 |
| Memory       | GC             | Manual / smart pointers |
| Compile time | Fast           | Slow                    |
| OOP          | No inheritance | Full OOP                |
| Performance  | Very fast      | Extremely fast          |

**Go is better for**

* Backend & cloud services
* Large teams (easy maintenance)

**C++ is better for**

* Game engines
* High-performance systems

---

### Go vs Java

| Feature        | Go              | Java    |
| -------------- | --------------- | ------- |
| Runtime        | Compiled binary | JVM     |
| Memory         | GC              | GC      |
| Startup time   | Fast            | Slower  |
| Concurrency    | Goroutines      | Threads |
| Code verbosity | Low             | High    |

**Why choose Go over Java**

* Faster startup
* Less boilerplate
* Lightweight concurrency

**Why choose Java**

* Mature ecosystem
* Enterprise frameworks (Spring)

**Interview Tip:**
*Go is simpler and more lightweight than Java for cloud-native apps.*

---

### Go vs Python

| Feature        | Go               | Python        |
| -------------- | ---------------- | ------------- |
| Speed          | Fast             | Slow          |
| Typing         | Static           | Dynamic       |
| Concurrency    | Excellent        | Limited (GIL) |
| Learning curve | Moderate         | Very easy     |
| Use cases      | Backend, systems | AI, scripting |

**Choose Go when**

* Performance matters
* High concurrency is needed

**Choose Python when**

* Rapid prototyping
* AI / ML / Data Science

**Interview Tip:**
*Python is productivity-first; Go is performance + simplicity.*

---

### Go vs JavaScript (Node.js)

| Feature       | Go               | Node.js         |
| ------------- | ---------------- | --------------- |
| Language type | Compiled         | Interpreted     |
| Concurrency   | Goroutines       | Event loop      |
| Performance   | Higher           | Moderate        |
| Typing        | Static           | Dynamic         |
| Use case      | Backend services | Full-stack apps |

**Why Go**

* Better CPU utilization
* Safer for large backend systems

**Why Node.js**

* Same language for frontend & backend
* Large ecosystem

---

### Go vs Rust

| Feature        | Go        | Rust                 |
| -------------- | --------- | -------------------- |
| Memory safety  | GC        | No GC (ownership)    |
| Learning curve | Easy      | Hard                 |
| Performance    | Very fast | Extremely fast       |
| Concurrency    | Simple    | Powerful but complex |

**Choose Go**

* Faster development
* Easier learning

**Choose Rust**

* Maximum performance
* No garbage collection

**Interview Tip:**
*Go prioritizes simplicity; Rust prioritizes correctness.*

---

## Decision Framework:

```
Need maximum performance + no GC?
Need low-level hardware access?
  → Rust/C

Need simplicity + good performance?
  → Go

Need data science/AI/ML?
  → Python is better

Need enterprise ecosystem?
  → Java

Need full-stack JavaScript?
Need to build UI-heavy apps?
  → Javascript/Node.js
```

---

## Final Summary

| Language   | Best For                      |
| ---------- | ----------------------------- |
| Go         | Backend, microservices, cloud |
| C          | OS, embedded systems          |
| C++        | Games, high-performance       |
| Java       | Enterprise apps               |
| Python     | AI, scripting                 |
| Rust       | Safe system programming       |
| JavaScript | Web apps                      |

---

</details>

## Practice Questions (Go Basics)

<details>
<summary>View contents</summary>

**Fill in the Blanks:**

1. Go is a __________ typed language, meaning types are known at compile time.
2. Go compiles to a __________ binary with no external dependencies.
3. The Go concurrency primitive is called a __________.
4. Go's memory is managed by automatic __________ collection.

**True/False:**

1. ⬜ Go requires a virtual machine like Java
2. ⬜ Go supports inheritance for object-oriented programming
3. ⬜ Go's concurrency is based on goroutines and channels
4. ⬜ Go is slower than Python for computational tasks

**Multiple Choice:**

1. What is the main advantage of Go over Java?
   - A) Larger ecosystem
   - B) Faster startup and lower memory usage
   - C) Better IDE support
   - D) More language features

2. Which scenario is Go NOT ideal for?
   - A) Building microservices
   - B) Machine learning research
   - C) Writing CLI tools
   - D) Creating web servers

---

**Answers:**

<details>
<summary>View answers</summary>

**Fill in the Blanks:**
1. statically
2. single
3. goroutine
4. garbage

**True/False:**
1. ❌ False (Go compiles to native code)
2. ❌ False (Go uses composition, not inheritance)
3. ✅ True
4. ❌ False (Go is significantly faster)

**Multiple Choice:**
1. **B** - Faster startup and lower memory usage
2. **B** - Machine learning research (Python is better)

</details>

</details>

---

## Go Setup

<details>
<summary>View contents</summary>

### 1. Install Go

Download and install Go from the **official website**:

👉 [https://go.dev/doc/install](https://go.dev/doc/install)

Why this matters:

* Installs the Go compiler (`go`)
* Installs standard tools (`go build`, `go test`, `go fmt`, etc.)
* Keeps your system aligned with official releases

After installation, verify:

```bash
go version
```

---

### 2. Environment Variables

Go uses environment variables to locate:

* Your workspace
* Compiled binaries

```bash
go env
```

**Important environment variables:**
- `GOROOT`: Go installation directory
- `GOPATH`: Workspace for Go code (legacy, less used with modules)
- `GOMODCACHE`: Module cache location

---

#### Environment Configuration

#### Bash Shell

```bash
# ~/.bash_profile or ~/.bashrc

# Go binary path
export PATH=$PATH:/usr/local/go/bin

# Personal Go binaries
export PATH=$PATH:$HOME/go/bin
```

Apply changes:
```bash
source ~/.bashrc
```

Why this matters:

* `$HOME/go/bin` allows running installed Go tools
* `/usr/local/go/bin` allows running the Go compiler

---

### 3. Update Go Version

Remove old version:

```bash
$ sudo rm -rf /usr/local/go
```

Then install the latest version from:
👉 [https://go.dev/doc/install](https://go.dev/doc/install)

---

### 4. VS Code Setup

**Required Extension:**
- **Go** (official): `golang.go`

**Recommended Extensions:**
- **Error Lens**: Inline error display
- **Better Comments**: Enhanced comment highlighting
- **Proto3**: for gRPC

**VS Code Settings (`settings.json`):**

```json
{
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
	"editor.quickSuggestions": {
      "other": "off",
      "comments": "off",
      "strings": "off"
    }
  },
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package",
  "go.testFlags": ["-v", "-race"],
  "go.buildTags": "",
  "go.toolsManagement.autoUpdate": true,
  "protoc": {
    "options": ["--proto_path=proto"]
  }
}
```

**Install Go Tools:**

Press `Cmd+Shift+P` (Mac) or `Ctrl+Shift+P` (Windows/Linux), then:
```
Go: Install/Update Tools
```

Select all tools and install.

---

</details>

## Project Example

<details>
<summary>View contents</summary>

**Modern Go uses modules** (not GOPATH):

```bash
# Create project directory (anywhere, not just in GOPATH)
mkdir myproject
cd myproject

# Initialize Go module
go mod init github.com/username/myproject

# This creates go.mod file
```

**Project Structure:**
```
myproject/
├── go.mod          # Module definition
├── go.sum          # Dependency checksums
├── main.go         # Entry point
├── internal/       # Private packages
│   └── util/
└── pkg/            # Public packages
    └── api/
```

**main.go:**
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

**Run it:**
```bash
go run main.go
# Output: Hello, World!
```

**Build it:**
```bash
go build
./hello
# Output: Hello, World!
```

Why Go Modules:

* Dependency versioning
* No GOPATH dependency
* Reproducible builds

`go.mod`

```go
module github.com/username/myproject // module path

go 1.18 // minimum Go version
```

---

</details>

## Essential Go Commands

<details>
<summary>View contents</summary>

```bash
# Run program
go run main.go

# Build binary
go build                    # Creates ./myproject
go build -o custom-name     # Custom binary name

# Install binary to $GOPATH/bin
go install

# Format code (ALWAYS use this)
go fmt ./...

# Static analysis
go vet ./...

# View documentation
go doc fmt.Println
go doc -http :8080  # Start local doc server

# Manage dependencies
go get package@version
go mod tidy         # Clean up dependencies
go mod download     # Download dependencies

# Testing
go test ./...
go test -v -cover ./...
go test -race ./...  # Race condition detection

# View documentation
go doc fmt.Println

# Install dependency or tool
go install golang.org/x/lint/golint@latest

# Run linter
golint
```

---

</details>

## Anatomy of a Go File

<details>
<summary>View contents</summary>

`main.go`

```go
// package name
/*
 Every go program needs at least one package main
 Go programs are organized into packages
 A package is a collection of source files
*/
package main

import (
  "fmt" // import built-in packages

  "github.com/foyez/hello/utils" // import custom packages
)

// Every go program needs one main function
// It's the entry point for the program where
// go starts executing the code
func main() {
  fmt.Println("Hello, World")
  fmt.Print(utils.Add(10, 20))
}
```

`utils/utils.go`

```go
package utils

// private variable - start with lowercase
// Can't be accessed from other packages
var s = "Hello"

// Public function - starting with capital letter
// Can be accessed from other packages
func Add(a int64, b int64) int64 {
  return a + b
}
```

Why capitalization matters:

* Uppercase → exported (public)
* Lowercase → unexported (private)

---

</details>

## Practice Questions (Essential Commands)

<details>
<summary>View contents</summary>

**Fill in the Blanks:**

1. The command to initialize a Go module is `go __________ init`.
2. The `go __________` command formats Go code automatically.
3. Go modules are defined in a file named __________.
4. The `go __________` command runs static analysis on code.

**True/False:**

1. ⬜ Go requires code to be in the GOPATH directory
2. ⬜ `go fmt` should be run before committing code
3. ⬜ `go build` creates an executable binary
4. ⬜ VS Code requires manual Go tool installation

**Multiple Choice:**

1. Which command runs tests with race detection?
   - A) `go test -v`
   - B) `go test -race`
   - C) `go test -cover`
   - D) `go test -bench`

2. Where should private packages be placed?
   - A) `pkg/`
   - B) `internal/`
   - C) `private/`
   - D) `src/`

---

**Answers**:

<details>
<summary>View answers</summary>

**Fill in the Blanks:**
1. mod
2. fmt
3. go.mod
4. vet

**True/False:**
1. ❌ False (modules work anywhere)
2. ✅ True (always format before commit)
3. ✅ True
4. ❌ False (VS Code can install them automatically)

**Multiple Choice:**
1. **B** - `go test -race`
2. **B** - `internal/` (cannot be imported by external packages)

</details>

---

</details>

## Go Keywords (All 25)

<details>
<summary>View contents</summary>

```txt
break        case        chan        const       continue

default     defer       else        fallthrough for

func         go          goto        if          import

interface    map         package     range       return

select       struct      switch      type        var
```

Why this matters:

* Go intentionally has a **small keyword set**
* Easier to learn
* Easier to read
* Less hidden behavior

---

</details>

## TDD with Go (Test-Driven Development)

Test-Driven Development (TDD) is a development approach where you **write tests before writing production code**.
The goal is to let tests **drive the design and behavior** of your code.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/tddWithGo)**

### 1. Write the test first

We start by writing a test for functionality that **does not exist yet**.

#### `hello_test.go`

```go
package hello

import "testing"

// exported if it begins with a capital letter
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Foyez")
		want := "Hello, Foyez"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
```

#### Key concepts

* Test files **must end with** `_test.go`
* Tests live in the **same package** as the code (`package hello`)
* `TestHello`:

  * Must start with `Test`
  * Takes `*testing.T`
* `t.Run`:

  * Creates a **subtest**
  * Makes tests easier to read and extend later

#### Run the test

```sh
go test
```

#### Output

```sh
./hello_test.go:7:10: undefined: Hello
```

This is **expected**.
The test fails because `Hello` does not exist yet.

This is the **first success in TDD**:

> The test correctly tells us what is missing.

---

### 2. Write the minimal code to make the test compile

Now we write the **smallest possible code** to satisfy the compiler.

#### `hello.go`

```go
package hello

func Hello(name string) string {
	return ""
}
```

#### Run the test again

```sh
go test
```

#### Output

```sh
hello_test.go:11: got "" want "Hello, Foyez"
```

The test **runs and fails**, which is exactly what we want.

At this stage:

* Compiler errors are gone
* Logic is still incorrect

---

### 3. Write enough code to make the test pass

Now we implement just enough logic to satisfy the test.

```go
func Hello(name string) string {
	return "Hello, " + name
}
```

#### Run the test

```sh
go test
```

#### Output

```sh
PASS
ok      hello   0.004s
```

🎉 The test passes.

---

### 4. Commit the code

Once tests are passing, **commit the working state**.

```sh
git commit -m "add Hello() - greeting to people"
```

Why commit now?

* You have a **green (passing) test**
* This is a stable checkpoint

---

### 5. Refactor (without changing behavior)

Refactoring means:

> Improve code structure **without changing behavior**

Since we already have tests, refactoring is **safe**.

#### `hello.go`

```go
package hello

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
```

#### Improve the test for readability and reuse

#### `hello_test.go`

```go
package hello

import "testing"

func TestHello(t *testing.T) {
	assertErrorMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Foyez")
		want := "Hello, Foyez"
		assertErrorMessage(t, got, want)
	})
}
```

#### Important details

* `testing.TB`:

  * Interface implemented by both `*testing.T` and `*testing.B`
* `t.Helper()`:

  * Marks this function as a helper
  * Error line numbers point to the **test**, not the helper

#### Run tests again

```sh
go test
```

```sh
PASS
ok      hello   0.004s
```

---

### 6. Amend the git commit

Since refactoring didn’t change behavior, we **amend the previous commit**.

```sh
git commit --amend
```

This keeps git history clean.

---

### TDD Workflow Summary

1. Write a test
2. Make the compiler pass
3. Run the test and see it fail
4. Write enough code to make it pass
5. Refactor
6. Repeat

---

### 7. Add a Benchmark test

Benchmarks measure **performance**, not correctness.

```go
func BenchmarkHello(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		Hello("Zayan")
	}
}
```

#### Notes

* `b.N` is automatically adjusted by Go
* `testing.Short()` allows skipping benchmarks in quick test runs

#### Run benchmark

```sh
go test -v --bench . --benchmem
```

#### Output

```sh
BenchmarkHello    2000000000          0.46 ns/op
```

This means:

* The function ran **2,000,000,000 times**
* Each call took **~0.46 nanoseconds**
* Extremely fast, as expected

---

### 8. Add Example tests (documentation + tests)

Example tests:

* Act as **tests**
* Appear in **Go documentation**
* Must match output exactly

```go
func ExampleHello() {
	greeting := Hello("Zayan")
	fmt.Println(greeting)
	// Output: Hello, Zayan
}

func ExampleHello_second() {
	greeting := Hello("Farah")
	fmt.Println(greeting)
	// Output: Hello, Farah
}
```

#### Run examples

```sh
go test -v
```

```sh
=== RUN   TestHello
=== RUN   TestHello/saying_hello_to_people
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/saying_hello_to_people (0.00s)
=== RUN   ExampleHello
--- PASS: ExampleHello (0.00s)
=== RUN   ExampleHello_second
--- PASS: ExampleHello_second (0.00s)
```

---

### Using `testify` for Cleaner Unit Tests in Go

Go’s standard `testing` package is powerful and minimal, but as tests grow, you often repeat:

* Equality checks
* Error handling
* Failure messages

[`testify`](https://github.com/stretchr/testify) solves this by providing:

* Better assertions
* Cleaner syntax
* More readable test failures

---

#### 1. Install Testify

Add `testify` to your project using Go modules:

```sh
go get github.com/stretchr/testify
```

This adds it to `go.mod` automatically.

---

#### 2. Why use `testify/assert`?

Compare this:

##### Standard library

```go
if got != want {
	t.Errorf("got %q want %q", got, want)
}
```

##### With `testify`

```go
assert.Equal(t, want, got)
```

Benefits:

* Less boilerplate
* Better failure output
* Easier to read and maintain

---

#### 3. Rewrite existing test using `testify/assert`

##### `hello_test.go`

```go
package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Foyez")
		want := "Hello, Foyez"

		assert.Equal(t, want, got)
	})
}
```

##### What `assert.Equal` does

* Compares expected (`want`) and actual (`got`)
* Automatically fails the test if they differ
* Prints a **clear diff-style error message**

---

#### 4. `assert` vs `require`

Testify provides two main assertion packages:

| Package   | Behavior                               |
| --------- | -------------------------------------- |
| `assert`  | Fails the test but continues execution |
| `require` | Fails the test and stops immediately   |

##### Example difference

```go
assert.Equal(t, want, got)
```

* Test continues even if this fails

```go
require.Equal(t, want, got)
```

* Test stops immediately if this fails

##### Use cases

* Use `assert` when multiple checks are independent
* Use `require` when later checks depend on earlier ones

---

#### 5. Using `require` in tests

```go
package hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Foyez")
		want := "Hello, Foyez"

		require.Equal(t, want, got)
	})
}
```

---

#### 6. Table-driven tests with Testify (recommended)

Table-driven tests are the **idiomatic Go way** to test multiple cases.

```go
func TestHello_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		input string
		want  string
	}{
		{
			name:  "normal name",
			input: "Foyez",
			want:  "Hello, Foyez",
		},
		{
			name:  "another name",
			input: "Zayan",
			want:  "Hello, Zayan",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Hello(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
```

##### Why this is powerful

* Easy to add new test cases
* No duplicated test logic
* Clear test intent

---

#### 7. Assertions commonly used in real projects

```go
assert.Equal(t, expected, actual)
assert.NotEqual(t, a, b)

assert.Nil(t, err)
assert.NotNil(t, value)

assert.True(t, condition)
assert.False(t, condition)

assert.Len(t, collection, 3)
assert.Contains(t, "hello world", "world")
```

These replace many `if` statements and manual error messages.

---

#### 8. Using Testify with Benchmarks (important note)

⚠️ **Do not use assertions inside benchmarks**

Benchmarks should measure **performance only**.

Correct benchmark (unchanged):

```go
func BenchmarkHello(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		Hello("Zayan")
	}
}
```

---

#### 9. Should you always use Testify?

##### Pros

* Cleaner and more expressive tests
* Better error messages
* Faster test writing

##### Cons

* External dependency
* Slightly less “pure stdlib”

##### Recommendation

* ✅ Use `testify` in **application code**
* ⚠️ Stdlib is fine for **very small libraries**

Most production Go teams use `testify`.

---

</details>

## Package Structure

**Every Go file starts with a package declaration:**

```go
package main  // Executable package (must have main function)

package mylib  // Library package
```

<details>
<summary><strong>View contents</strong></summary>

**Package Naming Rules (Google Style):**
- Use **lowercase**, **single-word** names
- Package name = directory name
- Avoid underscores or mixed caps

```
✅ Good:
   package http
   package json
   package user

❌ Bad:
   package HTTP
   package jsonParser
   package user_service
```

**Importing Packages:**

```go
package main

import (
	"fmt"           // Standard library
	"net/http"      // Standard library (nested)
	
	"github.com/gorilla/mux"  // External package
	"example.com/myproject/internal/util"  // Internal package
)
```

**Import Aliases:**

```go
import (
	"fmt"
	str "strings"  // Alias to avoid naming conflicts
)

func main() {
	str.ToUpper("hello")
}
```

**Import Side Effects (rare):**

```go
import _ "image/png"  // Registers PNG decoder without using package
```

---

</details>

## Visibility Rules

**Capitalization determines visibility:**

<details>
<summary><strong>View contents</strong></summary>

```go
package user

// Exported (public) - accessible from other packages
type User struct {
	Name string  // Exported field
	Age  int     // Exported field
}

// Unexported (private) - only accessible within this package
type session struct {
	token string  // Unexported field
}

// Exported function
func CreateUser(name string) *User {
	return &User{Name: name}
}

// Unexported function
func validateEmail(email string) bool {
	return true
}
```

**Key Rule:**
```
Uppercase first letter = Exported (public)
Lowercase first letter = Unexported (private)
```

---

</details>

## Variables and Constants

Variables store values that can be used and modified during program execution.

<details>
<summary><strong>View contents</strong></summary>

### Variable Declaration

**Method 1: Explicit Type**
```go
var name string = "Alice"
var age int = 30
var isActive bool = true
```

**Method 2: Type Inference**
```go
var name = "Alice"  // Type inferred as string
var age = 30        // Type inferred as int
```

**Method 3: Short Declaration (inside functions only)**
```go
name := "Alice"
age := 30
```

**Multiple Variable Declaration:**
```go
var x, y int = 1, 2
a, b := "hello", true

var (
	name   string = "Alice"
	age    int    = 30
	active bool   = true
)
```

### Zero Values

**Variables without initialization get zero values:**

```go
var i int       // 0
var f float64   // 0.0
var s string    // "" (empty string)
var b bool      // false
var p *int      // nil
var sl []int    // nil
var m map[string]int  // nil
```

**Real-World Example:**
```go
type Counter struct {
	value int  // Automatically 0
}

c := Counter{}
fmt.Println(c.value)  // Output: 0
```

### Constants

**Constants are immutable:**

```go
const Pi = 3.14159
const MaxConnections = 100

// Multiple constants
const (
	StatusOK    = 200
	StatusError = 500
)

// Typed constants
const MaxRetries int = 3

// iota (auto-incrementing)
const (
	Sunday = iota  // 0
	Monday         // 1
	Tuesday        // 2
)
```

**Why use constants:**
- Prevent accidental modification
- Improve code readability
- Enable compiler optimizations

---

</details>

## Data Types

Go is a **statically typed language**, meaning variable types are known at compile time.

<details>
<summary><strong>View contents</strong></summary>

### Basic Types

```go
// Integers
int, int8, int16, int32, int64      // Signed
uint, uint8, uint16, uint32, uint64 // Unsigned

// Floating-point
float32, float64

// Complex numbers
complex64, complex128

// String
string

// Boolean
bool

// Byte (alias for uint8)
byte

// Rune (alias for int32, represents Unicode code point)
rune
```

| Name        | Type Name                                                              | Examples                   |
| ----------- | ---------------------------------------------------------------------- | -------------------------- |
| **INTEGER** | int, int8, int16, int32, int64<br/>uint, uint8, uint16, uint32, uint64 | var age int = 20           |
| **FLOAT**   | float32, float64                                                       | var gpa float64 = 3.4      |
| **STRING**  | string                                                                 | var fruit string = "mango" |
| **BOOLEAN** | bool                                                                   | var adult bool = age > 18  |

---

**Platform-dependent types:**
- `int` and `uint` are 32 bits on 32-bit systems, 64 bits on 64-bit systems

---

**Boolean Operators:**

```go
&&  // AND
||  // OR
!   // NOT
< <= > >= == !=
```

```go
var adult bool = age >= 18
```

---

**Identify type:**

```go
reflect.TypeOf(6) // int
```

Useful for debugging and learning, not common in production code.

---

### Type Conversion

**Go requires explicit conversion:**

```go
var i int = 42
var f float64 = float64(i)  // Must convert explicitly
var u uint = uint(f)
var value float64 = float64(10) + 5.5 // 15.5

// ❌ This won't compile:
// var f float64 = i
```

**Real-World Example:**
```go
func calculateAverage(nums []int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))  // Must convert to float64
}
```

---

### Strings

**Strings are immutable sequences of bytes.**

```go
s := "Hello, 世界"

// Multi-line strings
s3 := `This is a
multi-line
string`
```

---

</details>

## Input and Output

This section explains how Go handles **output (printing)** and **input (reading data)** using the `fmt` and `os` packages.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/printing)**

### Printing

**Print:**

* Print output to the **stdout console**
* Return:

  * Number of bytes written
  * An error (`error`)

```go
import "fmt"

// Print without newline
fmt.Print("Hello")

// Print with newline
fmt.Println("Hello")

// Formatted print
fmt.Printf("Name: %s, Age: %d\n", "Alice", 30)
```

**Common Format Verbs:**

| Verb | Type | Example |
|------|------|---------|
| `%v` | Default format | `%v` → `{Name:Alice Age:30}` |
| `%+v` | Struct with field names | `%+v` → `{Name:Alice Age:30}` |
| `%#v` | Go-syntax representation | `%#v` → `main.User{Name:"Alice", Age:30}` |
| `%T` | Type | `%T` → `main.User` |
| `%d` | Integer | `%d` → `42` |
| `%f` | Float | `%f` → `3.141593`, `%.2f` → `3.14` |
| `%s` | String | `%s` → `hello` |
| `%q` | Quoted string | `%q` → `"hello"` |
| `%t` | Boolean | `%t` → `true` |
| `%p` | Pointer | `%p` → `0xc0000b2000` |

**Example:**
```go
type User struct {
	Name string
	Age  int
}

u := User{Name: "Alice", Age: 30}

fmt.Printf("%v\n", u)   // {Alice 30}
fmt.Printf("%+v\n", u)  // {Name:Alice Age:30}
fmt.Printf("%#v\n", u)  // main.User{Name:"Alice", Age:30}
fmt.Printf("%T\n", u)   // main.User
```

---

**String Formatting (without printing):**

```go
msg := fmt.Sprintf("User: %s, Age: %d", "Alice", 30)
// msg = "User: Alice, Age: 30"
```

---

### Reading Input

Go reads user input from **standard input (`stdin`)** using functions from the `fmt` package and utilities from `bufio`.

**Scan:**

```go
import "fmt"

var name string
fmt.Print("Enter your name: ")
fmt.Scan(&name)  // Reads until whitespace
fmt.Println("Hello,", name)
```

**Reading Lines (with bufio):**

```go
import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
	text = strings.TrimSpace(text)  // Remove newline
	
	fmt.Println("You entered:", text)
}
```

---

**Reading Multiple Values:**

```go
var a, b int
fmt.Print("Enter two numbers: ")
fmt.Scanf("%d %d", &a, &b)
fmt.Println("Sum:", a+b)
```

**Command-Line Arguments (`os.Args`):**

* Reads arguments passed from the command line
* Returns a slice of strings (`[]string`)

```go
import (
	"fmt"
	"os"
)

func main() {
	// os.Args[0] is program name
	// os.Args[1:] are actual arguments
	
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <arg>")
		os.Exit(1)
	}
	
	arg := os.Args[1]
	fmt.Println("Argument:", arg)
}
```

```bash
go run main.go hello
# Output: Argument: hello
```

---

### Less Common Input & Output

<details>
<summary><strong>View contents</strong></summary>

### `fmt.Fprint / Fprintln / Fprintf`

* Writes output to an io.Writer ((e.g., a file, `os.Stdout`, `os.Stderr`))
* Returns (bytesWritten, error)
* Used with files, network connections, HTTP responses

```go
import (
	"fmt"
	"os"
)

fmt.Fprint(writer, "Hello")           // no newline
fmt.Fprintln(writer, "Hello")         // with newline
fmt.Fprintf(writer, "Age: %d", age)   // formatted
```

Common writers:

```go
os.Stdout   // terminal output
os.Stderr   // error output
*os.File    // files
```

#### Example: Printing to stdout

```go
fmt.Fprint(os.Stdout, "Hello")
fmt.Fprintln(os.Stdout, "Hello")
fmt.Fprintf(os.Stdout, "Name: %s\n", "Alice")
```

#### Example: Printing to a file

```go
file, _ := os.Create("output.txt")
defer file.Close()

fmt.Fprintln(file, "Hello File")
fmt.Fprintf(file, "Score: %d\n", 100)
```

> **Use `fmt` for user-facing output (CLI, results, messages).**

---

## 2. Logging (`log` package)

* Opinionated logging for production
* Writes to stderr by default
* Automatically adds date & time
* Safe for concurrent use

Basic usage:

```go
log.Print("app started")
log.Println("user logged in")
log.Printf("user %s logged in", name)
```

Example output:

```
2025/01/01 10:15:30 app started
```

### Fatal & Panic Logs

```go
log.Fatal("error")  // prints + os.Exit(1), deferred funcs NOT run
log.Panic("error")  // prints + panic(), deferred funcs DO run
```

| Function  | Stops program | Deferred runs |
| --------- | ------------- | ------------- |
| log.Fatal | Yes           | No            |
| log.Panic | Yes           | Yes           |

---

### Customizing Logs

```go
log.SetPrefix("INFO: ")
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
```

Common flags:

```go
log.Ldate        // date
log.Ltime        // time
log.Lmicroseconds // Microsecond precision
log.Lshortfile   // file:line
log.Llongfile // Full file path
```

---

### Logging to a File

```go
file, err := os.OpenFile(
	"app.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY,
	0644,
)
if err != nil {
	log.Fatal(err)
}

log.SetOutput(file)
log.Println("logging to file")
```

> `log` uses `io.Writer` internally (same idea as `fmt.Fprint`).

---

### fmt vs log

```text
fmt → user output (CLI, results)
log → system/debug/error messages
```

| Feature    | fmt | log         |
| ---------- | --- | ----------- |
| Timestamp  | No  | Yes         |
| Stdout     | Yes | No (stderr) |
| Production | No  | Yes         |

---

## 3. Reading Input (`fmt.Fscan`)

* Reads from any io.Reader (stdin, file, network)
* Faster and more flexible than fmt.Scan

### Recommended Pattern

```go
in := bufio.NewReader(os.Stdin)

var n int
fmt.Fscan(in, &n)

fmt.Println("Number:", n)
```

Multiple values:

```go
var a, b int
fmt.Fscan(in, &a, &b)
fmt.Println(a + b)
```

Input:

```
10 20
```

---

### Scan vs Fscan

| Feature      | Scan       | Fscan         |
| ------------ | ---------- | ------------- |
| Input source | stdin only | any io.Reader |
| Speed        | slower     | faster        |
| Usage        | learning   | real-world    |

---

### When to Use What

```text
Simple learning      → fmt.Scan
Large input / CP     → bufio + fmt.Fscan
Files / network      → fmt.Fscan
CLI output           → fmt
Errors / debugging   → log
```

---

</details>

</details>

## Practice Questions (Packages, Variables, Data Types, etc.)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. A package name should be __________, single-word, and match the directory name.
2. Variables starting with an __________ letter are exported (public).
3. The __________ value for an uninitialized int is 0.
4. The `%__` format verb prints the type of a variable.

**True/False:**

1. ⬜ Go allows implicit type conversion between int and float64
2. ⬜ Strings in Go are immutable
3. ⬜ The `len()` function returns the number of Unicode characters in a string
4. ⬜ Variables declared with `:=` can only be used inside functions

**Multiple Choice:**

1. Which is the correct way to iterate over Unicode characters?
   - A) `for i := 0; i < len(s); i++`
   - B) `for i, c := range s`
   - C) `for c in s`
   - D) `for each c in s`

2. What is the zero value for a pointer?
   - A) 0
   - B) ""
   - C) nil
   - D) undefined

**Code Output:**

```go
var x int
var y = 10
z := 20

fmt.Printf("%d %d %d", x, y, z)
```

What is printed?

---

**Answers:**

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. lowercase
2. uppercase
3. zero
4. T

**True/False:**
1. ❌ False (explicit conversion required)
2. ✅ True
3. ❌ False (returns bytes, not characters)
4. ✅ True

**Multiple Choice:**
1. **B** - `for i, c := range s` (iterates over runes)
2. **C** - nil

**Code Output:**
```
0 10 20
```
(x gets zero value 0, y and z are initialized)

---

</details>

</details>

## Control Structures: If & Else

Go uses `if`, `else if`, and `else` for conditional logic.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/control)**

**Basic Syntax:**

```go
if condition {
	// code
}

if condition {
	// code
} else {
	// code
}

if condition1 {
	// code
} else if condition2 {
	// code
} else {
	// code
}
```

**Important:**
- **No parentheses** around condition
- **Braces are mandatory** (even for single-line blocks)

**Examples:**

```go
age := 20

if age >= 18 {
	fmt.Println("Adult")
} else {
	fmt.Println("Minor")
}

// With multiple conditions
if age < 13 {
	fmt.Println("Child")
} else if age < 18 {
	fmt.Println("Teenager")
} else {
	fmt.Println("Adult")
}
```

### if with Initialization

**Short variable declaration in if:**

```go
if err := doSomething(); err != nil {
	// err only exists in this if-else block
	return err
}
// err doesn't exist here
```

**Real-World Example:**

```go
import "os"

// Good: err scoped to if block
if file, err := os.Open("config.txt"); err != nil {
	fmt.Println("Error:", err)
} else {
	defer file.Close()
	// use file
}

// Bad: err pollutes outer scope
file, err := os.Open("config.txt")
if err != nil {
	// ...
}
```

**Google Style Guide Recommendation:**
- Use initialization in `if` to limit variable scope
- Reduces variable pollution in outer scope

---

</details>

## Control Structures: switch

Go’s `switch` is **more powerful and flexible** than many other languages.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/control)**

**Basic Switch:**

```go
day := "Monday"

switch day {
case "Monday":
	fmt.Println("Start of the week")
case "Friday":
	fmt.Println("End of the week")
case "Saturday", "Sunday":
	fmt.Println("Weekend!")
default:
	fmt.Println("Midweek")
}
```

**Key Differences from C/Java:**
- **No `break` needed** (no fall-through by default)
- Can switch on any comparable type
- Multiple values per case

**Switch Without Expression (replaces if-else chains):**

```go
age := 25

switch {
case age < 13:
	fmt.Println("Child")
case age < 18:
	fmt.Println("Teenager")
case age < 65:
	fmt.Println("Adult")
default:
	fmt.Println("Senior")
}
```

**Switch with Initialization:**

```go
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
	fmt.Println("Weekend")
default:
	fmt.Println("Weekday")
}
```

**Type Switch (checking interface types):**

```go
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}

describe(42)      // Integer: 42
describe("hello") // String: hello
```

**Fallthrough (explicit fall-through):**

```go
// Rare! Usually avoid fallthrough
n := 3

switch n {
case 1:
	fmt.Println("One")
	fallthrough  // Continues to next case
case 2:
	fmt.Println("Two")
case 3:
	fmt.Println("Three")
	fallthrough
default:
	fmt.Println("End")
}

// Output:
// Three
// End
```

---

</details>

## Loops

Go has **only one loop keyword**: `for`. But it supports multiple patterns.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/loops)**

### Traditional for Loop

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
// Output: 0 1 2 3 4
```

* Initialization; condition; increment

---

### While-style Loop

```go
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}
```

---

### Infinite Loop

```go
num := 1

for {
	num = num + 2

	if num == 5 {
		continue
	}

	fmt.Print(num)

	if num == 7 {
		break
	}
}
```

* `continue` skips to next iteration
* `break` exits the loop

---

### Basic for loop iteration (string bytes)

```go
var name = "Farah"

for i := 0; i < len(name); i++ {
	fmt.Println("Letter:", string(name[i]))
}
```

* Iterates over **bytes**, not Unicode characters

---

### Range-based Loop

**Basic for loop:**

```go
for i := range 5 {
	fmt.Println(i)
}
// Output: 0 1 2 3 4
```

**Slice/Array:**
```go
nums := []int{10, 20, 30}

for i, v := range nums {
	fmt.Printf("Index: %d, Value: %d\n", i, v)
}

// Ignore index
for _, v := range nums {
	fmt.Println(v)
}

// Only index
for i := range nums {
	fmt.Println(i)
}
```

**Map:**
```go
scores := map[string]int{"Alice": 90, "Bob": 85}

for name, score := range scores {
	fmt.Printf("%s: %d\n", name, score)
}

// Only keys
for name := range scores {
	fmt.Println(name)
}
```

* Map iteration order is **random**

---

**String (iterates over runes):**
```go
for i, r := range "Hello, 世界" {
	fmt.Printf("%d: %c\n", i, r)
}
```

**Channel:**
```go
ch := make(chan int)

go func() {
	ch <- 1
	ch <- 2
	close(ch)
}()

for val := range ch {
	fmt.Println(val)
}
```

* `range` receives values until channel is closed

#### Loop Control

**break:**
```go
for i := 0; i < 10; i++ {
	if i == 5 {
		break  // Exit loop
	}
	fmt.Println(i)
}
```

**continue:**
```go
for i := 0; i < 5; i++ {
	if i == 2 {
		continue  // Skip this iteration
	}
	fmt.Println(i)
}
// Output: 0 1 3 4
```

**Labeled break (breaking nested loops):**
```go
outer:
for i := 0; i < 3; i++ {
	for j := 0; j < 3; j++ {
		if i == 1 && j == 1 {
			break outer  // Breaks outer loop
		}
		fmt.Printf("(%d,%d) ", i, j)
	}
}
```

---

### Common Patterns

#### Early Return Pattern (Google Style)

**Preferred:**
```go
func processUser(id int) error {
	user, err := getUser(id)
	if err != nil {
		return err  // Early return
	}
	
	if !user.Active {
		return errors.New("user inactive")  // Early return
	}
	
	// Happy path - no nesting
	return saveUser(user)
}
```

**Avoid:**
```go
// Bad: Nested if statements
func processUser(id int) error {
	user, err := getUser(id)
	if err == nil {
		if user.Active {
			return saveUser(user)
		} else {
			return errors.New("user inactive")
		}
	}
	return err
}
```

#### Guard Clauses

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")  // Guard clause
	}
	return a / b, nil
}
```

---

</details>

## Practice Questions (Control Structures and Loops)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. In Go, switch statements do NOT require __________ to prevent fall-through.
2. The __________ keyword skips the current iteration of a loop.
3. To iterate over a map's keys and values, use `for __, __ := range map`.
4. An if statement with initialization limits variable __________ to the if block.

**True/False:**

1. ⬜ Parentheses around if conditions are optional in Go
2. ⬜ Go's for loop can function as a while loop
3. ⬜ Multiple values can be tested in a single case of a switch
4. ⬜ The range loop over strings iterates over bytes

**Multiple Choice:**

1. What does this loop print?
   ```go
   for i := 0; i < 3; i++ {
       if i == 1 {
           continue
       }
       fmt.Print(i)
   }
   ```
   - A) 0 1 2
   - B) 0 2
   - C) 1 2
   - D) 0 1

2. Which is the correct way to create an infinite loop?
   - A) `for (;;)`
   - B) `while (true)`
   - C) `for {}`
   - D) `loop {}`

**Code Challenge:**

Write a function that returns "Fizz" for multiples of 3, "Buzz" for multiples of 5, "FizzBuzz" for multiples of both, or the number as a string otherwise.

---

**Answers:**

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. break
2. continue
3. key, value
4. scope

**True/False:**
1. ❌ False (parentheses are not allowed)
2. ✅ True
3. ✅ True
4. ❌ False (iterates over runes/Unicode characters)

**Multiple Choice:**
1. **B** - 0 2 (skips 1 due to continue)
2. **C** - `for {}`

**Code Challenge:**
```go
func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
```

---

</details>

</details>

## Functions and Methods

Functions are reusable blocks of code that perform a specific task.
Go functions are **strongly typed**, support **multiple return values**, and treat functions as **first-class citizens**.

A **method** is a function with a **receiver**.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/funtions)**

### Function Basics

#### Function Declaration

**Basic Syntax:**
```go
func functionName(param1 type1, param2 type2) returnType {
	// function body
	return value
}
```

**Examples:**

```go
// No parameters, no return
func sayHello() {
	fmt.Println("Hello")
}

// With parameters and return
func add(a int, b int) int {
	return a + b
}

// Shortened parameter list (same type)
func multiply(a, b int) int {
	return a * b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
```

#### Multiple Return Values

**Common Pattern: (value, error)**

```go
import (
	"errors"
	"os"
)

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Usage
content, err := readFile("config.txt")
if err != nil {
	log.Fatal(err)
}
fmt.Println(content)
```

**Multiple Values:**
```go
func getUserInfo(id int) (string, int, bool) {
	return "Alice", 30, true
}

name, age, active := getUserInfo(1)
```

**Ignoring Return Values:**
```go
// Ignore error (NOT recommended in production)
data, _ := os.ReadFile("file.txt")

// Ignore all but one
_, _, active := getUserInfo(1)
```

---

### Named Return Values

**Named returns are pre-declared in the signature:**

```go
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return  // Naked return (returns result and err)
	}
	result = a / b
	return  // Naked return
}
```

**When to use:**
- ✅ Short functions (improves documentation)
- ✅ Error handling patterns
- ❌ Long functions (can be confusing)

**Google Style Recommendation:**
- Use named returns sparingly
- Prefer explicit returns in complex functions

**Example:**
```go
// Good: Clear and concise
func getData() (value int, err error) {
	if err = validate(); err != nil {
		return
	}
	value = fetchValue()
	return
}

// Better: Explicit returns for complex logic
func processData() (int, error) {
	if err := validate(); err != nil {
		return 0, err
	}
	
	value := fetchValue()
	return value, nil
}
```

---

### Variadic Functions (unknown number of arguments)

**Accept unlimited arguments:**

```go
func sum(nums ...int) int {
	total := 0
  // n - treated as slice
	for _, n := range nums {
		total += n
	}
	return total
}

// Usage
fmt.Println(sum(1, 2, 3))        // 6
fmt.Println(sum(1, 2, 3, 4, 5))  // 15

// Spread or unpack slice into variadic function
values := []int{1, 2, 3}
fmt.Println(sum(values...))      // 6
```

**Real-World Example:**

```go
import "fmt"

func logError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("[ERROR]", msg)
}

logError("User %s not found", "Alice")
logError("Invalid input: %d", 42)
```

**Rules:**
- Variadic parameters are treated as slices
- Variadic parameter must be **last**
- Only **one** variadic parameter per function

---

### Lexical scoping

Go functions are **lexically scoped**, meaning:

> A function can access variables defined in the same or outer block.

```go
var n1 = 5

func foo(n2 int) {
	n3 := 8
	fmt.Println(n1, n2, n3)
}
```

* `n1` → package scope
* `n2` → function parameter
* `n3` → function local variable

---

### Functions as Values

**Functions are first-class citizens:**

Functions can be:

* Assigned to variables
* Passed as arguments
* Returned from functions

---

**Assign function to variable:**

```go
// Assign function to variable
var add func(int, int) int = func(a, b int) int {
	return a + b
}

result := add(3, 4)  // 7
```

---

**Passing functions as arguments:**

```go
func apply(f func(int, int) int, a, b int) int {
	return f(a, b)
}

multiply := func(a, b int) int {
	return a * b
}

result := apply(multiply, 3, 4)  // 12
```

---

**Returning functions:**

```go
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// "factor" is in the closure of "func(x int) int"

double := multiplier(2)
triple := multiplier(3)

fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

---

### Closures

**Functions can capture variables from outer scope:**

```go
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// "count" is in the closure of "func() int"

c1 := counter()
fmt.Println(c1())  // 1
fmt.Println(c1())  // 2
fmt.Println(c1())  // 3

c2 := counter()
fmt.Println(c2())  // 1 (separate counter)
```

When functions are passed or returned, **their environment comes with them**.
`count` remains available even after `counter` finishes execution.

**Real-World Example (HTTP middleware):**

```go
import "net/http"

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
```

---

### Defer

A `defer` statement delays execution of a function **until the surrounding function returns**.

#### Key Rules

1. Deferred calls execute in **LIFO order** (stack)
2. Arguments are **evaluated immediately**
3. Execution happens **even after panic**

#### Common Uses of `defer`

* Closing files
* Unlocking mutexes
* Database cleanup
* Recovering from panics

---

**Defer executes a function call AFTER the surrounding function returns:**

```go
func processFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()  // Guaranteed to run when function exits
	
	// Work with file
	// Even if error occurs, file.Close() will be called
	return nil
}
```

**Multiple Defers (LIFO order):**

```go
func example() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Main")
}

// Output:
// Main
// Third
// Second
// First
```

**Defer with Arguments (evaluated immediately):**

```go
func example() {
	x := 10
	defer fmt.Println(x)  // x is evaluated NOW (10)
	x = 20
	fmt.Println(x)
}

// Output:
// 20
// 10
```

**Real-World Patterns:**

```go
// Resource cleanup
func readDatabase() error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	
	// Query database
	return nil
}

// Mutex unlock
func updateCounter(mu *sync.Mutex, counter *int) {
	mu.Lock()
	defer mu.Unlock()  // Always unlocks, even if panic
	
	*counter++
}

// Transaction rollback
func transfer(db *sql.DB, from, to int, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()  // Rollback if Commit is not called
	
	// ... perform transfer operations
	
	return tx.Commit()  // Overrides rollback
}
```

---

### Methods

A **method** is a function with a **receiver**.

**Syntax:**

```go
func (r ReceiverType) funcName(parameters) (results)
```

---

**Methods vs Functions:**

* Functions operate on values passed as arguments
* Methods are **called on a value of a specific type**

---

**Example:**

```go
type address struct {
	email   string
	zipCode int
}

type User struct {
	name string
	age  int
	address
}

// Function Version
func updateUserName(u *User, name string) {
	u.name = name
}

// Method Version
func (u *User) UpdateName(name string) {
	// (*u).name = name
	u.name = name
}

// Usage
func main() {
	user := User{
		name: "Manam",
		age: 25,
		address: address{
			email: "manam@email.com",
			zipCode: 34000,
		},
	}

	updateUserName(&user, "Chayon")

	// (&user).UpdateName("Chayon")
	user.UpdateName("Chayon")
}
```

---

**Methods are functions with a receiver:**

```go
type Rectangle struct {
	Width  float64
	Height float64
}

// Value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Usage
rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())  // 50

rect.Scale(2)
fmt.Println(rect.Area())  // 200 (modified)
```

**Why Methods Matter:**

* Better readability
* Encapsulation
* Enables interfaces
* Object-oriented style without classes

#### Value vs Pointer Receivers

**Value Receiver:**
- Receives a **copy** of the struct
- Cannot modify original struct
- Use when struct is **small** and method doesn't need to modify it

**Pointer Receiver:**
- Receives a **reference** to the struct
- Can modify original struct
- Use when:
  - Method needs to modify struct
  - Receiver (struct) is **large** (avoid copying)
  - Consistency (if any method uses pointer receiver, use it for all)

**Example:**

```go
type Counter struct {
	Value int
}

// Value receiver - doesn't modify original
func (c Counter) GetValue() int {
	return c.Value
}

// Pointer receiver - modifies original
func (c *Counter) Increment() {
	c.Value++
}

func main() {
	counter := Counter{Value: 0}
	
	counter.Increment()  // Go auto-converts to (&counter).Increment()
	fmt.Println(counter.GetValue())  // 1
}
```

**Google Style Recommendation:**
- **Consistency**: If any method uses pointer receiver, use pointer receivers for **all** methods on that type
- Prevents confusion about which methods modify the receiver

---

</details>

## Practice Questions (Functions and Methods)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. Functions in Go can return __________ values.
2. A __________ receiver is used when a method needs to modify the struct.
3. The `defer` keyword executes a function after the surrounding function __________.
4. A function that accepts unlimited arguments uses the __________ parameter syntax.

**True/False:**

1. ⬜ Named return values are always better than explicit returns
2. ⬜ Deferred functions execute in LIFO (last-in-first-out) order
3. ⬜ Value receivers receive a copy of the struct
4. ⬜ A function can have multiple variadic parameters

**Multiple Choice:**

1. What does this print?
   ```go
   func test() {
       defer fmt.Print("1")
       defer fmt.Print("2")
       fmt.Print("3")
   }
   test()
   ```
   - A) 123
   - B) 321
   - C) 312
   - D) 213

2. When should you use a pointer receiver?
   - A) Always
   - B) When the struct is small
   - C) When you need to modify the struct
   - D) Never

**Code Challenge:**

Write a function that takes a slice of integers and a filter function, then returns a new slice with only elements that pass the filter.

---

**Answers:**

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. multiple
2. pointer
3. returns
4. variadic

**True/False:**
1. ❌ False (use named returns sparingly)
2. ✅ True
3. ✅ True
4. ❌ False (only one, and it must be last)

**Multiple Choice:**
1. **C** - 312 (prints 3, then defers execute in reverse: 2, then 1)
2. **C** - When you need to modify the struct (also for large structs)

**Code Challenge:**
```go
func filter(nums []int, fn func(int) bool) []int {
	result := []int{}
	for _, n := range nums {
		if fn(n) {
			result = append(result, n)
		}
	}
	return result
}

// Usage
nums := []int{1, 2, 3, 4, 5, 6}
evens := filter(nums, func(n int) bool { return n%2 == 0 })
fmt.Println(evens)  // [2 4 6]
```

---

</details>

</details>

## Arrays

Arrays have a **fixed length** and store elements of the same type.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/arrays)**

### Declaring arrays

```go
// ARRAY
// [number]T
// A slice type has a specific length
var arr [3]float64
fmt.Println(arr) // [0 0 0]
```

* Length is part of the type
* Zero values are assigned automatically

---

### Accessing elements

```go
arr[1] = 23               // set element
element := arr[1]         // read element
fmt.Println(arr, element) // [0 23 0] 23
```

---

### Declaring and initializing

```go
scores := [3]float64{9, 1.5, 2.2}
fmt.Println(scores)
```

---

### Compiler-determined length

```go
arrNotMax := [...]int{2, 3, 4}
fmt.Println(arrNotMax, len(arrNotMax)) // [2 3 4] 3
```

---

### Slicing an array

```go
fruits := [5]string{"banana", "pear", "apple", "orange", "peach"}

splicedFruits := fruits[1:3]              // [pear apple]
splicedFruits2 := fruits[2:]              // [apple orange peach]
removeLastFruit := fruits[:len(fruits)-1] // [banana pear apple orange]
lastFruit := fruits[len(fruits)-1]        // peach

fmt.Println(splicedFruits, splicedFruits2, removeLastFruit, lastFruit)
fmt.Println(len(splicedFruits)) // 2
fmt.Println(cap(splicedFruits)) // 4
```

---

### Append behavior

```go
fruitsToAdd := append(splicedFruits, "cherry", "pineapple", "guava")
```

* If capacity is exceeded:

  * A **new underlying array** is allocated
  * Capacity usually **doubles**

---

### Prepend (not built-in)

```go
nums := []int{1, 2, 3}
nums = append([]int{0}, nums...)
```

---

### Multidimensional array

```go
multi := [2][3]int{{1, 2, 3}, {5, 6, 7}}
fmt.Println(multi) // [[1 2 3] [5 6 7]]
```

---

</details>

## Slices

Slices are **more flexible than arrays** and are used almost everywhere in Go.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/slices)**

### Declaring a slice

```go
// SLICE
// []T
// A slice type has no specific length

var mySlice []int
fmt.Println(mySlice) // []
```

❌ This will panic:

```go
// mySlice[0] = 1
```

Because the slice has **no allocated memory yet**.

---

## Slice internals

A slice has **three properties**:

* `ptr` – pointer to the underlying array
* `len` – number of elements
* `cap` – capacity of the underlying array

![image](https://user-images.githubusercontent.com/11992095/202870508-0739d792-8747-4e20-8cd2-0ffa888d5c08.png)

---

### Slice assignment copies metadata, not data

```go
var s = []int{1, 2, 3}
var s2 = s

s2[0] = 5

fmt.Println(s, s2) // [5 2 3] [5 2 3]
```

* Both slices point to the **same underlying array**
* Modifying one affects the other

---

## make() with slices

`make` initializes and allocates memory.

```go
// make([]T, len, cap)
s := make([]int, 0, 3)
```

⚠️ Important correction
This line is invalid:

```go
sliceWithMake[0] = 1
```

Because `len == 0`

Correct usage:

```go
s = append(s, 1)
```

---

### Capacity growth example

```go
for i := 0; i < 5; i++ {
	s = append(s, i)
	fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
}
```

Output:

```sh
cap 3, len 1, 0xc0000b2000
cap 3, len 2, 0xc0000b2000
cap 3, len 3, 0xc0000b2000
cap 6, len 4, 0xc0000b8000
cap 6, len 5, 0xc0000b8000
```

* Capacity grows
* Underlying array changes when capacity exceeded

---

### Unpack / spread slice

```go
var fruits = []string{"apple", "mango"}

// variable argument/variadic function
func addFruits(fruitsToAdd ...string) []string {
	// unpack or spread
	updatedFruits := append(fruits, fruitsToAdd...)
	return updatedFruits
}

addFruits("banana", "pineapple")
```

---

</details>

## Maps

In Go, a **map** is a built-in data type that stores **key–value pairs**.
It is similar to:

* Object (JavaScript)
* Dictionary (Python)
* HashMap (Java)

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/maps)**

### Map Basics

#### Declaring and creating a map

```go
var results map[string]float64 = make(map[string]float64) // create empty map
```

Key points:

* `map[string]float64` → key type is `string`, value type is `float64`
* `make` is required to allocate memory
* A `nil` map cannot be written to

---

#### Reading from a map

```go
fmt.Println(results["test"]) // 0
```

Important:

* If a key does **not exist**, Go returns the **zero value** of the value type
* For `float64`, zero value is `0`

---

#### Writing to a map

```go
results["foyez"] = 3.4
results["mithu"] = 3.5

fmt.Println(results) // map[foyez:3.4 mithu:3.5]
```

---

### Map literal (initialize with values)

```go
userEmails := map[int]string{
	1: "user1@email.com",
	2: "user2@email.com",
}
```

* Cleaner and more common
* `make` is implicit

---

#### Updating values

```go
userEmails[1] = "user12@email.com"
```

* If key exists → value updated
* If key does not exist → key created

---

### Checking if a key exists (`comma ok` idiom)

```go
emailOfSecondUser, ok := userEmails[2]
emailOfFourthUser, ok2 := userEmails[4]

fmt.Println(emailOfSecondUser, ok)  // user2@email.com true
fmt.Println(emailOfFourthUser, ok2) // "" false
```

Key points:

* `ok == true` → key exists
* `ok == false` → key does not exist
* Returned value is **zero value** if key is missing

---

#### Idiomatic existence check

```go
if email, ok := userEmails[2]; ok {
	fmt.Printf("%s exists\n", email)
} else {
	fmt.Println("email doesn't exist")
}
```

---

### Deleting from a map

```go
delete(userEmails, 1)
fmt.Println(userEmails) // map[2:user2@email.com]
```

Notes:

* Safe to delete a non-existing key
* No error is thrown

---

### Important Map Characteristics

* Map keys must be **comparable**

  * Allowed: `int`, `string`, `bool`, structs (if fields are comparable)
  * Not allowed: `slice`, `map`, `function`
* Map iteration order is **random**
* Maps are **reference types**

  * Assigning a map copies the reference, not the data

---

### Iterating Over a Map

```go
users := map[string]interface{}{
	"name":     "zayan",
	"age":      5,
	"religion": "islam",
}

for key, val := range users {
	fmt.Printf("%s -> %v\n", key, val)
}
```

#### Notes

* `interface{}` allows **mixed value types**
* `%v` prints the value in default format
* Order is **not guaranteed**

---

### When to Use `map[string]interface{}`

✅ Useful for:

* JSON-like data
* Dynamic or unknown schemas

⚠️ Avoid in:

* Strongly typed business logic
* Performance-critical code

---

</details>

## Strings

In Go, strings are:

* **Immutable sequences of bytes**
* **UTF-8 encoded**
* Internally a **read-only byte slice**

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/strings)**

**`len(string)` counts bytes, not characters:**

```go
len("hello") // 5
len("ক")     // 3
```

Because:

* Strings are UTF-8 encoded
* Non-ASCII characters use multiple bytes

---

**Indexing a string returns a byte:**

```go
p("hello"[1]) // 101
s := "Hello"

// Indexing returns bytes
fmt.Println(s[0])  // 72 (byte value, not character)

// To convert to character:
fmt.Println(string(s[0])) // E
```

Explanation:

* `'H'` → ASCII value `72`
* Type is `byte`, not `string`

---

**Concatenation & multi-line string:**

```go
s := "Hello"

// Concatenation
s2 := s + " - Go" // Hello - Go

// Multi-line strings
s3 := `This is a
multi-line
string`
```

**Iterating Strings Safely (Unicode):**

```go
s := "Hello, 世界"

// Iterate over bytes (WRONG for Unicode)
for i := 0; i < len(s); i++ {
	fmt.Printf("%c ", s[i])
}
// Output: H e l l o ,   ä ¸  ç  ...

// Iterate over runes (CORRECT for Unicode)
for i, r := range s {
	fmt.Printf("%d: %c\n", i, r)
}
// Output:
// 0: H
// 1: e
// 2: l
// 3: l
// 4: o
// 5: ,
// 6:  
// 7: 世
// 10: 界
```

* `range` iterates over **runes**, not bytes
* Correct way to handle Unicode

---

**String Immutability:**

❌ This is invalid:

```go
s := "hello"
s[0] = 'H' // compile-time error
```

✅ Correct approach:

```go
s = "Hello"
```

---

**Common String Operations:**

```go
import "strings"

strings.Contains("hello", "ll")        // true
strings.HasPrefix("hello", "he")       // true
strings.HasSuffix("hello", "lo")       // true
strings.Split("a,b,c", ",")            // ["a", "b", "c"]
strings.Join([]string{"a", "b"}, "-")  // "a-b"
strings.ToUpper("hello")               // "HELLO"
strings.Replace("hello", "l", "L", -1) // "heLLo"
strings.Replace("fooo", "o", "O", 2)   // "fOOo"

strings.Count("test", "t")             // 2
strings.Index("test", "t")             // 0
strings.LastIndex("test", "t")         // 3
strings.Repeat("a", 5)                 // aaaaa
len("hello")                           // 5
"hello"[1]                             // 101 (byte value of 'e')
"Let's" + " - Go"                      // Let's - Go
```

---

</details>

## Structs

A **struct** is a composite data type that groups together variables (fields) under one name.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/structs)**

### Why use structs?

* To represent **real-world entities**
* To organize related data
* To define **custom types**

---

### Basic Struct Definition

```go
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}
```

* `type User struct {}` defines a **new type**
* Field names starting with **capital letters are exported** (accessible outside the package)
* Field names starting with **lowercase letters are unexported**

---

### Creating a Struct Value

```go
user := User{
	ID:        1,
	FirstName: "Foyez",
	LastName:  "Ahmed",
	Email:     "foyez@email.com",
}

fmt.Println(user.FirstName) // Foyez
```

#### Important Notes

* This is called **named-field initialization** (recommended)
* It prevents bugs when fields are reordered
* Unspecified fields get **zero values**

---

### Zero Values of Struct Fields

```go
var user User
fmt.Println(user)
```

Output:

```go
{0 "" "" ""}
```

* `int` → `0`
* `string` → `""`
* `bool` → `false`
* slices/maps → `nil`

---

### Struct Pointers

You can create a pointer to a struct:

```go
user := &User{
	ID:        1,
	FirstName: "Foyez",
}
```

Access fields using **dot notation** (Go automatically dereferences):

```go
fmt.Println(user.FirstName)
```

---

</details>

## Pointers

A **pointer** stores the **memory address** of a variable instead of a copy of its value.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/pointers)**

### Pointer Syntax

```go
// Declare a pointer variable
var variableName *type

// Get the address of a variable
&variableName

// Dereference a pointer (get the value it points to)
*variableName
```

---

### Example: Pointer with Struct

```go
type person struct {
	firstName      string
	lastName       string
	favoriteSports []string
}
```

---

#### Main Function

```go
func main() {
	person := person{
		firstName: "Foyez",
		lastName:  "Ahmed",
		favoriteSports: []string{"Cricket"},
	}

	updateFirstName(&person, "Rumon")
	fmt.Println(person) // {Rumon Ahmed [Cricket]}

	updateFavoriteSports(person, "Football")
	fmt.Println(person) // {Rumon Ahmed [Football]}
}
```

---

### Updating Struct Fields Using Pointer

```go
func updateFirstName(p *person, newFirstName string) {
	fmt.Println(p)  // &{Foyez Ahmed [Cricket]}
	fmt.Println(&p) // address of pointer variable itself
	fmt.Println(*p) // {Foyez Ahmed [Cricket]}

	// (*p).firstName = newFirstName
	p.firstName = newFirstName
}
```

#### Key Concepts

* `p` is a pointer to `person`
* `*p` gives the actual struct value
* `p.firstName` works because **Go automatically dereferences pointers**

---

### Why This Works Without Pointer?

```go
func updateFavoriteSports(p person, sportName string) {
	p.favoriteSports[0] = sportName
}
```

#### Explanation

* `person` is passed **by value**
* BUT `slice` is a **reference type**
* Both `p.favoriteSports` and `person.favoriteSports` point to the **same underlying array**

---

### Value Types vs Reference Types (Very Important)

#### Value Types

```
int, float, string, bool, struct, array
```

* Passed **by value**
* A **copy** is created
* Changes inside a function **do NOT affect original data**
* Use **pointer** to modify original value

---

#### Reference Types

```
slice, map, channel, pointer, function
```

* Internally contain a **pointer**
* Passed **by value**, but reference same data
* Changes inside a function **DO affect original data**
* Pointer usually **not required**

---

### Call by Value (Default Behavior)

```go
type Person struct {
	name string
	age  int
}
```

#### Function Example

```go
func updateAge(p Person) {
	p.age = 20
	fmt.Println(p) // {Mithu 20}
}
```

#### Method Example

```go
func (p Person) updateAge() {
	p.age = 30
	fmt.Println(p) // {Mithu 30}
}
```

#### Main

```go
func main() {
	mithu := Person{name: "Mithu", age: 10}

	updateAge(mithu)
	fmt.Println(mithu) // {Mithu 10}

	mithu.updateAge()
	fmt.Println(mithu) // {Mithu 10}
}
```

#### Explanation

* `p` receives a **copy**
* Changes affect only the copy
* Original value remains unchanged

---

### Call by Reference (Using Pointers)

#### Function Example

```go
func updateAge(p *Person) {
	p.age = 20
	fmt.Println(*p) // {Mithu 20}
}
```

#### Method Example

```go
func (p *Person) updateAge() {
	p.age = 30
	fmt.Println(*p) // {Mithu 30}
}
```

#### Main

```go
func main() {
	mithu := Person{name: "Mithu", age: 10}

	updateAge(&mithu)
	fmt.Println(mithu) // {Mithu 20}

	mithu.updateAge()
	fmt.Println(mithu) // {Mithu 30}
}
```

---

### Important Go Feature: Automatic Addressing

```go
mithu.updateAge()
```

Even though `updateAge` expects `*Person`, Go automatically does:

```go
(&mithu).updateAge()
```

This makes pointer receivers **clean and safe to use**

---

### When to Use Pointer Receivers in Methods

Use pointer receivers when:

1. You want to **modify the receiver**
2. The struct is **large** (performance)
3. Consistency (recommended if any method uses pointer)

---

### Summary Cheat Sheet

| Concept                   | Behavior                |
| ------------------------- | ----------------------- |
| Struct                    | Collection of fields    |
| Pointer                   | Stores memory address   |
| Value Types               | Copied on function call |
| Reference Types           | Share underlying data   |
| Method (value receiver)   | Cannot modify original  |
| Method (pointer receiver) | Can modify original     |
| Slices                    | Reference type          |
| Maps                      | Reference type          |

---

</details>

## Error Handling

Go treats errors as **values**, not exceptions.
This design forces developers to **handle errors explicitly**, making programs more predictable and easier to reason about.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/errors)**

### Error

An **error** indicates that something went wrong, **but the program can usually continue**.

#### Characteristics

* Returned as the **last return value**
* Must be **checked explicitly**
* Does **not crash** the program by default

```go
type error interface {
	Error() string
}
```

Any type that implements `Error() string` satisfies the `error` interface.

---

#### Example

```go
err := funcReturnError()

if err != nil {
	fmt.Println(err.Error())
}
```

💡 **Idiomatic Go**

```go
if err != nil {
	return err
}
```

---

### Panic

A **panic** is a runtime error that:

* Immediately stops normal execution
* Unwinds the stack
* Executes deferred functions
* Crashes the program if not recovered

#### When to Panic

* Programmer mistakes (nil pointer dereference)
* Impossible states
* Unrecoverable errors

🚫 **Do NOT use panic for normal error handling**

---

#### Example

```go
panic(err.Error())
```

Common panic scenarios:

* Accessing out-of-bounds slice index
* Dereferencing a nil pointer
* Explicit `panic()`

---

### Error vs Panic

| Error             | Panic                    |
| ----------------- | ------------------------ |
| Recoverable       | Unrecoverable            |
| Returned as value | Stops program            |
| Must be handled   | Crashes if not recovered |
| Expected failures | Programming bugs         |

---

### Recover

`recover()` allows a program to **regain control after a panic**.

#### Important Rules

* Only works **inside a deferred function**
* Stops the panic and resumes normal execution
* Returns the value passed to `panic()`

---

#### Example

```go
func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
```

```go
func main() {
	defer recoverFromPanic()

	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 2 {
			panic("PANIC!")
		}
	}
}

// 0
// 1
// 2
// PANIC!
```

#### Explanation

* Panic happens at `i == 2`
* Deferred `recoverFromPanic()` executes
* Panic is handled → program does not crash

---

#### Best Practice

* Use `recover` only in **top-level goroutines**
* Never silently ignore panics

##### 1. “Use `recover` only in **top-level goroutines**”

**What it means:**

* A **goroutine** is a lightweight thread of execution in Go.
* Every goroutine has its own **call stack**. If a panic happens inside a goroutine **and isn’t recovered there**, it **kills the whole program**.
* “Top-level” here means the **main goroutine** or any goroutine that you explicitly control and launch.

**Example of good usage:**

```go
func main() {
    go safeGoroutine()
    fmt.Println("Main continues...")
}

func safeGoroutine() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in goroutine:", r)
        }
    }()

    panic("Something went wrong in goroutine")
}
```

* Panic happens in the goroutine
* `defer` + `recover` catches it **inside that goroutine**
* Program continues; main goroutine is safe

**Why “top-level”?**
If you try to `recover` **deep inside nested function calls**, it may not work as expected, because `recover` only stops **panics in the same goroutine where it’s deferred**.

---

##### 2. “Never silently ignore panics”

**What it means:**

* A panic usually indicates a **serious bug** or unexpected condition.
* If you recover from it **without logging, handling, or reporting**, you may **hide real issues**, making debugging very hard.

**Bad example:**

```go
defer func() {
    recover() // silently ignoring
}()
```

* You catch the panic but **do nothing**
* Bug remains, program continues in an inconsistent state

**Good example:**

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
        // optionally log to file, report error, clean resources
    }
}()
```

* Panic is **handled responsibly**
* Program can continue safely

---

**Rule of thumb:**

* Only use `recover` when you **can handle the panic meaningfully**, like:

  * Cleaning resources (closing files, network connections)
  * Logging errors
  * Gracefully shutting down a service

* Don’t use `recover` to **silently ignore bugs**, otherwise you’re hiding problems.

---

</details>

## Interfaces

**Interfaces** are one of the most powerful features in Go.
They allow you to define **behavior** without specifying how that behavior is implemented.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/interfaces)**

### Structs vs Interfaces

| Concept  | Struct                    | Interface                             |
| -------- | ------------------------- | ------------------------------------- |
| Purpose  | Holds **data/attributes** | Defines **behavior/methods**          |
| Contains | Fields                    | Method signatures (no implementation) |
| Usage    | Create instances          | Implemented by types implicitly       |
| Analogy  | “What an object **has**”  | “What an object **can do**”           |

---

### Defining an Interface

```go
type Shape2D interface {
	Area() float64
	Perimeter() float64
}
```

* `Shape2D` is an **interface type**
* Any type that has **both `Area()` and `Perimeter()` methods** automatically implements `Shape2D`
* **No explicit declaration needed** (unlike Java/C#)

---

### Implementing Interfaces with Structs

```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
```

✅ **Key Points**

* Both `Rectangle` and `Circle` implement `Shape2D`
* Go checks **struct methods** at compile time
* No need to explicitly declare `implements Shape2D`

---

### Using Interfaces as Function Parameters

```go
func fitInYard(s Shape2D) bool {
	return s.Area() > 200 && s.Perimeter() > 200
}
```

* `fitInYard` works with **any type** that implements `Shape2D`
* This is **polymorphism** in Go
* Allows **generic-like behavior** without generics

---

### Type Assertions

Sometimes you want to **access the concrete type** from an interface:

```go
func printShapeProps(s Shape2D) {
	if rect, ok := s.(Rectangle); ok {
		fmt.Printf("Height: %.2f, Width: %.2f\n", rect.Height, rect.Width)
	}
	if circle, ok := s.(Circle); ok {
		fmt.Printf("Radius: %.2f\n", circle.Radius)
	}
}
```

* `s.(Rectangle)` tries to assert that `s` is of type `Rectangle`
* `ok` is `true` if the assertion succeeds, `false` otherwise
* Safe: avoids panic

---

#### Example Usage

```go
func main() {
	circle := Circle{10}
	rectangle := Rectangle{10, 20}

	fmt.Println(fitInYard(circle))
	fmt.Println(fitInYard(rectangle))

	printShapeProps(rectangle) // Height: 20.00, Width: 10.00
	printShapeProps(circle)    // Radius: 10.00
}
```

---

### Empty Interface (`interface{}`)

* Defines **zero methods**
* Can hold **values of any type**
* Equivalent to `any` in TypeScript

```go
var people map[string]interface{} = make(map[string]interface{})

people["name"] = "Foyez"
people["age"] = 28

fmt.Printf("%#v %T\n", people["name"], people["name"]) // "Foyez" string
fmt.Printf("%#v %T", people["age"], people["age"])     // 28 int
```

**Use Cases**

* Storing **heterogeneous data**
* Functions like `fmt.Println` and JSON marshaling
* Type-agnostic containers

---

### Interfaces vs Structs Summary

| Feature      | Struct                            | Interface                      |
| ------------ | --------------------------------- | ------------------------------ |
| Stores       | Fields/data                       | Methods/behavior               |
| Implements   | N/A                               | Any type with matching methods |
| Inheritance  | No                                | Can embed interfaces           |
| Polymorphism | No                                | Yes                            |
| Example      | `Rectangle{Width: 10, Height: 5}` | `Shape2D`                      |

---

### Important Notes / Best Practices

1. **Implicit implementation**

   * A struct implements an interface **automatically** if it has required methods
   * No need to declare `implements`

2. **Pointers vs Values**

   * If a method has a **pointer receiver**, only a **pointer to struct** implements the interface

   ```go
   type MyInterface interface {
       Foo()
   }
   func (m *MyStruct) Foo() {}
   ```

   * `var x MyInterface = &MyStruct{}` ✅
   * `var x MyInterface = MyStruct{}` ❌

3. **Use interfaces for abstraction**

   * Don’t overuse empty interfaces; prefer **specific interfaces**

4. **Type switches** (advanced)

   * Safer than multiple type assertions

   ```go
   switch v := s.(type) {
   case Rectangle:
       fmt.Println("Rectangle", v.Width, v.Height)
   case Circle:
       fmt.Println("Circle", v.Radius)
   default:
       fmt.Println("Unknown shape")
   }
   ```

---

### Why Interfaces Matter

* Enable **polymorphism** without inheritance
* Decouple code: functions don’t need to know concrete types
* Allow **mocking/testing** by passing fake implementations

---

</details>

## Type Assertions

A **type assertion** extracts the **concrete value** from an interface.

<details>
<summary>View contents</summary>

### Syntax

```go
value, ok := interfaceValue.(ConcreteType)
```

---

### Example

```go
var foo interface{} = "Hello"

str := foo.(string)
fmt.Println(str) // "Hello"
```

---

### Unsafe Assertion (Causes Panic)

```go
num := foo.(int) // panic
fmt.Println(num)
```

---

### Safe Assertion (Recommended)

```go
num2, ok := foo.(int)
fmt.Println(num2, ok) // 0 false
```

---

### Why Type Assertions Are Needed

* Interfaces hide concrete types
* To access concrete behavior, you must assert
* Widely used in:

  * `error` handling
  * `fmt`
  * `context`
  * type switches

---

### Type Switch

```go
switch v := foo.(type) {
case string:
	fmt.Println("string:", v)
case int:
	fmt.Println("int:", v)
default:
	fmt.Println("unknown type")
}
```

</details>


## Go Type System & Relationships Diagram

<details>
<summary>View contents</summary>

```
              ┌────────────┐
              │  Struct    │
              │  (Data)    │
              │ FirstName, │
              │ LastName   │
              │ etc.       │
              └─────┬──────┘
                    │
         ┌──────────┴──────────┐
         │                     │
   Value Receiver          Pointer Receiver
  (func (s Struct) M())   (func (s *Struct) M())
         │                     │
         │                     │
         └──────────┬──────────┘
                    │
                Methods
                    │
         ┌──────────┴──────────┐
         │                     │
      Interface             Interface
  (defines required     (can accept pointer
     behavior)           receivers if needed)
         │
         ▼
   Any type that implements the
   interface methods
         │
         ▼
     Interface Value
     (holds concrete type + value/pointer)
         │
   ┌─────┴─────┐
   │           │
Type Assertion  Type Switch
(variable.(T))  switch v := variable.(type) {}
```

---

## Explanation of Each Part

### 1️⃣ Structs

* Hold **fields/data**
* Can have **methods** (value or pointer receiver)
* Example:

```go
type User struct {
    FirstName string
    LastName  string
}
```

---

### 2️⃣ Methods

* Methods are **functions tied to a struct**
* Can have **value receiver** `(u User)` or **pointer receiver** `(u *User)`
* Pointer receiver needed if you want to **modify the struct**

---

### 3️⃣ Interfaces

* Define **behavior only**
* Any type that has required methods **automatically implements the interface**
* Can accept **struct values or pointers**, depending on receiver type

```go
type Greeter interface {
    Greet() string
}
```

---

### 4️⃣ Interface Values

* Store **two things**:

  1. Concrete type (`User` or `*User`)
  2. Value/pointer to data
* Can pass **different types** to the same function using an interface

```go
func SayHello(g Greeter) {
    fmt.Println(g.Greet())
}
```

---

### 5️⃣ Type Assertions & Type Switches

* **Type assertion**: get the concrete type back from an interface

```go
u := g.(User)       // unsafe, panics if wrong
u, ok := g.(User)   // safe, returns ok=true/false
```

* **Type switch**: check multiple types safely

```go
switch v := g.(type) {
case User:
    fmt.Println("User", v.FirstName)
case Admin:
    fmt.Println("Admin", v.Level)
}
```

---

### 6️⃣ Pointers

* Important when:

  * You need to **modify a struct** in a method
  * Interface methods are defined on pointer receivers
* Go automatically handles:

```go
u := User{"Foyez", "Ahmed"}
u.UpdateName("Rumon") // auto-converts to (&u).UpdateName
```

---

## Visual Summary (Flow)

```
Struct (data)
    │
    ▼
Methods (value / pointer)
    │
    ▼
Interface (behavior)
    │
    ▼
Interface Value (type + pointer/value)
    │
    ├──> Type Assertion (recover concrete type)
    └──> Type Switch (branch on type)
```

---

✅ **Key Takeaways**

1. **Structs** = data
2. **Methods** = attach behavior to structs
3. **Interfaces** = abstract behavior (polymorphism)
4. **Pointer receivers** = modify data / satisfy interfaces
5. **Interface values** = hold concrete type + value/pointer
6. **Type assertions/switches** = access concrete type safely

---

</details>
 
## Generics

Generics (introduced in **Go 1.18**) allow you to write **type-safe, reusable code** without sacrificing compile-time checks.
 
<details>
<summary>View contents</summary>

### Problem: Using `interface{}` (Empty Interface)

```go
func isEqual(a, b interface{}) bool {
	return a == b
}

func main() {
	fmt.Println(isEqual(1, 1))   // true
	fmt.Println(isEqual(1, "1")) // ❌ NOT type-safe
}
```

#### Important Clarification

* This function **compiles**, but:

  * `interface{}` removes **type safety**
  * The compiler **cannot help you**
* You are allowing **any two values** to be compared, even if it makes no logical sense

💡 In reality:

```go
isEqual(1, "1") // returns false, but comparison itself is meaningless
```

The **real problem** is not panic here — it’s that:

* The compiler **should have rejected this code**
* Bugs move from **compile time → runtime**

---

### Why `interface{}` Is Dangerous Here

* You lose:

  * Compile-time guarantees
  * Meaningful comparisons
* You rely on **runtime behavior**
* Bugs become harder to detect

👉 This is exactly what generics were designed to fix.

---

### Solution: Generics (Type Parameters)

```go
func isEqual[T comparable](a, b T) bool {
	return a == b
}
```

#### What This Means

* `T` is a **type parameter**
* `comparable` is a **constraint**
* Only types that support `==` and `!=` are allowed

---

#### Usage

```go
func main() {
	fmt.Println(isEqual(1, 1))   // true
	fmt.Println(isEqual(1, "1")) // ❌ compile-time error
}
```

**Compiler error (GOOD thing):**

```
default type string of "1" does not match inferred type int for T
```

- This is **type safety**
- Bug caught **before the program runs**

---

### What Is `comparable`?

```go
func isEqual[T comparable](a, b T) bool
```

#### `comparable` (built-in constraint)

* Introduced in Go 1.18
* Allows types that can be compared using `==` and `!=`

Includes:

* `int`, `float`, `string`, `bool`
* pointers
* structs (if all fields are comparable)

Excludes:

* slices
* maps
* functions

---

### Constraints and Ordering

```go
import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Max returns the maximum of two values that implement the Ordered constraint.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
```

#### `constraints.Ordered`

* Allows:

  * `<`, `<=`, `>`, `>=`
* Works for:

  * integers
  * floats
  * strings

⚠️ Note:

* This package existed **before** Go finalized generics
* Today, many people define their **own constraints** instead

---

### Custom Constraints

```go
type Number interface {
	~int | ~int64 | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}
```

#### Why `~` (tilde)?

* Allows **underlying types**
* Makes generics work with custom type aliases

---

### Higher-Order Generic Functions

#### `Reduce`

```go
func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}
```

* `A` → element type
* `B` → accumulator/result type
* Extremely flexible
* Used for:

  * sums
  * concatenation
  * aggregation

---

#### `Find`

```go
func Find[A any](items []A, predict func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predict(v) {
			return v, true
		}
	}
	return
}
```

* Returns:

  * value
  * boolean indicating presence
* Avoids sentinel values (`nil`, `-1`, etc.)

---

#### `Filter`

```go
func Filter[A any](items []A, predict func(A) bool) []A {
	var founds []A

	for _, v := range items {
		if predict(v) {
			founds = append(founds, v)
		}
	}

	return founds
}
```

* Keeps type information
* No casting
* Fully compile-time safe

---

#### `Map`

```go
func Map[A, B any](items []A, modify func(A) B) []B {
	var modifiedItems []B
	for _, v := range items {
		modifiedItems = append(modifiedItems, modify(v))
	}
	return modifiedItems
}
```

* Transforms `[]A → []B`
* Classic functional pattern
* Impossible to do safely with `interface{}`

---

### Generics vs Interfaces

| Feature             | Interfaces                | Generics                |
| ------------------- | ------------------------- | ----------------------- |
| Purpose             | Abstraction over behavior | Abstraction over types  |
| Uses methods        | ✅                         | ❌                       |
| Compile-time safety | ✅                         | ✅                       |
| Polymorphism        | Dynamic                   | Static                  |
| Best for            | APIs, behavior            | Algorithms, collections |

---

#### Rule of Thumb

* Use **interfaces** when:

  * You care about *what something does*
* Use **generics** when:

  * You care about *what type something is*

---

### Generics vs `interface{}`

| Feature             | interface{} | Generics |
| ------------------- | ----------- | -------- |
| Type safety         | ❌           | ✅        |
| Compile-time checks | ❌           | ✅        |
| Casting required    | ✅           | ❌        |
| Performance         | Worse       | Better   |
| Readability         | Worse       | Better   |

---

### When NOT to Use Generics

❌ Don’t use generics when:

* You only have **one concrete type**
* Interfaces already solve the problem
* The generic version is harder to read

✅ Generics are powerful — but **not free**

---

### Mental Model

```
interface{}  → maximum flexibility, zero safety
interfaces   → behavior abstraction
generics     → type abstraction + safety
```

---

### How Go Generics Compile Internally

#### Important Truth (Core Idea)

👉 **Go generics are NOT implemented like C++ templates or TypeScript generics.**

Go uses a strategy called:

> **Monomorphization with sharing (a.k.a. dictionary passing)**

---

#### What That Means (Simplified)

When you write:

```go
func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}
```

And call it like:

```go
Max(10, 20)
Max(3.5, 2.1)
```

The compiler does **NOT** blindly generate infinite versions.

Instead, it does roughly this:

* Generates **specialized code per constraint group**
* Shares implementations when possible
* Passes **type-specific operations** (like `>`, `==`) via hidden dictionaries

---

#### Mental Model of Compilation

Conceptually, the compiler turns this:

```go
Max[int](10, 20)
```

Into something *like*:

```go
func Max_int(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

And:

```go
func Max_float64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
```

⚠️ **But**:

* Go does **not** literally generate one function per call
* It generates **shared instantiations** per type set

---

#### Why Go Did This

| Language | Strategy              | Result                          |
| -------- | --------------------- | ------------------------------- |
| C++      | Full monomorphization | Very fast, very large binaries  |
| Java     | Type erasure          | Smaller binaries, runtime casts |
| Go       | Hybrid                | Fast + controlled binary size   |

---

### Final Takeaway

* Generics solve problems that `interface{}` **never should**
* They move bugs from **runtime → compile time**
* They shine in:

  * collections
  * algorithms
  * utility libraries

---

</details>

## Concurrency

**Concurrency** means **dealing with multiple tasks at the same time**.

> Important distinction:
>
> * **Concurrency**: Structuring a program to handle many things at once
> * **Parallelism**: Actually running things at the same time on multiple CPU cores

A concurrent program **may or may not** run in parallel.

Go makes concurrency **easy to write and easy to reason about**.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/concurrency)**

### Why Concurrency Matters

Most real-world programs are naturally concurrent:

* Web servers handling many users
* APIs calling multiple services
* Background jobs + user requests
* Reading files while processing data
* Waiting on network responses

Without concurrency:

* Programs block unnecessarily
* CPU time is wasted
* Code becomes complex and fragile

---

### Key Concepts

#### 1 Process

* A **process** is an independent program in execution.
* It has its own memory, variables, and resources.
* Example: When you open a browser or run `go run main.go`, each runs as a separate process.
* Processes do **not share memory** by default. Communication between them requires mechanisms like IPC (Inter-Process Communication).

#### 2 Thread

* A **thread** is a smaller unit of execution within a process.
* Threads in the same process **share memory**, which allows faster communication but requires careful synchronization to avoid conflicts.
* Each process can have multiple threads. For example, a browser may use threads to render a page, handle network requests, and manage UI simultaneously.

#### 3 Goroutine

* A **goroutine** is Go’s lightweight abstraction over threads.
* They are **cheaper** than threads in terms of memory and startup cost.
* Created using the `go` keyword.
* Goroutines are managed by the Go runtime scheduler, which multiplexes many goroutines over a small number of OS threads.
* Unlike threads, you don’t manage goroutine lifecycle manually.

#### 4 Parallelism vs Concurrency

* **Concurrency**: Multiple tasks making progress at the same time (interleaved execution).

  * Example: Handling multiple network requests on a single CPU using goroutines.
* **Parallelism**: Multiple tasks running at the exact same time (requires multiple CPUs or cores).

  * Example: Rendering multiple images simultaneously using multiple CPU cores.

> Go supports both concurrency and parallelism via goroutines and GOMAXPROCS setting.

#### 5 Multithreading

* Traditional multithreading means creating multiple threads manually.
* Go abstracts this complexity using goroutines and a **scheduler**, so you rarely manage threads directly.
* OS threads are heavier than goroutines, so Go can run **thousands of goroutines** easily.

---

### Goroutines

A **goroutine** is a lightweight unit of execution managed by the Go runtime.

#### Key Properties

* Much lighter than OS threads
* Created with the `go` keyword
* Managed by Go, not the OS
* Can scale to **thousands or millions**

#### Example

```go
func sayHello() {
	fmt.Println("Hello")
}

func main() {
	go sayHello()
	fmt.Println("World")
}
```

Here:

* `sayHello()` runs concurrently
* `main()` continues immediately

⚠️ The program may exit before the goroutine runs.

---

### The Synchronization Problem

```go
func main() {
	go fmt.Println("Hello")
}
```

Why this is broken:

* `main()` exits immediately
* Goroutine may never execute

We need a way to **wait** for goroutines.

---

### WaitGroup (Synchronization)

* Ensures the main function waits for all goroutines to finish.

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // signals when goroutine is finished
    fmt.Println("Worker", id, "starting")
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait() // wait for all workers to finish
}
```

#### How It Works

* `Add(n)` → number of goroutines
* `Done()` → signals completion
* `Wait()` → blocks until all are done

Use `WaitGroup` when:

* You only need to wait
* No data needs to be passed back

---

### Channels – Communication Between Goroutines

* **Channels** are Go’s way of communicating safely between goroutines.
* Think of a channel as a **pipe**: one goroutine sends data in, another receives it.

> Go philosophy:
>
> **“Do not communicate by sharing memory; share memory by communicating.”**

#### Create a Channel

```go
ch := make(chan string)
```

#### Send and Receive

```go
go func() {
	ch <- "Hello"
}()

msg := <-ch
fmt.Println(msg)
```

Key behavior:

* Sending blocks until received
* Receiving blocks until sent
* Prevents data races

#### Example

```go
package main

import "fmt"

func worker(ch chan string) {
    ch <- "Work done!" // send message to channel
}

func main() {
    messageChannel := make(chan string) // create channel

    go worker(messageChannel) // start goroutine

    msg := <-messageChannel // receive message from channel
    fmt.Println(msg)
}
```

**Key points:**

* `chan string` is a channel that carries strings.
* `<-` operator is used to send (`ch <- value`) or receive (`value := <-ch`).

---

### Buffered vs Unbuffered Channels

#### Unbuffered Channel

* Communication blocks until both sender and receiver are ready.

```go
ch := make(chan int) // unbuffered
```

* Sender waits for receiver
* Strong synchronization

#### Buffered Channel

* Allows sending `n` values without waiting for receiver.

```go
ch := make(chan int, 2) // buffer of 2

ch <- 1
ch <- 2 // won't block because buffer is not full yet
```

* Holds up to 3 values
* Sender blocks only when buffer is full
* Useful for worker pools and queues

---

### Select Statement

* `select` allows goroutines to wait on multiple channels simultaneously.

```go
package main

import "fmt"

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() { ch1 <- "Message from ch1" }()
    go func() { ch2 <- "Message from ch2" }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
```

**Explanation:**

* Only one of the channels will proceed randomly if both are ready.
* Useful for timeouts and handling multiple sources of data.

Why `select` matters:

* Handles multiple concurrent inputs
* Avoids blocking
* Core for advanced concurrency patterns

---

### Mutex – Safe Access to Shared Memory

In Go, **goroutines can share memory**, but **concurrent access to shared data** can cause **race conditions**.  
A **mutex (mutual exclusion lock)** ensures that **only one goroutine** can access a critical section of code at a time.

Shared memory means multiple goroutines **read or write the same variable stored at a single memory location**.

```go
package main

import (
    "fmt"
    "sync"
)

var counter = 0
var mutex = &sync.Mutex{}

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    mutex.Lock()   // acquire lock
    counter++
    mutex.Unlock() // release lock
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}
```

- **What is “shared memory” here?**

```go
var counter = 0
```

👉 **All goroutines access the SAME `counter` variable**. This is called **shared memory**.

- **What REALLY happens without a mutex?**

You might think:

> “counter++ just adds 1, what’s the problem?”

But `counter++` is **NOT atomic**.

It actually means:

```
1. Read counter from memory
2. Add 1
3. Write result back to memory
```

Now imagine **two goroutines** running at the same time.

Timeline (WITHOUT mutex)

```
counter = 0

Goroutine A: read counter → 0
Goroutine B: read counter → 0

Goroutine A: add 1 → 1
Goroutine B: add 1 → 1

Goroutine A: write 1
Goroutine B: write 1
```

**Final value = 1 instead of 2**

This is called a **race condition**.

> The result depends on timing, which is unpredictable.

#### Real-World Analogy: Bank Account

Imagine a shared bank account:

* Balance = ₹1000
* Two ATMs withdraw ₹100 at the same time

**Without a lock:**

* ATM A reads balance → 1000
* ATM B reads balance → 1000
* Both write back → 900

💥 Lost ₹100

**With a lock (mutex):**

* ATM A locks the account
* ATM B waits
* ATM A updates balance → 900
* ATM A unlocks
* ATM B updates balance → 800

✅ Correct result

A mutex acts like a **lock on the bank account**.

---

#### ❌ Broken Version (No Mutex)

```go
package main

import (
	"fmt"
	"sync"
)

var balance = 1000

func withdraw(wg *sync.WaitGroup) {
	defer wg.Done()
	balance = balance - 100
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go withdraw(&wg)
	}

	wg.Wait()
	fmt.Println("Final Balance:", balance)
}
```

Expected:

```
1000 - (10 × 100) = 0
```

Actual (varies):

```
Final Balance: 300
Final Balance: 500
Final Balance: 100
```

⚠️ **Unpredictable = race condition**

#### ✅ Correct Version (With Mutex)

```go
package main

import (
	"fmt"
	"sync"
)

var balance = 1000
var mu sync.Mutex

func withdraw(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	balance = balance - 100
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go withdraw(&wg)
	}

	wg.Wait()
	fmt.Println("Final Balance:", balance)
}
```

Output:

```
Final Balance: 0
```

- Correct
- Deterministic
- Safe shared memory access

---

#### Why Mutex Works (Mental Model)

Think of a mutex as:

* 🚪 **One key**
* 🧍 **Only one goroutine can enter**
* ⏳ Others must wait
* 🔓 When unlocked, next goroutine enters

In Go:

```go
mu.Lock()   // enter critical section
// shared memory access
mu.Unlock() // exit
```

The protected code is called a **critical section**.

#### How Professionals Decide to Use Mutex

Use a mutex when:

* Multiple goroutines
* Writing OR reading shared data
* Data must stay consistent

Avoid mutex when:

* Data is immutable
* Each goroutine has its own data
* Channels can model the problem better

---

#### Common Problems with Mutexes

* Easy to forget `Unlock()`
* Can cause **deadlocks** if locks are mismanaged
* Harder to reason about in large codebases
* Can reduce performance due to blocking

Always prefer:

```go
defer mutex.Unlock()
```

right after `Lock()` to avoid mistakes.

---

#### Best Practices

* Use mutexes only to protect **small critical sections**
* Keep locked code minimal
* Avoid nested locks
* Use `sync.RWMutex` when reads are frequent and writes are rare

---

#### When to Avoid Mutexes

* When data can be **owned by a single goroutine**
* When communication patterns fit better
* When order of operations matters more than shared state

➡️ **Prefer channels when possible**

Go philosophy:

> “Don’t communicate by sharing memory; share memory by communicating.”

---

#### Debugging Tip

Use the race detector to catch race conditions:

```bash
go run -race main.go
```

This tool helps identify unsafe concurrent memory access during development.

### Data Races

A **data race** occurs when:

* Two goroutines access the same variable
* At least one writes
* No synchronization exists

#### Detect Races

```bash
go test -race
```

Go’s race detector is one of the best in any language.

---

### Common Patterns in Go Concurrency

1. **Worker Pool**

   * Manage a pool of goroutines to process tasks.
   * Fixed number of goroutines
   * Jobs sent through a channel
   * Results collected via another channel

2. **Fan-Out / Fan-In**

   * Multiple goroutines send data into a channel; main goroutine aggregates results.
   * Fan-out: distribute work to many goroutines
   * Fan-in: collect results into one channel

3. **Pipeline**

   * Data passes through multiple stages, each stage a goroutine.
   * Each stage runs in its own goroutine
   * Output of one stage feeds the next

---

### When to Use Concurrency

Use concurrency when:

* Tasks are independent
* Program waits on I/O
* Handling many users or requests

Avoid concurrency when:

* Code becomes harder to understand
* Performance gain is minimal

> **Concurrency is a tool, not a requirement.**

---

### Summary Table of Concepts

| Concept        | Definition                       | Go Feature/Example          |
| -------------- | -------------------------------- | --------------------------- |
| Process        | Program in execution             | `go run main.go`            |
| Thread         | Unit inside a process            | Managed by OS               |
| Goroutine      | Lightweight thread               | `go f()`                    |
| Concurrency    | Tasks making progress together   | Multiple goroutines         |
| Parallelism    | Tasks running simultaneously     | Multiple CPU cores          |
| Multithreading | Using multiple threads           | Managed automatically by Go |
| Channel        | Communication between goroutines | `chan`                      |
| Mutex          | Lock to protect shared memory    | `sync.Mutex`                |
| WaitGroup      | Wait for multiple goroutines     | `sync.WaitGroup`            |

---

### Tips for Learning Go Concurrency

1. Think in **goroutines and channels**, not threads.
2. Avoid sharing memory when possible; prefer **channel communication**.
3. Understand difference between **concurrency** and **parallelism**.
4. Practice small programs: worker pools, pipelines, fan-in/fan-out.
5. Always check for **race conditions** using `go run -race`.

### Example (Concurrent URL Crawler)

<details>
<summary>View contents</summary>

### What this crawler will do

* Take a list of URLs
* Fetch them concurrently
* Limit concurrency (so we don’t overload)
* Collect results safely
* Stop cleanly

---

## Stage 0: Sequential version (baseline)

```go
package main

import (
	"fmt"
	"net/http"
)

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(url, "->", resp.Status)
	return nil
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
	}

	for _, url := range urls {
		if err := fetch(url); err != nil {
			fmt.Println("error:", err)
		}
	}
}
```

### What you learn

* One URL at a time
* Slow, but predictable

---

## Stage 1: Goroutines (naive concurrency)

Now let’s make it concurrent **the wrong way** (on purpose).

```go
for _, url := range urls {
	go fetch(url)
}
```

Run it — **nothing prints or prints randomly**.

### Why?

* `main()` exits before goroutines finish
* Goroutines are cheap, but not magical

---

## Stage 2: WaitGroup (basic coordination)

We need to **wait** for all goroutines.

```go
import "sync"

var wg sync.WaitGroup

for _, url := range urls {
	wg.Add(1)
	go func(u string) {
		defer wg.Done()
		fetch(u)
	}(url)
}

wg.Wait()
```

### Key concepts

* `WaitGroup` = “don’t exit yet”
* Always pass loop variables explicitly (`u string`)

✅ Now all URLs fetch concurrently
❌ But still **unlimited concurrency**

---

## Stage 3: Worker Pool (real-world concurrency control)

This is where Go starts to shine.

### Idea

* Fixed number of workers
* URLs go into a channel
* Workers pull from the channel

### Architecture

```
URLs → jobs channel → workers → results
```

---

### Step 1: Jobs channel

```go
jobs := make(chan string)
```

---

### Step 2: Worker function

```go
func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d fetching %s\n", id, url)
		fetch(url)
	}
}
```

---

### Step 3: Start workers

```go
numWorkers := 3
var wg sync.WaitGroup

for i := 1; i <= numWorkers; i++ {
	wg.Add(1)
	go worker(i, jobs, &wg)
}
```

---

### Step 4: Send jobs

```go
for _, url := range urls {
	jobs <- url
}
close(jobs)

wg.Wait()
```

### What you just learned

* **Channels coordinate work**
* **Workers limit concurrency**
* This is how production crawlers work

---

## Stage 4: Results channel (no shared state)

Instead of printing inside workers, return data safely.

```go
type Result struct {
	URL    string
	Status string
	Err    error
}
```

### Worker sends results

```go
func worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d fetching %s\n", id, url)

		client := http.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Get(url)
		if err != nil {
			results <- Result{URL: url, Err: err}
			continue
		}
		resp.Body.Close()

		results <- Result{
			URL:    url,
			Status: resp.Status,
		}
	}
}
```

---

### Collect results in main

```go
results := make(chan Result)

go func() {
	wg.Wait()
	close(results)
}()

for r := range results {
	if r.Err != nil {
		fmt.Println("error:", r.URL, r.Err)
	} else {
		fmt.Println(r.URL, "->", r.Status)
	}
}
```

### Why this matters

* No mutexes
* No shared memory
* **Channels = ownership transfer**

---

## Final Project

📁 `main.go`

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result represents the outcome of fetching a URL
type Result struct {
	URL    string
	Status string
	Err    error
}

// worker pulls URLs from jobs channel and sends results
func worker(
	id int,
	jobs <-chan string,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d fetching %s\n", id, url)

		client := http.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Get(url)
		if err != nil {
			results <- Result{URL: url, Err: err}
			continue
		}

		resp.Body.Close()

		results <- Result{
			URL:    url,
			Status: resp.Status,
		}
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
		"https://httpbin.org/status/404",
		"https://invalid-url",
	}

	numWorkers := 3

	jobs := make(chan string)
	results := make(chan Result)

	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Close results when workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		if result.Err != nil {
			fmt.Printf("❌ %s error: %v\n", result.URL, result.Err)
			continue
		}
		fmt.Printf("✅ %s -> %s\n", result.URL, result.Status)
	}

	fmt.Println("Crawling finished")
}
```

---

## Why add `context.Context`?

Right now:

* If one URL hangs → it waits until timeout
* If main decides to stop → workers **keep running**
* No way to cancel everything at once

`context.Context` gives you:

* Global cancellation signal
* Timeouts / deadlines
* Clean shutdown of *all* goroutines

Think of it as a **broadcast “STOP” button**.

---

## Step 1: Create a cancellable context in `main`

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

This creates:

* `ctx` → passed to all workers
* `cancel()` → stops everything immediately

---

## Step 2: Use `http.NewRequestWithContext`

Instead of:

```go
client.Get(url)
```

Use:

```go
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := client.Do(req)
```

Now:

* If `ctx` is cancelled
* HTTP request is aborted instantly

---

## Step 3: Workers must *listen* for cancellation

Use `select`:

```go
select {
case <-ctx.Done():
	return
case url := <-jobs:
	// work
}
```

---

## Full version: Worker pool with context cancellation

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Worker int
	URL    string
	Status string
	Err    error
}

func fetch(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan string,
	results chan<- Result,
	client *http.Client,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return

		case url, ok := <-jobs:
			if !ok {
				return
			}

			resp, err := fetch(ctx, client, url)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}

			results <- Result{
				Worker: id,
				URL:    url,
				Status: resp.Status,
			}
		}
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
		"https://httpbin.org/status/404",
		"https://invalid-url",
	}

	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Entire crawl → max 30s
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second} // Each request → max 5s

	jobs := make(chan string)
	results := make(chan Result)

	numWorkers := 3
	var wg sync.WaitGroup

	// Start worker pool
	for i := range numWorkers {
		wg.Add(1)
		go worker(ctx, i+1, jobs, results, client, &wg)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Close results when workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consume results
	for r := range results {
		if r.Err != nil {
			fmt.Printf("❌ Worker %d %s error: %v\n", r.Worker, r.URL, r.Err)
			// example: cancel on first fatal error
			// cancel()
			continue
		}
		fmt.Printf("✅ Worker %d %s -> %s\n", r.Worker, r.URL, r.Status)
	}

	// example: cancel everything after 2 seconds
	// time.AfterFunc(1*time.Second, cancel)

	fmt.Println("Crawling finished")
}
```

---

## Mental model

* `context.Context` = **shared cancellation signal**
* `cancel()` = press the red button
* `select` = workers constantly check if they should stop
* HTTP requests respect context automatically

---

## Common mistakes (avoid these)

❌ Creating a new context inside each worker
❌ Forgetting to pass context to HTTP requests
❌ Not checking `ctx.Done()` in long loops

---

## When to use what

| Tool                  | Purpose                            |
| --------------------- | ---------------------------------- |
| `Client.Timeout`      | Max duration of a request          |
| `context.WithTimeout` | Cancel *everything* after deadline |
| `context.WithCancel`  | Manual cancellation                |

They **work together**, not instead of each other.

---

</details>

### Add rate limiting with time.Ticker

<details>
<summary>View contents</summary>

#### Goal

* Limit the number of HTTP requests per second (or per time interval)
* Keep workers concurrent but **polite** to the server
* Combine with **context cancellation** and **results channel**

---

#### Concept: `time.Ticker` as a rate limiter

```go
ticker := time.NewTicker(time.Second / 2) // 2 requests per second
defer ticker.Stop()
```

* Each tick = permission to send **one request**
* Workers **wait for a tick** before fetching

---

#### How to integrate with worker

Inside the worker:

```go
select {
case <-ctx.Done():
    return
case <-ticker.C:
    // allowed to fetch
}
```

This ensures **no more than X requests per second** globally.

---

#### Full example: Worker pool with context, results, and rate limiting

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Worker int
	URL    string
	Status string
	Err    error
}

func worker(ctx context.Context, id int, jobs <-chan string, results chan<- Result, client *http.Client, wg *sync.WaitGroup, ticker <-chan time.Time) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled\n", id)
			return

		case url, ok := <-jobs:
			if !ok {
				return
			}

			// Rate limiting: wait for a tick
			select {
			case <-ctx.Done():
				return
			case <-ticker:
				// allowed to fetch
			}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}

			resp, err := client.Do(req)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}
			resp.Body.Close()

			results <- Result{Worker: id, URL: url, Status: resp.Status}
		}
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second}

	numWorkers := 3
	jobs := make(chan string)
	results := make(chan Result)
	var wg sync.WaitGroup

	// Rate limiter: 1 request per second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i+1, jobs, results, client, &wg, ticker.C)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Close results when done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for r := range results {
		if r.Err != nil {
			fmt.Printf("❌ Worker %d %s error: %v\n", r.Worker, r.URL, r.Err)
			continue
		}
		fmt.Printf("✅ Worker %d %s -> %s\n", r.Worker, r.URL, r.Status)
	}

	fmt.Println("All done")
}
```

---

#### How it works

1. `ticker := time.NewTicker(1 * time.Second)`

   * produces **1 tick per second**

2. Each worker waits on `<-ticker` before sending a request

   * ensures **global request rate**
   * multiple workers share the same ticker → smooth traffic

3. `ctx.Done()` ensures clean cancellation

4. Results are sent via `results` channel — safe & observable

---

#### Mental model

```
jobs --> worker (waits for tick) --> fetch --> results
        ↑ ctx controls cancellation
```

* `ticker` = “permission slip” to fetch
* `context` = “STOP button”
* `results` = “what happened”

---

</details>

### Retry with exponential backoff

<details>
<summary>View contents</summary>

#### What is exponential backoff?

Exponential backoff is a retry strategy where:

* The wait time **doubles** after each failure
* Optional **jitter/randomness** to avoid thundering herd problems
* Stops after **max retries**

Example sequence:

```
retry 1 → wait 1s  
retry 2 → wait 2s  
retry 3 → wait 4s  
retry 4 → fail
```

---

#### Implementing a `fetchWithRetry` function

```go
func fetchWithRetry(ctx context.Context, client *http.Client, url string, maxRetries int) (*http.Response, error) {
	var resp *http.Response
	var err error
	backoff := 1 * time.Second

	for i := 0; i <= maxRetries; i++ {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err = client.Do(req)
		if err == nil {
			return resp, nil
		}

		// Retry only if context is not done
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(backoff):
			// increase backoff exponentially
			backoff *= 2
		}
	}

	return nil, err
}
```

✅ Notes:

* `ctx` ensures **cancellation propagates**
* `time.After(backoff)` sleeps between retries
* `backoff *= 2` doubles wait each retry
* `maxRetries` limits the total attempts

---

#### Integrate with worker

```go
func worker(
	ctx context.Context,
	id int,
	jobs <-chan string,
	results chan<- Result,
	client *http.Client,
	wg *sync.WaitGroup,
	ticker <-chan time.Time,
	maxRetries int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-jobs:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker: // rate limiting
			}

			resp, err := fetchWithRetry(ctx, client, url, maxRetries)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}
			resp.Body.Close()

			results <- Result{Worker: id, URL: url, Status: resp.Status}
		}
	}
}
```

---

#### Update `main` to pass `maxRetries`

```go
maxRetries := 3

for i := 0; i < numWorkers; i++ {
	wg.Add(1)
	go worker(ctx, i+1, jobs, results, client, &wg, ticker.C, maxRetries)
}
```

---

#### How it works together

```
jobs -> worker (wait for tick) -> fetchWithRetry -> results
        ↑ ctx cancels      ↑ maxRetries + backoff
```

* Workers respect **rate limiting** (`ticker`)
* Workers respect **cancellation** (`ctx`)
* Temporary failures are retried **without overwhelming the server**
* Results channel collects successes/errors

---

#### Optional improvement: add jitter

Exponential backoff + jitter prevents **all workers from retrying at the exact same time**:

```go
backoff := time.Duration(rand.Int63n(int64(base))) + base
```

* `rand.Int63n(base)` → random 0–base
* `base` doubles after each retry

---

</details>

</details>

## Web Servers

Go has a **powerful, production-ready HTTP server built into the standard library**.
You do **not** need external frameworks to build fast and scalable web servers.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/webServers)**

```go
package main

import (
 "fmt"
 "log"
 "net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
 fmt.Println("Home!")
 fmt.Fprint(w, "Home!")
}

func main() {
 http.HandleFunc("/", home)

 fmt.Println("Server is running on port :8080")
 log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Explanation:

* `http.ResponseWriter` → used to write the HTTP response
* `*http.Request` → contains request data (method, headers, body, URL)
* `fmt.Fprint` sends data back to the client
* `http.HandleFunc("/", home)`
  Registers the `/` route and associates it with the `home` handler
* `http.ListenAndServe(":8080", nil)`
  Starts an HTTP server on port `8080`
* `nil` → uses the default multiplexer (`http.DefaultServeMux`)
* `log.Fatal` → logs error and exits if server fails

### Key Concepts

#### 1. Routing

Go uses a **request multiplexer** (router) internally.

```go
http.HandleFunc("/about", aboutHandler)
```

#### 2. Concurrency

Each HTTP request is handled in its **own goroutine** automatically.

> You do not need to manage threads manually.

#### 3. Production Ready

* Fast
* Concurrent
* Secure by default
* Used internally by many Go services

---

### When to Use Frameworks

Use only when you need:

* Middleware chaining
* Advanced routing
* Request validation helpers

Otherwise, `net/http` is often enough.

</details>

## Working with files

Go provides simple and safe file handling via the `os` and `io` packages.

<details>
<summary>View contents</summary>

```go
import (
 "fmt"
 "os"
 "strings"
)

type names string[]

func main() {
 names := names{"Sohel", "Mithu", "Rupom"}
 names.saveToFile("my_names")
 fmt.Println(readNamesFromFiles("my_names"))
 removeFile("my_names")
}

func (n names) toString() string {
 return strings.Join(n, ",")
}

func (n names) saveToFile(filename string) error {
 return os.WriteFile(filename, []bytes(n.toString()), 0666)
}

func readNamesFromFile(filename string) names {
 bs, err := os.ReadFile(filename)
 if err != nil {
  fmt.Println("Error:", err)
  os.Exit(1)
 }

 return strings.Split(string(bs), ",")
}

func removeFile(filename string) {
 err := os.Remove(filename)
 if err != nil {
  fmt.Println("Error:", err)
 }
}
```

### Important Notes

* Always handle file errors
* File operations are blocking
* Avoid `os.Exit` in libraries (fine for small examples)
* Use streaming (`os.Open`) for large files

</details>
 
## Generate random numbers
 
<details>
<summary>View contents</summary>
 
```go
import (
 "math/rand"
 "time"
)
 
source := rand.NewSource(time.Now().UnixNano())
r := rand.New(source)
 
// genrate random number from 0 to n
r.Intn(8) // n = 8
```
 
</details>

## Effective Go in short <sup>[ref](https://go.dev/doc/effective_go)</sup>

<details>

<summary>View contents</summary>

### Formatting with `gofmt`

<details>

<summary>View contents</summary>

Go uses `gofmt` to automatically format code, ensuring consistency across projects. Instead of manually aligning comments or indentation, developers rely on `gofmt`.

#### Key Formatting Rules:

1. **Indentation**: Uses **tabs**, not spaces.
2. **Line Length**: No strict limit; wrap long lines if needed.
3. **Parentheses**: Avoid unnecessary parentheses in control structures (`if`, `for`, `switch`).

#### Example:

**Before `gofmt`:**
```go
type T struct {
    name string // name of the object
    value int // its value
}
```

**After `gofmt`:**
```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

To format code, run:
```sh
gofmt -w filename.go
```
Or use:
```sh
go fmt ./...
```
This keeps your code clean and readable without manual effort.

</details>

### Commentary

<details>

<summary>View contents</summary>

Go supports both **C-style block comments (`/* */`)** and **C++-style line comments (`//`)**. 

- **Line comments (`//`)** are the standard and widely used.
- **Block comments (`/* */`)** are mainly for package documentation or temporarily disabling code.

#### Example:

```go
package main

import "fmt"

// Greet prints a welcome message.
func Greet(name string) {
    fmt.Println("Hello,", name)
}

/* 
   This is a block comment.
   It is useful for package-level documentation or disabling large sections of code.
*/

func main() {
    Greet("Foyez") // Calling the Greet function
}
```

**Doc comments** (before functions, structs, or packages) serve as the primary documentation.

For more details, check: [`go doc`](https://pkg.go.dev/cmd/doc).

</details>

### Names

<details>

<summary>View contents</summary>

Go follows clear and consistent naming conventions for better readability and usability.

#### 1️⃣ **Package Names**
- Use **short, lowercase, single-word** names.
- The package name should match its directory name.

✅ **Example:**  
```go
import "bytes" // Use bytes.Buffer, not bytes_package.Buffer
```
```go
import "encoding/base64" // Imported as base64, not encodingBase64
```

#### 2️⃣ **Avoid Redundant Names**
- Names should be **concise and meaningful**.
- Use the **package name** as context.

✅ **Example:**  
```go
buf := bufio.NewReader(input) // Not bufio.NewBufReader
r := ring.New()               // Not ring.NewRing
```

#### 3️⃣ **Getters & Setters**
- Avoid `Get` in getter names.
- Use **exported (uppercase) methods** for getters, and `SetX` for setters.

✅ **Example:**  
```go
type User struct {
    owner string
}

func (u *User) Owner() string { return u.owner }   // Not GetOwner()
func (u *User) SetOwner(o string) { u.owner = o }  // Setter
```
```go
user := User{}
user.SetOwner("Foyez")
fmt.Println(user.Owner()) // Reads naturally
```

#### 4️⃣ **Interface Naming**
- **Single-method interfaces** should use `-er` suffix.

✅ **Example:**  
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

#### 5️⃣ **Use MixedCaps Instead of Underscores**
✅ **Example:**  
```go
type DataProcessor struct {}  // Not data_processor
func ProcessData() {}         // Not process_data()
```

Following these conventions makes Go code **clean, idiomatic, and easy to read**.

</details>

### Semicolons

<details>

<summary>View contents</summary>

Go automatically inserts semicolons (`;`) where needed, so they are **mostly invisible** in source code. The rule:  
👉 If a line ends with an **identifier, literal, or certain tokens (`return`, `break`, `}` etc.)**, Go **inserts a semicolon**.

#### ✅ **Examples:**
```go
x := 10  // Semicolon inserted automatically
fmt.Println(x)  // Semicolon inserted automatically
```

#### **Semicolons in Loops**
Semicolons are required **only in `for` loops**:
```go
for i := 0; i < 5; i++ {  // Semicolons required
    fmt.Println(i)
}
```

#### **Incorrect Brace Placement**
🚨 **Wrong:**
```go
if x > 0  // Semicolon inserted → Unexpected behavior!
{ 
    fmt.Println("Positive") 
}
```
✅ **Correct:**
```go
if x > 0 {  // No semicolon inserted
    fmt.Println("Positive")
}
```

**Conclusion:** Let Go handle semicolons **automatically**, except in `for` loops or multiple statements on one line.

</details>

### Control Structures

<details>

<summary>View contents</summary>

Go's control structures are **similar to C** but with key differences:  
✅ No `do` or `while`, just `for` loops  
✅ `switch` is more flexible  
✅ No parentheses `()` in conditions  
✅ Braces `{}` are **mandatory**  

---

#### **1. `if` Statement**  
Braces `{}` are **always required**.  
Optional **initialization** before condition.

✅ **Example:**  
```go
if x := getValue(); x > 0 {  
    fmt.Println("Positive:", x)
}
```
✅ **No need for `else` if returning early:**  
```go
f, err := os.Open("file.txt")  
if err != nil {
    return err
}
useFile(f)
```

---

#### **2. `for` Loop**  
Go has **only `for`**, which works like `for`, `while`, and `forever` loops.

✅ **Standard `for` loop:**  
```go
for i := 0; i < 5; i++ {  
    fmt.Println(i)
}
```
✅ **While-like loop:**  
```go
for x > 0 {  
    x--
}
```
✅ **Infinite loop:**  
```go
for {  
    fmt.Println("Running forever")
}
```

✅ **Loop over collections (array, slice, string, or map, or reading from a channel) using `range`:**  
```go
for i, val := range arr {  
    fmt.Println(i, val)
}

for key, value := range oldMap {
    newMap[key] = value
}
```
If you only need the first item in the range (the key or index), drop the second:
```go
for key := range myMap {
    if key.expired() {
        delete(myMap, key)
    }
}
```
If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:
```go
for _, val := range arr { fmt.Println(val) }
```

✅ Go's `range` loop **automatically decodes UTF-8** and provides **rune values** (Unicode code points) instead of bytes.

**Example:**  
```go
for pos, char := range "日本\x80語" { // \x80 is an invalid UTF-8 byte
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```
🛠 **Output:**  
```
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6  // Invalid UTF-8 replaced
character U+8A9E '語' starts at byte position 7
```

- `range` **extracts Unicode runes**, not raw bytes.
- Invalid UTF-8 sequences are replaced with `U+FFFD (�)`.
- **Positions refer to bytes, not runes**.

✅  **Go Has No Comma Operator (`++` & `--` are Statements)** \
In Go:
- ✅ `++` and `--` are **statements, not expressions**  
- 🚫 You **cannot** use them inside expressions like `x = y++ + z`  

**Reverse an Array Using Parallel Assignment:**  
```go
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]  // Swap elements
}
```

---

#### **3. `switch` Statement**  
✅ **No need for `break` (no fallthrough by default):**  
```go
switch x {
case 1:
    fmt.Println("One")
case 2, 3:
    fmt.Println("Two or Three")
default:
    fmt.Println("Other")
}
```
✅ **Can replace `if-else` chains:**  
```go
x := 5
switch {
case x >= 1 && x <= 4:
  fmt.Println("Between 1 to 4")
case x > 4:
  fmt.Println("Greater than 4")
default:
  fmt.Println("Zero")
}
```

✅ **Labeled `break` for loops:**  
```go
Loop:
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            if i+j > 5 {
                fmt.Println("Print only one time")
                break Loop // Exits outer loop
            }
        }
    }
```

---

#### **4. `type switch` (Detect Dynamic Type)**  
```go
var i any = "hello"

switch v := i.(type) {
case int:
    fmt.Println("Integer:", v)
case string:
    fmt.Println("String:", v)
default:
    fmt.Println("Unknown type")
}
```

---

</details>

### **Functions**

<details>

<summary>View contents</summary>

#### **1️⃣ Multiple Return Values**  
Go allows functions to return multiple values, avoiding in-band error returns (like `-1` in C).

✅ **Example:** `Write` method in `os` package  
```go
func (file *File) Write(b []byte) (n int, err error)
```
Returns:
- `n` → number of bytes written  
- `err` → error (if not all bytes were written)

✅ **Example:** Extracting an integer from a byte slice  
```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {}
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```
Usage:
```go
for i := 0; i < len(b); {
    x, i = nextInt(b, i)
    fmt.Println(x)
}
```

---

#### **2️⃣ Named Return Parameters**  
Return variables can be **named**, making code more readable.

✅ **Example:** Naming `value` and `nextPos`
```go
func nextInt(b []byte, pos int) (value, nextPos int) {
```
✅ **Example:** Simplifying `io.ReadFull`
```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return  // Uses named return variables
}
```

---

#### **3️⃣ `defer` Statement (Delayed Execution)**  
`defer` schedules a function to run **before the surrounding function exits**.

✅ **Example:** Ensuring a file closes  
```go
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // Ensures file closes before return

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf)
        result = append(result, buf[:n]...)
        if err == io.EOF {
            break
        } else if err != nil {
            return "", err
        }
    }
    return string(result), nil
}
```
**Why use `defer`?**  
✔ Guarantees cleanup (prevents leaks)  
✔ Keeps code cleaner (close near open)  

✅ **Example:** `defer` executes in LIFO order  
```go
for i := 0; i < 5; i++ {
    defer fmt.Print(i, " ")
}
// Output: 4 3 2 1 0
```

✅ **Example:** Function Tracing  
```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```
**Output:**  
```
entering: b  
in b  
entering: a  
in a  
leaving: a  
leaving: b  
```
**Why use `defer`?**  
✔ Ensures cleanup, even if function exits early  
✔ Improves readability and maintainability 🚀

</details>

### **Data**

<details>
<summary>View contents</summary>

**Memory Allocation: `new` vs. `make`**  

<details>
<summary>View contents</summary> 

Go provides two built-in functions for memory allocation:  
- **`new(T)`** → Allocates zeroed storage for a value of type `T` and returns `*T` (a pointer).  
- **`make(T, args)`** → Initializes slices, maps, and channels, returning `T` (not a pointer).  

---

### **Allocation with `new`**  
✅ **Usage:**  
- Allocates memory but does not initialize beyond zero values.  
- Returns a **pointer** to the allocated type.  

✅ **Example:** Allocating a struct  
```go
type SyncedBuffer struct {
    lock   sync.Mutex
    buffer bytes.Buffer
}

p := new(SyncedBuffer)  // Returns *SyncedBuffer
var v SyncedBuffer      // Direct allocation, ready to use
```
✔ `p` is a pointer (`*SyncedBuffer`)  
✔ `v` is a direct struct instance  

✅ **Example:** Allocating an integer  
```go
p := new(int) // *int with zero value (0)
fmt.Println(*p) // Prints 0
```
---

### **Constructors & Composite Literals**  
✅ **When zero values aren’t enough, use a constructor.**  

🔴 **Verbose version**  
```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    return f
}
```
✔ Can be **simplified** with **composite literals**  

🟢 **Optimized version**  
```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    return &File{fd, name, nil, 0} // Allocates and initializes
}
```

The fields of a composite literal are laid out in order and must all be present. However, by labeling the elements explicitly as field:value pairs, the initializers can appear in any order, with the missing ones left as their respective zero values. Thus we could say
```go
return &File{fd: fd, name: name}
```

✔ `&File{}` creates and returns an initialized struct  
✔ Equivalent to `new(File)` but with initialization  

✅ **Examples of composite literals:**  
```go
a := [...]string{0: "no error", 1: "Eio", 2: "invalid argument"} // Array
s := []string{"no error", "Eio", "invalid argument"}             // Slice
m := map[int]string{0: "no error", 1: "Eio", 2: "invalid argument"} // Map
```

---

### **Allocation with `make`**  
✅ **`make` initializes slices, maps, and channels**  
```go
v := make([]int, 10, 100) // Length 10, capacity 100
```
✔ Allocates an array of 100 integers  
✔ Creates a slice of length 10 referring to that array  

✅ **Key difference:**  
```go
var p *[]int = new([]int) // *p == nil (rarely useful)
var v  []int = make([]int, 100) // Allocates a slice of 100 ints
```
✔ `new([]int)` → Returns a pointer to an **empty** slice  
✔ `make([]int, 100)` → Returns a **usable** slice of 100 ints  

✅ **Idiomatic way:**  
```go
v := make([]int, 100) // Best practice
```
✔ Avoid unnecessary complexity  

✅ **Other `make` examples:**  
```go
ch := make(chan int, 10)    // Buffered channel
m := make(map[string]int)   // Empty map
```
✔ `make` initializes internal structures

</details>

### **Arrys**  

<details>
<summary>View contents</summary> 

Arrays in Go differ from C-style arrays:  
✔ **Arrays are values** → Assigning an array copies all elements.  
✔ **Passing to a function copies the array** (unless using a pointer).  
✔ **Array size is part of its type** → `[10]int` and `[20]int` are different types.  

---

#### **1️⃣ Array Declaration & Initialization**  
```go
var a [3]int            // Zero-initialized: [0, 0, 0]
b := [3]int{1, 2, 3}    // Explicit values
c := [...]int{4, 5, 6}  // Compiler determines size
```

---

#### **2️⃣ Copying Arrays**  
```go
a := [3]int{1, 2, 3}
b := a   // Creates a copy, modifying b won’t affect a
b[0] = 10
fmt.Println(a, b) // Output: [1 2 3] [10 2 3]
```

---

#### **3️⃣ Passing Arrays to Functions**  
🔴 **By Value (copying entire array)**  
```go
func ModifyArray(arr [3]int) {
    arr[0] = 100
}
a := [3]int{1, 2, 3}
ModifyArray(a)
fmt.Println(a)  // Output: [1 2 3] (unchanged)
```
  
🟢 **By Reference (using pointers for efficiency)**  
```go
func ModifyArray(arr *[3]int) {
    arr[0] = 100
}
a := [3]int{1, 2, 3}
ModifyArray(&a)
fmt.Println(a)  // Output: [100 2 3]
```

---

#### **4️⃣ Arrays vs. Slices**  
✅ **Use slices for flexibility & efficiency**  
```go
array := [3]float64{7.0, 8.5, 9.1}
slice := array[:]  // Convert to slice
```
✔ Slices are more idiomatic in Go  

---

</details>

### **Slices**  

<details>
<summary>View contents</summary> 

Slices provide a **dynamic, flexible view** over arrays and are the preferred way to handle collections in Go.  

✔ **Slices reference an underlying array** (modifications affect all references).  
✔ **Passing a slice to a function allows modifications** without explicit pointers.  
✔ **Slice length (`len`) and capacity (`cap`) can change dynamically.**  

---

#### **1️⃣ Declaring & Initializing Slices**  
```go
var s []int              // nil slice (len=0, cap=0)
s = []int{1, 2, 3}       // Initialized slice
t := make([]int, 5, 10)  // Slice with len=5, cap=10
```

---

#### **2️⃣ Slicing an Array**  
```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]  // Slice of [2, 3, 4]
```

🔹 **Slices hold references to arrays, so modifying `s` affects `arr`**  
```go
s[0] = 100
fmt.Println(arr) // Output: [1 100 3 4 5]
```

---

#### **3️⃣ Passing Slices to Functions**  
```go
func modify(s []int) {
    s[0] = 42
}
nums := []int{1, 2, 3}
modify(nums)
fmt.Println(nums) // Output: [42 2 3]
```
✔ No need to pass a pointer, as slices already reference the underlying array.  

---

#### **4️⃣ Appending to Slices**  
```go
s := []int{1, 2}
s = append(s, 3, 4, 5)  // Expands the slice
fmt.Println(s) // Output: [1 2 3 4 5]
```
✔ If capacity is exceeded, Go automatically **allocates new memory**.  

---

#### **5️⃣ Copying Slices**  
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
fmt.Println(dst) // Output: [1 2 3]
```
✔ `copy(dst, src)` copies **minimum of len(dst) and len(src)**.  

---

#### **6️⃣ Two-Dimensional Slices**  
🔹 **Independent inner slices:**  
```go
matrix := make([][]int, 3) // 3 rows
for i := range matrix {
    matrix[i] = make([]int, 4) // 4 columns
}
```
🔹 **Single allocation for efficiency:**  
```go
matrix := make([][]int, 3)
data := make([]int, 3*4) // 3 rows * 4 cols
for i := range matrix {
    matrix[i], data = data[:4], data[4:] // Slice the array into rows
}
```
✔ **More efficient** but less flexible than independent allocation.  

---

</details>

### **Maps** 

<details>
<summary>View contents</summary> 

Maps in Go provide an efficient **key-value** data structure with dynamic sizing.  

✔ **Keys must support equality (`==`)** (e.g., strings, ints, structs).  
✔ **Maps hold references** to an underlying structure (modifying inside a function affects the caller).  
✔ **Accessing a non-existent key returns the zero value** of the value type.  

---

#### **1️⃣ Declaring & Initializing Maps**  
```go
// Using map literals
timeZone := map[string]int{
    "UTC":  0,
    "EST": -5 * 3600,
    "CST": -6 * 3600,
}

// Using make()
users := make(map[string]int) // Empty map
```

---

#### **2️⃣ Accessing & Modifying Maps**  
```go
timeZone["PST"] = -8 * 3600  // Add key-value pair
offset := timeZone["EST"]    // Retrieve value
fmt.Println(offset)          // Output: -18000
```

❌ **Accessing a non-existent key returns the zero value** (for int, it's `0`).  

---

#### **3️⃣ Checking Key Existence (Comma-Ok Idiom)**  
```go
if offset, exists := timeZone["UTC"]; exists {
    fmt.Println("Offset:", offset)
}
```
✔ Helps differentiate between **"zero value"** and **"missing key"**.

To test for presence in the map without worrying about the actual value, you can use the blank identifier (_) in place of the usual variable for the value.

```go
_, exists := timeZone[tz]
```

---

#### **4️⃣ Deleting Keys**  
```go
delete(timeZone, "PST")  // Remove PST from map
```
✔ Safe to call **even if the key doesn’t exist**.  

---

#### **5️⃣ Using Maps as Sets**  
```go
attended := map[string]bool{"Ann": true, "Joe": true}
if attended["Ann"] {
    fmt.Println("Ann was at the meeting")
}
```
✔ **Maps with `bool` values** can function as sets.  

---

#### **6️⃣ Iterating Over a Map**  
```go
for key, value := range timeZone {
    fmt.Println("Key:", key, "Value:", value)
}
```

---

</details>

### **Printing**

<details>
<summary>View contents</summary> 

Go’s `fmt` package provides **formatted printing** similar to C’s `printf`, but with enhanced features.  

✔ `fmt.Print`, `fmt.Println`, `fmt.Printf` – Basic printing.  
✔ `fmt.Sprintf`, `fmt.Fprintf` – Return formatted strings.  
✔ `fmt.Fprint(os.Stdout, …)` – Print to `io.Writer`.  

---

#### **1️⃣ Basic Printing**  
```go
fmt.Print("Hello")      // Prints: Hello
fmt.Println("Hello")    // Prints: Hello\n (adds newline)
fmt.Printf("Age: %d\n", 25) // Prints: Age: 25
```
✔ `Println` adds spaces & newline.  
✔ `Printf` supports format specifiers (`%d`, `%s`, etc.).  

---

#### **2️⃣ Formatting Values (`Printf`)**  
```go
x := 255
fmt.Printf("%d %x %b\n", x, x, x) // Decimal, Hex, Binary
// Output: 255 ff 11111111

t := struct{ Name string; Age int }{"John", 30}
fmt.Printf("%v\n", t)   // {John 30} (default format)
fmt.Printf("%+v\n", t)  // {Name:John Age:30} (field names)
fmt.Printf("%#v\n", t)  // struct { Name string; Age int }{Name:"John", Age:30} (Go syntax)
```
✔ `%v` → default value representation.  
✔ `%+v` → shows struct field names.  
✔ `%#v` → prints Go syntax representation.  

---

#### **3️⃣ Printing Maps & Slices**  
```go
timeZone := map[string]int{"UTC": 0, "PST": -8 * 3600}
fmt.Printf("%v\n", timeZone)   // map[PST:-28800 UTC:0]
fmt.Printf("%#v\n", timeZone)  // map[string]int{"PST":-28800, "UTC":0}
```
✔ **Maps are sorted lexicographically by key**.  

---

#### **4️⃣ Printing Types & Quotes**  
```go
str := "Hello"
fmt.Printf("%T\n", str)   // string (prints type)
fmt.Printf("%q\n", str)   // "Hello" (quoted string)
fmt.Printf("%x\n", str)   // 48656c6c6f (hex encoding)
fmt.Printf("% x\n", str)  // 48 65 6c 6c 6f (spaced hex)
```
✔ `%T` → prints type.  
✔ `%q` → quoted string.  
✔ `%x` → hexadecimal representation.  

---

#### **5️⃣ Custom String Representation (`String()` Method)**  
```go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

p := Person{"Alice", 25}
fmt.Println(p) // Alice is 25 years old
```
✔ Implement `String()` for **custom print formatting**.  

---

#### **6️⃣ Variadic Functions (`...interface{}` & `...int`)**  
```go
func Min(values ...int) int {
    min := values[0]
    for _, v := range values {
        if v < min {
            min = v
        }
    }
    return min
}
fmt.Println(Min(3, 1, 4, 1, 5)) // 1
```
✔ `...int` → allows multiple `int` arguments.  

---

</details>

### **Append**

<details>
<summary>View contents</summary> 

Go’s built-in `append` function dynamically grows slices.  

#### **1️⃣ Syntax**  
```go
func append(slice []T, elements ...T) []T
```
- `T` is a generic type.  
- `append` returns a **new slice** because the underlying array **may change**.  

---

#### **2️⃣ Appending Elements**  
```go
x := []int{1, 2, 3}
x = append(x, 4, 5, 6)
fmt.Println(x) // [1 2 3 4 5 6]
```
✔ Adds multiple elements at once.  

---

#### **3️⃣ Appending a Slice (`...` Operator)**  
```go
x := []int{1, 2, 3}
y := []int{4, 5, 6}
x = append(x, y...) // Spreads `y` into `x`
fmt.Println(x) // [1 2 3 4 5 6]
```
✔ **`...` is required** to unpack a slice.  

---

#### **4️⃣ Capacity Expansion & Performance**  
```go
x := make([]int, 3, 5) // len=3, cap=5
x[0], x[1], x[2] = 1, 2, 3
x = append(x, 4, 5, 6) // Capacity exceeded → new array allocated
fmt.Println(x) // [1 2 3 4 5 6]
```
✔ If **capacity is exceeded, a new array is created**.  

---

#### **5️⃣ Efficient Slice Growth (Doubling Strategy)**  
```go
var s []int
for i := 0; i < 10; i++ {
    s = append(s, i)
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```
✔ **Capacity doubles dynamically** when needed.  

---

</details>

</details>


</details>

## Learning Resources

<details>
<summary>View contents</summary>

- [Build web application with golang](https://github.com/astaxie/build-web-application-with-golang) - `A golang ebook intro how to build a web with golang`
- [Go Patterns](https://github.com/tmrts/go-patterns) - `Curated list of Go design patterns, recipes and idioms`
- [Learn Go with Tests](https://github.com/quii/learn-go-with-tests) - `Learn Go with test-driven development`
- [Go for Javascript Developers](https://www.pazams.com/Go-for-Javascript-Developers/)
- [Creating a RESTful API With Golang](https://tutorialedge.net/golang/creating-restful-api-with-golang/)
- [Go Tour](https://tour.golang.org/list)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example](https://gobyexample.com/)
- [GOSAMPLES](https://gosamples.dev/)
- [Go Doc](https://golang.org/doc/)
- [Go Blog](https://blog.golang.org/)
- [Clean Go Article](https://github.com/Pungyeon/clean-go-article)
- [How To Code in Go](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go)
- [Google Style Guide](https://google.github.io/styleguide/go/)

### Video Tutorials

- [Go Tutorals - NerdCademy](https://youtube.com/playlist?list=PLujhHB_uAFJws6Vv5q1KDoaQ4YcpS9UOm)
- [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I) - `TechWorld with Nana`
- [Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU)
- [Learn Go Programming by Building 11 Projects – Full Course](https://www.youtube.com/watch?v=jFfo23yIWac)
- [Backend Master Class [Go + Postgres Docker + Kubernetes + gRPC]](https://youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)

</details>
