# 5. Functions

> Master functions, methods, closures, and advanced function patterns in Go.

**Navigation:** [← Control Flow](4-control-flow.md) | [README](./README.md) | [Next: Data Structures →](6-data-structures.md)

---

## Table of Contents

- [5.1 Function Basics](#51-function-basics)
- [5.2 Multiple Return Values](#52-multiple-return-values)
- [5.3 Variadic Functions](#53-variadic-functions)
- [5.4 Closures & Anonymous Functions](#54-closures--anonymous-functions)
- [5.5 Methods](#55-methods)
- [Practice Questions](#practice-questions)

---

## 5.1 Function Basics

Functions are reusable blocks of code that perform a specific task.

### Function Declaration

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

// Multiple parameters, different types
func greet(name string, age int) {
	fmt.Printf("Hello %s, you are %d years old\n", name, age)
}
```

### Calling Functions

```go
sayHello()                    // No return value
result := add(5, 3)          // Store return value
multiply(4, 5)               // Ignore return value
greet("Alice", 30)           // Multiple parameters
```

### Function Scope

```go
var n1 = 5  // Package scope

func foo(n2 int) {
	n3 := 8  // Function local variable
	fmt.Println(n1, n2, n3)
}

// n1 - accessible everywhere in package
// n2 - accessible only in foo
// n3 - accessible only in foo
```

---

## 5.2 Multiple Return Values

Go functions can return **multiple values** - a unique feature!

### Basic Multiple Returns

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
	log.Fatal(err)
}
fmt.Println(result)  // 5
```

### Common Pattern: (value, error)

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

### Multiple Values

```go
func getUserInfo(id int) (string, int, bool) {
	return "Alice", 30, true
}

name, age, active := getUserInfo(1)
fmt.Printf("Name: %s, Age: %d, Active: %v\n", name, age, active)
```

### Ignoring Return Values

```go
// Ignore error (NOT recommended in production)
data, _ := os.ReadFile("file.txt")

// Ignore all but one
_, _, active := getUserInfo(1)
```

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

## 5.3 Variadic Functions

**Accept unlimited arguments:**

```go
func sum(nums ...int) int {
	total := 0
	// nums is treated as []int slice
	for _, n := range nums {
		total += n
	}
	return total
}

// Usage
fmt.Println(sum(1, 2, 3))        // 6
fmt.Println(sum(1, 2, 3, 4, 5))  // 15

// Spread slice into variadic function
values := []int{1, 2, 3}
fmt.Println(sum(values...))      // 6
```

### Real-World Example

```go
import "fmt"

func logError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Println("[ERROR]", msg)
}

logError("User %s not found", "Alice")
logError("Invalid input: %d", 42)
```

### Rules

- Variadic parameters are treated as slices
- Variadic parameter must be **last**
- Only **one** variadic parameter per function

```go
// ✅ Valid
func myFunc(a int, b string, nums ...int) {}

// ❌ Invalid
func myFunc(nums ...int, a int) {}  // Must be last

// ❌ Invalid
func myFunc(nums1 ...int, nums2 ...int) {}  // Only one
```

---

## 5.4 Closures & Anonymous Functions

### Functions as Values

**Functions are first-class citizens:**

```go
// Assign function to variable
var add func(int, int) int = func(a, b int) int {
	return a + b
}

result := add(3, 4)  // 7
```

### Anonymous Functions

```go
// Declare and call immediately
func() {
	fmt.Println("Anonymous function")
}()

// With parameters
func(name string) {
	fmt.Println("Hello,", name)
}("Alice")
```

### Passing Functions as Arguments

```go
func apply(f func(int, int) int, a, b int) int {
	return f(a, b)
}

multiply := func(a, b int) int {
	return a * b
}

result := apply(multiply, 3, 4)  // 12
```

### Returning Functions

```go
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// "factor" is in the closure of returned function

double := multiplier(2)
triple := multiplier(3)

fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

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

// "count" is in the closure of returned function

c1 := counter()
fmt.Println(c1())  // 1
fmt.Println(c1())  // 2
fmt.Println(c1())  // 3

c2 := counter()
fmt.Println(c2())  // 1 (separate counter)
```

**Explanation:** `count` remains available even after `counter` finishes execution.

### Real-World Examples

**HTTP middleware:**
```go
import "net/http"

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
```

**Functional programming patterns:**
```go
// Map
func Map(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Usage
nums := []int{1, 2, 3, 4, 5}
squared := Map(nums, func(n int) int {
	return n * n
})
fmt.Println(squared)  // [1 4 9 16 25]
```

### Closure Gotcha (Loop Variables)

**❌ Wrong:**
```go
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)  // May print 5, 5, 5, 5, 5
	}()
}
```

**✅ Correct:**
```go
for i := 0; i < 5; i++ {
	go func(n int) {
		fmt.Println(n)  // Prints 0, 1, 2, 3, 4
	}(i)
}

// Or capture in local variable
for i := 0; i < 5; i++ {
	i := i  // Create new variable
	go func() {
		fmt.Println(i)  // Prints 0, 1, 2, 3, 4
	}()
}
```

---

## 5.5 Methods

A **method** is a function with a **receiver**.

### Method Declaration

**Syntax:**
```go
func (receiver ReceiverType) methodName(parameters) returnType {
	// method body
}
```

### Methods vs Functions

```go
type Rectangle struct {
	Width  float64
	Height float64
}

// Function version
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

// Method version
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Usage
rect := Rectangle{Width: 10, Height: 5}

// Function call
area1 := Area(rect)

// Method call
area2 := rect.Area()
```

### Value Receiver

```go
type Counter struct {
	Value int
}

// Value receiver - receives a copy
func (c Counter) GetValue() int {
	return c.Value
}

func (c Counter) Increment() {
	c.Value++  // Modifies copy, not original!
}

counter := Counter{Value: 0}
counter.Increment()
fmt.Println(counter.GetValue())  // 0 (unchanged!)
```

### Pointer Receiver

```go
// Pointer receiver - modifies original
func (c *Counter) Increment() {
	c.Value++  // Modifies original
}

counter := Counter{Value: 0}
counter.Increment()  // Go auto-converts to (&counter).Increment()
fmt.Println(counter.GetValue())  // 1 (modified!)
```

### When to Use Pointer Receivers

**Use pointer receivers when:**
- Method needs to modify the receiver
- Receiver is large (avoid copying)
- Consistency (if any method uses pointer, use for all)

**Use value receivers when:**
- Receiver is small
- Receiver is immutable
- Receiver is a value type (int, bool, etc.)

**Decision Matrix:**

| Scenario | Receiver Type |
|----------|---------------|
| Need to modify receiver | Pointer |
| Large struct | Pointer |
| Small struct, no modification | Value |
| Consistency (some methods use pointer) | Pointer (all) |

**Example:**
```go
type User struct {
	Name  string
	Email string
	// Large data fields...
}

// ✅ Consistent - all methods use pointer receivers
func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) GetName() string {
	return u.Name  // Still use pointer for consistency
}
```

### Methods on Non-Struct Types

```go
type MyInt int

func (m MyInt) IsEven() bool {
	return m%2 == 0
}

var num MyInt = 10
fmt.Println(num.IsEven())  // true
```

### Methods vs Functions - When to Use

**Use methods when:**
- Operation is specific to a type
- Building an API for a type
- Implementing interfaces

**Use functions when:**
- Operation works on multiple types
- Utility/helper functions
- Package-level operations

### Real-World Example

```go
type BankAccount struct {
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > b.Balance {
		return errors.New("insufficient funds")
	}
	b.Balance -= amount
	return nil
}

func (b BankAccount) GetBalance() float64 {
	return b.Balance
}

// Usage
account := &BankAccount{Balance: 1000}
account.Deposit(500)
account.Withdraw(200)
fmt.Printf("Balance: $%.2f\n", account.GetBalance())
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. Functions in Go can return __________ values.
2. A __________ receiver is used when a method needs to modify the struct.
3. Variadic parameters are treated as __________ inside the function.
4. Named return values allow __________ returns.
5. Functions are __________ citizens in Go.
6. A method is a function with a __________.

### True/False

1. ⬜ Named return values are always better than explicit returns
2. ⬜ A function can have multiple variadic parameters
3. ⬜ Value receivers receive a copy of the struct
4. ⬜ Go automatically converts value to pointer when calling methods
5. ⬜ Closures can capture variables from outer scope
6. ⬜ The variadic parameter must be the last parameter

### Multiple Choice

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

3. What is the output?
   ```go
   func counter() func() int {
       count := 0
       return func() int {
           count++
           return count
       }
   }
   c := counter()
   fmt.Println(c(), c())
   ```
   - A) 0 0
   - B) 1 1
   - C) 1 2
   - D) 0 1

4. Which is valid?
   - A) `func add(a, b int) int`
   - B) `func add(a int, b int) int`
   - C) Both A and B
   - D) Neither

### Code Challenge

Write a function `filter` that takes a slice of integers and a filter function, then returns a new slice with only elements that pass the filter.

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. multiple
2. pointer
3. slices
4. naked
5. first-class
6. receiver

**True/False:**
1. ❌ False (use named returns sparingly)
2. ❌ False (only one, and it must be last)
3. ✅ True
4. ✅ True
5. ✅ True
6. ✅ True

**Multiple Choice:**
1. **C** - 312 (prints 3, then defers execute in reverse: 2, then 1)
2. **C** - When you need to modify the struct (also for large structs)
3. **C** - 1 2 (count is captured and persists across calls)
4. **C** - Both A and B (shortened syntax is allowed)

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

greaterThan3 := filter(nums, func(n int) bool { return n > 3 })
fmt.Println(greaterThan3)  // [4 5 6]
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ Function declaration and calling  
✅ Multiple return values (common pattern: value, error)  
✅ Named return values and naked returns  
✅ Variadic functions (unlimited arguments)  
✅ Functions as first-class citizens  
✅ Closures and anonymous functions  
✅ Methods with value and pointer receivers  
✅ When to use methods vs functions  

**Next Steps:**
- Continue to [Chapter 6: Data Structures →](./6-data-structures.md)
- Practice writing functions with multiple return values
- Implement the filter function challenge above

---

**Navigation:** [← Control Flow](4-control-flow.md) | [README](./README.md) | [Next: Data Structures →](6-data-structures.md)