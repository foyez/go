# Golang

**[Go Playground](https://play.golang.org/)**

- Programming Language developed by Google in 2007
- Open-sourced in 2009

## Why Go

<details>
<summary>View contents</summary>

1. Strongly type language (Variable type has to be known at compile time)
2. Statically type language (The type of the variables cannot change at runtime)
3. Nearly as fast as C and C++
4. Readable as Python
5. Fast execution
6. Automatic garbage collection
7. Run on multiple cores and builtin to support concurrency

- **In Parallel:** Downloading, Uploading, Navigating at the same time
- **Multi-Threading:** Do multiple things at once, e.g., Watching, commenting in Youtube
- **Concurrency:** Dealing with lots of things at once but not necessarily run at the same time, e.g., Multiple users booking at the same time, Multiple users editing the same document

</details>

## Go vs JS

<details>
<summary>View contents</summary>

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

</details>

## Setup go

<details>

<summary>View contents</summary>

1. Download go installer & install it. link: [go installer](https://go.dev/doc/install)
2. Add environment variables into shell config

Bash shell

```bash
# ~/.bash_profile

# set the workspace path
export GOPATH=$HOME/go-workspace # change your path correctly!

# add the go bin path to be able to execute our programs
export PATH=$PATH:$GOPATH/bin
```

Fish shell

```bash
# ~/.config/fish/config.fish

# set the workspace path
set -x GOPATH $HOME/go-workspace # change your path correctly!

# add the go bin path to be able to execute our programs
set -x PATH $PATH /usr/local/go/bin $GOPATH/bin
```

3. Create workspace

```bash
$ mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin

# $GOPATH/src : Where your Go projects / programs are located
# $GOPATH/pkg : contains every package objects
# $GOPATH/bin : The compiled binaries home
```

4. Install godoc & run godoc

```bash
$ go install golang.org/x/tools/cmd/godoc@latest

# run godoc
$ godoc -http :8000

# go to: localhost:8000/pkg
# go to personal project: localhost:8000/pkg/project-name
```

5. Update the Go version

```sh
# Uninstall the exisiting version
$ sudo rm -rf /usr/local/go
```

Download the latest go installer & install it. Link: <https://go.dev/doc/install>

7. VS Code Setup

`settings.json`

```json
{
  "[go]": {
    "editor.quickSuggestions": {
      "other": "off",
      "comments": "off",
      "strings": "off"
    }
  },
  "go.toolsManagement.autoUpdate": true,
  "go.testFlags": ["-v", "-count=1"],
  "protoc": {
    "options": ["--proto_path=proto"]
  }
}
```

- Install `[Go](https://marketplace.visualstudio.com/items?itemName=golang.go)` extension
- Install `[vscode-proto3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)` extension for **gRPC**

</details>

## Go Directory Structure

<details>
<summary>View contents</summary>

```txt
$GOPATH
│
└───bin
│
└───pkg
│
└───src
    │
    │
    └───github.com
        │
        └───github_username
            │
            └───repo_name
```

</details>

## Create a go project

<details>
<summary>View contents</summary>

```bash
# created a directory called "hello"
$ mkdir hello

# change directory to "hello
$ cd hello

# create "main.go" file
$ touch main.go

# generate "go.mod" file
$ go mod init github.com/foyez/hello # module path can be repository you want to publish
```

`go.mod`

```go
module github.com/foyez/hello // name or module path

go 1.18 // go version
```

</details>

## Anatomy of a go file

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

- compiles and runs the code: `go run <file_name>`
- build the code: `go build main.go`

</details>

## Go Commands

<details>
<summary>View contents</summary>

```sh
# run a go program
$ go run main.go # go run <file_name>

# install go packages
$ go install

# create a binary file from go codes
$ go build

# format unindent go code
$ go fmt main.go

# shows go package directory tree
$ go list

# identify unused variables & errors
$ go vet

# show go documentation
go doc fmt.Println

# install third party library
$ go get golang.org/x/lint/golint

# linting go code
golint
```

</details>

## List of keywords

<details>
<summary>View contents</summary>

The list of all **25** keywords of Go language:

1. `break`
2. `case`
3. `chan`
4. `const`
5. `continue`
6. `default`
7. `defer`
8. `else`
9. `fallthrough`
10. `for`
11. `func`
12. `go`
13. `goto`
14. `if`
15. `import`
16. `interface`
17. `map`
18. `package`
19. `range`
20. `return`
21. `select`
22. `struct`
23. `switch`
24. `type`
25. `var`

</details>

## TDD with Go

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/tddWithGo)**

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

```sh
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

```sh
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

```sh
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

7. **Add Benchmark test**

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

run `go test -v --bench . --benchmem`

```sh
BenchmarkHello    2000000000          0.46 ns/op

// This means that the loop ran 2000000000 times at a speed of 0.46 ns per loop.
```

8. **Add example tests**

```go
func ExampleHello() {
 greeting := Hello("Zayan")
 fmt.Println(greeting)
 // Output: Hello, Zayan
}

func ExampleHello() {
 greeting := Hello("Farah")
 fmt.Println(greeting)
 // Output: Hello, Farah
}
```

run `go test -v`

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

</details>

## Printing and Getting user input

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/printing)**

### **Print**

```go
fmt.Print()
fmt.Println()
fmt.Printf()
```

- Prints output to the stdout console
- Returns number of bytes and an error
- (The error is generally not worried about)

```go
name := "Zohan"

fmt.Print("Hello, ", name, "\n")
fmt.Println("Hello,", name)
fmt.Printf("Hello, %s\n", name)
```

```go
 s := Student{
  ID:   1,
  Name: "John Doe",
 }
 fmt.Printf("%s\n", "Hello")              // string
 fmt.Printf("%d\n", -34)                  // decimal
 fmt.Printf("%+d\n", 4)                   // positive decimal
 fmt.Printf("%t\n", false)                // boolean
 fmt.Printf("%f, %.2f\n", 3.1416, 3.1416) // float
 fmt.Printf("%v\n", s)
 fmt.Printf("%+v\n", s)
 fmt.Printf("%T\n", s)
```

### Fprint

```go
fmt.Fprint()
fmt.Fprintln()
fmt.Fprintf()
```

- Prints the output to an external source (not in stdout console) (file, browser)
- Returns number of bytes, and any write error

### Sprint

```go
fmt.Sprint()
fmt.Sprintln()
fmt.Sprintf()
```

- Stores output on a character buffer
- Doesn't print to stdout console
- Returns the string

### Scan

```go
fmt.Scan()
fmt.Scanln()
fmt.Scanf()
```

- Read input from the stdout console
- Returns number of bytes and an error
- (The error is generally not worried about)

```go
var name string

fmt.Scan("Hello, ", &name)
fmt.Println("Hello,", name)

fmt.Scanf("Hello, %s", &name)
fmt.Println("Hello,", name)

fmt.Scanln("Hello, ", &name)
fmt.Println("Hello,", name)
```

### `os.Args`

- Read input from command-line
- User input starts from 1st index

```go
import (
 os
)

arguments := os.Args
```

```sh
go run main.go 10 20
# argouments[1] -> 10
# argouments[2] -> 20
```

</details>

## Types

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/types)**

| Name        | Type Name                                                                   | Examples                                  |
| ----------- | --------------------------------------------------------------------------- | ----------------------------------------- |
| **INTEGER** | int, int8, int16, int32, int64<br/>unint, unint8, unint16, unint32, unint64 | var age int = 20<br/>var count unint = -5 |
| **FLOAT**   | float32, float64                                                            | var gpa float64 = 3.4                     |
| **STRING**  | string                                                                      | var fruit string = "mango"                |
| **BOOLEAN** | bool<br/>&& <code>&#124;&#124;</code> ! < <= > >= == !=                     | true false<br/>var adult bool = age > 18  |

### Identify and convert type

```go
 // identify type
 reflect.TypeOf(6) // int

 // convert type
 float(10) + 5.5 // 15.5
```

</details>

## Variables

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/variables)**

```go
package main

import "fmt"

// var variableName type = value
// can declare outside and inside of a function
var name string = "Zayan"

func main() {
 // Infer variable type
 var age = 20

 // variables without assigning value
 // return default value
 // int: 0, float: 0.0, string: "", bool: false
 var salary int

 // value cannot be changed/re-assigned
 const birthPlace = "Bangladesh"

 // variables in only function
 funcVar := "can't declare outside of a function"

 // multiple variables
 one, two := 1, "two"

 fmt.Println(name, age, salary)
 fmt.Println(birthPlace)
 fmt.Println(funcVar)
 fmt.Println(one, two)
}
```

</details>

## Control Structures: If & Else

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/control)**

```go
package main

import (
 "fmt"
)

func main() {
 var age = 10

 if age < 18 {
  fmt.Println("younger")
 } else if age == 18 {
  fmt.Println("adult")
 } else {
  fmt.Println("elder")
 }

 if name := "Farah"; name != "Farhan" {
  fmt.Println("She is Farah")
 }
}
```

</details>

## Control Structures: switch

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/control)**

```go
package main

import "fmt"

func main() {
 // *****************************************
 switch city := "Cumilla"; city {
 case "Dhaka", "Cumilla", "Sylhet":
  fmt.Println("You live in", city)
 default:
  fmt.Println("You're not from around here")
 }

 // *****************************************
 var age int = 30

 switch {
 case age < 18:
  fmt.Println("young")
 case age > 18 && age <= 40:
  fmt.Println("adult")
 default:
  fmt.Println("elder")
 }

 // *****************************************
 var num int = 9

 switch {
 case num != 10:
  fmt.Println("Does not equal 10")
  fallthrough // check other case after matching this case
 case num < 10:
  fmt.Println("Less than 10")
 case num > 10:
  fmt.Println("Greater than 10")
 default:
  fmt.Println("Is 10")
 }
}
```

</details>

## Loops

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/loops)**

```go
package main

import (
 "fmt"
)

func main() {
 // *****************************************
 // BASIC FOR LOOP
 // *****************************************
 fmt.Println("Basic for loop")
 for i := 1; i <= 5; i++ {
  fmt.Print(i)
 }

 // *****************************************
 // SIMILAR TO WHILE LOOP
 // *****************************************
 fmt.Println("\nSimilar to while loop")
 j := 1

 for j <= 5 {
  fmt.Print(j)
  j++
 }

 // *****************************************
 // INFINITE LOOP
 // *****************************************
 fmt.Println("\nInfinite loop")
 num := 1

 for {
  num = num + 2

  if num == 7 {
   continue
  }

  fmt.Print(num)

  if num == 11 {
   break
  }
 }

 // ********************************************
 // BASIC FOR LOOP ITERATION (STRING, ARRAY,...)
 // ********************************************
 fmt.Println("\nBasic for loop iteration")
 var name = "Farah"

 for i := 0; i < len(name); i++ {
  fmt.Println("Letter:", string(name[i]))
 }

 // *****************************************
 // STRING ITERATION
 // *****************************************
 fmt.Println("\nString iteration")
 var myCity = "কুমিল্লা"

 for index, letter := range myCity {
  if index % 2 == 0 {
   fmt.Printf("Index: %d, Letter:%#U\n", index, letter)
  }
 }

 // *****************************************
 // SLICE OR ARRAY ITERATION
 // *****************************************
 fmt.Println("\nSlice or Array iteration")
 cities := []string{"Dhaka", "Cumilla"}

 for _, city := range cities {
  fmt.Printf("%s ", city)
 }

 // *****************************************
 // MAP ITERATION
 // *****************************************
 fmt.Println("\nMap iteration")
 results := map[string]float64{
  "Farah":   3.4,
  "Laaibah": 3.3,
  "Zayan":   3.5,
 }

 for key, value := range results {
  fmt.Println(key, value)
 }

 // *****************************************
 // CHANNEL ITERATION
 // *****************************************
 fmt.Println("\nChannel iteration")

 ch := make(chan int)
 go func() {
  ch <- 1
  ch <- 2
  close(ch)
 }()

 for n := range ch {
  fmt.Println(n)
 }
}
```

_[Loop guide](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/)_

</details>

## Functions

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/funtions)**

- Basic function

```go
func printAge() {
 fmt.Println(10)
}
```

- return type declaration

```go
func printAge(age int) int {
 return age
}
```

- return multiple values

```go
func printAge(age int) (string, int) {
 return "name", age
}

func main() {
 name, age = printAge(10)
}
```

- return named values

```go
func printAge(age1, age2 int) (ageOfBob, ageOfSally int) {
 ageOfBob = age1
 ageOfSally = age2
 return
}
```

- unknown number of arguments / variadic function

```go
func average(ages ...int) float64 {
 total := 0

 // ages - treated as slice
 for _, age := range ages {
  total += age
 }

 return float64(total) / float64(len(ages))
}

func main() {
 fmt.Println(average(10, 20, 32))

 nums := []int{10, 20, 32}
 // unpack or spread
 fmt.Println(average(nums...))
}
```

- go functions are `lexically scoped` means variables are accessible from the functions in the same block where the functions are defined

```go
var n1 = 5

func foo(n2 int) {
 n3 := 8
 fmt.Println(n1, n2, n3)
}
```

- function as first-class value (assigning as variable, pass as argument, return as value, etc.)

```go
func print(n int, fn func(int)) {
 fn(n)
}

print(6, func(val int) {
 fmt.Println(val) // 6
})
```

```go
func add(n1 int) func(int) int {
 fn := func(n2 int) int {
  return n1 + n2
 }
 return fn
}
// n1 is in the closure of fn()

sum := add(1)
fmt.Println(sum(5)) // 6
fmt.Println(sum(2)) // 3
```

When functions are passed/returned, their environment comes with them.

</details>

## Arrays

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/arrays)**

```go
// ARRAY
// [number]T
// A slice type has a specific length
// declare array
var arr [3]float64
fmt.Println(arr) // [0 0 0]

arr[1] = 23               // set element
element := arr[1]         // read element
fmt.Println(arr, element) // [0 23 0] 23

// declare and initialize
scores := [3]float64{9, 1.5, 2.2}
fmt.Println(scores)

// compiler figure out array length
arrNotMax := [...]int{2, 3, 4}
fmt.Println(arrNotMax, len(arrNotMax)) // [2 3 4] 3

// slice
fruits := [5]string{"banana", "pear", "apple", "orange", "peach"}
splicedFruits := fruits[1:3]              // [pear apple]
splicedFruits2 := fruits[2:]              // [apple orange peach]
removeLastFruit := fruits[:len(fruits)-1] // [banana pear apple orange]
lastFruit := fruits[len(fruits)-1]        // peach
fmt.Println(splicedFruits, splicedFruits2, removeLastFruit, lastFruit)
fmt.Println(len(splicedFruits)) // 2
fmt.Println(cap(splicedFruits)) // 4 (since starts from 1 and end index is 4)

// append
fruitsToAdd := append(splicedFruits, "cherry", "pineapple", "guava")
fmt.Println(splicedFruits, fruitsToAdd)             // [pear apple] [pear apple cherry pineapple guava]
fmt.Println(len(splicedFruits), cap(splicedFruits)) // 2 4
fmt.Println(len(fruitsToAdd), cap(fruitsToAdd)) // 5 8 (after crossing the previous capacity, the current capcity is doubled up)

// prepend
nums := []int{1,2,3}
nums = append([]int{0}, nums...)

// multidimensional array
multi := [2][3]int{{1, 2, 3}, {5, 6, 7}}
fmt.Println(multi) // [[1 2 3] [5 6 7]]
```

</details>

## Slices

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/slices)**

```go
// SLICE
// []T
// A slice type has no specific length

// declare a slice
var mySlice []int
fmt.Println(mySlice) // []

// mySlice[0] = 1 // occurs an error, since size is unknown
```

A slice has 3 properties:

- `ptr` - a pointer to the underlying array
- `len` - length of the slice - number of elements in the slice
- `cap` - capacity of the slice - length of the underlying array, which is also the maximum length the slice can take (until it grows)

![image](https://user-images.githubusercontent.com/11992095/202870508-0739d792-8747-4e20-8cd2-0ffa888d5c08.png)

source: https://gosamples.dev/capacity-and-length/

When we copy a slice it creates a new memeory location where it holds the same memory address of the underlying array, length & capacity. That's why, when modify the new copy of the slice, it also modify the old slice.

```go
var s = []int{1, 2, 3}
var s2 = s

s2[0] = 5

fmt.Println(s, s2) // [5 2 3], [5 2 3]
```

**Make**: make function "Initializes and allocates space in memory for a `slice`, `map`, or `channel`."

```go
// make([]T, len, cap)
s := make([]int, 0, 3)
sliceWithMake[0] = 1
fmt.Println(sliceWithMake)      // [1 0 0]

for i := 0; i < 5; i++ {
    s = append(s, i)
    fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
}
```

```sh
cap 3, len 1, 0xc0000b2000
cap 3, len 2, 0xc0000b2000
cap 3, len 3, 0xc0000b2000
cap 6, len 4, 0xc0000b8000 # larger capacity and a new pointer address
cap 6, len 5, 0xc0000b8000
```

- unpack/spread a slice

```go
var fruits = []string{"apple", "mango"}

// variable argument/vardiac function
func addFruits(fruitsToAdd ...string) []string {
 // unpack or spread
 updatedFruits := append(fruits, fruitsToAdd...)
 return updatedFruits
}

addFruits("banana", "pineapple") // [apple mango banana pineapple]
```

</details>

## Maps / Object / Dictionary

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/maps)**

1. Map operations

<details>
<summary>View codes</summary>

```go
var results map[string]float64 = make(map[string]float64) // create empty map

fmt.Println(results["test]) // 0

results["foyez"] = 3.4
results["mithu"] = 3.5

fmt.Println(results) // map[foyez:3.4 mithu:3.5]

// ***********************************************
userEmails := map[int]string{
 1: "user1@email.com",
 2: "user2@email.com",
}

userEmails[1] = "user12@email.com"
emailOfSecondUser, ok := userEmails[2]
emailOfFourthUser, ok2 := userEmails[4]

fmt.Println(userEmails)             // map[1:user12@email.com 2:user2@email.com]
fmt.Println(emailOfSecondUser, ok)  // user2@email.com true
fmt.Println(emailOfFourthUser, ok2) // false

if email, ok := userEmails[2]; ok {
 fmt.Printf("%s exists\n", email)
} else {
 fmt.Printf("%s doesn't exists\n", email)
}

delete(userEmails, 1)
fmt.Println(userEmails) // [2:user2@email.com]
```

</details>

2. Iterating map

<details>
<summary>View codes</summary>

```go
users := map[string]interface{}{
  "name":     "zayan",
  "age":      5,
  "religion": "islam",
}

for key, val := range users {
  fmt.Printf("%s -> %v\n", key, val) // name -> zayan, age -> 5, religion -> islam
}
```

</details>


</details>

## Strings

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/strings)**

```go
package main

import (
 "fmt"
 s "strings"
)

var p = fmt.Println

func main() {
 p(s.Contains("test", "es"))        // true
 p(s.Count("test", "t"))            // 2
 p(s.HasPrefix("test", "te"))       // true
 p(s.HasSuffix("test", "st"))       // true
 p(s.Index("test", "t"))            // 0
 p(s.LastIndex("test", "t"))        // 3
 p(s.Join([]string{"a", "b"}, "-")) // a-b
 p(s.Repeat("a", 5))                // aaaaa
 p(s.Replace("fooo", "o", "O", -1)) // fOOO
 p(s.Replace("fooo", "o", "O", 2))  // fOOo
 p(s.Split("a-b-c", "-"))           // [a b c]
 p(s.ToLower("TEST"))               // test
 p(s.ToUpper("test"))               // TEST
 p(len("hello"))                    // 5
 p("hello"[1])                      // 1
}
```

</details>

## Structs

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/structs)**

```go
type User struct {
 ID        int
 FirstName string
 LastName  string
 Email     string
}

user := User{ID: 1, FirstName: "Foyez", LastName: "Ahmed", Email: "foyez@email.com"}

fmt.Println(user.FirstName) // Foyez
```

</details>

## Pointers

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/pointers)**

**Pointer:** a variable that holds the **memory location** of a variable instead of a copy of its value.

```go
// Declare a pointer variable
var variableName *type

// Access to the variable address
&variableName

// Access to the variable value
*variableName
```

```go
type person struct {
 firstName      string
 lastName       string
 faboriteSports []string
}

func main() {
 person := person{
  firstName: "Foyez",
  lastName:  "Ahmed",
  faboriteSports: []string{"Cricket"}
 }

 updateFirstName(&person, "Rumon")
 fmt.Println(person) // {Foyez Ahmed [Cricket]}
 updateFavoriteSports(person, "Football")
 fmt.Println(person) // {Foyez Ahmed [Football]}
}

func updateFirstName(p *person, newFirstName string) {
 fmt.Println(p)  // &{Foyez Ahmed [Cricket]}
 fmt.Println(&p) // 0xc00000e028
 fmt.Println(*p) // {Foyez Ahmed [Cricket]}

 // (*p).firstName = newFirstName
 p.firstName = newFirstName
}

func updateFavoriteSports(p person, sportName string) {
 p.favoriteSports[0] = sportName
}
```

**Value Types:** `int`, `float`, `string`, `bool`, `structs`, `array`

> Have to use pointer to update these types of variables

**Reference Types:** `slices`, `maps`, `channels`, `pointers`, `functions`

> Don't need to use pointer to update these types of variables

**Call by Value:**

- Passed arguments or receiver are copied to parameters
- Modifying parameters or receiver has no effect outside of the function or the method

```go
type Person struct {
 name string
 age  int
}

func updateAge(p Person) {
 p.age = 20
 fmt.Println(p) // {Mithu 20}
}

func (p Person) updateAge() {
 p.age = 30
 fmt.Println(p) // {Mithu 30}
}

func main() {
 mithu := Person{name: "Mithu", age: 10}

 updateAge(mithu)
 fmt.Println(mithu) // {Mithu 10}

 mithu.updateAge()
 fmt.Println(mithu) // {Mithu 10}
}
```

**Call by Reference:**

- Pass pointer (memory location) as arguments or receiver
- Modifying parameters or receiver has effect outside of the function or the method

```go
type Person struct {
 name string
 age  int
}

func updateAge(p *Person) {
 p.age = 20
 fmt.Println(*p) // {Mithu 20}
}

func (p *Person) updateAge() {
 p.age = 30
 fmt.Println(*p) // {Mithu 30}
}

func main() {
 mithu := Person{name: "Mithu", age: 10}

 updateAge(&mithu)
 fmt.Println(mithu) // {Mithu 20}

 mithu.updateAge()
 fmt.Println(mithu) // {Mithu 30}
}
```

</details>

## Error Handling

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/errors)**

#### Error

- indicates that something bad happened, but it might be possible to continue running the program.
- i.e: A function that intentionally returns an error if something goes wrong

#### Panic

- happen at run time
- something happened that was fatal to the program and program stops execution
- ex: Trying to open a file that doesn't exist

```go
type error interface {
 Error() string
}

err := funcReturnError()
fmt.Println(err.Error())
panic(err.Error())
```

#### Defer

A defer statement defers the execution of a function until the surrounding function completes. Typically used for cleanup activities. Arguments of a deffered call are evaluted immediately.

```go
func main(){
 let country := "Bangladesh"

 defer fmt.Println(country)
 defer fmt.Println("love")
 country = "Australia"

 fmt.Println("I")
}

// I
// love
// Bangladesh
```

#### Recover

- **Panic** is called during a run time error and fatally kill the program
- **Recover** tells Go what to do when a panic happens (returns what was passed to panic)
- Recover must be paired with **defer**, which will fire even after a panic

```go
func recoverFromPanic() {
 if r := recover(); r != nil {
  fmt.Println(r)
 }
}

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

</details>

## Methods

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/methods)**
Syntax of method

```go
func (r ReceiverType) funcName(parameters) (results)
```

#### Methods vs Functions

- The difference between a method and a function is that instead of accepting an argument as struct, we're calling a method on an instance of that struct.

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

func main() {
 user := User{
  name: "Manam",
  age: 25,
  address: address{
   email: "manam@email.com",
   zipCode: 34000
  },
 }

 updateUserName(&user, "Chayon")

 // (&user).UpdateName("Chayon")
 user.UpdateName("Chayon")
}

func updateUserName(u *User, name string) {
 u.name = name
}

// func (receiverName ReceiverType) MethodName(args)
// When a method is called on a variable of that type,
// we get the reference to its data via the receiverName variable.
func (u *User) UpdateName(name string) {
 // (*u).name = name
 u.name = name
}
```

**When should we make the pointer receiver type of a method?**

1. When the receiver type uses a large amount of memory, otherwise the receiver will be copied with a large amount of data which is costly.
2. When the method must modify the data in the object of the receiver type.

**Good practices:**

1. All methods of a type should have pointer receivers, or
2. All methods of a type should have non-pointer receivers

</details>

## Type Assertions

<details>
<summary>View contents</summary>

> Type assertions is used to assert the type of a given variable. It provides access to an interface value's underlying concrete value.

```go
// assertedVariable, ok := variable.(Type)

var foo interface{} = "Hello"

str := foo.(string)
fmt.Println(str) // "Hello"

num := foo.(int) // panic
fmt.Println(num)

num2, ok := foo.(int)
fmt.Println(num2, ok) // 0, false
```

</details>

## Interfaces

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/interfaces)**

**Structs:** define a set of attributes on a type, e.g.: a user has a `FirstName` and a `LastName`, it is the type of User.

**Interfaces:** define a set of method signatures (name, parameters & return types), NOT the implementation.

```go
type Shape2D interface {
 Area() float64
 Perimeter() float64
}

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

func fitInYard(s Shape2D) bool {
 return s.Area() > 200 && s.Perimeter() > 200
}

func printShapeProps(s Shape2D) {
 if rect, ok := s.(Rectangle); ok {
  fmt.Printf("Height: %.2f, Width: %.2f\n", rect.Height, rect.Width)
 }
 if circle, ok := s.(Circle); ok {
  fmt.Printf("Radius: %.2f\n", circle.Radius)
 }
}

func main() {
 circle := Circle{10}
 rectangle := Rectangle{10, 20}

 fmt.Println(fitInYard(circle))
 fmt.Println(fitInYard(rectangle))

 printShapeProps(rectangle) // Height: 20.00, Width: 10.00
 printShapeProps(circle) // Radius: 10.00
}
```

#### Empty Interface

```go
interface{}
```

- Specifies zero methods
- An empty interface may hold values of any type
- Like _any_ type in Typescript

```go
var people map[string]interface{} = make(map[string]interface{})

people["name"] = "Foyez"
people["age"] = 28

fmt.Printf("%#v %T\n", people["name"], people["name"]) // "Foyez" string
fmt.Printf("%#v %T", people["age"], people["age"])     // 28 int
```

</details>
 
## Generics
 
<details>
<summary>View contents</summary>
 
Suppose, we write a function that accepts string or integer as arguments
 
```go
func isEqual(a, b interface{}) bool {
 return a == b
}

func main() {
fmt.Println(isEqual(1, 1)) // true
fmt.Println(isEqual(1, "1")) // true
}

```

Here, though the empty interface `interface{}` gives us the flexibility to pass string or integer type, it don't provide us type-safety. Because we can't compare a number with a string. This means the compiler can't help us and we're instead more likely to have runtime errors.

To solve this issue, we can use generics which give us flexibility and type-safety at the same time.

```go
func isEqual[T comparable](a, b T) bool {
 return a == b
}

func main() {
 fmt.Println(isEqual(1, 1))   // true
 fmt.Println(isEqual(1, "1")) // default type string of "1" does not match inferred type int for T
}
```

> `comparable` is a built-in Go constraint introduced in Go 1.18. It is used to denote types that can be compared for equality using == and !=.

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
```

> The `constraints.Ordered` constraint comes from the `golang.org/x/exp/constraints` package and is used to denote types that support ordering operations like <, <=, >, and >=.

**Implementation of `reduce()`, `find()`, `filter()` & `map()`:**

```go
func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
 var result = initialValue
 for _, x := range collection {
  result = accumulator(result, x)
 }
 return result
}

func Find[A any](items []A, predict func(A) bool) (value A, found bool) {
 for _, v := range items {
  if predict(v) {
   return v, true
  }
 }
 return
}

func Filter[A any](items []A, predict func(A) bool) []A {
 var founds []A

 for _, v := range items {
  if predict(v) {
   founds = append(founds, v)
  }
 }

 return founds
}

func Map[A, B any](items []A, modify func(A) B) []B {
 var modifiedItems []B
 for _, v := range items {
  modifiedItems = append(modifiedItems, modify(v))
 }
 return modifiedItems
}
```

Reference: [Golang Generics Are Here! - Golang Beyond the Basics](https://www.youtube.com/watch?v=P2CQWeZZ--4)

</details>

## Concurrency

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/concurrency)**

### Goroutines

- A **Goroutine** is a lightweight thread manged by the Go runtime
- Implemented by adding the `go` keyword before executing a function
- Tells go to spin up a new thread to do that thing

```go

```

</details>

## Web Servers

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

</details>

## Working with files

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

### Video Tutorials

- [Go Tutorals - NerdCademy](https://youtube.com/playlist?list=PLujhHB_uAFJws6Vv5q1KDoaQ4YcpS9UOm)
- [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I) - `TechWorld with Nana`
- [Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU)
- [Learn Go Programming by Building 11 Projects – Full Course](https://www.youtube.com/watch?v=jFfo23yIWac)
- [Backend Master Class [Go + Postgres Docker + Kubernetes + gRPC]](https://youtube.com/playlist?list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE)

</details>
