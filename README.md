# Golang

## Why Go

<details>
<summary>View contents</summary>

1. Fast compile times
2. Ease to development
3. Fast execution
4. Automatic garbage collection

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

## Anatomy of a go file

<details>
<summary>View contents</summary>

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

</details>

## Create a go project

<details>
<summary>View contents</summary>

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

	// variables in only function
	funcVar := "can't declare outside of a function"

	// multiple variables
	one, two := 1, "two"

	fmt.Println(name, age, salary)
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

// multidimensional array
multi := [2][3]int{{1, 2, 3}, {5, 6, 7}}
fmt.Println(multi) // [[1 2 3] [5 6 7]]
```

</details>
