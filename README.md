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

</details>

---

### Practice Questions (Go Basics)

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

</details>

---

### Practice Questions (Essential Commands)

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

## Package Structure

A **package** in Go is a way to **organize code**.
It groups related variables, functions, types, and constants together.

Think of a package like a **folder of related code**.

<details>
<summary><strong>View contents</strong></summary>

**Examples of built-in packages:**

* `fmt` → printing and formatting
* `math` → math functions
* `net/http` → web servers and HTTP

**Every Go file starts with a package declaration:**

```go
package main  // Executable package (must have main function)

package mylib  // Library package
```

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

**Special package: `main`:**

* `package main` is used for **executable programs**
* It must contain a `main()` function
* Other packages are **reusable libraries**

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

### Creating your own Go package

**Example project structure:**

```
myproject/
│
├── go.mod
├── main.go
└── mathutils/
    └── mathutils.go
```

---

**Step 1: Create the package file**

**mathutils/mathutils.go**

```go
package mathutils

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}
```

### Important rules

* The folder name **does not have to** match the package name, but it usually should
* Functions starting with a **capital letter** are **exported** (public)
* Lowercase names are **private** to the package

---

**2. Importing and using a package**

`go.mod`

```go
module myproject

go 1.22
```

`main.go`

```go
package main

import (
	"fmt"
	"myproject/mathutils"
)

func main() {
	fmt.Println(mathutils.Add(5, 3))
	fmt.Println(mathutils.Subtract(10, 4))
}
```

### Key points

* Import path = **module name + folder name**
* Exported functions must start with a **capital letter**
* Use `packageName.FunctionName`

---

</details>

## Visibility Rules

- **If a name starts with a capital letter, it is exported (public).**
- **If a name starts with a lowercase letter, it is unexported (private).**

This applies to:

* functions
* variables
* constants
* structs
* struct fields
* interfaces
* methods

<details>
<summary><strong>View contents</strong></summary>

**Exported (visible outside the package):**

```go
package mathutils

// Exported (public) - accessible from other packages
func Add(a, b int) int {
	return a + b
}

// Unexported (private) - only accessible within this package
func validateNumber(num string) bool {
	return true
}

var Pi = 3.14 // Exported variable
var a = 10 // Unexported variable
```

Can be used in another package:

```go
mathutils.Add(2, 3)
fmt.Println(mathutils.Pi)

mathutils.validateNumber("3") // ❌ compile error
mathutils.a // ❌ compile error
```

---

**Struct field visibility (very common pitfall):**

```go
// Exported (public) - accessible from other packages
type User struct {
	Name  string // Exported field
	age   int    // Unexported field
}
```

From another package:

```go
user.Name = "Alice" // ✅
user.age = 20       // ❌ compile error
```

⚠️ This is **extremely important** when using:

* JSON encoding
* databases
* reflection

Example:

```go
type User struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}
```

`age` will **NOT** appear in JSON output.

---

**Best practices:**

- Export only what users need
- Keep internal helpers lowercase
- Use comments for exported names (required for docs)

Example:

```go
// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}
```

---

**Quick summary:**

* **Capital letter = public**
* **Lowercase letter = private**
* Visibility is **package-level**
* Struct fields must be exported to be accessible or encoded

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

### Type Parameters (Go 1.18+)

Go supports **generic programming** with type parameters.

**Quick Example:**

```go
// Generic function
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Works with any comparable type
fmt.Println(Max(10, 20))           // int: 20
fmt.Println(Max(3.14, 2.71))       // float64: 3.14
fmt.Println(Max("hello", "world")) // string: world
```

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

</details>

</details>

---

### Practice Questions (Packages, Variables, Data Types, etc.)

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

</details>

---

### Practice Questions (Control Structures and Loops)

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

</details>

---

### Practice Questions (Functions and Methods)

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

## Data Structures

### Arrays

Arrays have a **fixed length** and store elements of the same type.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/arrays)**

**Fixed-length sequences of elements:**

```go
// Declaration
var arr [5]int  // Array of 5 integers, zero-initialized
fmt.Println(arr) // [0 0 0]

// Initialization
arr1 := [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3, 4}  // Compiler counts length

// Access elements
arr1[0] = 10 // set element
value := arr1[0] // read element
```

**Important Characteristics:**
- **Length is part of the type**: `[3]int` and `[5]int` are different types
- **Arrays are values**: Assigning copies all elements
- **Passing to functions copies the array**

```go
func modify(arr [3]int) {
	arr[0] = 100  // Modifies copy, not original
}

a := [3]int{1, 2, 3}
modify(a)
fmt.Println(a)  // [1 2 3] (unchanged)
```

**Use pointer to modify:**
```go
func modify(arr *[3]int) {
	arr[0] = 100  // Modifies original
}

a := [3]int{1, 2, 3}
modify(&a)
fmt.Println(a)  // [100 2 3]
```

**Compiler-determined length:**

```go
arrNotMax := [...]int{2, 3, 4}
fmt.Println(arrNotMax, len(arrNotMax)) // [2 3 4] 3
```

**Slicing an array:**

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

**Append behavior:**

```go
fruitsToAdd := append(splicedFruits, "cherry", "pineapple", "guava")
```

* If capacity is exceeded:

  * A **new underlying array** is allocated
  * Capacity usually **doubles**

**Prepend (not built-in):**

```go
nums := []int{1, 2, 3}
nums = append([]int{0}, nums...)
```

**Multidimensional array:**

```go
multi := [2][3]int{{1, 2, 3}, {5, 6, 7}}
fmt.Println(multi) // [[1 2 3] [5 6 7]]
```

**When to Use Arrays:**
- Known, fixed size at compile time
- Performance-critical code (avoid heap allocation)
- **Generally prefer slices** for flexibility

</details>

---

### Slices

Slices are **more flexible than arrays** and are used almost everywhere in Go.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/slices)**

**Dynamic, flexible view into arrays:**

```go
// Declaration (creates nil slice)
var s []int

// mySlice[0] = 1 // ❌ compile-error
// Because the slice has no allocated memory yet

// Initialization
s1 := []int{1, 2, 3}
s2 := make([]int, 5)      // Length 5, capacity 5
s3 := make([]int, 5, 10)  // Length 5, capacity 10

// Slicing an array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2 3 4]
```

- `make` initializes with `zero-value` and allocates memory.
- `make([]T, len, cap)`

---

#### Slice Internals

**A slice has three components:**
- **Pointer** to underlying array
- **Length** (number of elements)
- **Capacity** (max elements before reallocation)

![image](https://user-images.githubusercontent.com/11992095/202870508-0739d792-8747-4e20-8cd2-0ffa888d5c08.png)

```go
// make([]T, len, cap)
s := make([]int, 3, 5)
fmt.Println(len(s))  // 3 (length)
fmt.Println(cap(s))  // 5 (capacity)
```

---

#### Append and Capacity

```go
s := []int{1, 2, 3}
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// len=3 cap=3 [1 2 3]

s = append(s, 4)
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// len=4 cap=6 [1 2 3 4] (capacity doubled)

s = append(s, 5, 6, 7)
fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// len=7 cap=12 [1 2 3 4 5 6 7]
```

**Key Points:**
- `append` may allocate a **new array** if capacity exceeded
- Always use result: `s = append(s, value)`
- Capacity typically doubles when exceeded

---

#### Slice Operations

**Slicing:**
```go
s := []int{1, 2, 3, 4, 5}

s[1:4]    // [2 3 4] (index 1 to 3)
s[:3]     // [1 2 3] (start to index 2)
s[2:]     // [3 4 5] (index 2 to end)
s[:]      // [1 2 3 4 5] (full slice)
```

**Copying:**
```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
fmt.Println(dst)  // [1 2 3]

// Copy returns number of elements copied
n := copy(dst, src)
fmt.Println(n)  // 3
```

**Unpack / spread slice:**

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

**Deleting elements (no built-in delete):**
```go
s := []int{1, 2, 3, 4, 5}

// Remove element at index 2
i := 2
s = append(s[:i], s[i+1:]...)
fmt.Println(s)  // [1 2 4 5]
```

**Inserting elements:**
```go
s := []int{1, 2, 5}
i := 2
value := 3

s = append(s[:i], append([]int{value}, s[i:]...)...)
fmt.Println(s)  // [1 2 3 5]
```

---

#### Generic Slice Operations (Go 1.18+)

**Common generic slice functions:**

```go
// Generic Map
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Generic Filter
func Filter[T any](slice []T, fn func(T) bool) []T {
    result := []T{}
    for _, v := range slice {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

// Usage
nums := []int{1, 2, 3, 4, 5}
doubled := Map(nums, func(n int) int { return n * 2 })
evens := Filter(nums, func(n int) bool { return n%2 == 0 })
```

---

#### Slice Gotchas

**Slices share underlying array:**
```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:3]  // [2 3]
slice2 := original[2:4]  // [3 4]

slice1[1] = 100
fmt.Println(original)  // [1 2 100 4 5]
fmt.Println(slice2)    // [100 4]
```

**Full slice expression (limit capacity):**
```go
s := []int{1, 2, 3, 4, 5}
slice := s[1:3:3]  // [low:high:max] limits capacity

fmt.Println(len(slice))  // 2
fmt.Println(cap(slice))  // 2 (not 4)

// Now append creates new array (doesn't affect s)
slice = append(slice, 100)
fmt.Println(s)      // [1 2 3 4 5] (unchanged)
fmt.Println(slice)  // [2 3 100]
```

</details>

---

### Maps

In Go, a **map** is a built-in data structure that stores **key–value pairs**.
It is similar to:

* Object (JavaScript)
* Dictionary (Python)
* HashMap (Java)

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/maps)**

**Hash tables mapping keys to values:**

```go
// Declaration (creates nil map)
var m map[string]int

// Initialization
m1 := make(map[string]int)
m2 := map[string]int{
	"Alice": 25,
	"Bob":   30,
}

// Set values
// If key exists → value updated
// If key does not exist → key created
m1["Alice"] = 25

// Get values
age := m1["Alice"]

// If a key does not exist, Go returns the zero value of the value type
age := m1["John] // 0

// Check existence
age, exists := m1["Alice"]
if exists {
	fmt.Println("Found:", age)
}

// Delete key
// Safe to delete a non-existing key
// No error is thrown
delete(m1, "Alice")

// Length
fmt.Println(len(m1))
```

#### Important Map Characteristics

* Map keys must be **comparable**

  * Allowed: `int`, `string`, `bool`, structs (if fields are comparable)
  * Not allowed: `slice`, `map`, `function`
* Map iteration order is **random**
* Maps are **reference types**

  * Assigning a map copies the reference, not the data

#### Map Iteration

```go
scores := map[string]int{
	"Alice": 90,
	"Bob":   85,
	"Carol": 95,
}

// Iterate over keys and values
for name, score := range scores {
	fmt.Printf("%s: %d\n", name, score)
}

// Only keys
for name := range scores {
	fmt.Println(name)
}
```

**⚠️ Map iteration order is random!**

```go
for i := 0; i < 3; i++ {
	for name := range scores {
		fmt.Print(name, " ")
	}
	fmt.Println()
}

// Different output each run:
// Bob Carol Alice
// Alice Bob Carol
// Carol Alice Bob
```

**To iterate in order, sort keys:**
```go
import "sort"

names := make([]string, 0, len(scores))
for name := range scores {
	names = append(names, name)
}
sort.Strings(names)

for _, name := range names {
	fmt.Printf("%s: %d\n", name, scores[name])
}
```

#### Generic Map Operations (Go 1.18+)

**Type-safe map operations:**

```go
// Generic Keys function
func Keys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

// Generic Values function
func Values[K comparable, V any](m map[K]V) []V {
    values := make([]V, 0, len(m))
    for _, v := range m {
        values = append(values, v)
    }
    return values
}

// Usage
m := map[string]int{"a": 1, "b": 2, "c": 3}
keys := Keys(m)     // []string
values := Values(m) // []int
```

#### Map Characteristics

**Zero value is nil:**
```go
var m map[string]int
// m["key"] = 1  // Panic! Cannot write to nil map

m = make(map[string]int)  // Must initialize
m["key"] = 1  // OK
```

**Maps are reference types:**
```go
func modify(m map[string]int) {
	m["key"] = 100  // Modifies original map
}

scores := map[string]int{"key": 1}
modify(scores)
fmt.Println(scores)  // map[key:100]
```

**Checking for missing keys:**
```go
value := m["missing"]  // Returns zero value (0 for int)

// Distinguish between "key exists with value 0" and "key missing"
value, exists := m["missing"]
if !exists {
	fmt.Println("Key not found")
}
```

**Using maps as sets:**
```go
visited := make(map[string]bool)
visited["page1"] = true
visited["page2"] = true

if visited["page1"] {
	fmt.Println("Already visited")
}

// More memory-efficient: use struct{} (zero bytes)
visited2 := make(map[string]struct{})
visited2["page1"] = struct{}{}

if _, exists := visited2["page1"]; exists {
	fmt.Println("Already visited")
}
```

</details>

---

### Structs

A **struct** is a composite data type that groups together variables (fields) under one name.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/structs)**

**Custom composite types:**

```go
type Person struct {
	Name string // Exported field (accessible outside the package)
	Age  int
}

// Create struct
p1 := Person{Name: "Alice", Age: 30} // named-field initialization (recommended)
p2 := Person{"Bob", 25}  // Positional (avoid in production)

// Zero value
var p3 Person  // Name: "", Age: 0

// Access fields
fmt.Println(p1.Name)
p1.Age = 31
```

#### Anonymous Structs

```go
// One-off use (e.g., JSON marshaling)
person := struct {
	Name string
	Age  int
}{
	Name: "Alice",
	Age:  30,
}
```

#### Struct Embedding (Composition)

**Go doesn't have inheritance; use composition:**

```go
type Address struct {
	City    string
	Country string
}

type Person struct {
	Name    string
	Age     int
	Address Address  // Nested struct
}

// Usage
p := Person{
	Name: "Alice",
	Age:  30,
	Address: Address{
		City:    "New York",
		Country: "USA",
	},
}

fmt.Println(p.Address.City)  // New York
```

**Embedded fields (promoted):**

```go
type Person struct {
	Name string
	Age  int
	Address  // Embedded (no field name)
}

p := Person{
	Name: "Alice",
	Age:  30,
	Address: Address{
		City:    "New York",
		Country: "USA",
	},
}

// Promoted fields accessible directly
fmt.Println(p.City)  // New York (promoted from Address)
fmt.Println(p.Address.City)  // Also valid
```

#### Struct Tags

**Metadata for reflection (used by JSON, XML, etc.):**

```go
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`  // Never serialize
	Age      int    `json:"age,string"`
}

u := User{Name: "Alice", Email: "alice@example.com", Age: 30}
data, _ := json.Marshal(u)
fmt.Println(string(data))
// {"name":"Alice","email":"alice@example.com","age":"30"}
```

**Common tags:**
- `json:"fieldName,omitempty"`: Omit if zero value
- `json:"-"`: Ignore field
- `xml:"fieldName"`
- `yaml:"fieldName"`
- `db:"column_name"`

#### Generic Structs (Go 1.18+)

**Structs can be generic:**

```go
// Generic Pair
type Pair[T, U any] struct {
    First  T
    Second U
}

// Usage
p1 := Pair[string, int]{First: "age", Second: 30}
p2 := Pair[int, string]{First: 1, Second: "one"}

// Generic Stack
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage
stack := Stack[int]{}
stack.Push(1)
stack.Push(2)
val, ok := stack.Pop() // 2, true
```

</details>

---

### Pointers

A **pointer** stores the **memory address** of a variable instead of a copy of its value.

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/pointers)**

**Pointer Syntax:**

```go
// Declare a pointer variable
var variableName *type

// Get the address of a variable
&variableName

// Dereference a pointer (get the value it points to)
*variableName
```

**Variables that store memory addresses:**

```go
x := 10
p := &x  // p holds the address of x

fmt.Println(p)   // 0xc000014078 (memory address)
fmt.Println(*p)  // 10 (dereference to get value)

*p = 20  // Modify value through pointer
fmt.Println(x)  // 20
```

#### Why Use Pointers?

**1. Modify function parameters:**
```go
func increment(x *int) {
	*x++
}

count := 0
increment(&count)
fmt.Println(count)  // 1
```

**2. Avoid copying large structs:**
```go
type LargeStruct struct {
	Data [1000]int
}

// Bad: Copies entire struct
func process(s LargeStruct) {
	// ...
}

// Good: Passes pointer (8 bytes)
func process(s *LargeStruct) {
	// ...
}
```

**3. Nullable values:**
```go
type User struct {
	Name  string
	Email *string  // Optional field
}

u := User{Name: "Alice"}
if u.Email != nil {
	fmt.Println(*u.Email)
}
```

#### Pointers vs Values

```go
type Counter struct {
	Value int
}

// Value receiver (doesn't modify original)
func (c Counter) Get() int {
	return c.Value
}

// Pointer receiver (modifies original)
func (c *Counter) Increment() {
	c.Value++
}

c := Counter{Value: 0}
c.Increment()  // Go auto-converts to (&c).Increment()
fmt.Println(c.Get())  // 1
```

**Google Style: Pointer vs Value Receivers**

Use pointer receivers when:
- Method modifies the receiver
- Receiver is large struct (avoid copying)
- Consistency (if any method uses pointer, use it everywhere)

Use value receivers when:
- Receiver is small (basic types, small structs)
- Receiver is immutable
- Receiver is a map, function, or channel (already reference types)

---

**Behavior of data types:**

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

</details>

---

#### Practice Questions (Data structures)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. The built-in function `__________` is used to add elements to a slice.
2. Maps in Go are __________ types, meaning modifications affect the original.
3. A struct can be embedded in another struct to achieve __________.
4. The `&` operator gets the __________ of a variable.

**True/False:**

1. ⬜ Arrays in Go can change size after declaration
2. ⬜ Map iteration order is guaranteed to be consistent
3. ⬜ Writing to a nil map causes a runtime panic
4. ⬜ Struct fields can be accessed through pointers using the dot notation

**Multiple Choice:**

1. What happens when you append to a slice beyond its capacity?
   - A) Runtime error
   - B) Original array is resized
   - C) New underlying array is allocated
   - D) Append fails silently

2. What is the zero value of a map?
   - A) Empty map `{}`
   - B) `nil`
   - C) Panic
   - D) `0`

**Code Output:**

```go
m := map[string]int{"a": 1, "b": 2}
m2 := m
m2["c"] = 3
fmt.Println(len(m))
```

What is printed?

---

### Answers

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. append
2. reference
3. composition
4. address

**True/False:**
1. ❌ False (arrays are fixed size)
2. ❌ False (iteration order is random)
3. ✅ True
4. ✅ True (Go auto-dereferences)

**Multiple Choice:**
1. **C** - New underlying array is allocated
2. **B** - `nil`

**Code Output:**
```
3
```
(Maps are reference types; `m` and `m2` point to same underlying data)

</details>

</details>

---

## Object-Oriented Programming

### Interfaces

**Interfaces** allow you to define **behavior** without specifying how that behavior is implemented.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/interfaces)**

**Interfaces define behavior (method sets):**

```go
type Shape interface {
	Area() float64
	Perimeter() float64
}
```

**Any type that implements all methods satisfies the interface:**

```go
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Rectangle implements Shape (implicitly)
```

#### Interface Satisfaction

**No explicit declaration needed:**

```go
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Both Circle and Rectangle implement Shape
func printInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}
	
	printInfo(rect)  // Works
	printInfo(circ)  // Works
}
```

#### Empty Interface

**`interface{}` or `any` (Go 1.18+) accepts any type:**

```go
func printAnything(v interface{}) {
	fmt.Println(v)
}

printAnything(42)
printAnything("hello")
printAnything([]int{1, 2, 3})
```

**Use cases:**
- Generic containers before Go 1.18
- JSON unmarshaling
- Reflection

**⚠️ Avoid overuse:** Prefer specific interfaces or generics

---

### Type Constraints (Go 1.18+)

**Interfaces can now define type constraints for generics:**

```go
// Traditional interface (methods)
type Stringer interface {
    String() string
}

// Type constraint (union of types)
type Number interface {
    int | int64 | float64
}

// Constraint with underlying types
type Integer interface {
    ~int | ~int64
}

// Usage in generic function
func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

Sum([]int{1, 2, 3})        // ✅ Works
Sum([]float64{1.5, 2.5})   // ✅ Works
// Sum([]string{"a", "b"})  // ❌ Compile error
```

**Key Difference:**
- **Method-based interfaces** → Describe behavior (runtime polymorphism)
- **Type constraint interfaces** → Limit generic types (compile-time)

---

### Type Assertions

**Extract concrete type from interface:**

```go
var i any = "hello"

// Type assertion
s := i.(string)
fmt.Println(s)  // "hello"

// Safe assertion (comma-ok idiom)
s, ok := i.(string)
if ok {
	fmt.Println(s)
} else {
	fmt.Println("Not a string")
}

// Unsafe assertion (panics if wrong type)
n := i.(int)  // Panic!
```

```go
func printShapeProps(s Shape) {
	if rect, ok := s.(Rectangle); ok {
		fmt.Printf("Height: %.2f, Width: %.2f\n", rect.Height, rect.Width)
	}
	if circle, ok := s.(Circle); ok {
		fmt.Printf("Radius: %.2f\n", circle.Radius)
	}
}

circle := Circle{10}
rectangle := Rectangle{10, 20}

printShapeProps(rectangle) // Height: 20.00, Width: 10.00
printShapeProps(circle)    // Radius: 10.00
```

#### Type Switch

**Check multiple types:**

```go
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	case nil:
		fmt.Println("Nil")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

describe(42)       // Integer: 42
describe("hello")  // String: hello
describe(true)     // Boolean: true
```

---

### Generic Interfaces (Go 1.18+)

**Interfaces themselves can be generic:**

```go
type Container[T any] interface {
    Add(item T)
    Get(index int) (T, bool)
    Size() int
}

// Implementation
type List[T any] struct {
    items []T
}

func (l *List[T]) Add(item T) {
    l.items = append(l.items, item)
}

func (l *List[T]) Get(index int) (T, bool) {
    if index < 0 || index >= len(l.items) {
        var zero T
        return zero, false
    }
    return l.items[index], true
}

func (l *List[T]) Size() int {
    return len(l.items)
}

// Usage
var container Container[string] = &List[string]{}
container.Add("hello")
container.Add("world")
```

---

### Common Interfaces

#### Stringer (custom string representation)

```go
type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

p := Person{Name: "Alice", Age: 30}
fmt.Println(p)  // Alice (30 years old)
```

#### Reader and Writer

**IO operations:**

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Many types implement these:
// - os.File
// - bytes.Buffer
// - strings.Reader
// - net.Conn
```

**Example:**
```go
import (
	"io"
	"os"
	"strings"
)

func readAll(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	return string(data), err
}

// Works with any Reader
file, _ := os.Open("file.txt")
content, _ := readAll(file)

reader := strings.NewReader("Hello, World!")
content, _ = readAll(reader)
```

#### Error Interface

```go
type error interface {
	Error() string
}

// Custom error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validate(email string) error {
	if !strings.Contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Message: "must contain @",
		}
	}
	return nil
}
```

---

### Interface Best Practices

**1. Small interfaces (Google Style)**

```go
// Good: Single method
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Bad: Too many methods
type FileManager interface {
	Open() error
	Close() error
	Read() ([]byte, error)
	Write([]byte) error
	Delete() error
	Rename(string) error
}
```

**2. Accept interfaces, return structs**

```go
// Good
func ProcessData(r io.Reader) (*Result, error) {
	// ...
}

// Avoid
func ProcessData(f *os.File) (*Result, error) {
	// ...
}
```

**3. Define interfaces where used**

```go
// consumer.go
type DataStore interface {
	Get(key string) (string, error)
}

func ProcessUser(store DataStore, userID string) {
	// ...
}

// implementation.go
type RedisStore struct {
	// ...
}

func (r *RedisStore) Get(key string) (string, error) {
	// ...
}
```

---

### Polymorphism Example

**Real-world payment processing:**

```go
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionID string) error
}

type CreditCard struct {
	Number string
	CVV    string
}

func (c CreditCard) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via credit card\n", amount)
	return nil
}

func (c CreditCard) RefundPayment(transactionID string) error {
	fmt.Println("Refunding credit card payment")
	return nil
}

type PayPal struct {
	Email string
}

func (p PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via PayPal\n", amount)
	return nil
}

func (p PayPal) RefundPayment(transactionID string) error {
	fmt.Println("Refunding PayPal payment")
	return nil
}

// Single function works with any payment method
func checkout(processor PaymentProcessor, amount float64) error {
	return processor.ProcessPayment(amount)
}

func main() {
	card := CreditCard{Number: "1234-5678", CVV: "123"}
	paypal := PayPal{Email: "user@example.com"}
	
	checkout(card, 100.50)    // Works
	checkout(paypal, 200.75)  // Works
}
```

</details>

---

#### Practice Questions (Interfaces)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. An interface in Go is satisfied __________ when a type implements all its methods.
2. The empty interface `interface{}` can hold values of __________ type.
3. A type assertion extracts the __________ type from an interface value.
4. The `__________` interface allows custom string representation via the String() method.

**True/False:**

1. ⬜ Go requires explicit declaration to implement an interface
2. ⬜ Type assertions can cause a panic if the type is incorrect
3. ⬜ A single type can implement multiple interfaces
4. ⬜ Interfaces should typically have many methods for flexibility

**Multiple Choice:**

1. What is the safest way to perform a type assertion?
   - A) `x := i.(string)`
   - B) `x, ok := i.(string)`
   - C) `x := string(i)`
   - D) `x := i.toString()`

2. Which pattern is recommended by Google Style Guide?
   - A) Large interfaces with many methods
   - B) Accept interfaces, return structs
   - C) Return interfaces from functions
   - D) Define interfaces in implementation packages

**Code Challenge:**

Create a `Logger` interface with methods `Info(msg string)` and `Error(msg string)`. Implement it for both `ConsoleLogger` and `FileLogger`.

---

### Answers

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. implicitly
2. any
3. concrete
4. Stringer

**True/False:**
1. ❌ False (implicit satisfaction)
2. ✅ True (use comma-ok idiom for safety)
3. ✅ True
4. ❌ False (prefer small, focused interfaces)

**Multiple Choice:**
1. **B** - `x, ok := i.(string)` (comma-ok idiom)
2. **B** - Accept interfaces, return structs

**Code Challenge:**
```go
type Logger interface {
	Info(msg string)
	Error(msg string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Info(msg string) {
	fmt.Println("[INFO]", msg)
}

func (c ConsoleLogger) Error(msg string) {
	fmt.Println("[ERROR]", msg)
}

type FileLogger struct {
	file *os.File
}

func (f FileLogger) Info(msg string) {
	f.file.WriteString("[INFO] " + msg + "\n")
}

func (f FileLogger) Error(msg string) {
	f.file.WriteString("[ERROR] " + msg + "\n")
}

// Usage
func logMessage(logger Logger, msg string) {
	logger.Info(msg)
}

consoleLog := ConsoleLogger{}
logMessage(consoleLog, "Application started")
```

---

</details>

</details>

## Error Handling

Go treats errors as **values**, not exceptions.
This design forces developers to **handle errors explicitly**, making programs more predictable and easier to reason about.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/errors)**

### Error Basics

**Errors are values:**

```go
// Built-in error interface
type error interface {
	Error() string
}

// Creating errors
import "errors"

err1 := errors.New("something went wrong")
err2 := fmt.Errorf("user %s not found", "Alice")
```

**Standard error handling pattern:**

```go
result, err := doSomething()
if err != nil {
	// Handle error
	return err
}
// Use result
```

---

### Error Handling Patterns

#### Early Return (Google Style)

**Preferred:**
```go
func processUser(id int) error {
	user, err := getUser(id)
	if err != nil {
		return fmt.Errorf("getting user: %w", err)
	}
	
	if !user.Active {
		return errors.New("user is inactive")
	}
	
	if err := saveUser(user); err != nil {
		return fmt.Errorf("saving user: %w", err)
	}
	
	return nil
}
```

**Avoid:**
```go
// Bad: Nested if statements
func processUser(id int) error {
	user, err := getUser(id)
	if err == nil {
		if user.Active {
			err = saveUser(user)
			if err == nil {
				return nil
			}
		}
	}
	return err
}
```

#### Error Wrapping (Go 1.13+)

**Wrap errors to preserve context:**

```go
import (
	"errors"
	"fmt"
)

func readConfig() error {
	_, err := os.ReadFile("config.json")
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}
	return nil
}

func initialize() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("initialization failed: %w", err)
	}
	return nil
}

// Error chain preserved:
// initialization failed: reading config: open config.json: no such file or directory
```

**Unwrapping errors:**

```go
err := initialize()
if errors.Is(err, os.ErrNotExist) {
	fmt.Println("Config file doesn't exist")
}

var pathErr *os.PathError
if errors.As(err, &pathErr) {
	fmt.Println("Path error:", pathErr.Path)
}
```

---

### Custom Errors

**Struct-based errors:**

```go
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Message: "must contain @ symbol",
		}
	}
	return nil
}

// Usage
err := validateEmail("invalid")
if err != nil {
	var valErr ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("Validation failed for %s: %s\n", valErr.Field, valErr.Message)
	}
}
```

**Sentinel errors (predefined errors):**

```go
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

func getUser(id int) (*User, error) {
	// ...
	return nil, ErrNotFound
}

// Check with errors.Is
user, err := getUser(123)
if errors.Is(err, ErrNotFound) {
	fmt.Println("User not found")
}
```

---

### Panic and Recover

**Panic stops normal execution:**

```go
func mustConnect(url string) *Connection {
	conn, err := connect(url)
	if err != nil {
		panic(err)  // Rarely used; prefer returning errors
	}
	return conn
}
```

**When to panic:**
- **Programmer errors** (nil pointer dereference, out-of-bounds)
- **Initialization failures** in `init()` or `main()`
- **Never in library code** (return errors instead)

**Recovering from panic:**

```go
func safeExecute() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	
	panic("something went wrong")
	fmt.Println("This won't execute")
}

safeExecute()
fmt.Println("Program continues")

// Output:
// Recovered from panic: something went wrong
// Program continues
```

**Real-world example (HTTP server):**

```go
func handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v", r)
			http.Error(w, "Internal Server Error", 500)
		}
	}()
	
	// Handler logic that might panic
}
```

---

### Error Handling Best Practices

**1. Return errors, don't panic**

```go
// Good
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Bad
func divide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
```

**2. Add context when wrapping**

```go
// Good: Clear error chain
func loadConfig() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return fmt.Errorf("loading config file: %w", err)
	}
	
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}
	
	return nil
}

// Error message:
// loading config file: parsing config: yaml: line 5: mapping values are not allowed
```

**3. Handle errors at appropriate level**

```go
// Low-level: Return errors
func fetchData() ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetching data: %w", err)
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// High-level: Decide what to do with errors
func processData() {
	data, err := fetchData()
	if err != nil {
		log.Printf("Failed to fetch data: %v", err)
		// Retry, use cache, or return error
		return
	}
	// Process data
}
```

**4. Use named error types for specific handling**

```go
type NetworkError struct {
	Op  string
	Err error
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("network error during %s: %v", e.Op, e.Err)
}

func (e NetworkError) Unwrap() error {
	return e.Err
}

// Usage
err := doRequest()
var netErr NetworkError
if errors.As(err, &netErr) {
	// Retry on network errors
	time.Sleep(time.Second)
	err = doRequest()
}
```

</details>

---

### Practice Questions (Error Handling)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. The `fmt.Errorf` function with `%w` verb is used to __________ errors.
2. The `errors.__________()` function checks if an error matches a specific value.
3. The `__________` keyword stops normal execution and begins panicking.
4. A `defer` statement with `recover()` can __________ from a panic.

**True/False:**

1. ⬜ Panics should be used for normal error handling in Go
2. ⬜ Error wrapping with %w preserves the error chain
3. ⬜ Library code should return errors, not panic
4. ⬜ The errors.As function is used to check error values

**Multiple Choice:**

1. What is the recommended pattern for error handling?
   - A) Nested if-else statements
   - B) Early return with error checks
   - C) Panic for all errors
   - D) Ignore errors

2. When should you use panic?
   - A) For all error conditions
   - B) For validation errors
   - C) For unrecoverable initialization failures
   - D) For user input errors

**Code Challenge:**

Write a function that validates a user struct and returns a custom error type with details about which field failed validation.

---

### Answers

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. wrap
2. Is
3. panic
4. recover

**True/False:**
1. ❌ False (return errors instead)
2. ✅ True
3. ✅ True
4. ❌ False (errors.As is for type assertions; errors.Is checks values)

**Multiple Choice:**
1. **B** - Early return with error checks
2. **C** - For unrecoverable initialization failures

**Code Challenge:**
```go
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s - %s", e.Field, e.Message)
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) Validate() error {
	if u.Name == "" {
		return ValidationError{Field: "name", Message: "cannot be empty"}
	}
	if !strings.Contains(u.Email, "@") {
		return ValidationError{Field: "email", Message: "must contain @"}
	}
	if u.Age < 0 || u.Age > 150 {
		return ValidationError{Field: "age", Message: "must be between 0 and 150"}
	}
	return nil
}

// Usage
user := User{Name: "", Email: "test@example.com", Age: 25}
if err := user.Validate(); err != nil {
	var valErr ValidationError
	if errors.As(err, &valErr) {
		fmt.Printf("Field: %s, Error: %s\n", valErr.Field, valErr.Message)
	}
}
```

---

</details>

</details>

## Generics

Generics (introduced in **Go 1.18**) allow you to write **type-safe, reusable code** without sacrificing compile-time checks.
 
<details>
<summary>View contents</summary>

**Before Generics (Go < 1.18):**

```go
// Separate functions for each type
func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func MaxFloat(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}

func MaxString(a, b string) string {
    if a > b {
        return a
    }
    return b
}

// Or use interface{} (type erasure)
func Max(a, b interface{}) interface{} {
    // Type assertions needed
    // Runtime type checking
    // Not type-safe!
}
```

**With Generics (Go 1.18+):**

```go
func Max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Type-safe usage
fmt.Println(Max(10, 20))        // int: 20
fmt.Println(Max(3.14, 2.71))    // float64: 3.14
fmt.Println(Max("hello", "world")) // string: world
```

---

### Why Generics?

**Problems Generics Solve:**

1. **Code Duplication**
   ```go
   // Before: Multiple implementations
   func ReverseIntSlice(s []int) []int { /* ... */ }
   func ReverseStringSlice(s []string) []string { /* ... */ }
   
   // After: Single generic implementation
   func Reverse[T any](s []T) []T { /* ... */ }
   ```

2. **Type Safety**
   ```go
   // Before: interface{} loses type information
   func First(slice []interface{}) interface{} {
       return slice[0] // Caller must type assert
   }
   
   // After: Type-safe
   func First[T any](slice []T) T {
       return slice[0] // Type preserved
   }
   ```

3. **Performance**
   ```go
   // interface{} has allocation overhead
   // Generics compile to efficient code
   ```

---

### Real-World Analogy

```
Think of generics like a universal remote control:

Without Generics:
- TV Remote (only controls TVs)
- DVD Remote (only controls DVDs)
- Stereo Remote (only controls stereos)
→ Need separate remote for each device

With Generics:
- Universal Remote [Device] (controls any device)
→ One remote, many devices
```

---

## Type Parameters

### Syntax

```go
func FunctionName[T TypeConstraint](param T) T {
    // Use T as a type
}
```

**Components:**
- `[T TypeConstraint]` - Type parameter declaration
- `T` - Type parameter (placeholder for actual type)
- `TypeConstraint` - What types T can be

---

### Single Type Parameter

```go
func Print[T any](value T) {
    fmt.Println(value)
}

// Usage
Print[int](42)
Print[string]("hello")
Print(42)        // Type inference
Print("hello")   // Type inference
```

---

### Multiple Type Parameters

```go
func Pair[K comparable, V any](key K, value V) map[K]V {
    return map[K]V{key: value}
}

// Usage
m1 := Pair[string, int]("age", 30)
m2 := Pair("name", "Alice")  // Type inference
```

---

### Type Inference

**Go can often infer type parameters:**

```go
func Identity[T any](val T) T {
    return val
}

// Explicit
x := Identity[int](42)

// Inferred (preferred)
x := Identity(42)
```

**When type inference works:**
```go
func Add[T int | float64](a, b T) T {
    return a + b
}

Add(1, 2)        // ✅ Infers T=int
Add(1.5, 2.5)    // ✅ Infers T=float64
// Add(1, 2.5)   // ❌ Compile error (mixed types)
```

**When explicit types needed:**
```go
func Make[T any]() T {
    var zero T
    return zero
}

// Must be explicit (no parameters to infer from)
x := Make[int]()
s := Make[string]()
```

---

## Type Constraints

### Built-in Constraints

**1. `any` (alias for `interface{}`)**

```go
// Accepts any type
func Print[T any](v T) {
    fmt.Println(v)
}
```

**2. `comparable`**

```go
// Types that support == and !=
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {  // Requires comparable
            return true
        }
    }
    return false
}

// Works with:
Contains([]int{1, 2, 3}, 2)           // ✅
Contains([]string{"a", "b"}, "a")     // ✅

// Doesn't work with:
// Contains([][]int{{1}}, []int{1})   // ❌ Slices not comparable
```

---

### Custom Constraints (Interface-based)

**Type Union Constraints:**

```go
type Number interface {
    int | int64 | float64
}

func Sum[T Number](numbers []T) T {
    var total T
    for _, n := range numbers {
        total += n
    }
    return total
}

// Usage
Sum([]int{1, 2, 3})           // 6
Sum([]float64{1.5, 2.5})      // 4.0
// Sum([]string{"a", "b"})    // ❌ Compile error
```

**With Approximation (`~`)**

```go
// Without ~: Only exact types
type Integer interface {
    int | int64
}

// With ~: Underlying types too
type Integer interface {
    ~int | ~int64
}

type MyInt int

func Double[T Integer](n T) T {
    return n * 2
}

var x MyInt = 5
fmt.Println(Double(x))  // ✅ Works with ~ in constraint
```

**Why `~` Matters:**

```go
type Celsius float64
type Fahrenheit float64

type Temperature interface {
    ~float64  // Includes Celsius, Fahrenheit
}

func Average[T Temperature](temps []T) T {
    var sum T
    for _, t := range temps {
        sum += t
    }
    return sum / T(len(temps))
}

// Works with custom types
temps := []Celsius{0, 10, 20}
avg := Average(temps)  // ✅
```

---

### Method Constraints

```go
// Types must implement String() method
type Stringer interface {
    String() string
}

func PrintAll[T Stringer](items []T) {
    for _, item := range items {
        fmt.Println(item.String())  // Guaranteed to exist
    }
}

// Example type
type Person struct {
    Name string
}

func (p Person) String() string {
    return p.Name
}

// Usage
people := []Person{{"Alice"}, {"Bob"}}
PrintAll(people)  // ✅ Person implements Stringer
```

---

### Combined Constraints

```go
// Multiple constraints
type Ordered interface {
    ~int | ~int64 | ~float64 | ~string
}

type Printable interface {
    String() string
}

// Type must satisfy both
type OrderedPrintable interface {
    Ordered
    Printable
}
```

---

## Generic Functions

### Simple Generic Functions

**1. Identity Function**

```go
func Identity[T any](val T) T {
    return val
}

x := Identity(42)      // int
s := Identity("hello") // string
```

**2. Swap Function**

```go
func Swap[T any](a, b T) (T, T) {
    return b, a
}

x, y := Swap(1, 2)
fmt.Println(x, y)  // 2 1
```

**3. First Element**

```go
func First[T any](slice []T) (T, bool) {
    if len(slice) == 0 {
        var zero T
        return zero, false
    }
    return slice[0], true
}

// Usage
val, ok := First([]int{1, 2, 3})
fmt.Println(val, ok)  // 1 true

val2, ok2 := First([]string{})
fmt.Println(val2, ok2)  // "" false
```

---

### Generic Algorithms

**1. Map (Transform Slice)**

```go
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Usage
nums := []int{1, 2, 3, 4, 5}

// Square numbers
squares := Map(nums, func(n int) int {
    return n * n
})
// [1, 4, 9, 16, 25]

// Convert to strings
strs := Map(nums, func(n int) string {
    return fmt.Sprintf("num-%d", n)
})
// ["num-1", "num-2", "num-3", "num-4", "num-5"]
```

**2. Filter**

```go
func Filter[T any](slice []T, fn func(T) bool) []T {
    result := []T{}
    for _, v := range slice {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

// Usage
nums := []int{1, 2, 3, 4, 5, 6}

// Get even numbers
evens := Filter(nums, func(n int) bool {
    return n%2 == 0
})
// [2, 4, 6]

// Get numbers > 3
large := Filter(nums, func(n int) bool {
    return n > 3
})
// [4, 5, 6]
```

**3. Reduce**

```go
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
    result := initial
    for _, v := range slice {
        result = fn(result, v)
    }
    return result
}

// Usage
nums := []int{1, 2, 3, 4, 5}

// Sum
sum := Reduce(nums, 0, func(acc, n int) int {
    return acc + n
})
// 15

// Product
product := Reduce(nums, 1, func(acc, n int) int {
    return acc * n
})
// 120

// Concatenate strings
words := []string{"Hello", " ", "World"}
sentence := Reduce(words, "", func(acc, word string) string {
    return acc + word
})
// "Hello World"
```

**4. Find**

```go
func Find[T any](slice []T, fn func(T) bool) (T, bool) {
    for _, v := range slice {
        if fn(v) {
            return v, true
        }
    }
    var zero T
    return zero, false
}

// Usage
nums := []int{1, 2, 3, 4, 5}

// Find first even number
even, found := Find(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(even, found)  // 2 true

// Find number > 10
large, found := Find(nums, func(n int) bool {
    return n > 10
})
fmt.Println(large, found)  // 0 false
```

---

### Variadic Generic Functions

```go
func Min[T Ordered](values ...T) T {
    if len(values) == 0 {
        panic("Min requires at least one value")
    }
    
    min := values[0]
    for _, v := range values[1:] {
        if v < min {
            min = v
        }
    }
    return min
}

// Usage
fmt.Println(Min(5, 2, 8, 1, 9))      // 1
fmt.Println(Min(3.14, 2.71, 1.41))   // 1.41
fmt.Println(Min("zebra", "apple"))   // "apple"
```

---

## Generic Types

### Generic Structs

**1. Generic Stack**

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

// Usage
stack := Stack[int]{}
stack.Push(1)
stack.Push(2)
stack.Push(3)

val, ok := stack.Pop()
fmt.Println(val, ok)  // 3 true
```

**2. Generic Queue**

```go
type Queue[T any] struct {
    items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

// Usage
queue := Queue[string]{}
queue.Enqueue("first")
queue.Enqueue("second")

val, ok := queue.Dequeue()
fmt.Println(val, ok)  // "first" true
```

**3. Generic Pair**

```go
type Pair[T, U any] struct {
    First  T
    Second U
}

func NewPair[T, U any](first T, second U) Pair[T, U] {
    return Pair[T, U]{First: first, Second: second}
}

// Usage
p1 := Pair[string, int]{"age", 30}
p2 := NewPair("name", "Alice")
p3 := NewPair(1, "one")
```

**4. Generic Optional (Like Option in Rust)**

```go
type Optional[T any] struct {
    value   T
    present bool
}

func Some[T any](value T) Optional[T] {
    return Optional[T]{value: value, present: true}
}

func None[T any]() Optional[T] {
    return Optional[T]{present: false}
}

func (o Optional[T]) IsPresent() bool {
    return o.present
}

func (o Optional[T]) Get() (T, bool) {
    return o.value, o.present
}

func (o Optional[T]) OrElse(defaultValue T) T {
    if o.present {
        return o.value
    }
    return defaultValue
}

// Usage
opt1 := Some(42)
val, ok := opt1.Get()
fmt.Println(val, ok)  // 42 true

opt2 := None[int]()
val2 := opt2.OrElse(0)
fmt.Println(val2)  // 0
```

**5. Generic Result (Error Handling)**

```go
type Result[T any] struct {
    value T
    err   error
}

func Ok[T any](value T) Result[T] {
    return Result[T]{value: value}
}

func Err[T any](err error) Result[T] {
    return Result[T]{err: err}
}

func (r Result[T]) IsOk() bool {
    return r.err == nil
}

func (r Result[T]) Unwrap() (T, error) {
    return r.value, r.err
}

// Usage
func Divide(a, b float64) Result[float64] {
    if b == 0 {
        return Err[float64](errors.New("division by zero"))
    }
    return Ok(a / b)
}

result := Divide(10, 2)
if result.IsOk() {
    val, _ := result.Unwrap()
    fmt.Println(val)  // 5
}
```

---

### Generic Slices (Type Aliases)

```go
type Numbers[T int | float64] []T

func (n Numbers[T]) Sum() T {
    var total T
    for _, v := range n {
        total += v
    }
    return total
}

// Usage
nums := Numbers[int]{1, 2, 3, 4, 5}
fmt.Println(nums.Sum())  // 15

prices := Numbers[float64]{9.99, 19.99, 29.99}
fmt.Println(prices.Sum())  // 59.97
```

---

### Generic Maps

```go
type Cache[K comparable, V any] struct {
    data map[K]V
    mu   sync.RWMutex
}

func NewCache[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{
        data: make(map[K]V),
    }
}

func (c *Cache[K, V]) Set(key K, value V) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

// Usage
cache := NewCache[string, int]()
cache.Set("age", 30)
val, ok := cache.Get("age")
fmt.Println(val, ok)  // 30 true
```

---

## Generic Interfaces

### Simple Generic Interface

```go
type Container[T any] interface {
    Add(item T)
    Get(index int) (T, bool)
    Size() int
}

// Implementation
type List[T any] struct {
    items []T
}

func (l *List[T]) Add(item T) {
    l.items = append(l.items, item)
}

func (l *List[T]) Get(index int) (T, bool) {
    if index < 0 || index >= len(l.items) {
        var zero T
        return zero, false
    }
    return l.items[index], true
}

func (l *List[T]) Size() int {
    return len(l.items)
}

// Usage
var c Container[string] = &List[string]{}
c.Add("hello")
c.Add("world")
```

---

### Constraint Interface with Methods

```go
type Numeric interface {
    ~int | ~int64 | ~float64
}

type Addable[T Numeric] interface {
    Add(other T) T
}

type Counter[T Numeric] struct {
    value T
}

func (c Counter[T]) Add(other T) T {
    return c.value + other
}

func Sum[T Numeric, A Addable[T]](items []A) T {
    var total T
    for _, item := range items {
        total = item.Add(0)  // Access Add method
    }
    return total
}
```

---

## Advanced Constraints

### constraints Package

**Official constraints package:**

```go
import "golang.org/x/exp/constraints"

// Signed integers
type Signed interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned integers
type Unsigned interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// All integers
type Integer interface {
    Signed | Unsigned
}

// Floating point
type Float interface {
    ~float32 | ~float64
}

// All numbers
type Ordered interface {
    Integer | Float | ~string
}
```

**Usage:**

```go
import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func Sum[T constraints.Integer | constraints.Float](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}
```

---

### Type Sets

```go
// Empty type set (no types satisfy this)
type Impossible interface {
    int
    string  // Can't be both int AND string
}

// Union type set
type IntOrString interface {
    int | string  // Can be int OR string
}

// Method + type constraint
type StringerInt interface {
    ~int
    String() string
}
```

---

## Real-World Examples

### 1. Generic Linked List

```go
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type LinkedList[T any] struct {
    head *Node[T]
    size int
}

func (l *LinkedList[T]) Add(value T) {
    newNode := &Node[T]{Value: value}
    
    if l.head == nil {
        l.head = newNode
    } else {
        current := l.head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    l.size++
}

func (l *LinkedList[T]) ToSlice() []T {
    result := make([]T, 0, l.size)
    current := l.head
    for current != nil {
        result = append(result, current.Value)
        current = current.Next
    }
    return result
}

// Usage
list := &LinkedList[int]{}
list.Add(1)
list.Add(2)
list.Add(3)
fmt.Println(list.ToSlice())  // [1 2 3]
```

---

### 2. Generic Binary Tree

```go
type TreeNode[T constraints.Ordered] struct {
    Value T
    Left  *TreeNode[T]
    Right *TreeNode[T]
}

type BinaryTree[T constraints.Ordered] struct {
    root *TreeNode[T]
}

func (t *BinaryTree[T]) Insert(value T) {
    t.root = t.insertNode(t.root, value)
}

func (t *BinaryTree[T]) insertNode(node *TreeNode[T], value T) *TreeNode[T] {
    if node == nil {
        return &TreeNode[T]{Value: value}
    }
    
    if value < node.Value {
        node.Left = t.insertNode(node.Left, value)
    } else if value > node.Value {
        node.Right = t.insertNode(node.Right, value)
    }
    
    return node
}

func (t *BinaryTree[T]) InOrder() []T {
    result := []T{}
    t.inOrder(t.root, &result)
    return result
}

func (t *BinaryTree[T]) inOrder(node *TreeNode[T], result *[]T) {
    if node == nil {
        return
    }
    t.inOrder(node.Left, result)
    *result = append(*result, node.Value)
    t.inOrder(node.Right, result)
}

// Usage
tree := &BinaryTree[int]{}
tree.Insert(5)
tree.Insert(3)
tree.Insert(7)
tree.Insert(1)
fmt.Println(tree.InOrder())  // [1 3 5 7]
```

---

### 3. Generic Set

```go
type Set[T comparable] struct {
    data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
    return &Set[T]{
        data: make(map[T]struct{}),
    }
}

func (s *Set[T]) Add(item T) {
    s.data[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
    delete(s.data, item)
}

func (s *Set[T]) Contains(item T) bool {
    _, exists := s.data[item]
    return exists
}

func (s *Set[T]) Size() int {
    return len(s.data)
}

func (s *Set[T]) ToSlice() []T {
    result := make([]T, 0, len(s.data))
    for k := range s.data {
        result = append(result, k)
    }
    return result
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
    result := NewSet[T]()
    for k := range s.data {
        result.Add(k)
    }
    for k := range other.data {
        result.Add(k)
    }
    return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
    result := NewSet[T]()
    for k := range s.data {
        if other.Contains(k) {
            result.Add(k)
        }
    }
    return result
}

// Usage
set1 := NewSet[int]()
set1.Add(1)
set1.Add(2)
set1.Add(3)

set2 := NewSet[int]()
set2.Add(2)
set2.Add(3)
set2.Add(4)

union := set1.Union(set2)
fmt.Println(union.ToSlice())  // [1 2 3 4] (order may vary)

intersection := set1.Intersection(set2)
fmt.Println(intersection.ToSlice())  // [2 3] (order may vary)
```

---

### 4. Generic HTTP Response Wrapper

```go
type Response[T any] struct {
    Data    T      `json:"data,omitempty"`
    Error   string `json:"error,omitempty"`
    Success bool   `json:"success"`
}

func SuccessResponse[T any](data T) Response[T] {
    return Response[T]{
        Data:    data,
        Success: true,
    }
}

func ErrorResponse[T any](err string) Response[T] {
    return Response[T]{
        Error:   err,
        Success: false,
    }
}

// Usage in HTTP handlers
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    type User struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }
    
    user := User{ID: 1, Name: "Alice"}
    response := SuccessResponse(user)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
    type User struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    }
    
    users := []User{
        {ID: 1, Name: "Alice"},
        {ID: 2, Name: "Bob"},
    }
    
    response := SuccessResponse(users)
    json.NewEncoder(w).Encode(response)
}
```

---

## Best Practices

### 1. When to Use Generics

**✅ Good Use Cases:**

```go
// Data structures
type Stack[T any] struct { /* ... */ }

// Algorithms
func Sort[T constraints.Ordered](slice []T) { /* ... */ }

// Utility functions
func Map[T, U any](slice []T, fn func(T) U) []U { /* ... */ }

// API wrappers
type Response[T any] struct { /* ... */ }
```

**❌ Avoid Generics When:**

```go
// When interface{} is simpler and performance doesn't matter
func Print(v interface{}) { fmt.Println(v) }

// When only used with one type
func ProcessInts(nums []int) { /* ... */ }  // Just use int

// When logic differs significantly per type
// Better to write separate functions
```

---

### 2. Naming Conventions

```go
// Single letter for simple cases
func Identity[T any](v T) T { return v }

// Descriptive for complex cases
func Transform[Input, Output any](in Input, fn func(Input) Output) Output

// Common conventions:
// T - Type
// K - Key
// V - Value
// E - Element
// R - Result
```

---

### 3. Keep Constraints Minimal

```go
// ❌ Too restrictive
func Process[T int](value T) { /* ... */ }

// ✅ More flexible
func Process[T constraints.Integer](value T) { /* ... */ }

// ✅ Even better if you only need addition
func Process[T interface{ ~int | ~int64 }](value T) { /* ... */ }
```

---

### 4. Don't Overuse Type Parameters

```go
// ❌ Unnecessary
func Add[T int](a, b T) T {
    return a + b
}

// ✅ Just use int
func Add(a, b int) int {
    return a + b
}

// ✅ Generic only when supporting multiple types
func Add[T constraints.Integer | constraints.Float](a, b T) T {
    return a + b
}
```

---

### 5. Prefer Interfaces for Behavior

```go
// ✅ Good: Use interface for behavior
type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

func Sort(data Sorter) { /* ... */ }

// Instead of overly generic
func Sort[T any](data []T, less func(T, T) bool) { /* ... */ }
```

</details>

---

### Practice Questions (Generics)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. The `any` constraint is an alias for __________.
2. Types that support `==` and `!=` must use the __________ constraint.
3. The `~` symbol in constraints matches the __________ type.
4. Type parameters are declared in __________ brackets.

**True/False:**

1. ⬜ Generics were introduced in Go 1.18
2. ⬜ Type parameters can always be inferred
3. ⬜ Slices are comparable and can use the `comparable` constraint
4. ⬜ Generic types can have methods

**Multiple Choice:**

1. Which constraint allows comparison with `>` and `<`?
   - A) `any`
   - B) `comparable`
   - C) `constraints.Ordered`
   - D) `interface{}`

2. What does `~int` mean in a constraint?
   - A) Not int
   - B) int only
   - C) int and types with int as underlying type
   - D) Approximately int

**Code Challenge:**

Write a generic function `Chunk` that splits a slice into chunks of specified size.

```go
func Chunk[T any](slice []T, size int) [][]T {
    // Your implementation
}

// Example usage:
nums := []int{1, 2, 3, 4, 5, 6, 7}
chunks := Chunk(nums, 3)
// Result: [[1, 2, 3], [4, 5, 6], [7]]
```

---

## Answers

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. interface{}
2. comparable
3. underlying
4. square

**True/False:**
1. ✅ True
2. ❌ False (sometimes explicit types needed)
3. ❌ False (slices are not comparable)
4. ✅ True

**Multiple Choice:**
1. **C** - `constraints.Ordered`
2. **C** - int and types with int as underlying type

**Code Challenge:**

```go
func Chunk[T any](slice []T, size int) [][]T {
    if size <= 0 {
        return nil
    }
    
    result := [][]T{}
    for i := 0; i < len(slice); i += size {
        end := i + size
        if end > len(slice) {
            end = len(slice)
        }
        result = append(result, slice[i:end])
    }
    return result
}

// Usage
nums := []int{1, 2, 3, 4, 5, 6, 7}
chunks := Chunk(nums, 3)
fmt.Println(chunks)  // [[1 2 3] [4 5 6] [7]]

words := []string{"a", "b", "c", "d", "e"}
wordChunks := Chunk(words, 2)
fmt.Println(wordChunks)  // [[a b] [c d] [e]]
```

</details>

</details>

---

### Interview Questions (Generics)

<details>
<summary><strong>View contents</strong></summary>

### Q1: What problem do generics solve in Go?

**Answer:**

Generics solve three main problems:

1. **Code Duplication**
   ```go
   // Before: Separate functions for each type
   func MaxInt(a, b int) int
   func MaxFloat(a, b float64) float64
   
   // After: One generic function
   func Max[T constraints.Ordered](a, b T) T
   ```

2. **Type Safety**
   ```go
   // Before: interface{} loses type info
   func First(slice []interface{}) interface{}
   
   // After: Type-safe
   func First[T any](slice []T) T
   ```

3. **Performance**
   ```go
   // interface{} has allocation overhead
   // Generics compile to specialized code (monomorphization)
   ```

---

### Q2: What's the difference between `any` and `comparable` constraints?

**Answer:**

```go
// any: Accepts ANY type
func Print[T any](v T) {
    fmt.Println(v)  // Only fmt.Println works
}

// comparable: Only types that support == and !=
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {  // Requires comparable
            return true
        }
    }
    return false
}

// Works:
Contains([]int{1, 2, 3}, 2)        // ✅ int is comparable
Contains([]string{"a"}, "a")       // ✅ string is comparable

// Doesn't work:
Contains([][]int{{1}}, []int{1})   // ❌ Slices not comparable
```

---

### Q3: What does the `~` (tilde) mean in type constraints?

**Answer:**

**`~` includes the underlying type:**

```go
// Without ~
type Integer interface {
    int | int64
}

// With ~
type Integer interface {
    ~int | ~int64
}

type MyInt int

func Double[T Integer](n T) T {
    return n * 2
}

var x MyInt = 5

// Without ~: Compile error (MyInt not in constraint)
// With ~: Works! (MyInt's underlying type is int)
fmt.Println(Double(x))
```

**Use cases:**
- Supporting custom types (like `type UserID int`)
- Making constraints more flexible
- Working with type aliases

---

### Q4: Can generic types have methods?

**Answer:**

**Yes! Generic types can have methods:**

```go
type Stack[T any] struct {
    items []T
}

// Method on generic type
func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage
stack := Stack[int]{}
stack.Push(1)
stack.Push(2)
val, ok := stack.Pop()  // 2 true
```

**Important:** Methods themselves cannot have additional type parameters:

```go
// ❌ Not allowed
func (s *Stack[T]) Transform[U any](fn func(T) U) Stack[U]

// ✅ Type parameters must be on the type itself
type Stack[T any] struct { /* ... */ }
```

---

### Q5: When should you NOT use generics?

**Answer:**

**Avoid generics when:**

1. **Single concrete type**
   ```go
   // ❌ Unnecessary
   func ProcessInts[T int](nums []T) { /* ... */ }
   
   // ✅ Just use int
   func ProcessInts(nums []int) { /* ... */ }
   ```

2. **Interface is sufficient**
   ```go
   // ✅ Better with interface
   func Print(v interface{ String() string }) {
       fmt.Println(v.String())
   }
   
   // Not: func Print[T Stringer](v T)
   ```

3. **Performance doesn't matter**
   ```go
   // ✅ Simpler
   func Log(v interface{}) {
       fmt.Println(v)
   }
   ```

4. **Logic differs per type**
   ```go
   // ❌ Don't force generics
   func Process[T any](v T) {
       switch any(v).(type) {
       case int: /* ... */
       case string: /* ... */
       }
   }
   
   // ✅ Separate functions
   func ProcessInt(v int) { /* ... */ }
   func ProcessString(v string) { /* ... */ }
   ```

---

</details>


## Concurrency

**Concurrency** means **dealing with multiple tasks at the same time**.

**Key Distinction:**

```
Concurrency: Dealing with many things at once (structure). Structuring a program to handle many things at once
Parallelism: Doing many things at once (execution). Actually running things at the same time on multiple CPU cores
```

A concurrent program **may or may not** run in parallel.

Go makes concurrency **easy to write and easy to reason about**.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/concurrency)**

**Example:**

```
Concurrency (one chef, multiple dishes):
- Chef starts soup
- While soup simmers, chef prepares salad
- While salad chills, chef checks soup
- Interleaved execution

Parallelism (multiple chefs):
- Chef 1 makes soup
- Chef 2 makes salad
- Chef 3 makes dessert
- Simultaneous execution
```

**Go enables both:**
- Goroutines provide **concurrency**
- Runtime scheduler provides **parallelism** (on multi-core systems)

---

### Key Concepts

**1. Process**

* A **process** is an independent program in execution.
* It has its own memory, variables, and resources.
* Example: When you open a browser or run `go run main.go`, each runs as a separate process.
* Processes do **not share memory** by default. Communication between them requires mechanisms like IPC (Inter-Process Communication).

**2. Thread**

* A **thread** is a smaller unit of execution within a process.
* Threads in the same process **share memory**, which allows faster communication but requires careful synchronization to avoid conflicts.
* Each process can have multiple threads. For example, a browser may use threads to render a page, handle network requests, and manage UI simultaneously.

**3. Goroutine**

* A **goroutine** is Go’s lightweight abstraction over threads.
* They are **cheaper** than threads in terms of memory and startup cost.
* Created using the `go` keyword.
* Goroutines are managed by the Go runtime scheduler, which multiplexes many goroutines over a small number of OS threads.
* Unlike threads, you don’t manage goroutine lifecycle manually.

**4. Parallelism vs Concurrency**

* **Concurrency**: Multiple tasks making progress at the same time (interleaved execution).

  * Example: Handling multiple network requests on a single CPU using goroutines.
* **Parallelism**: Multiple tasks running at the exact same time (requires multiple CPUs or cores).

  * Example: Rendering multiple images simultaneously using multiple CPU cores.

> Go supports both concurrency and parallelism via goroutines and GOMAXPROCS setting.

**5. Multithreading**

* Traditional multithreading means creating multiple threads manually.
* Go abstracts this complexity using goroutines and a **scheduler**, so you rarely manage threads directly.
* OS threads are heavier than goroutines, so Go can run **thousands of goroutines** easily.

---

### Goroutines

A **goroutine** is a lightweight unit of execution managed by the Go runtime.

**Lightweight threads managed by Go runtime:**

```go
func main() {
	go sayHello()  // Launches in new goroutine
	
	fmt.Println("Main function")
	time.Sleep(time.Second)  // Wait for goroutine
}

func sayHello() {
	fmt.Println("Hello from goroutine")
}

// Output (order may vary):
// Main function
// Hello from goroutine
```

**Key Characteristics:**
- **Cheap**: ~2KB stack size (vs ~2MB for OS threads)
- **Fast**: Creation takes microseconds
- **Scalable**: Can run millions of goroutines

**Anonymous goroutines:**

```go
go func() {
	fmt.Println("Anonymous goroutine")
}()

// With parameters
name := "Alice"
go func(n string) {
	fmt.Println("Hello,", n)
}(name)  // Pass name by value
```

**⚠️ Common Mistake (closure over loop variable):**

```go
// Wrong: All goroutines see final value
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)  // May print 5, 5, 5, 5, 5
	}()
}

// Correct: Pass as parameter
for i := 0; i < 5; i++ {
	go func(n int) {
		fmt.Println(n)  // Prints 0, 1, 2, 3, 4
	}(i)
}
```

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

**Wait for goroutines to complete:**

```go
import "sync"

func main() {
	var wg sync.WaitGroup
	
	for i := 0; i < 5; i++ {
		wg.Add(1)  // Increment counter
		
		go func(n int) {
			defer wg.Done()  // Decrement when done
			fmt.Println("Worker", n)
		}(i)
	}
	
	wg.Wait()  // Block until counter is 0
	fmt.Println("All workers done")
}
```

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

**How It Works:**

* `Add(n)` → number of goroutines
* `Done()` → signals completion
* `Wait()` → blocks until all are done

Use `WaitGroup` when:

* You only need to wait
* No data needs to be passed back

**Real-world example (concurrent API calls):**

```go
func fetchURLs(urls []string) {
	var wg sync.WaitGroup
	
	for _, url := range urls {
		wg.Add(1)
		
		go func(u string) {
			defer wg.Done()
			
			resp, err := http.Get(u)
			if err != nil {
				log.Printf("Error fetching %s: %v", u, err)
				return
			}
			defer resp.Body.Close()
			
			fmt.Printf("%s: %s\n", u, resp.Status)
		}(url)
	}
	
	wg.Wait()
}
```

---

### Channels – Communication Between Goroutines

* **Channels** are Go’s way of communicating safely between goroutines.
* Think of a channel as a **pipe**: one goroutine sends data in, another receives it.

**Communication between goroutines:**

```go
// Create channel
ch := make(chan int)

// Send value (blocks until received)
ch <- 42

// Receive value (blocks until sent)
value := <-ch
```

Key behavior:

* Sending blocks until received
* Receiving blocks until sent
* Prevents data races

**Unbuffered channels (synchronous):**

* Communication blocks until both sender and receiver are ready.

```go
func main() {
	ch := make(chan string)
	
	go func() {
		ch <- "Hello"  // Blocks until received
	}()
	
	msg := <-ch  // Blocks until sent
	fmt.Println(msg)
}
```

* Sender waits for receiver
* Strong synchronization

**Buffered channels:**

* Allows sending `n` values without waiting for receiver.

```go
// Buffer of 3
ch := make(chan int, 3)

ch <- 1  // Doesn't block
ch <- 2  // Doesn't block
ch <- 3  // Doesn't block
// ch <- 4  // Blocks (buffer full)

fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
```

* Holds up to 3 values
* Sender blocks only when buffer is full
* Useful for worker pools and queues

**Closing channels:**

```go
ch := make(chan int, 3)

ch <- 1
ch <- 2
ch <- 3
close(ch)  // No more values will be sent

// Receive until closed
for val := range ch {
	fmt.Println(val)
}

// Check if closed
val, ok := <-ch
if !ok {
	fmt.Println("Channel closed")
}
```

**⚠️ Important Rules:**
- **Only sender closes** channels
- Receiving from closed channel returns zero value
- Sending to closed channel **panics**
- Closing nil channel **panics**

---

### Channel Patterns

#### Worker Pool

```go
func workerPool(jobs <-chan int, results chan<- int, numWorkers int) {
	var wg sync.WaitGroup
	
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		
		go func(id int) {
			defer wg.Done()
			
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(time.Second)
				results <- job * 2
			}
		}(i)
	}
	
	wg.Wait()
	close(results)
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start worker pool
	go workerPool(jobs, results, 3)
	
	// Send jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
	
	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
```

#### Fan-Out, Fan-In

**Fan-Out (distribute work):**

```go
func fanOut(input <-chan int, workers int) []<-chan int {
	outputs := make([]<-chan int, workers)
	
	for i := 0; i < workers; i++ {
		out := make(chan int)
		outputs[i] = out
		
		go func() {
			defer close(out)
			for val := range input {
				out <- val * 2
			}
		}()
	}
	
	return outputs
}
```

**Fan-In (merge results):**

```go
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	
	for _, ch := range channels {
		wg.Add(1)
		
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}
	
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}
```

#### Pipeline

```go
// Stage 1: Generate numbers
func generator(nums ...int) <-chan int {
	out := make(chan int)
	
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	
	return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)
	
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	
	return out
}

// Stage 3: Sum numbers
func sum(in <-chan int) int {
	total := 0
	for n := range in {
		total += n
	}
	return total
}

// Usage
nums := generator(1, 2, 3, 4, 5)
squared := square(nums)
result := sum(squared)
fmt.Println(result)  // 55
```

---

### Select Statement

* `select` allows goroutines to wait on multiple channels simultaneously.

**Multiplex channel operations:**

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "Message from ch1" }()
go func() { ch2 <- "Message from ch2" }()

select {
case msg := <-ch1:
	fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
	fmt.Println("Received from ch2:", msg)
case ch3 <- "hello":
	fmt.Println("Sent to ch3")
default:
	fmt.Println("No channel ready")
}
```

**Timeout pattern:**

```go
import "time"

select {
case result := <-ch:
	fmt.Println("Received:", result)
case <-time.After(time.Second):
	fmt.Println("Timeout!")
}
```

**Non-blocking receive:**

```go
select {
case msg := <-ch:
	fmt.Println("Received:", msg)
default:
	fmt.Println("No message available")
}
```

**Real-world example (context cancellation):**

```go
import "context"

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker cancelled")
			return
		default:
			// Do work
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Working...")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	go worker(ctx)
	
	time.Sleep(3 * time.Second)
}
```

---

### Mutex (Safe Access to Shared Memory)

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

**Protect shared state:**

```go
import "sync"

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	counter := &SafeCounter{}
	
	for i := 0; i < 1000; i++ {
		go counter.Increment()
	}
	
	time.Sleep(time.Second)
	fmt.Println(counter.Value())  // 1000
}
```

**RWMutex (read/write lock):**

```go
type Cache struct {
	mu    sync.RWMutex
	data  map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()  // Multiple readers allowed
	defer c.mu.RUnlock()
	
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()  // Exclusive write access
	defer c.mu.Unlock()
	
	c.data[key] = value
}
```

**Google Style: Prefer channels over mutexes**

```
Use channels when:
- Passing ownership of data
- Distributing work
- Communicating async results

Use mutexes when:
- Protecting shared state
- Performance critical (mutexes are faster)
- Simple in-memory cache
```

---

### Race Detection

**Find data races at runtime:**

```bash
go test -race ./...
go run -race main.go
go build -race
```

**Example race condition:**

```go
// RACE CONDITION!
var counter int

func increment() {
	counter++  // Read-modify-write (not atomic)
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	
	time.Sleep(time.Second)
	fmt.Println(counter)  // Unpredictable result
}
```

**Fix with mutex:**

```go
var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}
```

**Fix with atomic:**

```go
import "sync/atomic"

var counter int64

func increment() {
	atomic.AddInt64(&counter, 1)
}
```

---

### Context (Cancellation and Deadlines)

**Manage goroutine lifecycles:**

```go
import "context"

// Create contexts
ctx := context.Background()  // Root context
ctx, cancel := context.WithCancel(ctx)  // Cancellable
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)  // Timeout
ctx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))  // Deadline
ctx = context.WithValue(ctx, "key", "value")  // Carry values
```

**HTTP request example:**

```go
func fetchData(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	return io.ReadAll(resp.Body)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	data, err := fetchData(ctx, "https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(data))
}
```

**Worker with context:**

```go
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	for i := 0; i < 3; i++ {
		go worker(ctx, i)
	}
	
	time.Sleep(2 * time.Second)
	cancel()  // Cancel all workers
	
	time.Sleep(time.Second)
}
```

---

**Small Project to Practice Concurrency: [Link](https://github.com/foyez/go/tree/main/codes/concurrency)**

</details>

---

### Practice Questions

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. A __________ is a lightweight thread managed by the Go runtime.
2. Channels can be buffered or __________, affecting their blocking behavior.
3. The `sync.__________` type is used to wait for goroutines to complete.
4. The `__________` statement allows waiting on multiple channel operations.

**True/False:**

1. ⬜ Goroutines are more expensive than OS threads
2. ⬜ Only the sender should close a channel
3. ⬜ Mutexes are always better than channels for concurrency
4. ⬜ The -race flag can detect data races at compile time

**Multiple Choice:**

1. What happens if you send to a closed channel?
   - A) Blocks forever
   - B) Returns zero value
   - C) Panic
   - D) Nothing

2. When should you use context?
   - A) Never
   - B) For all goroutines
   - C) For goroutines that need cancellation or deadlines
   - D) Only in HTTP handlers

**Code Challenge:**

Create a worker pool that processes jobs concurrently with a configurable number of workers, and gracefully shuts down on context cancellation.

---

### Answers

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. goroutine
2. unbuffered
3. WaitGroup
4. select

**True/False:**
1. ❌ False (goroutines are much cheaper)
2. ✅ True
3. ❌ False (prefer channels; use mutexes for shared state)
4. ❌ False (detects at runtime with -race flag)

**Multiple Choice:**
1. **C** - Panic
2. **C** - For goroutines that need cancellation or deadlines

**Code Challenge:**
```go
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			results <- job * 2
		}
	}
}

func workerPool(ctx context.Context, numWorkers int) {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, id, jobs, results)
		}(i)
	}
	
	// Send jobs
	go func() {
		for i := 1; i <= 20; i++ {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			case jobs <- i:
			}
		}
		close(jobs)
	}()
	
	// Close results after workers done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	workerPool(ctx, 3)
}
```

---

</details>

</details>

## Testing

<details>
<summary><strong>View contents</strong></summary>

### Test Basics

**Test files end with `_test.go`:**

```go
// math.go
package math

func Add(a, b int) int {
	return a + b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
```

**Run tests:**

```bash
go test                  # Current package
go test ./...            # All packages
go test -v               # Verbose
go test -run TestAdd     # Run specific test
go test -cover           # Coverage report
go test -race            # Race detection
```

---

### Table-Driven Tests

**Idiomatic Go testing pattern:**

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
		{"with zero", 5, 0, 5},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", 
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
```

**Benefits:**
- Easy to add new cases
- Self-documenting
- Subtests for better output

---

### Testing Generic Functions (Go 1.18+)

**Test generic functions with multiple types:**

```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func TestMax(t *testing.T) {
    // Test with int
    t.Run("int", func(t *testing.T) {
        result := Max(10, 20)
        if result != 20 {
            t.Errorf("Max(10, 20) = %d; want 20", result)
        }
    })
    
    // Test with float64
    t.Run("float64", func(t *testing.T) {
        result := Max(3.14, 2.71)
        if result != 3.14 {
            t.Errorf("Max(3.14, 2.71) = %f; want 3.14", result)
        }
    })
    
    // Test with string
    t.Run("string", func(t *testing.T) {
        result := Max("apple", "banana")
        if result != "banana" {
            t.Errorf("Max(apple, banana) = %s; want banana", result)
        }
    })
}
```

**Table-driven test for generics:**

```go
func TestMaxTableDriven(t *testing.T) {
    t.Run("int", func(t *testing.T) {
        tests := []struct {
            name     string
            a, b     int
            expected int
        }{
            {"positive", 10, 20, 20},
            {"negative", -5, -10, -5},
            {"equal", 5, 5, 5},
        }
        
        for _, tt := range tests {
            t.Run(tt.name, func(t *testing.T) {
                result := Max(tt.a, tt.b)
                if result != tt.expected {
                    t.Errorf("got %d; want %d", result, tt.expected)
                }
            })
        }
    })
    
    // Similar tests for float64, string, etc.
}
```

---

### Test Helpers

**Use `t.Helper()` to mark helper functions:**

```go
func assertEqual(t *testing.T, got, want int) {
	t.Helper()  // Error line points to caller, not this line
	
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assertEqual(t, result, 5)  // Error points here if fails
}
```

---

### Mocking and Interfaces

**Test through interfaces:**

```go
// production code
type UserStore interface {
	GetUser(id int) (*User, error)
}

type UserService struct {
	store UserStore
}

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.store.GetUser(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

// test code
type MockUserStore struct {
	users map[int]*User
}

func (m *MockUserStore) GetUser(id int) (*User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func TestGetUserName(t *testing.T) {
	mock := &MockUserStore{
		users: map[int]*User{
			1: {Name: "Alice"},
		},
	}
	
	service := &UserService{store: mock}
	
	name, err := service.GetUserName(1)
	if err != nil {
		t.Fatal(err)
	}
	
	if name != "Alice" {
		t.Errorf("got %q; want %q", name, "Alice")
	}
}
```

---

### Benchmarks

**Measure performance:**

```go
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

func BenchmarkAddParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Add(2, 3)
		}
	})
}
```

**Run benchmarks:**

```bash
go test -bench .              # All benchmarks
go test -bench BenchmarkAdd   # Specific benchmark
go test -bench . -benchmem    # Memory allocation stats
go test -bench . -cpuprofile=cpu.prof  # CPU profiling
```

**Output:**
```
BenchmarkAdd-8          1000000000               0.25 ns/op
BenchmarkAddParallel-8  5000000000               0.24 ns/op
```

---

### Example Tests (Documentation)

**Tests that appear in godoc:**

```go
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

func ExampleAdd_negative() {
	result := Add(-2, -3)
	fmt.Println(result)
	// Output: -5
}
```

**Requirements:**
- Function name: `Example<FunctionName>`
- `// Output:` comment required
- Output must match exactly (including whitespace)

---

### Test Coverage

**Generate coverage report:**

```bash
go test -cover                          # Basic coverage
go test -coverprofile=coverage.out      # Save report
go tool cover -html=coverage.out        # View in browser
go test -coverprofile=coverage.out -covermode=atomic  # Atomic mode
```

**Coverage directives:**

```go
// TestMain runs before/after all tests
func TestMain(m *testing.M) {
	// Setup
	code := m.Run()
	// Teardown
	os.Exit(code)
}
```

---

### Testing Best Practices (Google Style)

**1. Test behavior, not implementation**

```go
// Good: Tests public API
func TestUserService_CreateUser(t *testing.T) {
	service := NewUserService()
	user, err := service.CreateUser("Alice", "alice@example.com")
	
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "Alice" {
		t.Errorf("got name %q; want %q", user.Name, "Alice")
	}
}

// Bad: Tests internal implementation
func TestUserService_createUserInternal(t *testing.T) {
	// ...
}
```

**2. Use table-driven tests**

**3. Keep tests independent**

```go
// Bad: Tests depend on order
func TestCreateUser(t *testing.T) {
	// Creates user with ID 1
}

func TestGetUser(t *testing.T) {
	// Assumes user 1 exists from previous test
}

// Good: Each test sets up its own state
func TestGetUser(t *testing.T) {
	user := createTestUser(t)
	result := getUser(user.ID)
	// ...
}
```

**4. Fail fast with t.Fatal**

```go
func TestProcess(t *testing.T) {
	data, err := loadData()
	if err != nil {
		t.Fatal(err)  // Stop test if data can't be loaded
	}
	
	// Continue with tests that need data
}
```

---

### TDD with Go (Test-Driven Development)

Test-Driven Development (TDD) is a development approach where you **write tests before writing production code**.
The goal is to let tests **drive the design and behavior** of your code.

<details>
<summary><strong>View contents</strong></summary>

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

</details>

</details>

---

### Practice Questions (Testing)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. Test files in Go must end with `__________`.
2. The `t.__________()` method marks a function as a test helper.
3. Benchmark functions must start with the word `__________`.
4. The `// __________:` comment is required for Example tests.

**True/False:**

1. ⬜ Tests must be in the same package as the code they test
2. ⬜ Table-driven tests are the idiomatic way to test in Go
3. ⬜ Benchmarks measure the performance of code
4. ⬜ Test files are compiled into the final binary

**Multiple Choice:**

1. What does `t.Run()` do?
   - A) Runs all tests
   - B) Creates a subtest
   - C) Runs benchmarks
   - D) Generates coverage

2. When should you use `t.Fatal()` instead of `t.Error()`?
   - A) Always
   - B) Never
   - C) When continuing the test makes no sense
   - D) For benchmarks

---

### Answers

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. _test.go
2. Helper
3. Benchmark
4. Output

**True/False:**
1. ❌ False (can be in `<package>_test` for black-box testing)
2. ✅ True
3. ✅ True
4. ❌ False (excluded from production builds)

**Multiple Choice:**
1. **B** - Creates a subtest
2. **C** - When continuing the test makes no sense

---

</details>

</details>

## Web Servers

Go has a **powerful, production-ready HTTP server built into the standard library**.
You do **not** need external frameworks to build fast and scalable web servers.

<details>
<summary><strong>View contents</strong></summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/webServers)**

### HTTP Server Basics

**Simple HTTP server:**

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}
```

**Key Components:**
- `http.ResponseWriter`: Write response
- `*http.Request`: Read request
- `http.HandleFunc`: Register route
- `http.ListenAndServe`: Start server

---

### Request Handling

**Reading request data:**

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// Method
	fmt.Println(r.Method)  // GET, POST, etc.
	
	// URL and path
	fmt.Println(r.URL.Path)
	
	// Query parameters
	name := r.URL.Query().Get("name")
	
	// Headers
	userAgent := r.Header.Get("User-Agent")
	
	// Body (for POST/PUT)
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	
	// Form data
	r.ParseForm()
	email := r.FormValue("email")
}
```

**Setting response:**

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// Set status code
	w.WriteHeader(http.StatusOK)  // 200
	
	// Set headers
	w.Header().Set("Content-Type", "application/json")
	
	// Write body
	fmt.Fprintf(w, `{"message": "success"}`)
}
```

---

### Routing

**ServeMux (built-in router):**

```go
func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users/", usersHandler)  // Trailing slash matches /users/*
	mux.HandleFunc("/api/v1/products", productsHandler)
	
	http.ListenAndServe(":8080", mux)
}
```

**Custom handler (implements http.Handler):**

```go
type APIHandler struct {
	db *sql.DB
}

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handle request
}

func main() {
	handler := &APIHandler{db: db}
	http.Handle("/api/", handler)
	http.ListenAndServe(":8080", nil)
}
```

---

### Middleware

**Function signature:**

```go
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before request
		next.ServeHTTP(w, r)
		// After request
	})
}
```

**Logging middleware:**

```go
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		next.ServeHTTP(w, r)
		
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	
	// Wrap with middleware
	handler := loggingMiddleware(mux)
	
	http.ListenAndServe(":8080", handler)
}
```

**Authentication middleware:**

```go
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
```

**Chaining middleware:**

```go
func chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", apiHandler)
	
	handler := chain(mux,
		loggingMiddleware,
		authMiddleware,
		corsMiddleware,
	)
	
	http.ListenAndServe(":8080", handler)
}
```

---

### JSON Handling

**Encoding (struct to JSON):**

```go
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
```

**Decoding (JSON to struct):**

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Process user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
```

---

### REST API Example

```go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Phone", Price: 599.99},
}

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /api/products", listProducts)
	mux.HandleFunc("GET /api/products/{id}", getProduct)
	mux.HandleFunc("POST /api/products", createProduct)
	mux.HandleFunc("PUT /api/products/{id}", updateProduct)
	mux.HandleFunc("DELETE /api/products/{id}", deleteProduct)
	
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")  // Go 1.22+
	
	// Find product
	for _, p := range products {
		if strconv.Itoa(p.ID) == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	
	http.Error(w, "Product not found", http.StatusNotFound)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	product.ID = len(products) + 1
	products = append(products, product)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
```

---

### File Serving

**Static files:**

```go
func main() {
	// Serve files from ./static directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", homeHandler)
	
	http.ListenAndServe(":8080", nil)
}
```

**File upload:**

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)  // 10 MB max
	
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	// Save file
	dst, err := os.Create(filepath.Join("./uploads", header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	
	io.Copy(dst, file)
	
	fmt.Fprintf(w, "File uploaded successfully")
}
```

</details>

---

### Practice Questions (Web Servers)

<details>
<summary><strong>View contents</strong></summary>

**Fill in the Blanks:**

1. The `http.__________` function registers a handler for a route.
2. Middleware functions wrap an `http.__________` to add functionality.
3. The `__________` package is used to encode and decode JSON.
4. The `http.FileServer` function is used to serve __________ files.

**True/False:**

1. ⬜ Every HTTP handler must return a value
2. ⬜ Middleware can modify both requests and responses
3. ⬜ The http package supports HTTP/2 by default
4. ⬜ ServeMux is thread-safe

**Code Challenge:**

Create a simple REST API for a todo list with endpoints to create, list, and delete todos using JSON.

---

### Answers

<details>
<summary><strong>View answers</strong></summary>

**Fill in the Blanks:**
1. HandleFunc
2. Handler
3. json
4. static

**True/False:**
1. ❌ False (handlers write to ResponseWriter)
2. ✅ True
3. ✅ True (since Go 1.6)
4. ✅ True

**Code Challenge:**
```go
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	todos  = []Todo{}
	nextID = 1
	mu     sync.Mutex
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /todos", listTodos)
	mux.HandleFunc("POST /todos", createTodo)
	mux.HandleFunc("DELETE /todos/{id}", deleteTodo)
	
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	mu.Lock()
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	mu.Unlock()
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	
	mu.Lock()
	defer mu.Unlock()
	
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	
	http.Error(w, "Todo not found", http.StatusNotFound)
}
```

---

</details>

</details>

## Working with files

Go provides simple and safe file handling via the `os` and `io` packages.

<details>
<summary><strong>View contents</strong></summary>

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

## Best Practices (Google Style Guide)

<details>
<summary><strong>View contents</strong></summary>

### Naming Conventions

**Packages:**
- **Lowercase, single word**
- No underscores or mixed caps
- Name = directory name

```go
✅ package user
✅ package http
✅ package json

❌ package userService
❌ package user_service
❌ package Users
```

**Variables:**
- **CamelCase** for exported
- **camelCase** for unexported
- Short names for small scopes
- Descriptive names for large scopes

```go
✅ var userCount int
✅ var u User  // Short scope
✅ var db *sql.DB

❌ var UserCount int  // Unexported
❌ var user_count int
```

**Functions:**
- **CamelCase** for exported
- **camelCase** for unexported
- Verb-based names

```go
✅ func GetUser(id int) (*User, error)
✅ func validateEmail(email string) bool

❌ func get_user(id int) (*User, error)
```

**Constants:**
- **CamelCase** (not SCREAMING_CASE)

```go
✅ const MaxRetries = 3
✅ const defaultTimeout = time.Second

❌ const MAX_RETRIES = 3
❌ const DEFAULT_TIMEOUT = time.Second
```

---

### Code Organization

**Package structure:**

```
myapp/
├── cmd/
│   └── myapp/
│       └── main.go      # Application entry point
├── internal/            # Private application code
│   ├── user/
│   │   ├── user.go
│   │   └── user_test.go
│   └── auth/
│       ├── auth.go
│       └── auth_test.go
├── pkg/                 # Public library code
│   └── api/
│       └── api.go
├── go.mod
└── go.sum
```

**File organization:**
- Group related types/functions
- Limit file size (~500 lines)
- One concept per file

---

### Error Handling

**Always check errors:**

```go
// Good
f, err := os.Open("file.txt")
if err != nil {
	return err
}
defer f.Close()

// Bad
f, _ := os.Open("file.txt")
```

**Wrap errors with context:**

```go
// Good
if err := process(); err != nil {
	return fmt.Errorf("processing data: %w", err)
}

// Bad
if err := process(); err != nil {
	return err  // No context
}
```

**Handle errors at appropriate level:**

```go
// Low-level: Return errors
func readFile() ([]byte, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}
	return data, nil
}

// High-level: Decide what to do
func loadConfig() {
	data, err := readFile()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Process data
}
```

---

### Concurrency

**Use goroutines sparingly:**

```go
// Good: Clear benefit
go processInBackground()

// Bad: Unnecessary complexity
go fmt.Println("hello")
```

**Synchronize with channels or WaitGroups:**

```go
// Good
func processConcurrently(items []Item) {
	var wg sync.WaitGroup
	for _, item := range items {
		wg.Add(1)
		go func(it Item) {
			defer wg.Done()
			process(it)
		}(item)
	}
	wg.Wait()
}

// Bad: No synchronization
func processConcurrently(items []Item) {
	for _, item := range items {
		go process(item)
	}
	// Returns immediately!
}
```

**Prefer channels for communication:**

```go
// Good
func generator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

// Use
for val := range generator() {
	fmt.Println(val)
}
```

---

### Function Design

**Short functions:**
- Aim for <50 lines
- One responsibility
- Easy to test

**Clear parameters:**

```go
// Good
func CreateUser(name, email string, age int) (*User, error)

// Bad: Too many parameters
func CreateUser(name, email, address, city, country, phone string, age, zip int) (*User, error)

// Better: Use struct
type UserInput struct {
	Name    string
	Email   string
	Address string
	Age     int
}

func CreateUser(input UserInput) (*User, error)
```

**Return errors last:**

```go
// Good
func Process() (result string, err error)

// Bad
func Process() (err error, result string)
```

---

### Comments

**Package comments:**

```go
// Package user provides user management functionality.
// It handles user creation, authentication, and authorization.
package user
```

**Exported identifiers:**

```go
// User represents a user account.
type User struct {
	// ID is the unique identifier.
	ID int
	
	// Name is the full name.
	Name string
}

// CreateUser creates a new user account.
func CreateUser(name string) (*User, error) {
	// ...
}
```

**Avoid obvious comments:**

```go
// Bad
i := 0  // Set i to 0
i++     // Increment i

// Good (when needed)
// Retry 3 times with exponential backoff
for i := 0; i < 3; i++ {
	// ...
}
```

---

### Testing

**Test file naming:**

```go
user.go       → user_test.go
validator.go  → validator_test.go
```

**Test function naming:**

```go
func TestFunctionName(t *testing.T)
func TestStructName_MethodName(t *testing.T)
func BenchmarkFunctionName(b *testing.B)
func ExampleFunctionName()
```

**Use table-driven tests:**

```go
func TestProcess(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"valid input", "test", "TEST", false},
		{"empty input", "", "", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Process(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Process() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

---

### Performance

**Use pointers for large structs:**

```go
// Good
func ProcessUser(u *User) error

// Bad: Copies entire struct
func ProcessUser(u User) error
```

**Preallocate slices when size is known:**

```go
// Good
items := make([]Item, 0, expectedSize)

// Bad
var items []Item
```

**Use sync.Pool for frequently allocated objects:**

```go
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func process() {
	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()
	
	// Use buf
}
```

---

### Common Pitfalls

**1. Loop variable capture:**

```go
// Wrong
for _, item := range items {
	go func() {
		process(item)  // All goroutines see last value
	}()
}

// Correct
for _, item := range items {
	go func(it Item) {
		process(it)
	}(item)
}
```

**2. Nil map/slice:**

```go
// Wrong
var m map[string]int
m["key"] = 1  // Panic!

// Correct
m := make(map[string]int)
m["key"] = 1
```

**3. Slice append:**

```go
// Wrong
func addItem(items []Item, item Item) {
	items = append(items, item)  // Doesn't modify caller's slice
}

// Correct
func addItem(items []Item, item Item) []Item {
	return append(items, item)
}
```

---

</details>

## Interview Questions

<details>
<summary><strong>View contents</strong></summary>

## Basics & Fundamentals

### Q1: What is Go and why was it created?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

Go (Golang) is a statically typed, compiled programming language designed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, and open-sourced in 2009.

**Design Goals:**
- **Simplicity**: Easy to learn and read (25 keywords)
- **Performance**: Fast compilation and execution
- **Concurrency**: Built-in support via goroutines
- **Productivity**: Strong tooling and fast development

**Why it was created:**
- Address shortcomings of C++ (complexity, slow compilation)
- Provide better concurrency than Java
- Maintain performance while improving developer experience
- Handle large-scale distributed systems at Google

**Real-world impact:**
- Docker, Kubernetes (container orchestration)
- Terraform (infrastructure as code)
- Prometheus (monitoring)
- etcd (distributed key-value store)

</details>

---

### Q2: Explain the difference between `var`, `:=`, and `const`.

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

```go
// var - Explicit declaration
var name string = "Alice"
var age = 30              // Type inference
var count int             // Zero value (0)

// := - Short variable declaration (function scope only)
city := "New York"        // Type inferred

// const - Immutable value (compile-time constant)
const MaxRetries = 3
const Pi = 3.14159
```

**Key Differences:**

| Feature | var | := | const |
|---------|-----|-------|-------|
| **Scope** | Package or function | Function only | Package or function |
| **Type inference** | Yes | Yes | Yes |
| **Zero value** | Yes | No | N/A |
| **Reassignable** | Yes | Yes | No |
| **Compile-time** | No | No | Yes |

**Common Pitfall:**
```go
// Wrong: := cannot be used at package level
package main
name := "Alice"  // Compile error!

// Correct
var name = "Alice"
```

</details>

---

### Q3: What are the basic data types in Go?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Numeric Types:**
```go
// Integers
int, int8, int16, int32, int64       // Signed
uint, uint8, uint16, uint32, uint64  // Unsigned

// Floating-point
float32, float64

// Complex numbers
complex64, complex128

// Aliases
byte  // alias for uint8
rune  // alias for int32 (Unicode code point)
```

**Other Types:**
```go
string  // Immutable sequence of bytes
bool    // true or false
```

**Important Notes:**
- `int` size is **platform-dependent** (32-bit on 32-bit systems, 64-bit on 64-bit)
- **No implicit type conversion** - must be explicit
- **Zero values**: 0 for numbers, "" for strings, false for bool

**Example:**
```go
var i int = 42
var f float64 = float64(i)  // Must convert explicitly
// var f float64 = i         // Compile error!
```

</details>

---

### Q4: Explain how strings work in Go. Are they mutable?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Key Characteristics:**
- **Immutable**: Cannot be changed after creation
- **UTF-8 encoded** byte sequences
- **len()** returns **byte count**, not character count

```go
s := "Hello, 世界"

// Length in BYTES, not characters
fmt.Println(len(s))  // 13 (not 9!)

// Indexing returns bytes
fmt.Println(s[0])    // 72 (byte value of 'H')

// Immutable
s[0] = 'h'  // Compile error!

// Concatenation creates new string
s2 := s + " Go"  // New string allocated
```

**Iterating over strings:**

```go
// Wrong: Iterates over bytes
for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i])  // Breaks on Unicode
}

// Correct: Iterates over runes (Unicode code points)
for _, r := range s {
    fmt.Printf("%c ", r)  // Works with Unicode
}
```

**Why immutable?**
- **Thread safety**: Can be shared across goroutines
- **Performance**: String literals can be reused
- **Simplicity**: No unexpected modifications

</details>

---

### Q5: What is the difference between `nil` slice and empty slice?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

```go
// Nil slice
var s1 []int
fmt.Println(s1 == nil)     // true
fmt.Println(len(s1))       // 0
fmt.Println(cap(s1))       // 0

// Empty slice
s2 := []int{}
fmt.Println(s2 == nil)     // false
fmt.Println(len(s2))       // 0
fmt.Println(cap(s2))       // 0

s3 := make([]int, 0)
fmt.Println(s3 == nil)     // false
```

**Visual Representation:**

```
Nil slice:
┌──────┬─────┬─────┐
│ ptr  │ len │ cap │
│ nil  │  0  │  0  │
└──────┴─────┴─────┘

Empty slice:
┌──────┬─────┬─────┐
│ ptr  │ len │ cap │
│ 0x.. │  0  │  0  │ (points to empty array)
└──────┴─────┴─────┘
```

**When to use which?**

```go
// Nil slice - Default zero value
var users []User  // Represents "no data"

// Empty slice - Explicitly empty collection
users := []User{}  // Represents "initialized but empty"

// JSON marshaling difference
json.Marshal(nil)     // null
json.Marshal([]int{}) // []
```

**Best Practice:**
- Use **nil slice** as zero value (simpler, same behavior)
- Check `len(s) == 0` instead of `s == nil` (works for both)

</details>

---

### Q6: Explain the difference between arrays and slices.

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Arrays - Fixed Size:**

```go
// Size is part of the type
var a1 [3]int       // Different type from [5]int
a2 := [3]int{1, 2, 3}
a3 := [...]int{1, 2, 3}  // Compiler counts

// Arrays are VALUES (copied on assignment)
b := a2
b[0] = 100
fmt.Println(a2)  // [1 2 3] (unchanged)
fmt.Println(b)   // [100 2 3]
```

**Slices - Dynamic Size:**

```go
// Reference to underlying array
s1 := []int{1, 2, 3}
s2 := make([]int, 3, 5)  // len=3, cap=5

// Slices are REFERENCES
s3 := s1
s3[0] = 100
fmt.Println(s1)  // [100 2 3] (modified!)
fmt.Println(s3)  // [100 2 3]
```

**Comparison Table:**

| Feature | Array | Slice |
|---------|-------|-------|
| **Size** | Fixed | Dynamic |
| **Type** | [N]T | []T |
| **Passed to function** | Copied | Reference |
| **Resizable** | No | Yes (via append) |
| **Common use** | Rare | Very common |

**When to use arrays:**
- Fixed-size data (SHA256 hash: `[32]byte`)
- Performance-critical code (avoid allocations)
- **Generally, prefer slices**

</details>

---

### Q7: What is the zero value in Go? Give examples.

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Zero values are automatically assigned to variables without initialization.**

```go
var i int           // 0
var f float64       // 0.0
var s string        // ""
var b bool          // false
var p *int          // nil
var slice []int     // nil
var m map[string]int // nil
var ch chan int     // nil
var fn func()       // nil
var iface interface{} // nil

type Person struct {
    Name string  // ""
    Age  int     // 0
}
var person Person  // {Name: "", Age: 0}
```

**Why zero values matter:**

```go
// Safe defaults
type Counter struct {
    count int  // Starts at 0 (useful!)
}

c := Counter{}
fmt.Println(c.count)  // 0 (no initialization needed)

// Nil checks
var m map[string]int
if m == nil {
    m = make(map[string]int)
}
```

**Common Pitfall:**

```go
var m map[string]int
m["key"] = 1  // PANIC! Cannot write to nil map

// Must initialize
m = make(map[string]int)
m["key"] = 1  // OK
```

</details>

---

### Q8: Explain `defer`. What happens with multiple defer statements?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**defer delays function execution until surrounding function returns.**

```go
func readFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()  // Executes when readFile returns
    
    // Work with file
    // Even if error occurs, file.Close() is called
    return nil
}
```

**Multiple defers execute in LIFO (Last In, First Out):**

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

**Arguments are evaluated immediately:**

```go
func example() {
    x := 10
    defer fmt.Println(x)  // x evaluated NOW (10)
    x = 20
    fmt.Println(x)
}

// Output:
// 20
// 10
```

**Common Use Cases:**

```go
// 1. Resource cleanup
defer file.Close()
defer db.Close()

// 2. Unlock mutex
mu.Lock()
defer mu.Unlock()

// 3. Recover from panic
defer func() {
    if r := recover(); r != nil {
        log.Println("Recovered:", r)
    }
}()

// 4. Timing function execution
start := time.Now()
defer func() {
    fmt.Println("Duration:", time.Since(start))
}()
```

**Real-world Example:**

```go
func processTransaction(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()  // Rollback if not committed
    
    // Perform operations
    if err := operation1(tx); err != nil {
        return err  // Auto-rollback
    }
    
    if err := operation2(tx); err != nil {
        return err  // Auto-rollback
    }
    
    return tx.Commit()  // Commit overrides rollback
}
```

</details>

---

### Q9: What are init functions and when are they called?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**`init()` functions run automatically before `main()`.**

```go
package main

import "fmt"

func init() {
    fmt.Println("Init 1")
}

func init() {
    fmt.Println("Init 2")
}

func main() {
    fmt.Println("Main")
}

// Output:
// Init 1
// Init 2
// Main
```

**Execution Order:**

1. **Package-level variables** initialized
2. **All `init()` functions** in the order they appear
3. **Dependencies first** (imported packages before current)
4. **`main()` function** (if it exists)

**Multi-package example:**

```go
// package config
var DBConnection string

func init() {
    DBConnection = "postgres://..."
    fmt.Println("Config initialized")
}

// package main
import (
    "config"
)

func init() {
    fmt.Println("Main package initialized")
}

func main() {
    fmt.Println("Main function")
    fmt.Println(config.DBConnection)
}

// Output:
// Config initialized
// Main package initialized
// Main function
// postgres://...
```

**Common Use Cases:**

```go
// 1. Package initialization
func init() {
    rand.Seed(time.Now().UnixNano())
}

// 2. Register drivers (database, image formats)
import _ "github.com/lib/pq"  // Runs init() only

func init() {
    sql.Register("postgres", &Driver{})
}

// 3. Validation
var config Config

func init() {
    if err := loadConfig(&config); err != nil {
        panic("Failed to load config: " + err.Error())
    }
}
```

**Best Practices:**
- ✅ Keep `init()` simple and deterministic
- ✅ Avoid side effects if possible
- ❌ Avoid complex logic (hard to test)
- ❌ Don't use `init()` for dependency injection

</details>

---

### Q10: What is the difference between `new()` and `make()`?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Both allocate memory, but serve different purposes.**

**`new(T)` - Returns `*T` (pointer to zeroed memory):**

```go
// new allocates zeroed storage
p := new(int)
fmt.Println(*p)  // 0

type Person struct {
    Name string
    Age  int
}

p2 := new(Person)
fmt.Printf("%+v\n", p2)  // &{Name: Age:0}
```

**`make(T)` - Returns initialized `T` (slices, maps, channels only):**

```go
// make creates and initializes
s := make([]int, 5, 10)    // len=5, cap=10
m := make(map[string]int)  // Empty map
ch := make(chan int, 5)    // Buffered channel
```

**Comparison:**

| Feature | new(T) | make(T) |
|---------|--------|---------|
| **Returns** | `*T` (pointer) | `T` (value) |
| **Works with** | Any type | Slice, map, channel only |
| **Initialization** | Zero value | Ready to use |
| **Use case** | Rare | Common |

**Examples:**

```go
// Slice
var s1 *[]int = new([]int)  // *s1 == nil (useless!)
var s2 []int = make([]int, 0) // Ready to use

// Map
var m1 *map[string]int = new(map[string]int)  // *m1 == nil
var m2 map[string]int = make(map[string]int)  // Ready to use

// Struct
type Counter struct {
    count int
}

c1 := new(Counter)    // Returns *Counter
c2 := &Counter{}      // Equivalent (preferred)
```

**Best Practice:**

```go
// Prefer composite literals over new()
// new(T)
p := new(Person)
p.Name = "Alice"

// Composite literal (better)
p := &Person{
    Name: "Alice",
}
```

</details>

---

## Data Structures

### Q11: How do maps work internally in Go?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Maps are hash tables implemented using buckets.**

**Internal Structure:**

```
Hash Table
┌─────────────┐
│   Buckets   │
├─────────────┤
│ [0] → [k,v] │
│ [1] → [k,v] │
│ [2] → [k,v] │
│     ...     │
│ [7] → [k,v] │
└─────────────┘
     ↓
Overflow buckets (if needed)
```

**Key Characteristics:**

1. **Hash Function**: Converts key to bucket index
2. **Buckets**: Each bucket holds ~8 key-value pairs
3. **Load Factor**: When ~6.5 entries per bucket, map grows (2x)
4. **Collision Handling**: Chaining via overflow buckets

**Example:**

```go
m := make(map[string]int)
m["Alice"] = 25

// Internally:
// 1. hash("Alice") → bucket index
// 2. Store in bucket[index]
// 3. If bucket full, create overflow bucket
```

**Performance:**

| Operation | Average | Worst Case |
|-----------|---------|------------|
| **Get** | O(1) | O(n) |
| **Set** | O(1) | O(n) |
| **Delete** | O(1) | O(n) |

**Important Notes:**

```go
// Maps are NOT thread-safe
var m = make(map[string]int)

// Race condition!
go func() { m["key"] = 1 }()
go func() { m["key"] = 2 }()

// Use sync.Map or mutex
var mu sync.Mutex
mu.Lock()
m["key"] = 1
mu.Unlock()

// Map iteration order is RANDOM
for k, v := range m {
    fmt.Println(k, v)  // Order not guaranteed!
}

// Maps cannot be compared
m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}
// m1 == m2  // Compile error!
```

</details>

---

### Q12: Explain slice capacity and how `append` works.

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Slice Structure:**

```go
type slice struct {
    ptr *array  // Pointer to underlying array
    len int     // Number of elements
    cap int     // Capacity of underlying array
}
```

**Example:**

```go
s := make([]int, 3, 5)
// ptr → [0, 0, 0, _, _]
// len = 3
// cap = 5

fmt.Println(len(s))  // 3
fmt.Println(cap(s))  // 5
```

**How `append` Works:**

```go
s := []int{1, 2, 3}
fmt.Printf("len=%d cap=%d\n", len(s), cap(s))  // len=3 cap=3

s = append(s, 4)
fmt.Printf("len=%d cap=%d\n", len(s), cap(s))  // len=4 cap=6

// What happened?
// 1. Old capacity (3) exceeded
// 2. New array allocated (capacity doubled to 6)
// 3. Old elements copied to new array
// 4. New element added
```

**Capacity Growth Strategy:**

```go
// Capacity doubles when small (<1024), then grows by 25%
s := []int{}
for i := 0; i < 10; i++ {
    s = append(s, i)
    fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}

// Output:
// len=1 cap=1
// len=2 cap=2
// len=3 cap=4
// len=4 cap=4
// len=5 cap=8
// len=6 cap=8
// len=7 cap=8
// len=8 cap=8
// len=9 cap=16
// len=10 cap=16
```

**Appending Multiple Values:**

```go
s1 := []int{1, 2, 3}
s2 := []int{4, 5, 6}

// Append single values
s1 = append(s1, 4, 5, 6)

// Append slice (must unpack)
s1 = append(s1, s2...)  // ... unpacks slice
```

**Common Pitfall (Shared Underlying Array):**

```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[0:3]  // [1, 2, 3]
slice2 := original[2:5]  // [3, 4, 5]

slice1[2] = 100

fmt.Println(original)  // [1 2 100 4 5]
fmt.Println(slice1)    // [1 2 100]
fmt.Println(slice2)    // [100 4 5] - Affected!
```

**Solution - Limit Capacity:**

```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[0:3:3]  // len=3, cap=3 (limits capacity)

slice1 = append(slice1, 100)  // New array allocated
fmt.Println(original)  // [1 2 3 4 5] - Unchanged!
```

</details>

---

### Q13: What is the difference between passing by value and by reference in Go?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Go always passes by VALUE, but some types contain references.**

**Value Types (Copied):**

```go
// Basic types, arrays, structs
func modifyInt(x int) {
    x = 100  // Modifies copy
}

func modifyArray(arr [3]int) {
    arr[0] = 100  // Modifies copy
}

func modifyStruct(p Person) {
    p.Name = "Bob"  // Modifies copy
}

x := 10
modifyInt(x)
fmt.Println(x)  // 10 (unchanged)
```

**Reference Types (Contain Pointers):**

```go
// Slices, maps, channels
func modifySlice(s []int) {
    s[0] = 100  // Modifies original (via pointer)
}

func modifyMap(m map[string]int) {
    m["key"] = 100  // Modifies original
}

s := []int{1, 2, 3}
modifySlice(s)
fmt.Println(s)  // [100 2 3] (modified!)
```

**Using Pointers for True Reference Passing:**

```go
func modifyPerson(p *Person) {
    p.Name = "Bob"  // Modifies original
}

person := Person{Name: "Alice"}
modifyPerson(&person)
fmt.Println(person.Name)  // Bob
```

**Visual Representation:**

```
Value Types:
┌─────────┐        ┌─────────┐
│ Caller  │ Copy   │Function │
│  x=10   │───────>│  x=100  │
└─────────┘        └─────────┘
   (unchanged)

Slice (Reference Type):
┌─────────┐        ┌─────────┐
│ Caller  │ Copy   │Function │
│ slice ──┼───────>│ slice ──┼─> Same array
└─────────┘        └─────────┘
   (modified!)

Pointer:
┌─────────┐        ┌─────────┐
│ Caller  │ Copy   │Function │
│  p ─────┼───────>│  p ─────┼─> Same struct
└─────────┘        └─────────┘
   (modified!)
```

**Summary Table:**

| Type | Passed As | Modifications Visible? |
|------|-----------|------------------------|
| int, float, bool | Value | No |
| string | Value | No (immutable anyway) |
| array | Value | No |
| struct | Value | No |
| slice | Value (but contains pointer) | Yes (elements) |
| map | Value (but contains pointer) | Yes |
| channel | Value (but contains pointer) | Yes |
| pointer | Value (of address) | Yes |

</details>

---

### Q14: How would you implement a set in Go?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Go doesn't have a built-in set, but maps work perfectly.**

**Method 1: `map[T]bool`**

```go
type StringSet map[string]bool

func NewStringSet() StringSet {
    return make(StringSet)
}

func (s StringSet) Add(item string) {
    s[item] = true
}

func (s StringSet) Remove(item string) {
    delete(s, item)
}

func (s StringSet) Contains(item string) bool {
    return s[item]
}

func (s StringSet) Size() int {
    return len(s)
}

// Usage
set := NewStringSet()
set.Add("apple")
set.Add("banana")
set.Add("apple")  // Duplicate ignored

fmt.Println(set.Size())           // 2
fmt.Println(set.Contains("apple")) // true
```

**Method 2: `map[T]struct{}` (More Memory Efficient)**

```go
type StringSet map[string]struct{}

func (s StringSet) Add(item string) {
    s[item] = struct{}{}  // struct{} uses 0 bytes
}

func (s StringSet) Contains(item string) bool {
    _, exists := s[item]
    return exists
}
```

**Set Operations:**

```go
type StringSet map[string]struct{}

// Union
func (s StringSet) Union(other StringSet) StringSet {
    result := make(StringSet)
    for k := range s {
        result[k] = struct{}{}
    }
    for k := range other {
        result[k] = struct{}{}
    }
    return result
}

// Intersection
func (s StringSet) Intersection(other StringSet) StringSet {
    result := make(StringSet)
    for k := range s {
        if _, exists := other[k]; exists {
            result[k] = struct{}{}
        }
    }
    return result
}

// Difference
func (s StringSet) Difference(other StringSet) StringSet {
    result := make(StringSet)
    for k := range s {
        if _, exists := other[k]; !exists {
            result[k] = struct{}{}
        }
    }
    return result
}

// Usage
s1 := StringSet{"a": {}, "b": {}, "c": {}}
s2 := StringSet{"b": {}, "c": {}, "d": {}}

union := s1.Union(s2)           // {a, b, c, d}
intersection := s1.Intersection(s2)  // {b, c}
diff := s1.Difference(s2)       // {a}
```

**Generic Set (Go 1.18+):**

```go
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
    return make(Set[T])
}

func (s Set[T]) Add(item T) {
    s[item] = struct{}{}
}

func (s Set[T]) Contains(item T) bool {
    _, exists := s[item]
    return exists
}

// Usage
intSet := NewSet[int]()
intSet.Add(1)
intSet.Add(2)

stringSet := NewSet[string]()
stringSet.Add("hello")
```

</details>

---

### Q15: Explain struct embedding and composition in Go.

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Go doesn't have inheritance; it uses composition.**

**Basic Embedding:**

```go
type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Age     int
    Address  // Embedded field (anonymous)
}

// Usage
p := Person{
    Name: "Alice",
    Age:  30,
    Address: Address{
        City:    "New York",
        Country: "USA",
    },
}

// Promoted fields
fmt.Println(p.City)  // New York (promoted from Address)
fmt.Println(p.Address.City)  // Also valid
```

**Method Promotion:**

```go
type Engine struct {
    Power int
}

func (e Engine) Start() {
    fmt.Println("Engine started")
}

type Car struct {
    Model string
    Engine  // Embedded
}

// Engine methods promoted to Car
car := Car{Model: "Tesla", Engine: Engine{Power: 500}}
car.Start()  // Engine started (promoted method)
```

**Multiple Embedding:**

```go
type Logger struct{}

func (l Logger) Log(msg string) {
    fmt.Println("[LOG]", msg)
}

type Database struct{}

func (d Database) Query(sql string) {
    fmt.Println("Executing:", sql)
}

type Service struct {
    Logger
    Database
}

// Service has both Log and Query methods
service := Service{}
service.Log("Starting service")
service.Query("SELECT * FROM users")
```

**Name Conflicts:**

```go
type A struct {
    Name string
}

func (a A) Print() {
    fmt.Println("A:", a.Name)
}

type B struct {
    Name string
}

func (b B) Print() {
    fmt.Println("B:", b.Name)
}

type C struct {
    A
    B
}

c := C{
    A: A{Name: "Alice"},
    B: B{Name: "Bob"},
}

// c.Print()  // Ambiguous! Compile error
c.A.Print()   // OK: A: Alice
c.B.Print()   // OK: B: Bob

// c.Name  // Ambiguous! Compile error
fmt.Println(c.A.Name)  // OK: Alice
fmt.Println(c.B.Name)  // OK: Bob
```

**Interface Implementation via Embedding:**

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(s string)
}

type FileReader struct{}

func (f FileReader) Read() string {
    return "file content"
}

type FileWriter struct{}

func (f FileWriter) Write(s string) {
    fmt.Println("Writing:", s)
}

// ReadWriter via embedding
type ReadWriter struct {
    FileReader
    FileWriter
}

// ReadWriter automatically implements both interfaces!
func process(rw ReadWriter) {
    content := rw.Read()
    rw.Write(content)
}
```

**Comparison with Inheritance:**

```
Inheritance (Java/C++):
- "is-a" relationship
- Tight coupling
- Polymorphism via inheritance

Composition (Go):
- "has-a" relationship
- Loose coupling
- Polymorphism via interfaces
```

</details>

---

## Concurrency

### Q16: What are goroutines and how do they differ from threads?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Goroutines are lightweight threads managed by Go runtime.**

**Key Differences:**

| Feature | Goroutine | OS Thread |
|---------|-----------|-----------|
| **Stack Size** | ~2KB (grows dynamically) | ~2MB (fixed) |
| **Creation Cost** | Microseconds | Milliseconds |
| **Switching Cost** | Nanoseconds | Microseconds |
| **Managed By** | Go runtime | Operating system |
| **Scalability** | Millions possible | Thousands practical |

**Example:**

```go
// Creating a goroutine (cheap!)
go func() {
    fmt.Println("Hello from goroutine")
}()

// Creating OS thread (expensive)
thread := Thread.new { puts "Hello from thread" }
```

**How Goroutines Work:**

```
Go Runtime (M:N Scheduler)
┌────────────────────────────┐
│    Goroutines (G)          │
│    G1 G2 G3 G4 G5 ...      │
└────────────────────────────┘
         ↓
┌────────────────────────────┐
│    Logical Processors (P)  │
│    P1  P2  P3              │
└────────────────────────────┘
         ↓
┌────────────────────────────┐
│    OS Threads (M)          │
│    M1  M2  M3              │
└────────────────────────────┘
```

**Scheduling:**

```go
// GOMAXPROCS controls logical processors
runtime.GOMAXPROCS(4)  // Use 4 cores

// Goroutines are multiplexed onto OS threads
for i := 0; i < 1000000; i++ {
    go func(n int) {
        // Work
    }(i)
}
// Only a few OS threads needed!
```

**Synchronization:**

```go
func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Println("Worker", id)
        }(i)
    }
    
    wg.Wait()  // Wait for all goroutines
}
```

**Why Goroutines are Better:**

1. **Resource Efficient**: Can run millions concurrently
2. **Fast Context Switching**: Managed in user space
3. **Dynamic Stack**: Grows and shrinks as needed
4. **Integrated with Language**: Channels, select, etc.

</details>

---

### Q17: Explain channels and how they work. What's the difference between buffered and unbuffered channels?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Channels enable communication between goroutines.**

**Unbuffered Channels (Synchronous):**

```go
ch := make(chan int)  // Unbuffered

// Sender blocks until receiver ready
go func() {
    ch <- 42  // Blocks until received
}()

value := <-ch  // Blocks until sent
fmt.Println(value)  // 42
```

**Buffered Channels (Asynchronous):**

```go
ch := make(chan int, 3)  // Buffer size 3

// Can send without receiver (until buffer full)
ch <- 1  // Doesn't block
ch <- 2  // Doesn't block
ch <- 3  // Doesn't block
// ch <- 4  // Blocks (buffer full)

fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
```

**Visual Representation:**

```
Unbuffered Channel:
Sender → │ │ → Receiver
         (no buffer)
         ↓
    Synchronous handoff

Buffered Channel (size 3):
Sender → │1│2│3│ → Receiver
         └─┴─┴─┘
         (buffer)
         ↓
    Asynchronous up to capacity
```

**Comparison Table:**

| Feature | Unbuffered | Buffered |
|---------|-----------|----------|
| **Synchronization** | Yes | Until buffer full |
| **Blocking send** | Always | Only when full |
| **Blocking receive** | Always | Only when empty |
| **Use case** | Synchronization | Decoupling, smoothing |

**Closing Channels:**

```go
ch := make(chan int, 3)

// Send values
ch <- 1
ch <- 2
ch <- 3
close(ch)  // Close when done

// Receive until closed
for val := range ch {
    fmt.Println(val)
}

// Check if closed
val, ok := <-ch
if !ok {
    fmt.Println("Channel closed")
}
```

**Important Rules:**

```go
// ✅ Only sender closes
close(ch)

// ❌ Don't close twice
close(ch)
close(ch)  // PANIC!

// ❌ Don't send on closed channel
ch <- 1  // PANIC!

// ✅ Receive from closed channel returns zero value
val := <-ch  // 0 (if int channel)
```

**Directional Channels:**

```go
// Send-only channel
func producer(ch chan<- int) {
    ch <- 42
    // val := <-ch  // Compile error!
}

// Receive-only channel
func consumer(ch <-chan int) {
    val := <-ch
    // ch <- 42  // Compile error!
}

func main() {
    ch := make(chan int)
    go producer(ch)
    go consumer(ch)
}
```

**Real-World Example:**

```go
// Worker pool with buffered channel
func workerPool(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        // Process job
        time.Sleep(time.Second)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for i := 0; i < 3; i++ {
        go workerPool(jobs, results)
    }
    
    // Send jobs
    for i := 1; i <= 10; i++ {
        jobs <- i
    }
    close(jobs)
    
    // Collect results
    for i := 1; i <= 10; i++ {
        fmt.Println(<-results)
    }
}
```

</details>

---

### Q18: What is the `select` statement and when do you use it?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**`select` multiplexes multiple channel operations.**

**Basic Syntax:**

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case ch3 <- value:
    fmt.Println("Sent to ch3")
default:
    fmt.Println("No channel ready")
}
```

**Use Case 1: Timeout Pattern**

```go
select {
case result := <-ch:
    fmt.Println("Received:", result)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout!")
}
```

**Use Case 2: Non-blocking Channel Operations**

```go
// Non-blocking send
select {
case ch <- value:
    fmt.Println("Sent")
default:
    fmt.Println("Channel full, skipping")
}

// Non-blocking receive
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No message available")
}
```

**Use Case 3: Context Cancellation**

```go
func worker(ctx context.Context, ch <-chan int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker cancelled")
            return
        case job := <-ch:
            fmt.Println("Processing", job)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    ch := make(chan int)
    
    go worker(ctx, ch)
    
    ch <- 1
    ch <- 2
    
    cancel()  // Cancel worker
    time.Sleep(time.Second)
}
```

**Use Case 4: Fan-In (Merge Channels)**

```go
func fanIn(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)
    
    go func() {
        defer close(out)
        for {
            select {
            case val, ok := <-ch1:
                if !ok {
                    ch1 = nil  // Disable this case
                    continue
                }
                out <- val
            case val, ok := <-ch2:
                if !ok {
                    ch2 = nil
                    continue
                }
                out <- val
            }
            
            // Both channels closed
            if ch1 == nil && ch2 == nil {
                return
            }
        }
    }()
    
    return out
}
```

**Random Selection:**

```go
// If multiple cases ready, select chooses randomly
ch1 := make(chan int, 1)
ch2 := make(chan int, 1)

ch1 <- 1
ch2 <- 2

select {
case v := <-ch1:
    fmt.Println("ch1:", v)
case v := <-ch2:
    fmt.Println("ch2:", v)
}
// Output: ch1: 1 OR ch2: 2 (random!)
```

**Complete Example (Web Scraper with Timeout):**

```go
func fetchURL(url string, ch chan<- string) {
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("Error: %v", err)
        return
    }
    defer resp.Body.Close()
    
    ch <- fmt.Sprintf("Status: %s", resp.Status)
}

func main() {
    urls := []string{
        "https://google.com",
        "https://github.com",
        "https://stackoverflow.com",
    }
    
    ch := make(chan string)
    
    for _, url := range urls {
        go fetchURL(url, ch)
    }
    
    for range urls {
        select {
        case result := <-ch:
            fmt.Println(result)
        case <-time.After(5 * time.Second):
            fmt.Println("Timeout!")
            return
        }
    }
}
```

</details>

---

### Q19: What is a race condition and how do you prevent it?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**A race condition occurs when multiple goroutines access shared data concurrently, and at least one modifies it.**

**Example Race Condition:**

```go
var counter int

func increment() {
    counter++  // Read-Modify-Write (not atomic!)
}

func main() {
    for i := 0; i < 1000; i++ {
        go increment()
    }
    
    time.Sleep(time.Second)
    fmt.Println(counter)  // Unpredictable! (not 1000)
}
```

**Why This Happens:**

```
counter++ is actually three operations:
1. Read counter (e.g., 5)
2. Add 1 (5 + 1 = 6)
3. Write back (counter = 6)

Goroutine A: Read(5) → Add(6) → Write(6)
Goroutine B:     Read(5) → Add(6) → Write(6)
                      ↑
               Both read 5, both write 6
               Lost one increment!
```

**Detection:**

```bash
# Race detector (runtime tool)
go run -race main.go
go test -race ./...
go build -race
```

**Solution 1: Mutex**

```go
var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}

// Or use defer
func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}
```

**Solution 2: Atomic Operations**

```go
import "sync/atomic"

var counter int64

func increment() {
    atomic.AddInt64(&counter, 1)
}
```

**Solution 3: Channels (Share Memory by Communicating)**

```go
func counterServer(ch chan int) {
    counter := 0
    for range ch {
        counter++
    }
    fmt.Println("Final count:", counter)
}

func main() {
    ch := make(chan int)
    go counterServer(ch)
    
    for i := 0; i < 1000; i++ {
        ch <- 1
    }
    
    close(ch)
    time.Sleep(time.Second)
}
```

**Solution 4: sync.Map (Concurrent Map)**

```go
var m sync.Map

// Set
m.Store("key", "value")

// Get
value, ok := m.Load("key")

// Delete
m.Delete("key")

// No race conditions!
```

**Comparison:**

| Method | Use When | Performance |
|--------|----------|-------------|
| **Mutex** | Protecting any shared state | Good |
| **Atomic** | Simple counters, flags | Excellent |
| **Channels** | Communication patterns | Good |
| **sync.Map** | Concurrent map access | Good |

**Real-World Example (Safe Counter):**

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func main() {
    counter := &SafeCounter{}
    
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Println(counter.Value())  // Always 1000!
}
```

</details>

---

### Q20: Explain context in Go. When and why would you use it?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Context carries deadlines, cancellation signals, and request-scoped values across API boundaries.**

**Creating Contexts:**

```go
// Root context (never cancelled)
ctx := context.Background()

// Cancellable context
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Timeout context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Deadline context
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Value context (request-scoped data)
ctx = context.WithValue(ctx, "userID", 123)
```

**Use Case 1: HTTP Request Cancellation**

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()  // Request context
    
    // Do work that respects cancellation
    result, err := processRequest(ctx)
    if err != nil {
        if ctx.Err() == context.Canceled {
            http.Error(w, "Request cancelled", 499)
            return
        }
        http.Error(w, err.Error(), 500)
        return
    }
    
    w.Write([]byte(result))
}

func processRequest(ctx context.Context) (string, error) {
    // Check cancellation periodically
    select {
    case <-ctx.Done():
        return "", ctx.Err()
    default:
        // Continue processing
    }
    
    // Perform work
    return "result", nil
}
```

**Use Case 2: Database Query with Timeout**

```go
func getUser(ctx context.Context, id int) (*User, error) {
    // Context passed to DB query
    row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", id)
    
    var user User
    err := row.Scan(&user.ID, &user.Name)
    if err != nil {
        return nil, err
    }
    
    return &user, nil
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    user, err := getUser(ctx, 123)
    if err != nil {
        if err == context.DeadlineExceeded {
            fmt.Println("Query timeout!")
        }
        return
    }
    
    fmt.Println(user)
}
```

**Use Case 3: Goroutine Cancellation**

```go
func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            // Do work
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    // Start workers
    for i := 0; i < 3; i++ {
        go worker(ctx, i)
    }
    
    // Let them work
    time.Sleep(2 * time.Second)
    
    // Cancel all workers
    cancel()
    
    time.Sleep(time.Second)
}
```

**Use Case 4: Request-Scoped Values**

```go
type key string

const userIDKey key = "userID"

func withUserID(ctx context.Context, id int) context.Context {
    return context.WithValue(ctx, userIDKey, id)
}

func getUserID(ctx context.Context) (int, bool) {
    id, ok := ctx.Value(userIDKey).(int)
    return id, ok
}

func handler(w http.ResponseWriter, r *http.Request) {
    userID := 123  // From auth middleware
    
    ctx := withUserID(r.Context(), userID)
    
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    if userID, ok := getUserID(ctx); ok {
        fmt.Println("Processing request for user:", userID)
    }
}
```

**Best Practices:**

```go
// ✅ Pass context as first parameter
func DoSomething(ctx context.Context, arg string) error

// ❌ Don't store context in structs
type Server struct {
    ctx context.Context  // Bad!
}

// ✅ Use context.TODO() if unsure
ctx := context.TODO()

// ✅ Always call cancel (even if not needed)
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()

// ❌ Don't pass nil context
DoSomething(nil, "arg")  // Bad!

// ✅ Pass context.Background() instead
DoSomething(context.Background(), "arg")
```

**Context Errors:**

```go
switch ctx.Err() {
case context.Canceled:
    // Context was cancelled
case context.DeadlineExceeded:
    // Timeout occurred
case nil:
    // No error
}
```

</details>

## Error Handling

### Q21: How does Go handle errors? Why doesn't it have exceptions?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Go treats errors as values, not exceptions.**

**Error Handling Pattern:**

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)
```

**Why No Exceptions?**

1. **Explicit error handling** - Forces developers to handle errors
2. **No hidden control flow** - Code is easier to follow
3. **No performance overhead** - No stack unwinding
4. **Simpler code** - No try/catch/finally blocks

**Comparison:**

```java
// Java (Exception-based)
try {
    result = divide(10, 0);
    System.out.println(result);
} catch (DivisionByZeroException e) {
    System.err.println(e);
}

// Go (Error-based)
result, err := divide(10, 0)
if err != nil {
    log.Println(err)
    return
}
fmt.Println(result)
```

**Error Wrapping (Go 1.13+):**

```go
func readConfig() error {
    _, err := os.ReadFile("config.json")
    if err != nil {
        return fmt.Errorf("reading config: %w", err)
    }
    return nil
}

func main() {
    err := readConfig()
    if err != nil {
        // Unwrap to check original error
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("Config file doesn't exist")
        }
    }
}
```

**Custom Errors:**

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func validateEmail(email string) error {
    if !strings.Contains(email, "@") {
        return ValidationError{
            Field:   "email",
            Message: "must contain @ symbol",
        }
    }
    return nil
}

// Check error type
err := validateEmail("invalid")
var valErr ValidationError
if errors.As(err, &valErr) {
    fmt.Printf("Field: %s, Message: %s\n", valErr.Field, valErr.Message)
}
```

</details>

---

### Q22: Explain `panic` and `recover`. When should you use them?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**`panic` stops normal execution; `recover` catches panics.**

**Basic Usage:**

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    
    panic("something went wrong")
    fmt.Println("This won't execute")
}

// Output:
// Recovered from panic: something went wrong
```

**When to Use Panic:**

```go
// ✅ Programmer errors (can't continue)
func mustConnect(url string) *Connection {
    conn, err := connect(url)
    if err != nil {
        panic("failed to connect: " + err.Error())
    }
    return conn
}

// ✅ Initialization failures
func init() {
    if err := loadConfig(); err != nil {
        panic("config required: " + err.Error())
    }
}

// ❌ Don't panic for normal errors
func getUser(id int) (*User, error) {
    if id < 0 {
        panic("invalid user ID")  // Bad! Return error instead
    }
    return nil, errors.New("user not found")  // Good
}
```

**When to Use Recover:**

```go
// ✅ HTTP server (prevent one request from crashing server)
func handler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Panic in handler: %v", r)
            http.Error(w, "Internal Server Error", 500)
        }
    }()
    
    // Handler logic that might panic
}

// ✅ Goroutine safety
func safeGoroutine(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Goroutine panic: %v", r)
            }
        }()
        fn()
    }()
}
```

**Complete Example:**

```go
func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
            debug.PrintStack()
        }
    }()
    
    // Simulating panic
    arr := []int{1, 2, 3}
    fmt.Println(arr[10])  // Panic: index out of range
}

func main() {
    riskyOperation()
    fmt.Println("Program continues")
}

// Output:
// Recovered: runtime error: index out of range [10] with length 3
// [stack trace]
// Program continues
```

**Best Practices:**

```go
// ✅ Return errors, don't panic
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// ❌ Don't panic in library code
func divide(a, b float64) float64 {
    if b == 0 {
        panic("division by zero")  // Bad for library
    }
    return a / b
}

// ✅ Recover only in known safe spots
func processRequest() {
    defer func() {
        recover()  // Recover and log
    }()
    // Process
}

// ❌ Don't silently recover
func processRequest() {
    defer func() {
        recover()  // Swallows all errors!
    }()
}
```

</details>

---

## Interfaces & Polymorphism

### Q23: How do interfaces work in Go? What is implicit interface satisfaction?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Interfaces define behavior (method sets); types satisfy them implicitly.**

**Interface Definition:**

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

**Implicit Satisfaction:**

```go
type Rectangle struct {
    Width, Height float64
}

// No explicit "implements Shape" needed!
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Rectangle automatically satisfies Shape interface
func printInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    printInfo(rect)  // Works!
}
```

**Why Implicit is Better:**

```
Java/C# (Explicit):
class Rectangle implements Shape { ... }
- Tight coupling
- Must know interface at type definition
- Hard to retrofit existing types

Go (Implicit):
type Rectangle struct { ... }
func (r Rectangle) Area() float64 { ... }
- Loose coupling
- Can satisfy interfaces defined later
- Easy to adapt existing types
```

**Interface Composition:**

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Composed interface
type ReadWriter interface {
    Reader
    Writer
}

// Any type with Read() and Write() satisfies ReadWriter
type File struct{}

func (f File) Read(p []byte) (int, error) {
    return 0, nil
}

func (f File) Write(p []byte) (int, error) {
    return 0, nil
}

var rw ReadWriter = File{}  // Works!
```

**Empty Interface:**

```go
// interface{} accepts ANY type
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

PrintAnything(42)
PrintAnything("hello")
PrintAnything([]int{1, 2, 3})

// Go 1.18+: "any" is alias for interface{}
func PrintAnything(v any) {
    fmt.Println(v)
}
```

**Interface Values:**

```go
var s Shape = Rectangle{Width: 10, Height: 5}

// Interface value holds (type, value) pair
// Type: Rectangle
// Value: {Width: 10, Height: 5}

fmt.Printf("%T\n", s)  // main.Rectangle
fmt.Printf("%v\n", s)  // {10 5}
```

</details>

---

### Q24: What is the difference between a value receiver and a pointer receiver?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Value receivers operate on copies; pointer receivers modify originals.**

**Value Receiver:**

```go
type Counter struct {
    Count int
}

func (c Counter) Increment() {
    c.Count++  // Modifies copy, not original
}

func main() {
    counter := Counter{Count: 0}
    counter.Increment()
    fmt.Println(counter.Count)  // 0 (unchanged!)
}
```

**Pointer Receiver:**

```go
func (c *Counter) Increment() {
    c.Count++  // Modifies original
}

func main() {
    counter := Counter{Count: 0}
    counter.Increment()  // Go auto-converts to (&counter).Increment()
    fmt.Println(counter.Count)  // 1 (modified!)
}
```

**Decision Matrix:**

| Use Pointer Receiver When | Use Value Receiver When |
|---------------------------|-------------------------|
| Method modifies receiver | Receiver is small (≤ few ints) |
| Receiver is large struct | Receiver is immutable |
| Want consistency (if any method uses pointer) | Receiver is value type |
| Method called on nil receiver is valid | Don't need to modify |

**Interface Satisfaction:**

```go
type Incrementer interface {
    Increment()
}

type Counter struct {
    Count int
}

func (c *Counter) Increment() {
    c.Count++
}

// Pointer receiver method
var i Incrementer = &Counter{}  // ✅ Works
// var i Incrementer = Counter{}  // ❌ Doesn't work!

// Why? Counter doesn't have Increment(), only *Counter does
```

**Best Practice (Google Style):**

```go
// ✅ Consistent - all methods use pointer receivers
type User struct {
    Name  string
    Email string
}

func (u *User) SetName(name string) {
    u.Name = name
}

func (u *User) SetEmail(email string) {
    u.Email = email
}

func (u *User) GetName() string {
    return u.Name  // Still use pointer for consistency
}

// ❌ Mixed - confusing
func (u User) GetName() string {
    return u.Name
}

func (u *User) SetName(name string) {
    u.Name = name
}
```

**Large Struct Example:**

```go
type LargeStruct struct {
    Data [1000]int
}

// ❌ Bad: Copies 8KB on every call
func (l LargeStruct) Process() {
    // ...
}

// ✅ Good: Passes 8-byte pointer
func (l *LargeStruct) Process() {
    // ...
}
```

</details>

---

### Q25: Explain type assertion and type switch.

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Type assertion extracts concrete type from interface; type switch checks multiple types.**

**Type Assertion:**

```go
var i interface{} = "hello"

// Unsafe assertion (panics if wrong type)
s := i.(string)
fmt.Println(s)  // hello

n := i.(int)  // PANIC!

// Safe assertion (comma-ok idiom)
s, ok := i.(string)
if ok {
    fmt.Println("String:", s)
} else {
    fmt.Println("Not a string")
}

n, ok := i.(int)
fmt.Println(n, ok)  // 0 false
```

**Type Switch:**

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    case bool:
        fmt.Println("Boolean:", v)
    case nil:
        fmt.Println("Nil")
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

describe(42)       // Integer: 42
describe("hello")  // String: hello
describe(true)     // Boolean: true
describe(nil)      // Nil
describe(3.14)     // Unknown type: float64
```

**Real-World Example (JSON Unmarshaling):**

```go
func processValue(data interface{}) {
    switch v := data.(type) {
    case map[string]interface{}:
        fmt.Println("Object with keys:", len(v))
        for key, val := range v {
            fmt.Printf("  %s: %v\n", key, val)
        }
    case []interface{}:
        fmt.Println("Array with", len(v), "elements")
    case string:
        fmt.Println("String:", v)
    case float64:
        fmt.Println("Number:", v)
    case bool:
        fmt.Println("Boolean:", v)
    case nil:
        fmt.Println("Null")
    }
}

var data interface{}
json.Unmarshal([]byte(`{"name": "Alice", "age": 30}`), &data)
processValue(data)
```

**Error Type Checking:**

```go
_, err := os.Open("file.txt")

// Check for specific error type
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("Path error:", pathErr.Path)
}

// Check for specific error value
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File doesn't exist")
}
```

</details>

---

## Performance & Optimization

### Q26: How do you optimize Go code for performance?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Key optimization strategies:**

**1. Use Profiling (Don't Guess!):**

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench .
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench .
go tool pprof mem.prof

# Benchmark code
go test -bench . -benchmem
```

**2. Preallocate Slices:**

```go
// ❌ Bad: Multiple allocations
func bad(n int) []int {
    var result []int
    for i := 0; i < n; i++ {
        result = append(result, i)
    }
    return result
}

// ✅ Good: Single allocation
func good(n int) []int {
    result := make([]int, 0, n)
    for i := 0; i < n; i++ {
        result = append(result, i)
    }
    return result
}

// Benchmark results:
// bad:  15000 ns/op  57344 B/op  10 allocs/op
// good:  3000 ns/op   8192 B/op   1 allocs/op
```

**3. Use sync.Pool for Reusable Objects:**

```go
var bufPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func process() {
    buf := bufPool.Get().(*bytes.Buffer)
    defer bufPool.Put(buf)
    buf.Reset()
    
    // Use buf
}
```

**4. Avoid Unnecessary Allocations:**

```go
// ❌ Bad: Creates new strings
func bad(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s  // Allocates new string each time
    }
    return result
}

// ✅ Good: Single allocation
func good(strs []string) string {
    var builder strings.Builder
    for _, s := range strs {
        builder.WriteString(s)
    }
    return builder.String()
}
```

**5. Use Pointer Receivers for Large Structs:**

```go
type LargeStruct struct {
    Data [1000]int
}

// ❌ Bad: Copies 8KB
func (l LargeStruct) Process() {}

// ✅ Good: 8-byte pointer
func (l *LargeStruct) Process() {}
```

**6. Batch Channel Operations:**

```go
// ❌ Bad: Send one at a time
for i := 0; i < 1000; i++ {
    ch <- i
}

// ✅ Good: Batch sends
batch := make([]int, 0, 100)
for i := 0; i < 1000; i++ {
    batch = append(batch, i)
    if len(batch) == 100 {
        ch <- batch
        batch = batch[:0]
    }
}
```

**7. Use strconv for Conversions:**

```go
// ❌ Bad: Slower
s := fmt.Sprintf("%d", num)

// ✅ Good: Faster
s := strconv.Itoa(num)

// Benchmark:
// Sprintf: 150 ns/op  16 B/op  2 allocs/op
// Itoa:     50 ns/op   8 B/op  1 allocs/op
```

**8. Avoid Defer in Hot Paths:**

```go
// ❌ Bad: defer has overhead
func hotPath() {
    mu.Lock()
    defer mu.Unlock()
    
    // Simple operation
}

// ✅ Good: Manual unlock
func hotPath() {
    mu.Lock()
    // Simple operation
    mu.Unlock()
}

// Use defer for complex functions with multiple returns
```

</details>

---

### Q27: What is escape analysis in Go?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Escape analysis determines whether a variable goes on stack or heap.**

**Stack vs Heap:**

```
Stack:
- Fast allocation/deallocation
- Automatic cleanup
- Limited size
- Function-local

Heap:
- Slower allocation
- Garbage collected
- Large size
- Survives function return
```

**Example 1: Stack Allocation**

```go
func stackAlloc() {
    x := 42  // Allocated on stack
    fmt.Println(x)
}  // x destroyed when function returns
```

**Example 2: Heap Allocation (Escapes)**

```go
func heapAlloc() *int {
    x := 42
    return &x  // x escapes to heap (survives return)
}
```

**Check Escape Analysis:**

```bash
go build -gcflags="-m" main.go

# Output:
# ./main.go:5: moved to heap: x
# ./main.go:6: &x escapes to heap
```

**Common Escape Scenarios:**

```go
// 1. Return pointer
func escape1() *int {
    x := 42
    return &x  // Escapes
}

// 2. Store in global
var global *int

func escape2() {
    x := 42
    global = &x  // Escapes
}

// 3. Send to channel
func escape3() {
    x := 42
    ch <- &x  // Escapes
}

// 4. Assign to interface
func escape4() {
    x := 42
    var i interface{} = x  // May escape
}

// 5. Large structs
type Large struct {
    Data [10000]int
}

func escape5() Large {
    return Large{}  // May escape if too large for stack
}
```

**Optimization Tips:**

```go
// ✅ Keep on stack when possible
func noEscape() {
    x := 42
    process(x)  // Passed by value, no escape
}

// ❌ Forces heap allocation
func escapes() {
    x := 42
    process(&x)  // Pointer escapes
}

// ✅ Use value receivers for small types
type Point struct {
    X, Y int
}

func (p Point) Distance() float64 {  // Value receiver
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}
```

</details>

---

### Q28: How does the Go garbage collector work?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Go uses a concurrent, tri-color mark-and-sweep GC.**

**GC Algorithm:**

```
1. Mark Phase (Concurrent):
   - Roots (globals, stack) marked
   - Traverse object graph
   - Mark reachable objects

2. Sweep Phase (Concurrent):
   - Free unmarked objects
   - Return memory to allocator

3. Tri-Color Marking:
   - White: Unreachable (garbage)
   - Gray: Reachable, not scanned
   - Black: Reachable, scanned
```

**GC Tuning:**

```go
import "runtime/debug"

// Set GC target percentage
debug.SetGCPercent(100)  // Default

// Force GC
runtime.GC()

// Get GC stats
var stats runtime.MemStats
runtime.ReadMemStats(&stats)
fmt.Printf("Alloc: %v MB\n", stats.Alloc/1024/1024)
fmt.Printf("TotalAlloc: %v MB\n", stats.TotalAlloc/1024/1024)
fmt.Printf("NumGC: %v\n", stats.NumGC)
```

**Minimize GC Pressure:**

```go
// ✅ Reuse objects
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

// ✅ Preallocate slices
data := make([]int, 0, expectedSize)

// ✅ Avoid unnecessary pointers
type Node struct {
    Value int
    // Left  *Node  // Escapes to heap
    // Right *Node
}

// ❌ Avoid frequent small allocations
for i := 0; i < 1000000; i++ {
    _ = new(int)  // Bad!
}
```

**GC Pacing:**

```
GOGC=100 (default):
- GC runs when heap doubles
- Live heap: 10 MB → GC at 20 MB

GOGC=200:
- GC runs when heap triples
- Live heap: 10 MB → GC at 30 MB
- Less frequent GC, more memory

GOGC=50:
- GC runs at 50% growth
- Live heap: 10 MB → GC at 15 MB
- More frequent GC, less memory
```

</details>

---

## Testing

### Q29: How do you write unit tests in Go?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Go has built-in testing framework in `testing` package.**

**Basic Test:**

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

**Table-Driven Tests (Recommended):**

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -2, -3, -5},
        {"zero", 0, 5, 5},
        {"mixed", -2, 3, 1},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d; want %d", result, tt.expected)
            }
        })
    }
}
```

**Test Helpers:**

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // Mark as helper
    
    if got != want {
        t.Errorf("got %d; want %d", got, want)
    }
}

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    assertEqual(t, result, 5)  // Error points to this line
}
```

**Setup and Teardown:**

```go
func TestMain(m *testing.M) {
    // Setup
    setup()
    
    // Run tests
    code := m.Run()
    
    // Teardown
    teardown()
    
    os.Exit(code)
}

func setup() {
    // Database connection, etc.
}

func teardown() {
    // Cleanup
}
```

**Mocking with Interfaces:**

```go
// Production code
type UserStore interface {
    GetUser(id int) (*User, error)
}

type UserService struct {
    store UserStore
}

func (s *UserService) GetUserName(id int) (string, error) {
    user, err := s.store.GetUser(id)
    if err != nil {
        return "", err
    }
    return user.Name, nil
}

// Test code
type MockUserStore struct {
    users map[int]*User
}

func (m *MockUserStore) GetUser(id int) (*User, error) {
    user, ok := m.users[id]
    if !ok {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func TestGetUserName(t *testing.T) {
    mock := &MockUserStore{
        users: map[int]*User{
            1: {Name: "Alice"},
        },
    }
    
    service := &UserService{store: mock}
    
    name, err := service.GetUserName(1)
    if err != nil {
        t.Fatal(err)
    }
    
    if name != "Alice" {
        t.Errorf("got %q; want %q", name, "Alice")
    }
}
```

**Run Tests:**

```bash
go test                # Current package
go test ./...          # All packages
go test -v             # Verbose
go test -run TestAdd   # Specific test
go test -cover         # Coverage
go test -race          # Race detection
```

</details>

---

### Q30: What are benchmarks and how do you write them?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Benchmarks measure performance of code.**

**Basic Benchmark:**

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

**Run Benchmark:**

```bash
go test -bench .
go test -bench . -benchmem  # Include memory stats
go test -bench . -benchtime=10s  # Run for 10 seconds
```

**Output:**

```
BenchmarkAdd-8    1000000000    0.25 ns/op    0 B/op    0 allocs/op
                  └─ iterations └─ time/op    └─ bytes └─ allocations
```

**Table-Driven Benchmarks:**

```go
func BenchmarkFibonacci(b *testing.B) {
    tests := []struct {
        name string
        n    int
    }{
        {"small", 10},
        {"medium", 20},
        {"large", 30},
    }
    
    for _, tt := range tests {
        b.Run(tt.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                Fibonacci(tt.n)
            }
        })
    }
}
```

**Reset Timer:**

```go
func BenchmarkWithSetup(b *testing.B) {
    // Expensive setup
    data := generateLargeDataset()
    
    b.ResetTimer()  // Don't include setup in benchmark
    
    for i := 0; i < b.N; i++ {
        process(data)
    }
}
```

**Parallel Benchmarks:**

```go
func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            // Code to benchmark
        }
    })
}
```

**Comparing Benchmarks:**

```bash
# Save baseline
go test -bench . > old.txt

# Make changes...

# Compare
go test -bench . > new.txt
benchcmp old.txt new.txt
```

</details>

---

## Web Servers

### Q31: How do you build a RESTful API in Go?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Use `net/http` package (or frameworks like Gin, Echo).**

**Basic REST API:**

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: 1, Name: "Alice", Email: "alice@example.com"},
    {ID: 2, Name: "Bob", Email: "bob@example.com"},
}

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("GET /api/users", listUsers)
    mux.HandleFunc("GET /api/users/{id}", getUser)
    mux.HandleFunc("POST /api/users", createUser)
    mux.HandleFunc("PUT /api/users/{id}", updateUser)
    mux.HandleFunc("DELETE /api/users/{id}", deleteUser)
    
    log.Fatal(http.ListenAndServe(":8080", mux))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.PathValue("id"))  // Go 1.22+
    
    for _, user := range users {
        if user.ID == id {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    
    http.Error(w, "User not found", http.StatusNotFound)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    user.ID = len(users) + 1
    users = append(users, user)
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.PathValue("id"))
    
    for i, user := range users {
        if user.ID == id {
            var updated User
            if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            
            updated.ID = id
            users[i] = updated
            
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updated)
            return
        }
    }
    
    http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.PathValue("id"))
    
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    
    http.Error(w, "User not found", http.StatusNotFound)
}
```

</details>

---

### Q32: Explain middleware in Go HTTP servers.

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Middleware wraps handlers to add cross-cutting concerns.**

**Middleware Pattern:**

```go
func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Before request
        next.ServeHTTP(w, r)
        // After request
    })
}
```

**Logging Middleware:**

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        next.ServeHTTP(w, r)
        
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

**Auth Middleware:**

```go
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        
        if token != "Bearer valid-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}
```

**Chaining Middleware:**

```go
func chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middleware) - 1; i >= 0; i-- {
        h = middleware[i](h)
    }
    return h
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/", apiHandler)
    
    handler := chain(mux,
        loggingMiddleware,
        authMiddleware,
        corsMiddleware,
    )
    
    http.ListenAndServe(":8080", handler)
}
```

</details>

---

## Best Practices

### Q33: What are Go's code organization best practices?

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

**Follow standard project layout:**

```
myproject/
├── cmd/
│   └── myapp/
│       └── main.go           # Application entry point
├── internal/                 # Private application code
│   ├── config/
│   ├── database/
│   └── handlers/
├── pkg/                      # Public library code
│   └── api/
├── test/                     # Additional test files
├── scripts/                  # Build and deployment scripts
├── go.mod
└── go.sum
```

**Package Naming:**

```go
// ✅ Good
package user
package http
package json

// ❌ Bad
package userService
package http_handler
package JSON
```

**Function Naming:**

```go
// ✅ Good
func GetUser(id int) (*User, error)
func NewClient() *Client
func IsValid() bool

// ❌ Bad
func get_user(id int) (*User, error)
func createNewClient() *Client
```

</details>

---

## Coding Challenges

### Q34: Implement a LRU Cache in Go.

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

```go
package main

import (
    "container/list"
    "fmt"
)

type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
}

type entry struct {
    key   int
    value int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*list.Element),
        list:     list.New(),
    }
}

func (c *LRUCache) Get(key int) int {
    if elem, ok := c.cache[key]; ok {
        c.list.MoveToFront(elem)
        return elem.Value.(*entry).value
    }
    return -1
}

func (c *LRUCache) Put(key, value int) {
    if elem, ok := c.cache[key]; ok {
        c.list.MoveToFront(elem)
        elem.Value.(*entry).value = value
        return
    }
    
    elem := c.list.PushFront(&entry{key, value})
    c.cache[key] = elem
    
    if c.list.Len() > c.capacity {
        oldest := c.list.Back()
        if oldest != nil {
            c.list.Remove(oldest)
            delete(c.cache, oldest.Value.(*entry).key)
        }
    }
}

func main() {
    cache := Constructor(2)
    
    cache.Put(1, 1)
    cache.Put(2, 2)
    fmt.Println(cache.Get(1))  // 1
    cache.Put(3, 3)            // Evicts key 2
    fmt.Println(cache.Get(2))  // -1 (not found)
    cache.Put(4, 4)            // Evicts key 1
    fmt.Println(cache.Get(1))  // -1 (not found)
    fmt.Println(cache.Get(3))  // 3
    fmt.Println(cache.Get(4))  // 4
}
```

</details>

---

### Q35: Implement a rate limiter using Go.

**Answer:**

<details>
<summary><strong>View answers</strong></summary>

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type RateLimiter struct {
    tokens   int
    capacity int
    mu       sync.Mutex
    lastTime time.Time
}

func NewRateLimiter(capacity int) *RateLimiter {
    return &RateLimiter{
        tokens:   capacity,
        capacity: capacity,
        lastTime: time.Now(),
    }
}

func (r *RateLimiter) Allow() bool {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    now := time.Now()
    elapsed := now.Sub(r.lastTime)
    
    // Refill tokens (1 per second)
    tokensToAdd := int(elapsed.Seconds())
    if tokensToAdd > 0 {
        r.tokens = min(r.capacity, r.tokens+tokensToAdd)
        r.lastTime = now
    }
    
    if r.tokens > 0 {
        r.tokens--
        return true
    }
    
    return false
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    limiter := NewRateLimiter(5)  // 5 requests per second
    
    for i := 0; i < 10; i++ {
        if limiter.Allow() {
            fmt.Println("Request allowed")
        } else {
            fmt.Println("Request denied - rate limit exceeded")
        }
        time.Sleep(100 * time.Millisecond)
    }
}
```

</details>

---

## Generics (Go 1.18+)

### Q36: What are generics in Go and why were they added?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

Generics allow writing functions and types that work with any type while maintaining type safety.

**Before Generics:**
```go
// Code duplication
func MaxInt(a, b int) int { /* ... */ }
func MaxFloat(a, b float64) float64 { /* ... */ }

// Or type erasure (not type-safe)
func Max(a, b interface{}) interface{} { /* ... */ }
```

**With Generics:**
```go
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

**Why added:**
1. Eliminate code duplication
2. Maintain type safety (unlike interface{})
3. Improve performance (no boxing/unboxing)
4. Enable type-safe data structures

</details>

---

### Q37: Explain the difference between `any` and `comparable` constraints.

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

```go
// any: Accepts ANY type (alias for interface{})
func Print[T any](v T) {
    fmt.Println(v)
}

// comparable: Only types that support == and !=
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {  // Requires ==
            return true
        }
    }
    return false
}

// Works:
Contains([]int{1, 2, 3}, 2)         // ✅
Contains([]string{"a", "b"}, "a")   // ✅

// Doesn't work:
Contains([][]int{{1}}, []int{1})    // ❌ Slices not comparable
```

**Comparable types:**
- Basic types (int, string, bool, etc.)
- Pointers
- Structs (if all fields comparable)
- Arrays (if element type comparable)

**Not comparable:**
- Slices
- Maps
- Functions

</details>

---

### Q38: What does the `~` (tilde) mean in type constraints?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**`~` matches the underlying type:**

```go
type Integer interface {
    ~int | ~int64  // Matches int, int64, AND their custom types
}

type MyInt int
type UserID int64

func Double[T Integer](n T) T {
    return n * 2
}

var x MyInt = 5
var id UserID = 100

fmt.Println(Double(x))   // ✅ Works (MyInt's underlying type is int)
fmt.Println(Double(id))  // ✅ Works (UserID's underlying type is int64)
```

**Without `~`:**
```go
type Integer interface {
    int | int64  // Only exact types
}

// Double(x)   // ❌ Compile error
// Double(id)  // ❌ Compile error
```

**Use case:** Supporting custom types (like type aliases).

</details>

---

### Q39: Can you write a generic Stack in Go?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
    return len(s.items)
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
intStack.Push(3)

val, ok := intStack.Pop()
fmt.Println(val, ok)  // 3 true

stringStack := Stack[string]{}
stringStack.Push("hello")
stringStack.Push("world")
```

</details>

---

### Q40: When should you NOT use generics?

**Answer:**

<details>
<summary><strong>View answer</strong></summary>

**Avoid generics when:**

1. **Only one concrete type needed**
   ```go
   // ❌ Unnecessary
   func ProcessInts[T int](nums []T) { /* ... */ }
   
   // ✅ Just use int
   func ProcessInts(nums []int) { /* ... */ }
   ```

2. **Interface is sufficient**
   ```go
   // ✅ Better with interface
   func Print(v fmt.Stringer) {
       fmt.Println(v.String())
   }
   ```

3. **Logic differs significantly per type**
   ```go
   // ❌ Don't force generics
   func Process[T any](v T) {
       switch any(v).(type) {
       case int: /* completely different logic */
       case string: /* completely different logic */
       }
   }
   
   // ✅ Separate functions
   func ProcessInt(v int) { /* ... */ }
   func ProcessString(v string) { /* ... */ }
   ```

4. **Performance doesn't matter**
   ```go
   // ✅ Simpler with interface{}
   func Log(v any) {
       fmt.Println(v)
   }
   ```

**When TO use generics:**
- Data structures (Stack, Queue, Tree)
- Generic algorithms (Sort, Filter, Map)
- Type-safe wrappers (Result, Optional)
- Avoiding code duplication

---

</details>

</details>

## Resources

<details>
<summary><strong>View contents</strong></summary>

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
