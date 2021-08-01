# Golang

## Why Go

1. Fast compile times
2. Ease to development
3. Fast execution
4. Automatic garbage collection

## Go vs JS

1. TYPING

- Go: Strongly typed (String, Float, Int, Byte, Struct...)
- JS: Dynamically typed

2. STRUCTURES

- Go: Structs, Pointers, Methods, Interfaces
- JS: ES6 classes

3. ERROR HANDLING

- Go: Explicit (sad path won't handle itself)
- JS: Built-in

4. MULTI-TASKING

- Go: Multi-Threaded (Concurrency, Goroutines, Sync)
- JS: Single-Threaded (Callbacks, async await, sagas, sadness)

5. OPINIONATED-NESS

- Go: Strong Opinions (Convention, built-in tooling and linters)
- JS: Fluid Opinions (Subjective to the mood that day)

## Anatomy of a go file

```go
// package name
// Every go program needs at least one package main
package main

// import built-in packages
import "fmt"

// Every go program needs one main function
// It's the entry point for the program where
// go starts executing the code
func main() {
  fmt.Println("Hello, World")
}
```

## Create a go project

```bash
# created a directory called "hello"
$ mkdir hello

# change directory to "hello
$ cd hello

# generate "go.mod" file
$ go mod init hello

# create "hello.go" file
$ touch hello.go
```

## TDD with Go

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/tddWithGo)**

1 **Write the test first**

`hello_test.go`

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

run `go test`

```
./hello_test.go:7:10: undefined: Hello
```

2. **Write the minimal amount of code for the test to run and check the failing test output**

```go
package hello

func Hello(name string) string {
	return ""
}
```

run `go test`

```
hello_test.go:11: got "" want "Hello, Foyez"
```

3. **Write enough code to make it pass**

```go
func Hello(name string) string {
  return "Hello, " + name
}
```

run `go test`

```
PASS
ok      hello   0.004s
```

4. **Commit the code**

```git
git commit "add Hello() - greeting to people"
```

5. **Refactor**

`hello.go`

```go
package hello

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
```

`hello_test.go`

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

run `go test`

```
PASS
ok      hello   0.004s
```

6. **Amend git commit**

```git
git commit --amend
```

#### TDD workflow

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor

7. **Add test examples**

```go
func ExampleHello() {
	greeting := Hello("Zayan")
	fmt.Println(greeting)
	// Output: Hello, Zayan
}
```

run `go test -v`

```
=== RUN   TestHello
=== RUN   TestHello/saying_hello_to_people
--- PASS: TestHello (0.00s)
    --- PASS: TestHello/saying_hello_to_people (0.00s)
=== RUN   ExampleHello
--- PASS: ExampleHello (0.00s)
```

## Printing with fmt

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/printing)**

#### **Print**

```
fmt.Print()
fmt.Println()
fmt.Printf()
```

- Prints output to the stdout console
- Returns number of bytes and an error
- (The error is generally not worried about)

#### Fprint

```
fmt.Fprint()
fmt.Fprintln()
fmt.Fprintf()
```

- Prints the output to an external source (not in stdout console) (file, browser)
- Returns number of bytes, and any write error

#### Sprint

```
fmt.Sprint()
fmt.Sprintln()
fmt.Sprintf()
```

- Stores output on a character buffer
- Doesn't print to stdout console
- Returns the string
