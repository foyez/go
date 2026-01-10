# 2. Fundamentals

> Master Go's core building blocks: packages, variables, data types, and operators.

**Navigation:** [← Introduction](1-introduction.md) | [README](./README.md) | [Next: Strings & I/O →](3-strings-io.md)

---

## Table of Contents

- [2.1 Package Structure](#21-package-structure)
- [2.2 Variables & Constants](#22-variables--constants)
- [2.3 Data Types](#23-data-types)
- [2.4 Type Conversion](#24-type-conversion)
- [2.5 Operators](#25-operators)
- [Practice Questions](#practice-questions)

---

## 2.1 Package Structure

A **package** in Go is a way to **organize code**. It groups related variables, functions, types, and constants together.

**Think of a package like a folder of related code.**

### Package Basics

**Examples of built-in packages:**
- `fmt` → printing and formatting
- `math` → math functions
- `net/http` → web servers and HTTP

**Every Go file starts with a package declaration:**

```go
package main  // Executable package (must have main function)

package mylib  // Library package
```

### Package Naming Rules (Google Style)

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

### Special Package: `main`

- `package main` is used for **executable programs**
- It must contain a `main()` function
- Other packages are **reusable libraries**

### Importing Packages

```go
package main

import (
	"fmt"           // Standard library
	"net/http"      // Standard library (nested)
	
	"github.com/gorilla/mux"  // External package
	"example.com/myproject/internal/util"  // Internal package
)
```

### Import Aliases

```go
import (
	"fmt"
	str "strings"  // Alias to avoid naming conflicts
)

func main() {
	str.ToUpper("hello")
}
```

### Import Side Effects (rare)

```go
import _ "image/png"  // Registers PNG decoder without using package
```

### Creating Your Own Package

**Example project structure:**

```
myproject/
│
├── go.mod
├── main.go
└── mathutils/
    └── mathutils.go
```

**Step 1: Create the package file**

**mathutils/mathutils.go:**

```go
package mathutils

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}
```

### Important Rules

- The folder name **does not have to** match the package name, but it usually should
- Functions starting with a **capital letter** are **exported** (public)
- Lowercase names are **private** to the package

**Step 2: Importing and using a package**

**go.mod:**

```go
module myproject

go 1.22
```

**main.go:**

```go
package main

import (
	"fmt"
	"myproject/mathutils"
)

func main() {
	fmt.Println(mathutils.Add(5, 3))      // 8
	fmt.Println(mathutils.Subtract(10, 4)) // 6
}
```

### Key Points

- Import path = **module name + folder name**
- Exported functions must start with a **capital letter**
- Use `packageName.FunctionName`

---

## 2.2 Variables & Constants

Variables store values that can be used and modified during program execution.

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
var sl []int     // nil
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
	Wednesday      // 3
	Thursday       // 4
	Friday         // 5
	Saturday       // 6
)
```

**Why use constants:**
- Prevent accidental modification
- Improve code readability
- Enable compiler optimizations

**Real-World Example with iota:**
```go
const (
	_ = iota // Ignore 0
	KB = 1 << (10 * iota)  // 1 << 10 = 1024
	MB                      // 1 << 20 = 1048576
	GB                      // 1 << 30 = 1073741824
)

fmt.Println(KB)  // 1024
fmt.Println(MB)  // 1048576
fmt.Println(GB)  // 1073741824
```

---

## 2.3 Data Types

Go is a **statically typed language**, meaning variable types are known at compile time.

### Basic Types

```go
// Integers (Signed)
int, int8, int16, int32, int64

// Integers (Unsigned)
uint, uint8, uint16, uint32, uint64

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

### Type Summary Table

| Category | Type Names | Examples |
|----------|-----------|----------|
| **INTEGER** | int, int8, int16, int32, int64<br/>uint, uint8, uint16, uint32, uint64 | var age int = 20 |
| **FLOAT** | float32, float64 | var gpa float64 = 3.4 |
| **STRING** | string | var fruit string = "mango" |
| **BOOLEAN** | bool | var adult bool = age > 18 |

### Platform-Dependent Types

- `int` and `uint` are **32 bits** on 32-bit systems, **64 bits** on 64-bit systems
- Most common: use `int` for integers and `float64` for decimals

### Boolean Operators

```go
&&  // AND
||  // OR
!   // NOT
< <= > >= == !=
```

```go
var adult bool = age >= 18
var canVote bool = adult && age >= 18
```

### Identify Type

```go
import "reflect"

reflect.TypeOf(6)        // int
reflect.TypeOf(3.14)     // float64
reflect.TypeOf("hello")  // string
```

**Note:** Useful for debugging and learning, not common in production code.

### Integer Types Details

```go
var i8 int8 = 127       // -128 to 127
var i16 int16 = 32767   // -32768 to 32767
var i32 int32 = 2147483647
var i64 int64 = 9223372036854775807

var u8 uint8 = 255      // 0 to 255
var u16 uint16 = 65535  // 0 to 65535
```

**When to use specific sizes:**
- `int8`, `int16`, etc. → Optimize memory (rare)
- `int` → Default choice (most common)
- `uint` → Non-negative numbers

### Float Types Details

```go
var f32 float32 = 3.14
var f64 float64 = 3.141592653589793

// Float precision
fmt.Printf("%.2f\n", 3.14159)  // 3.14
fmt.Printf("%.5f\n", 3.14159)  // 3.14159
```

**Default:** Use `float64` unless memory is a concern.

### Rune (Unicode Code Point)

```go
var r rune = 'A'         // 65 (ASCII value)
var r2 rune = '世'        // 19990 (Unicode)

fmt.Printf("%c\n", r)    // A
fmt.Printf("%c\n", r2)   // 世
```

### Byte vs Rune

```go
var b byte = 'A'   // uint8, 0-255 (ASCII)
var r rune = '世'   // int32, Unicode code point
```

---

## 2.4 Type Conversion

**Go requires explicit conversion** - no implicit conversions!

### Basic Conversions

```go
var i int = 42
var f float64 = float64(i)  // Must convert explicitly
var u uint = uint(f)

// ❌ This won't compile:
// var f float64 = i
```

### String Conversions

```go
import "strconv"

// Int to String
i := 42
s := strconv.Itoa(i)        // "42"

// String to Int
s := "42"
i, err := strconv.Atoi(s)   // 42, nil
if err != nil {
	// Handle error
}

// Float to String
f := 3.14
s := strconv.FormatFloat(f, 'f', 2, 64)  // "3.14"

// String to Float
s := "3.14"
f, err := strconv.ParseFloat(s, 64)  // 3.14, nil
```

### Real-World Example

```go
func calculateAverage(nums []int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))  // Must convert to float64
}

nums := []int{90, 85, 95}
avg := calculateAverage(nums)
fmt.Printf("Average: %.2f\n", avg)  // Average: 90.00
```

### Type Assertion (for interfaces)

```go
var i interface{} = "hello"

// Unsafe (panics if wrong type)
s := i.(string)
fmt.Println(s)  // "hello"

// Safe (comma-ok idiom)
s, ok := i.(string)
if ok {
	fmt.Println(s)  // "hello"
}

n, ok := i.(int)
fmt.Println(n, ok)  // 0 false
```

---

## 2.5 Operators

### Arithmetic Operators

```go
a := 10
b := 3

fmt.Println(a + b)  // 13 (Addition)
fmt.Println(a - b)  // 7  (Subtraction)
fmt.Println(a * b)  // 30 (Multiplication)
fmt.Println(a / b)  // 3  (Division - integer)
fmt.Println(a % b)  // 1  (Modulus)

// Floating-point division
f1 := 10.0
f2 := 3.0
fmt.Println(f1 / f2)  // 3.3333...
```

### Comparison Operators

```go
a := 10
b := 20

fmt.Println(a == b)  // false (Equal to)
fmt.Println(a != b)  // true  (Not equal to)
fmt.Println(a < b)   // true  (Less than)
fmt.Println(a > b)   // false (Greater than)
fmt.Println(a <= b)  // true  (Less than or equal)
fmt.Println(a >= b)  // false (Greater than or equal)
```

### Logical Operators

```go
a := true
b := false

fmt.Println(a && b)  // false (AND)
fmt.Println(a || b)  // true  (OR)
fmt.Println(!a)      // false (NOT)
```

### Assignment Operators

```go
a := 10

a += 5   // a = a + 5  → 15
a -= 3   // a = a - 3  → 12
a *= 2   // a = a * 2  → 24
a /= 4   // a = a / 4  → 6
a %= 4   // a = a % 4  → 2
```

### Increment/Decrement

```go
a := 10

a++  // a = a + 1  → 11
a--  // a = a - 1  → 10

// ❌ NOT allowed:
// b := a++  // Compile error
// ++a       // Compile error (only postfix)
```

**Note:** Go only supports **postfix** increment/decrement.

### Bitwise Operators

```go
a := 5   // 0101 in binary
b := 3   // 0011 in binary

fmt.Println(a & b)   // 1  (AND)     0101 & 0011 = 0001
fmt.Println(a | b)   // 7  (OR)      0101 | 0011 = 0111
fmt.Println(a ^ b)   // 6  (XOR)     0101 ^ 0011 = 0110
fmt.Println(a << 1)  // 10 (Left shift)   0101 << 1 = 1010
fmt.Println(a >> 1)  // 2  (Right shift)  0101 >> 1 = 0010
```

### Operator Precedence

From **highest to lowest**:

```
1. Unary:       !  +  -  ^  *  &  <-
2. Multiply:    *  /  %  <<  >>  &  &^
3. Add:         +  -  |  ^
4. Compare:     ==  !=  <  <=  >  >=
5. Logical AND: &&
6. Logical OR:  ||
```

**Example:**
```go
result := 2 + 3 * 4  // 14 (not 20)
// Multiplication has higher precedence

result := (2 + 3) * 4  // 20
// Parentheses override precedence
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. A package name should be __________, single-word, and match the directory name.
2. Variables starting with an __________ letter are exported (public).
3. The __________ value for an uninitialized int is 0.
4. The `%__` format verb prints the type of a variable.
5. Go requires __________ type conversion between int and float64.
6. The `:=` operator can only be used __________ functions.

### True/False

1. ⬜ Go allows implicit type conversion between int and float64
2. ⬜ Constants can be modified after declaration
3. ⬜ The `int` type size depends on the platform (32-bit or 64-bit)
4. ⬜ Variables declared with `:=` can only be used inside functions
5. ⬜ `iota` is used for auto-incrementing constant values
6. ⬜ Go supports both prefix (++a) and postfix (a++) increment

### Multiple Choice

1. What is the zero value of a string?
   - A) nil
   - B) 0
   - C) ""
   - D) undefined

2. Which operator has the highest precedence?
   - A) +
   - B) *
   - C) ==
   - D) &&

3. How do you convert an int to a string?
   - A) string(42)
   - B) strconv.Itoa(42)
   - C) fmt.Sprintf("%d", 42)
   - D) Both B and C

4. What does `const _ = iota` do?
   - A) Creates a constant with value 0
   - B) Skips the value 0
   - C) Causes an error
   - D) Resets iota

### Code Output

```go
const (
	A = iota
	B
	C = iota * 2
	D
)
fmt.Println(A, B, C, D)
```

What is printed?

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. lowercase
2. uppercase
3. zero
4. T
5. explicit
6. inside

**True/False:**
1. ❌ False (explicit conversion required)
2. ❌ False (constants are immutable)
3. ✅ True (32-bit on 32-bit systems, 64-bit on 64-bit systems)
4. ✅ True
5. ✅ True
6. ❌ False (only postfix a++ is supported)

**Multiple Choice:**
1. **C** - "" (empty string)
2. **B** - * (multiplication has higher precedence than +)
3. **D** - Both B and C work (strconv.Itoa is more efficient)
4. **B** - Skips the value 0 (blank identifier)

**Code Output:**
```
0 1 4 6
```

Explanation:
- A = 0 (iota starts at 0)
- B = 1 (iota auto-increments)
- C = 4 (iota is 2, multiplied by 2)
- D = 6 (iota is 3, multiplied by 2)

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ How to organize code with packages  
✅ Package naming conventions and visibility rules  
✅ Variable declaration methods (var, :=)  
✅ Zero values and constants  
✅ All Go data types (int, float, string, bool, etc.)  
✅ Type conversion requirements  
✅ Operators (arithmetic, comparison, logical, bitwise)  

**Next Steps:**
- Continue to [Chapter 3: Strings & I/O →](3-strings-io.md)
- Practice creating packages and using different data types
- Complete the practice questions above

---

**Navigation:** [← Introduction](1-introduction.md) | [README](./README.md) | [Next: Strings & I/O →](3-strings-io.md)