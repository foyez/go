# 4. Control Flow

> Master conditional statements, loops, and control flow in Go.

**Navigation:** [← Strings & I/O](3-strings-io.md) | [README](./README.md) | [Next: Functions →](5-functions.md)

---

## Table of Contents

- [4.1 if-else Statements](#41-if-else-statements)
- [4.2 switch Statements](#42-switch-statements)
- [4.3 Loops](#43-loops)
- [4.4 defer Statement](#44-defer-statement)
- [Practice Questions](#practice-questions)

---

## 4.1 if-else Statements

Go uses `if`, `else if`, and `else` for conditional logic.

### Basic Syntax

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

### Examples

```go
age := 20

if age >= 18 {
	fmt.Println("Adult")
} else {
	fmt.Println("Minor")
}

// Multiple conditions
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

### Early Return Pattern (Google Style)

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

### Guard Clauses

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")  // Guard clause
	}
	return a / b, nil
}
```

---

## 4.2 switch Statements

Go's `switch` is **more powerful and flexible** than many other languages.

### Basic Switch

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

### Switch Without Expression (replaces if-else chains)

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

### Switch with Initialization

```go
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
	fmt.Println("Weekend")
default:
	fmt.Println("Weekday")
}
```

### Type Switch (checking interface types)

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

### Fallthrough (explicit fall-through)

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

**Note:** `fallthrough` is rarely used. Default no-fallthrough is safer.

---

## 4.3 Loops

Go has **only one loop keyword**: `for`. But it supports multiple patterns.

### Traditional for Loop

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
// Output: 0 1 2 3 4
```

**Components:**
- Initialization: `i := 0`
- Condition: `i < 5`
- Post statement: `i++`

### While-style Loop

```go
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}
```

### Infinite Loop

```go
num := 1

for {
	num = num + 2

	if num == 5 {
		continue
	}

	fmt.Println(num)

	if num == 7 {
		break
	}
}
```

**Loop control:**
- `continue` - Skip to next iteration
- `break` - Exit loop

### Range-based Loop (for Slices/Arrays)

```go
nums := []int{10, 20, 30}

// Index and value
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

### Range-based Loop (for Maps)

```go
scores := map[string]int{"Alice": 90, "Bob": 85}

// Key and value
for name, score := range scores {
	fmt.Printf("%s: %d\n", name, score)
}

// Only keys
for name := range scores {
	fmt.Println(name)
}
```

**Important:** Map iteration order is **random**!

### Range-based Loop (for Strings)

```go
// Iterates over runes (Unicode characters)
for i, r := range "Hello, 世界" {
	fmt.Printf("%d: %c\n", i, r)
}
```

### Range-based Loop (for Channels)

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

**Note:** `range` receives values until channel is closed.

### Basic for Loop (Go 1.22+)

```go
// Simple iteration
for i := range 5 {
	fmt.Println(i)
}
// Output: 0 1 2 3 4
```

### Loop Control

**break:**
```go
for i := 0; i < 10; i++ {
	if i == 5 {
		break  // Exit loop
	}
	fmt.Println(i)
}
// Output: 0 1 2 3 4
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
// Output: (0,0) (0,1) (0,2) (1,0)
```

### Common Loop Patterns

**Iterate backwards:**
```go
for i := 10; i >= 0; i-- {
	fmt.Println(i)
}
```

**Skip by 2:**
```go
for i := 0; i < 10; i += 2 {
	fmt.Println(i)  // 0 2 4 6 8
}
```

**Iterate with condition:**
```go
for i := 0; i*i < 100; i++ {
	fmt.Println(i)
}
```

### Real-World Examples

**Processing items:**
```go
users := []User{
	{Name: "Alice", Age: 30},
	{Name: "Bob", Age: 25},
	{Name: "Carol", Age: 35},
}

for _, user := range users {
	if user.Age < 30 {
		continue  // Skip users under 30
	}
	
	if err := processUser(user); err != nil {
		log.Printf("Error processing %s: %v", user.Name, err)
		continue  // Skip to next user
	}
	
	fmt.Printf("Processed: %s\n", user.Name)
}
```

**Finding first match:**
```go
func findUser(users []User, name string) *User {
	for _, user := range users {
		if user.Name == name {
			return &user  // Found, exit early
		}
	}
	return nil  // Not found
}
```

---

## 4.4 defer Statement

A `defer` statement delays execution of a function **until the surrounding function returns**.

### Basic defer

```go
func example() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

// Output:
// Hello
// World
```

### Key Rules

1. **Deferred calls execute in LIFO order** (stack)
2. **Arguments are evaluated immediately**
3. **Execution happens even after panic**

### Multiple Defers (LIFO order)

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

### defer with Arguments (evaluated immediately)

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

### Common Use Cases

**1. Resource Cleanup:**
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

**2. Unlock Mutex:**
```go
func updateCounter(mu *sync.Mutex, counter *int) {
	mu.Lock()
	defer mu.Unlock()  // Always unlocks, even if panic
	
	*counter++
}
```

**3. Transaction Rollback:**
```go
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

**4. Recover from Panic:**
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

// Output: Recovered from panic: something went wrong
```

**5. Timing Function Execution:**
```go
func measureTime() {
	start := time.Now()
	defer func() {
		fmt.Println("Duration:", time.Since(start))
	}()
	
	// Function work
	time.Sleep(time.Second)
}

// Output: Duration: 1s
```

### defer Gotchas

**Defer in loops (be careful!):**
```go
// ❌ Bad: All defers execute after function returns
func badExample() {
	for i := 0; i < 5; i++ {
		file, _ := os.Open(fmt.Sprintf("file%d.txt", i))
		defer file.Close()  // All files stay open until function ends!
	}
}

// ✅ Good: Use a helper function
func goodExample() {
	for i := 0; i < 5; i++ {
		processFile(fmt.Sprintf("file%d.txt", i))
	}
}

func processFile(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()  // Closes when processFile returns
	
	// Work with file
}
```

**Defer with named returns:**
```go
func example() (result int) {
	defer func() {
		result++  // Can modify named return value
	}()
	
	result = 5
	return  // Returns 6 (modified by defer)
}
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. In Go, switch statements do NOT require __________ to prevent fall-through.
2. The __________ keyword skips the current iteration of a loop.
3. Braces are __________ (mandatory/optional) for if statements in Go.
4. Deferred functions execute in __________ order (FIFO/LIFO).
5. The __________ keyword exits a loop immediately.
6. A switch without an expression acts like a chain of __________.

### True/False

1. ⬜ Parentheses around if conditions are optional in Go
2. ⬜ Go's for loop can function as a while loop
3. ⬜ Multiple values can be tested in a single case of a switch
4. ⬜ The range loop over strings iterates over bytes
5. ⬜ Map iteration order is guaranteed to be consistent
6. ⬜ defer executes before the function returns

### Multiple Choice

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

3. What is the output?
   ```go
   defer fmt.Print("1")
   defer fmt.Print("2")
   fmt.Print("3")
   ```
   - A) 123
   - B) 321
   - C) 312
   - D) 213

4. How do you iterate over map keys only?
   - A) `for k, _ := range m`
   - B) `for k := range m`
   - C) `for _, k := range m`
   - D) `for range m`

### Code Challenge

Write a function `fizzBuzz(n int) string` that returns:
- "Fizz" for multiples of 3
- "Buzz" for multiples of 5
- "FizzBuzz" for multiples of both
- The number as a string otherwise

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. break
2. continue
3. mandatory
4. LIFO
5. break
6. if-else statements

**True/False:**
1. ❌ False (parentheses are not allowed)
2. ✅ True
3. ✅ True
4. ❌ False (iterates over runes/Unicode characters)
5. ❌ False (order is random)
6. ❌ False (executes when function returns, not before)

**Multiple Choice:**
1. **B** - 0 2 (skips 1 due to continue)
2. **C** - `for {}`
3. **C** - 312 (prints 3, then defers execute in reverse: 2, then 1)
4. **B** - `for k := range m`

**Code Challenge:**
```go
import "strconv"

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

// Usage
for i := 1; i <= 20; i++ {
	fmt.Println(fizzBuzz(i))
}
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ if-else statements with initialization  
✅ Early return pattern and guard clauses  
✅ switch statements (basic, without expression, type switch)  
✅ for loops (traditional, while-style, infinite, range)  
✅ Loop control (break, continue, labels)  
✅ defer statement and common patterns  

**Next Steps:**
- Continue to [Chapter 5: Functions →](./5-functions.md)
- Practice control flow patterns
- Complete the FizzBuzz challenge above

---

**Navigation:** [← Strings & I/O](3-strings-io.md) | [README](./README.md) | [Next: Functions →](5-functions.md)