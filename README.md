# Golang

**[Go Playground](https://play.golang.org/)**

- Programming Language developed by Google in 2007
- Open-sourced in 2009

## Why Go

<details>
<summary>View contents</summary>

1. Run on multiple cores and builtin to support concurrency
2. Fast compile times
3. Ease to development
4. Fast execution
5. Automatic garbage collection

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
// Every go program needs at least one package main
// Go programs are organized into packages
// A package is a collection of source files
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

```
BenchmarkHello   	2000000000	         0.46 ns/op

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

```
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

## Printing with fmt

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/printing)**

#### **Print**

```
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
fmt.Printf("Hello, %v\n", name)
```

```sh
Hello, Zohan
Hello, Zohan
Hello, Zohan
```

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

#### Identify and convert type

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

	for _, age := range ages {
		total += age
	}

	return float64(total) / float64(len(ages))
}

func main() {
	fmt.Println(average(10, 20, 32))
}
```

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

// create a slice using make function
// dynamically-sized arrays
// make([]T, len, cap)
sliceWithMake := make([]int, 3, 10)
fmt.Println(sliceWithMake)      // [0 0 0]
fmt.Println(len(sliceWithMake)) // 3
fmt.Println(cap(sliceWithMake)) // 5
```

**Make**: make function "Initialize and allocates space in memory for a slice, map, or channel."
	
```go
s := make([]int, 0, 3)
	
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

```go
var fruits = []string{"apple", "mango"}

// varArg
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

```go
var results map[string]float64 = make(map[string]float64) // create empty map

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
	
**Value Types:** `int`, `float`, `string`, `bool`, `structs`
> Have to use pointer to update these types of variables
	
**Reference Types:** `slices`, `maps`, `channels`, `pointers`, `functions`
> Don't need to use pointer to update these types of variables

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

A defer statement defers the execution of a function until the surrounding function returns.

```go
func main(){
	defer fmt.Println("Bangladesh")
	defer fmt.Println("love")
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
	
Good practices:
1. All methods of a type should have pointer receivers, or
2. All methods of a type should have non-pointer receivers

</details>

## Interfaces

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/interfaces)**

**Structs:** define a set of attributes on a type, e.g.: a user has a `FirstName` and a `LastName`, it is the type of User.

**Interfaces:** describe a set of behaviors that also define a type, e.g. a user can change the `FirstName` is a type of that interface.

```go
// Describer has a Describe method
type Human interface {
	Describe() string
}

// Any struct that has a method called describe is also a type of Describer

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

func (u User) Describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (g *Group) Describe() string {
	if len(g.users) > 2 {
		g.spaceAvailable = false
	}

	desc := fmt.Sprintf("Users: %d, Newest User: %s %s, Accepting New User: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Foyez", LastName: "Ahmed", Email: "foyez@email.com"}
	user2 := User{ID: 2, FirstName: "Manam", LastName: "Ahmed", Email: "manam@email.com"}

	group := Group{
		role:           "admin",
		users:          []User{user, user2},
		newestUser:     user2,
		spaceAvailable: true,
	}

	describeSomething := func(human Human) {
		desc := human.Describe()
		fmt.Println(desc)
	}

	describeSomething(user) // Name: Foyez Ahmed, Email: foyez@email.com
	describeSomething(&group) // Users: 2, Newest User: Manam Ahmed, Accepting New User: true
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

## Concurrency

<details>
<summary>View contents</summary>

**[You can find all the code for this section here](https://github.com/foyez/go/tree/main/codes/concurrency)**

#### Goroutines

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
- [Go Doc](https://golang.org/doc/)
- [Go Blog](https://blog.golang.org/)
- [Clean Go Article](https://github.com/Pungyeon/clean-go-article)
- [How To Code in Go](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go)

### Video Tutorials

- [Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I) - `TechWorld with Nana`
- [Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU)
- [Learn Go Programming by Building 11 Projects – Full Course](https://www.youtube.com/watch?v=jFfo23yIWac)

</details>
