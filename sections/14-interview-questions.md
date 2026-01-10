# 14. Interview Preparation

> Common Go interview questions with detailed answers.

**Navigation:** [← Best Practices](13-best-practices.md) | [README](./README.md) | [Next: Resources →](15-resources.md)

---

## Table of Contents

**Questions:**
1. [What is Go and why was it created?](#1-what-is-go-and-why-was-it-created)
2. [What are goroutines?](#2-what-are-goroutines)
3. [What is the difference between goroutines and threads?](#3-what-is-the-difference-between-goroutines-and-threads)
4. [What are channels?](#4-what-are-channels)
5. [Difference between buffered and unbuffered channels?](#5-difference-between-buffered-and-unbuffered-channels)
6. [What is a WaitGroup?](#6-what-is-a-waitgroup)
7. [What is an interface?](#7-what-is-an-interface)
8. [How does Go handle errors?](#8-how-does-go-handle-errors)
9. [What is defer?](#9-what-is-defer)
10. [What is panic and recover?](#10-what-is-panic-and-recover)
11. [What is the difference between make and new?](#11-what-is-the-difference-between-make-and-new)
12. [What are slices and how do they differ from arrays?](#12-what-are-slices-and-how-do-they-differ-from-arrays)
13. [What is a map?](#13-what-is-a-map)
14. [What is the zero value?](#14-what-is-the-zero-value)
15. [What are struct tags?](#15-what-are-struct-tags)
16. [What is embedding?](#16-what-is-embedding)
17. [What is the difference between value and pointer receivers?](#17-what-is-the-difference-between-value-and-pointer-receivers)
18. [What is type assertion?](#18-what-is-type-assertion)
19. [What is the select statement?](#19-what-is-the-select-statement)
20. [What is Context?](#20-what-is-context)
21. [What is a mutex?](#21-what-is-a-mutex)
22. [What is the difference between Mutex and RWMutex?](#22-what-is-the-difference-between-mutex-and-rwmutex)
23. [What are race conditions?](#23-what-are-race-conditions)
24. [How do you detect race conditions?](#24-how-do-you-detect-race-conditions)
25. [What is init function?](#25-what-is-init-function)
26. [What is the main function?](#26-what-is-the-main-function)
27. [How does Go handle package initialization?](#27-how-does-go-handle-package-initialization)
28. [What are goroutine leaks?](#28-what-are-goroutine-leaks)
29. [How do you prevent goroutine leaks?](#29-how-do-you-prevent-goroutine-leaks)
30. [What is the difference between = and :=?](#30-what-is-the-difference-between--and-)
31. [Can you return multiple values from a function?](#31-can-you-return-multiple-values-from-a-function)
32. [What are variadic functions?](#32-what-are-variadic-functions)
33. [What is a closure?](#33-what-is-a-closure)
34. [What are generics in Go?](#34-what-are-generics-in-go)
35. [How do you write tests in Go?](#35-how-do-you-write-tests-in-go)
36. [What is table-driven testing?](#36-what-is-table-driven-testing)
37. [How do you benchmark code?](#37-how-do-you-benchmark-code)
38. [What is the Go scheduler?](#38-what-is-the-go-scheduler)
39. [What is garbage collection in Go?](#39-what-is-garbage-collection-in-go)
40. [What are some Go best practices?](#40-what-are-some-go-best-practices)

---

## 1. What is Go and why was it created?

**Answer:**

Go (Golang) is an **open-source, statically typed, compiled programming language** created by Google in 2007 and released in 2009.

**Created by:** Robert Griesemer, Rob Pike, and Ken Thompson

**Why it was created:**
- **Simplicity**: Reduce complexity in large codebases
- **Performance**: Fast compilation and execution (near C/C++ performance)
- **Concurrency**: Built-in support for concurrent programming
- **Modern needs**: Address challenges of multicore processors and networked systems

**Key characteristics:**
- Statically typed with type inference
- Garbage collected
- Built-in concurrency (goroutines, channels)
- Fast compilation
- Simple syntax (only 25 keywords)
- Strong standard library

---

## 2. What are goroutines?

**Answer:**

A **goroutine** is a lightweight thread managed by the Go runtime.

**Key points:**
- Extremely lightweight (~2KB initial stack)
- Multiplexed onto OS threads by Go scheduler
- Can run thousands/millions concurrently
- Started with the `go` keyword

**Example:**
```go
func main() {
	go sayHello()  // Starts goroutine
	
	time.Sleep(time.Second)
}

func sayHello() {
	fmt.Println("Hello from goroutine")
}
```

**Why goroutines vs threads:**
- Much lighter weight (2KB vs 1-2MB)
- Faster creation and destruction
- Managed by Go runtime, not OS
- M:N scheduling (M goroutines on N threads)

---

## 3. What is the difference between goroutines and threads?

**Answer:**

| Feature | Goroutines | OS Threads |
|---------|-----------|------------|
| **Size** | ~2KB initial stack | 1-2MB fixed stack |
| **Creation cost** | Very cheap | Expensive |
| **Scheduling** | Go runtime (cooperative) | OS kernel (preemptive) |
| **Context switching** | Fast (~200ns) | Slow (~1-2μs) |
| **Scalability** | Millions possible | Thousands max |
| **Management** | Automatic by runtime | Manual or via thread pools |

**Example demonstrating scalability:**
```go
func main() {
	// Can easily create 100,000 goroutines
	for i := 0; i < 100000; i++ {
		go func(n int) {
			time.Sleep(time.Second)
			fmt.Println(n)
		}(i)
	}
	
	time.Sleep(2 * time.Second)
}
```

**Interview tip:** Emphasize that goroutines are scheduled by the Go runtime using M:N scheduling (M goroutines on N OS threads).

---

## 4. What are channels?

**Answer:**

**Channels** are typed conduits for communication between goroutines.

**Key points:**
- Enable safe data sharing between goroutines
- Follow CSP (Communicating Sequential Processes) model
- Can be buffered or unbuffered
- Provide synchronization

**Example:**
```go
func main() {
	ch := make(chan int)
	
	go func() {
		ch <- 42  // Send to channel
	}()
	
	value := <-ch  // Receive from channel
	fmt.Println(value)
}
```

**Channel operations:**
- `ch <- value` - Send
- `value := <-ch` - Receive
- `close(ch)` - Close channel

**Interview tip:** Explain that channels solve the "don't communicate by sharing memory, share memory by communicating" principle.

---

## 5. Difference between buffered and unbuffered channels?

**Answer:**

**Unbuffered channels:**
- Capacity: 0
- Send blocks until receive
- Receive blocks until send
- Provides synchronization

```go
ch := make(chan int)  // Unbuffered
```

**Buffered channels:**
- Capacity: > 0
- Send blocks only when full
- Receive blocks only when empty
- Allows asynchronous communication

```go
ch := make(chan int, 3)  // Buffered, capacity 3
```

**Comparison:**
```go
// Unbuffered - synchronous
ch1 := make(chan int)
go func() {
	ch1 <- 1  // Blocks until received
}()
fmt.Println(<-ch1)

// Buffered - asynchronous (up to capacity)
ch2 := make(chan int, 2)
ch2 <- 1  // Doesn't block
ch2 <- 2  // Doesn't block
// ch2 <- 3  // Would block (buffer full)
```

**When to use each:**
- **Unbuffered**: When you need synchronization
- **Buffered**: When you know capacity and want to reduce blocking

---

## 6. What is a WaitGroup?

**Answer:**

**WaitGroup** waits for a collection of goroutines to finish.

**Methods:**
- `Add(n)` - Increment counter by n
- `Done()` - Decrement counter by 1
- `Wait()` - Block until counter is 0

**Example:**
```go
import "sync"

func main() {
	var wg sync.WaitGroup
	
	for i := 0; i < 3; i++ {
		wg.Add(1)  // Increment counter
		go func(n int) {
			defer wg.Done()  // Decrement when done
			fmt.Println(n)
		}(i)
	}
	
	wg.Wait()  // Block until all done
	fmt.Println("All goroutines finished")
}
```

**Common mistake:**
```go
// ❌ Wrong: Add inside goroutine (race condition)
for i := 0; i < 3; i++ {
	go func() {
		wg.Add(1)  // Bad!
		defer wg.Done()
	}()
}

// ✅ Correct: Add before goroutine
for i := 0; i < 3; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
}
```

---

## 7. What is an interface?

**Answer:**

An **interface** is a type that specifies a set of method signatures.

**Key points:**
- Defines behavior, not data
- Implemented implicitly (no "implements" keyword)
- Enables polymorphism
- Empty interface (`interface{}` or `any`) accepts any type

**Example:**
```go
type Writer interface {
	Write([]byte) (int, error)
}

type FileWriter struct {
	filename string
}

// Implements Writer implicitly
func (f FileWriter) Write(data []byte) (int, error) {
	// Write to file
	return len(data), nil
}

func save(w Writer, data []byte) {
	w.Write(data)
}

// Usage
fw := FileWriter{filename: "data.txt"}
save(fw, []byte("hello"))  // Works!
```

**Interview tip:** Emphasize implicit implementation - you don't declare that a type implements an interface, it just does if it has the methods.

---

## 8. How does Go handle errors?

**Answer:**

Go uses **explicit error values** instead of exceptions.

**Key points:**
- Errors are values (implement `error` interface)
- Functions return `(value, error)`
- Check errors immediately
- Use `fmt.Errorf` with `%w` to wrap errors (Go 1.13+)

**Example:**
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
fmt.Println(result)
```

**Error wrapping:**
```go
func processFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("process file %s: %w", path, err)
	}
	// ... process data
	return nil
}
```

**Error checking:**
```go
// errors.Is() - check sentinel errors
if errors.Is(err, os.ErrNotExist) {
	// Handle file not found
}

// errors.As() - extract error type
var pathErr *os.PathError
if errors.As(err, &pathErr) {
	fmt.Println("Path:", pathErr.Path)
}
```

---

## 9. What is defer?

**Answer:**

**defer** delays execution of a function until the surrounding function returns.

**Key points:**
- Executes in LIFO order (stack)
- Arguments evaluated immediately
- Runs even if panic occurs
- Common for cleanup (close files, unlock mutexes)

**Example:**
```go
func example() {
	defer fmt.Println("Third")
	defer fmt.Println("Second")
	fmt.Println("First")
}
// Output: First, Second, Third
```

**Common use cases:**
```go
// 1. Close file
func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()  // Guaranteed to run
	
	// Work with file
	return nil
}

// 2. Unlock mutex
func updateCounter() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

// 3. Recover from panic
func safeExecute() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()
	
	// Code that might panic
}
```

**Interview tip:** Explain LIFO order and that defer is used for cleanup to ensure resources are released.

---

## 10. What is panic and recover?

**Answer:**

**panic** stops normal execution and begins unwinding the stack.  
**recover** catches a panic and resumes execution.

**panic:**
- Used for unrecoverable errors
- Stops current function
- Executes deferred functions
- Propagates up the stack

**recover:**
- Only works inside deferred functions
- Returns the panic value
- Allows cleanup and graceful shutdown

**Example:**
```go
func main() {
	safeDivide(10, 0)
	fmt.Println("Program continues")
}

func safeDivide(a, b float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	
	if b == 0 {
		panic("division by zero")
	}
	
	fmt.Println(a / b)
}

// Output:
// Recovered from panic: division by zero
// Program continues
```

**When to use panic:**
- Programming errors (bugs)
- Initialization failures
- Unrecoverable situations

**When NOT to use panic:**
- Expected errors (use error return instead)
- Input validation
- Business logic errors

**Interview tip:** Emphasize that panic should be rare - use errors for expected failures.

---

## 11. What is the difference between make and new?

**Answer:**

| Feature | `new(T)` | `make(T, args)` |
|---------|----------|-----------------|
| **Purpose** | Allocates memory | Initializes + allocates |
| **Returns** | `*T` (pointer) | `T` (value) |
| **Zero value** | Returns zeroed memory | Returns initialized value |
| **Used for** | Any type | Slices, maps, channels only |

**new(T):**
```go
p := new(int)
fmt.Println(*p)  // 0

type Person struct {
	Name string
	Age  int
}
p := new(Person)
fmt.Println(p)  // &{ 0}
```

**make(T):**
```go
// Slice
s := make([]int, 5, 10)  // length 5, capacity 10

// Map
m := make(map[string]int)

// Channel
ch := make(chan int, 5)
```

**Key difference:**
```go
// new returns pointer to zero value
var p *[]int = new([]int)  // p points to nil slice

// make returns initialized value
var s []int = make([]int, 0)  // s is empty slice, not nil
```

**Interview tip:** `new` allocates zeroed memory and returns a pointer. `make` initializes slices, maps, and channels and returns the initialized value.

---

## 12. What are slices and how do they differ from arrays?

**Answer:**

**Arrays:**
- Fixed size (part of type)
- Value type (copied on assignment)
- Size known at compile time

**Slices:**
- Dynamic size
- Reference type (shares underlying array)
- More flexible and commonly used

**Comparison:**
```go
// Array
var arr [5]int = [5]int{1, 2, 3, 4, 5}
arr2 := arr  // Copies entire array

// Slice
slice := []int{1, 2, 3, 4, 5}
slice2 := slice  // Shares underlying array

slice[0] = 100
fmt.Println(slice2[0])  // 100 (shared)
```

**Slice internals:**
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
fmt.Println(len(s))  // 3
fmt.Println(cap(s))  // 5

s = append(s, 1, 2)
fmt.Println(len(s))  // 5
fmt.Println(cap(s))  // 5

s = append(s, 3)  // Exceeds capacity
fmt.Println(cap(s))  // 10 (doubled)
```

**Interview tip:** Emphasize that slices are references to underlying arrays, and capacity grows when exceeded (typically doubles).

---

## 13. What is a map?

**Answer:**

A **map** is an unordered collection of key-value pairs (hash table).

**Key points:**
- Keys must be comparable (==, !=)
- Reference type
- Iteration order is random
- Zero value is `nil`

**Example:**
```go
// Declaration
m := make(map[string]int)

// Literal
scores := map[string]int{
	"Alice": 90,
	"Bob":   85,
}

// Add/Update
scores["Carol"] = 95

// Get
score := scores["Alice"]

// Check existence
if score, ok := scores["Dave"]; ok {
	fmt.Println(score)
} else {
	fmt.Println("Not found")
}

// Delete
delete(scores, "Bob")
```

**Important:**
```go
// ❌ nil map (read-only)
var m map[string]int
// m["key"] = 1  // Panic!

// ✅ Initialized map
m := make(map[string]int)
m["key"] = 1  // OK
```

**Thread safety:**
```go
// Maps are NOT thread-safe
// Use sync.RWMutex or sync.Map for concurrent access
```

---

## 14. What is the zero value?

**Answer:**

The **zero value** is the default value for uninitialized variables.

**Zero values by type:**
```go
var i int        // 0
var f float64    // 0.0
var s string     // ""
var b bool       // false
var p *int       // nil
var slice []int   // nil
var m map[string]int  // nil
var ch chan int  // nil
var fn func()    // nil
```

**Struct zero values:**
```go
type Person struct {
	Name string
	Age  int
}

var p Person
fmt.Println(p)  // { 0}
```

**Benefits:**
- No undefined behavior
- Predictable initialization
- Safe defaults

**Example:**
```go
type Counter struct {
	value int  // Automatically 0
}

func (c *Counter) Increment() {
	c.value++  // Works immediately
}

c := Counter{}  // Zero value
c.Increment()
fmt.Println(c.value)  // 1
```

---

## 15. What are struct tags?

**Answer:**

**Struct tags** provide metadata about struct fields.

**Common uses:**
- JSON encoding/decoding
- XML marshaling
- Database column mapping
- Validation rules

**Example:**
```go
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"`  // Never serialize
	CreatedAt time.Time `json:"created_at"`
}

// Usage
user := User{
	ID:    1,
	Name:  "Alice",
	Email: "alice@example.com",
}

json, _ := json.Marshal(user)
fmt.Println(string(json))
// {"id":1,"name":"Alice","email":"alice@example.com","created_at":"..."}
```

**Multiple tags:**
```go
type Product struct {
	Name  string `json:"name" xml:"name" db:"product_name"`
	Price float64 `json:"price" xml:"price" db:"price"`
}
```

**Validation tags (using validator library):**
```go
type User struct {
	Email string `validate:"required,email"`
	Age   int    `validate:"gte=0,lte=130"`
}
```

---

## 16. What is embedding?

**Answer:**

**Embedding** is composition by including one struct or interface in another.

**Struct embedding:**
```go
type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Address  // Embedded (anonymous field)
}

p := Person{
	Name: "Alice",
	Address: Address{
		Street: "123 Main St",
		City:   "NYC",
	},
}

// Can access embedded fields directly (promoted)
fmt.Println(p.City)    // NYC
fmt.Println(p.Street)  // 123 Main St

// Or explicitly
fmt.Println(p.Address.City)  // NYC
```

**Interface embedding:**
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader  // Embedded
	Writer  // Embedded
}
```

**Method promotion:**
```go
type Engine struct{}
func (e Engine) Start() { fmt.Println("Engine started") }

type Car struct {
	Engine  // Embedded
}

car := Car{}
car.Start()  // Promoted method
```

**Interview tip:** Embedding is Go's alternative to inheritance - it's composition, not inheritance.

---

## 17. What is the difference between value and pointer receivers?

**Answer:**

| Feature | Value Receiver | Pointer Receiver |
|---------|---------------|------------------|
| **Modifies receiver** | No | Yes |
| **Copying** | Copies receiver | Passes pointer |
| **Usage** | Small, immutable types | Large structs, need modification |

**Value receiver:**
```go
type Counter struct {
	Value int
}

func (c Counter) Increment() {
	c.Value++  // Modifies copy, not original!
}

c := Counter{Value: 0}
c.Increment()
fmt.Println(c.Value)  // 0 (unchanged)
```

**Pointer receiver:**
```go
func (c *Counter) Increment() {
	c.Value++  // Modifies original
}

c := Counter{Value: 0}
c.Increment()
fmt.Println(c.Value)  // 1 (modified)
```

**When to use each:**

**Use pointer receiver:**
- Method modifies receiver
- Receiver is large (avoid copying)
- Consistency (if one method uses pointer, all should)

**Use value receiver:**
- Small receiver
- Receiver is immutable
- Receiver is a value type (int, bool, etc.)

**Interview tip:** Pointer receivers are more common. For consistency, if any method needs a pointer receiver, use pointer receivers for all methods on that type.

---

## 18. What is type assertion?

**Answer:**

**Type assertion** extracts the concrete value from an interface.

**Syntax:**
```go
value := interfaceValue.(ConcreteType)
```

**Unsafe (panics if wrong):**
```go
var i interface{} = "hello"

s := i.(string)
fmt.Println(s)  // "hello"

// n := i.(int)  // Panic!
```

**Safe (comma-ok idiom):**
```go
s, ok := i.(string)
if ok {
	fmt.Println("String:", s)
} else {
	fmt.Println("Not a string")
}

n, ok := i.(int)
if ok {
	fmt.Println("Int:", n)
} else {
	fmt.Println("Not an int")  // This prints
}
```

**Type switch:**
```go
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

describe(42)      // Integer: 42
describe("hello") // String: hello
```

---

## 19. What is the select statement?

**Answer:**

**select** waits on multiple channel operations.

**Key points:**
- Like switch but for channels
- Blocks until one case can proceed
- Randomly chooses if multiple ready
- `default` makes it non-blocking

**Basic example:**
```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "one" }()
go func() { ch2 <- "two" }()

select {
case msg1 := <-ch1:
	fmt.Println("Received:", msg1)
case msg2 := <-ch2:
	fmt.Println("Received:", msg2)
}
```

**Timeout pattern:**
```go
select {
case result := <-ch:
	fmt.Println(result)
case <-time.After(1 * time.Second):
	fmt.Println("Timeout!")
}
```

**Non-blocking:**
```go
select {
case msg := <-ch:
	fmt.Println(msg)
default:
	fmt.Println("No message")
}
```

---

## 20. What is Context?

**Answer:**

**Context** carries deadlines, cancellation signals, and request-scoped values.

**Key uses:**
- Cancellation propagation
- Timeout management
- Request-scoped values

**Creating contexts:**
```go
// Background context
ctx := context.Background()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// With value
ctx := context.WithValue(context.Background(), "key", "value")
```

**Example:**
```go
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
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

**Best practices:**
- Pass as first parameter
- Don't store in structs
- Always call cancel()

---

## 21. What is a mutex?

**Answer:**

A **mutex** (mutual exclusion) protects shared resources from concurrent access.

**Key points:**
- Ensures only one goroutine accesses critical section
- Prevents race conditions
- Must unlock after use (use defer)

**Example:**
```go
import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
```

**Usage:**
```go
counter := &Counter{}
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter.Increment()
	}()
}

wg.Wait()
fmt.Println(counter.Value())  // 1000 (correct)
```

**Interview tip:** Always unlock mutex (use defer), and keep critical sections small.

---

## 22. What is the difference between Mutex and RWMutex?

**Answer:**

| Feature | Mutex | RWMutex |
|---------|-------|---------|
| **Locks** | Exclusive only | Read or Write locks |
| **Concurrent reads** | No | Yes |
| **Use case** | Simple protection | Read-heavy workloads |

**Mutex:**
```go
var mu sync.Mutex

mu.Lock()
// Critical section
mu.Unlock()
```

**RWMutex:**
```go
var mu sync.RWMutex

// Read lock (multiple readers allowed)
mu.RLock()
value := data
mu.RUnlock()

// Write lock (exclusive)
mu.Lock()
data = newValue
mu.Unlock()
```

**Example:**
```go
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
```

**Interview tip:** Use RWMutex when reads far outnumber writes.

---

## 23. What are race conditions?

**Answer:**

A **race condition** occurs when multiple goroutines access shared data concurrently and at least one modifies it.

**Example of race condition:**
```go
var counter int

func increment() {
	counter++  // Not atomic! (read, modify, write)
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	
	time.Sleep(time.Second)
	fmt.Println(counter)  // Likely < 1000
}
```

**Why it happens:**
```
Goroutine 1: Read counter (10) -> Increment (11) -> Write (11)
Goroutine 2:     Read counter (10) -> Increment (11) -> Write (11)
Result: counter = 11 (should be 12)
```

**Solutions:**

**1. Mutex:**
```go
var mu sync.Mutex

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}
```

**2. Channels:**
```go
func main() {
	ch := make(chan int)
	counter := 0
	
	go func() {
		for {
			counter++
			ch <- counter
		}
	}()
	
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
```

**3. atomic package:**
```go
import "sync/atomic"

var counter int64

func increment() {
	atomic.AddInt64(&counter, 1)
}
```

---

## 24. How do you detect race conditions?

**Answer:**

Use the **race detector** with the `-race` flag.

**Running with race detector:**
```bash
go run -race main.go
go test -race
go build -race
```

**Example:**
```go
// main.go
var counter int

func main() {
	go func() { counter++ }()
	go func() { counter++ }()
	time.Sleep(time.Second)
	fmt.Println(counter)
}
```

**Output with -race:**
```
==================
WARNING: DATA RACE
Write at 0x... by goroutine 6:
  main.main.func1()
      /path/main.go:8 +0x3c

Previous write at 0x... by goroutine 7:
  main.main.func2()
      /path/main.go:9 +0x3c
==================
```

**Interview tip:** Always run tests with `-race` flag before production deployment.

---

## 25. What is init function?

**Answer:**

The **init function** runs automatically when a package is initialized.

**Key points:**
- Runs before main()
- Can have multiple init() in same package
- Executes in declaration order
- Cannot be called explicitly
- No parameters or return values

**Example:**
```go
package main

import "fmt"

var config Config

func init() {
	fmt.Println("Init 1")
	config = loadConfig()
}

func init() {
	fmt.Println("Init 2")
	setupDatabase()
}

func main() {
	fmt.Println("Main")
}

// Output:
// Init 1
// Init 2
// Main
```

**Common uses:**
- Initialize package-level variables
- Register drivers (database, image formats)
- One-time setup

**Example (database driver):**
```go
import _ "github.com/lib/pq"  // Runs init() to register driver
```

---

## 26. What is the main function?

**Answer:**

The **main function** is the entry point of a Go program.

**Requirements:**
- Must be in `package main`
- Must be named `main`
- No parameters
- No return values
- Only one per program

**Example:**
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

**Exit codes:**
```go
import "os"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
```

**Interview tip:** Explain that main() is where program execution begins, after all init() functions complete.

---

## 27. How does Go handle package initialization?

**Answer:**

**Initialization order:**

1. **Import dependencies** (depth-first)
2. **Initialize package-level variables**
3. **Run init() functions**
4. **Repeat for importing package**

**Example:**
```go
// package a
package a

var A = "A"

func init() {
	fmt.Println("Init A")
}

// package b (imports a)
package b

import "a"

var B = "B"

func init() {
	fmt.Println("Init B")
}

// main (imports b)
package main

import "b"

func init() {
	fmt.Println("Init main")
}

func main() {
	fmt.Println("Main")
}

// Output:
// Init A
// Init B
// Init main
// Main
```

**Variable initialization:**
```go
var a = b + 1  // Depends on b
var b = f()    // Computed first

func f() int { return 1 }

// Order: b initialized, then a
```

---

## 28. What are goroutine leaks?

**Answer:**

A **goroutine leak** occurs when a goroutine is started but never terminates, consuming resources.

**Common causes:**

**1. Channel operations that never complete:**
```go
// ❌ Leaks goroutine
func leak() {
	ch := make(chan int)
	go func() {
		ch <- 1  // Blocks forever (no receiver)
	}()
}
```

**2. Goroutines waiting on conditions that never happen:**
```go
// ❌ Leaks goroutine
func leak() {
	done := make(chan bool)
	go func() {
		// Work
		done <- true
	}()
	// Forgot to receive from done
}
```

**3. No cancellation mechanism:**
```go
// ❌ Leaks goroutine
func leak() {
	go func() {
		for {
			// Infinite loop, no way to stop
			time.Sleep(time.Second)
		}
	}()
}
```

**Interview tip:** Goroutine leaks are memory leaks in Go. Always ensure goroutines can terminate.

---

## 29. How do you prevent goroutine leaks?

**Answer:**

**Strategies to prevent leaks:**

**1. Use buffered channels:**
```go
func noleak() {
	ch := make(chan int, 1)  // Buffered
	go func() {
		ch <- 1  // Won't block even if not received
	}()
}
```

**2. Use context for cancellation:**
```go
func noleak(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return  // Exit when canceled
			default:
				// Work
			}
		}
	}()
}
```

**3. Use done channel:**
```go
func noleak() {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// Work
			}
		}
	}()
	
	// Later...
	close(done)  // Signal goroutine to stop
}
```

**4. Use timeout:**
```go
func noleak() {
	ch := make(chan int)
	go func() {
		select {
		case ch <- compute():
		case <-time.After(time.Second):
			return  // Timeout
		}
	}()
}
```

---

## 30. What is the difference between = and :=?

**Answer:**

| Operator | Name | Usage |
|----------|------|-------|
| `=` | Assignment | Assigns to existing variable |
| `:=` | Short declaration | Declares AND assigns new variable |

**= (Assignment):**
```go
var x int
x = 10  // Assignment to existing variable

var y int = 20  // Declaration with assignment
```

**:= (Short declaration):**
```go
x := 10  // Declares x and assigns 10
// Equivalent to: var x int = 10
```

**Rules for :=:**
- Only inside functions (not at package level)
- Must declare at least one new variable
- Type is inferred

**Examples:**
```go
// ✅ Valid
x := 10
y := "hello"

// ❌ Invalid: outside function
package main
x := 10  // Compile error

// ✅ Valid: redeclaration if at least one new
x := 10
x, y := 20, 30  // OK: y is new

// ❌ Invalid: no new variables
x := 10
x := 20  // Compile error
```

---

## 31. Can you return multiple values from a function?

**Answer:**

**Yes**, Go supports multiple return values.

**Example:**
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
fmt.Println(result)
```

**Named return values:**
```go
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return  // Naked return
	}
	result = a / b
	return
}
```

**Common patterns:**
- `(value, error)` - Most common
- `(value, bool)` - Comma-ok idiom
- `(value1, value2, value3)` - Multiple values

**Interview tip:** Multiple return values are idiomatic Go, especially `(value, error)` pattern.

---

## 32. What are variadic functions?

**Answer:**

**Variadic functions** accept unlimited arguments of the same type.

**Syntax:**
```go
func functionName(param ...Type) {
	// param is []Type inside function
}
```

**Example:**
```go
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Usage
fmt.Println(sum(1, 2, 3))        // 6
fmt.Println(sum(1, 2, 3, 4, 5))  // 15

// Spread slice
values := []int{1, 2, 3}
fmt.Println(sum(values...))      // 6
```

**Rules:**
- Variadic parameter must be last
- Only one variadic parameter per function

**Real-world example:**
```go
func Printf(format string, args ...interface{}) {
	// Format string with any number of arguments
}

fmt.Printf("Name: %s, Age: %d\n", "Alice", 30)
```

---

## 33. What is a closure?

**Answer:**

A **closure** is a function that references variables from outside its body.

**Example:**
```go
func counter() func() int {
	count := 0
	return func() int {
		count++  // Captures count from outer scope
		return count
	}
}

c1 := counter()
fmt.Println(c1())  // 1
fmt.Println(c1())  // 2
fmt.Println(c1())  // 3

c2 := counter()
fmt.Println(c2())  // 1 (separate instance)
```

**Real-world use case:**
```go
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Closure captures 'next'
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
```

**Common pitfall:**
```go
// ❌ Wrong: all closures capture same i
for i := 0; i < 3; i++ {
	go func() {
		fmt.Println(i)  // May print 3, 3, 3
	}()
}

// ✅ Correct: pass as parameter
for i := 0; i < 3; i++ {
	go func(n int) {
		fmt.Println(n)  // 0, 1, 2
	}(i)
}
```

---

## 34. What are generics in Go?

**Answer:**

**Generics** (Go 1.18+) allow functions and types to work with any type.

**Basic example:**
```go
func Print[T any](value T) {
	fmt.Println(value)
}

Print(42)
Print("hello")
Print(true)
```

**Generic data structure:**
```go
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)

stringStack := Stack[string]{}
stringStack.Push("hello")
```

**Type constraints:**
```go
func Sum[T int | float64](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}
```

**Interview tip:** Generics eliminate code duplication while maintaining type safety.

---

## 35. How do you write tests in Go?

**Answer:**

Go has built-in testing via the `testing` package.

**Test file structure:**
- Files end with `_test.go`
- Functions start with `Test`
- Take `*testing.T` parameter

**Example:**

**math.go:**
```go
package math

func Add(a, b int) int {
	return a + b
}
```

**math_test.go:**
```go
package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}
}
```

**Running tests:**
```bash
go test
go test -v        # Verbose
go test ./...     # All packages
go test -run TestAdd  # Specific test
```

**Test methods:**
```go
t.Error("message")     // Log and continue
t.Errorf("format", args)
t.Fatal("message")     // Log and stop
t.Fatalf("format", args)
t.Skip("message")      // Skip test
```

---

## 36. What is table-driven testing?

**Answer:**

**Table-driven tests** test multiple cases using a data structure.

**Example:**
```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
		{"zero", 0, 0, 0},
		{"mixed", -5, 10, 5},
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
- Easy to add test cases
- Clear test data
- Better coverage
- Subtests with t.Run()

**Interview tip:** Table-driven tests are a Go best practice.

---

## 37. How do you benchmark code?

**Answer:**

Use benchmark functions with `*testing.B`.

**Example:**
```go
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
```

**Running benchmarks:**
```bash
go test -bench=.
go test -bench=BenchmarkAdd
go test -bench=. -benchmem  # Include memory
```

**Output:**
```
BenchmarkAdd-8   1000000000   0.25 ns/op
```

**Comparing implementations:**
```go
func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "a"
		}
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 100; j++ {
			builder.WriteString("a")
		}
	}
}
```

---

## 38. What is the Go scheduler?

**Answer:**

The **Go scheduler** multiplexes goroutines onto OS threads.

**Key concepts:**

**M:N scheduling:**
- M goroutines run on N OS threads
- Goroutines are lightweight (2KB stack)
- Threads are heavyweight (1-2MB stack)

**Components:**
- **G** - Goroutine
- **M** - Machine (OS thread)
- **P** - Processor (logical CPU, scheduling context)

**Work stealing:**
- Idle threads steal work from busy threads
- Balances load across cores

**Example:**
```
GOMAXPROCS=4 (4 logical processors)

G1, G2, G3, G4... (thousands of goroutines)
      ↓
P1    P2    P3    P4  (4 processors)
↓     ↓     ↓     ↓
M1    M2    M3    M4  (4 OS threads)
```

**Interview tip:** Emphasize that Go scheduler is cooperative, not preemptive, and uses work-stealing for load balancing.

---

## 39. What is garbage collection in Go?

**Answer:**

Go uses **automatic garbage collection** to reclaim unused memory.

**Key points:**
- Concurrent mark-and-sweep
- Low latency (<1ms pauses)
- Generational (focuses on young objects)
- Runs in background

**GC process:**
1. **Mark** - Find reachable objects
2. **Sweep** - Reclaim unreachable objects

**Controlling GC:**
```go
import "runtime"

// Force GC (rarely needed)
runtime.GC()

// Set GC percentage
debug.SetGCPercent(100)  // Default

// View GC stats
var stats runtime.MemStats
runtime.ReadMemStats(&stats)
fmt.Println("Alloc:", stats.Alloc)
```

**Best practices:**
- Reuse objects (sync.Pool)
- Avoid allocations in hot paths
- Preallocate slices/maps when size known

**Interview tip:** Go's GC is designed for low latency, not throughput.

---

## 40. What are some Go best practices?

**Answer:**

**Key Go best practices:**

**1. Error handling:**
```go
// Always check errors
if err != nil {
	return fmt.Errorf("context: %w", err)
}
```

**2. Use gofmt:**
```bash
go fmt ./...
```

**3. Accept interfaces, return structs:**
```go
func Process(r io.Reader) *Result {
	// Implementation
}
```

**4. Keep interfaces small:**
```go
type Reader interface {
	Read([]byte) (int, error)
}
```

**5. Use defer for cleanup:**
```go
file, err := os.Open("file.txt")
if err != nil {
	return err
}
defer file.Close()
```

**6. Avoid goroutine leaks:**
```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()
```

**7. Use table-driven tests:**
```go
tests := []struct{...}{...}
for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {...})
}
```

**8. Preallocate slices:**
```go
items := make([]Item, 0, expectedSize)
```

**Interview tip:** Mention "Effective Go" and Go Code Review Comments as authoritative sources.

---

## Summary

In this part, you learned:

✅ Advanced concurrency (mutex, race conditions)  
✅ Initialization (init, main, package order)  
✅ Goroutine management (leaks, prevention)  
✅ Language features (closures, variadic, generics)  
✅ Testing and benchmarking  
✅ Runtime internals (scheduler, GC)  
✅ Best practices  

**Interview Preparation Tips:**
- Practice explaining concepts clearly
- Use concrete examples
- Know when to use each feature
- Understand trade-offs
- Be familiar with standard library

**Next Steps:**
- Continue to [Chapter 15: Resources →](15-resources.md)
- Practice coding challenges
- Review Go proverbs

---

**Navigation:** [← Best Practices](14-best-practices.md) | [README](./README.md) | [Next: Resources →](15-resources.md)