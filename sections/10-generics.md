# 10. Generics

> Master type parameters and generic programming in Go 1.18+.

**Navigation:** [← Concurrency](9-concurrency.md) | [README](./README.md) | [Next: Testing →](11-testing.md)

---

## Table of Contents

- [10.1 Generics Basics](#101-generics-basics)
- [10.2 Type Constraints](#102-type-constraints)
- [10.3 Generic Data Structures](#103-generic-data-structures)
- [10.4 Generic Functions](#104-generic-functions)
- [10.5 Constraints Package](#105-constraints-package)
- [10.6 Best Practices](#106-best-practices)
- [Practice Questions](#practice-questions)

---

## 10.1 Generics Basics

**Generics** allow you to write functions and types that work with any type.

**Available since:** Go 1.18 (March 2022)

### Before Generics

**Problem: Code duplication**

```go
// Need separate functions for each type
func SumInts(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumFloats(nums []float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum
}
```

### With Generics

**Solution: Single generic function**

```go
func Sum[T int | float64](nums []T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

// Usage
ints := []int{1, 2, 3, 4, 5}
fmt.Println(Sum(ints))  // 15

floats := []float64{1.5, 2.5, 3.5}
fmt.Println(Sum(floats))  // 7.5
```

### Generic Syntax

```go
func FunctionName[T TypeConstraint](param T) T {
	// function body
}

type TypeName[T TypeConstraint] struct {
	field T
}
```

**Components:**
- `[T TypeConstraint]` - Type parameter list
- `T` - Type parameter (like a variable for types)
- `TypeConstraint` - Limits what types T can be

---

## 10.2 Type Constraints

**Constraints** specify which types are allowed for type parameters.

### any Constraint

**Accepts any type:**

```go
func Print[T any](value T) {
	fmt.Println(value)
}

Print(42)
Print("hello")
Print(true)
Print([]int{1, 2, 3})
```

**Note:** `any` is an alias for `interface{}`

### Union Constraints (|)

**Multiple allowed types:**

```go
func Add[T int | float64 | string](a, b T) T {
	return a + b
}

fmt.Println(Add(1, 2))           // 3
fmt.Println(Add(1.5, 2.5))       // 4.0
fmt.Println(Add("Hello, ", "World"))  // Hello, World
```

### comparable Constraint

**Types that support == and !=:**

```go
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

nums := []int{1, 2, 3, 4, 5}
fmt.Println(Contains(nums, 3))  // true

words := []string{"go", "rust", "python"}
fmt.Println(Contains(words, "go"))  // true
```

### Interface Constraints

```go
type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}
```

### Custom Constraints

```go
type Number interface {
	int | int64 | float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}
```

### Approximate Constraint (~)

**Includes underlying type:**

```go
type MyInt int

type Integer interface {
	~int  // Allows int and types based on int
}

func Double[T Integer](n T) T {
	return n * 2
}

var x int = 5
fmt.Println(Double(x))  // 10

var y MyInt = 5
fmt.Println(Double(y))  // 10 (works because MyInt is based on int)
```

---

## 10.3 Generic Data Structures

### Generic Stack

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

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
fmt.Println(intStack.Pop())  // 2, true

stringStack := Stack[string]{}
stringStack.Push("hello")
stringStack.Push("world")
fmt.Println(stringStack.Pop())  // world, true
```

### Generic Queue

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

func (q *Queue[T]) Size() int {
	return len(q.items)
}
```

### Generic Binary Tree

```go
type TreeNode[T comparable] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (n *TreeNode[T]) Insert(value T) {
	if value == n.Value {
		return
	}
	
	if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode[T]{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode[T]{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *TreeNode[T]) Contains(value T) bool {
	if n == nil {
		return false
	}
	
	if value == n.Value {
		return true
	} else if value < n.Value {
		return n.Left.Contains(value)
	} else {
		return n.Right.Contains(value)
	}
}
```

### Generic LinkedList

```go
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
}

func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList[T]) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}
```

---

## 10.4 Generic Functions

### Map Function

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
squared := Map(nums, func(n int) int {
	return n * n
})
fmt.Println(squared)  // [1 4 9 16 25]

words := []string{"hello", "world"}
lengths := Map(words, func(s string) int {
	return len(s)
})
fmt.Println(lengths)  // [5 5]
```

### Filter Function

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
evens := Filter(nums, func(n int) bool {
	return n%2 == 0
})
fmt.Println(evens)  // [2 4 6]
```

### Reduce Function

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
sum := Reduce(nums, 0, func(acc, n int) int {
	return acc + n
})
fmt.Println(sum)  // 15

product := Reduce(nums, 1, func(acc, n int) int {
	return acc * n
})
fmt.Println(product)  // 120
```

### Find Function

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
first, found := Find(nums, func(n int) bool {
	return n > 3
})
if found {
	fmt.Println("Found:", first)  // Found: 4
}
```

### Min/Max Functions

```go
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Usage
fmt.Println(Min(5, 3))        // 3
fmt.Println(Max(5, 3))        // 5
fmt.Println(Min(1.5, 2.5))    // 1.5
fmt.Println(Max("abc", "xyz")) // xyz
```

---

## 10.5 Constraints Package

**golang.org/x/exp/constraints** provides common constraints.

```go
import "golang.org/x/exp/constraints"

// Ordered: supports <, <=, >, >=
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}

// Signed: signed integers
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned: unsigned integers
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer: all integers
type Integer interface {
	Signed | Unsigned
}

// Float: floating point
type Float interface {
	~float32 | ~float64
}

// Complex: complex numbers
type Complex interface {
	~complex64 | ~complex128
}
```

### Using constraints Package

```go
import "golang.org/x/exp/constraints"

func Abs[T constraints.Signed | constraints.Float](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

fmt.Println(Abs(-5))    // 5
fmt.Println(Abs(-3.14)) // 3.14
```

---

## 10.6 Best Practices

### When to Use Generics

**✅ Use generics for:**
- Data structures (Stack, Queue, Tree)
- Algorithms (Sort, Search, Map, Filter)
- Utility functions working on multiple types

**❌ Don't use generics for:**
- Simple functions (unnecessary complexity)
- When interface{} is sufficient
- Method implementations (generics not allowed on methods)

### Type Inference

```go
func Print[T any](value T) {
	fmt.Println(value)
}

// Explicit type parameter
Print[int](42)

// Type inference (preferred)
Print(42)
```

**Prefer type inference when possible.**

### Avoid Over-Generalization

**❌ Too generic:**
```go
func Process[T any](data T) T {
	// What can you do with any?
	return data
}
```

**✅ Better:**
```go
type Processable interface {
	Process() string
}

func Process[T Processable](data T) string {
	return data.Process()
}
```

### Named Type Parameters

**Use descriptive names:**

```go
// ✅ Good
func Transform[Input, Output any](data Input, fn func(Input) Output) Output {
	return fn(data)
}

// ❌ Less clear
func Transform[T, U any](data T, fn func(T) U) U {
	return fn(data)
}
```

### Real-World Example: Result Type

```go
type Result[T any] struct {
	value T
	err   error
}

func NewResult[T any](value T, err error) Result[T] {
	return Result[T]{value: value, err: err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) Unwrap() (T, error) {
	return r.value, r.err
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}
	return r.value
}

// Usage
func divide(a, b float64) Result[float64] {
	if b == 0 {
		return NewResult(0.0, errors.New("division by zero"))
	}
	return NewResult(a/b, nil)
}

result := divide(10, 2)
fmt.Println(result.UnwrapOr(0))  // 5

result = divide(10, 0)
fmt.Println(result.UnwrapOr(0))  // 0
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. Generics were introduced in Go version __________.
2. The __________ constraint accepts any type.
3. The __________ constraint allows types that support == and !=.
4. The __________ symbol allows types with a specific underlying type.
5. Type parameters are specified in __________ brackets.
6. The constraints package provides the __________ interface for ordered types.

### True/False

1. ⬜ Generics were available since Go 1.0
2. ⬜ Type inference is always required for generic functions
3. ⬜ Methods can have type parameters
4. ⬜ The `any` constraint is an alias for `interface{}`
5. ⬜ Union constraints use the | operator
6. ⬜ Generics eliminate all need for interface{}

### Multiple Choice

1. What does the ~ symbol mean in constraints?
   - A) Negation
   - B) Approximate (includes underlying type)
   - C) Optional
   - D) Reference

2. Which constraint allows comparison with < and >?
   - A) comparable
   - B) any
   - C) constraints.Ordered
   - D) interface{}

3. What is the output?
   ```go
   func Identity[T any](v T) T {
       return v
   }
   fmt.Println(Identity(42))
   ```
   - A) Compile error
   - B) 42
   - C) nil
   - D) 0

4. Can methods have type parameters?
   - A) Yes, always
   - B) No, only functions
   - C) Yes, but only on pointer receivers
   - D) Only in Go 1.20+

### Code Challenge

Create a generic `Pair[T, U]` type that:
- Stores two values of potentially different types
- Has `First() T` and `Second() U` methods
- Has a `Swap() Pair[U, T]` method

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. 1.18
2. any
3. comparable
4. ~ (tilde)
5. square [ ]
6. Ordered

**True/False:**
1. ❌ False (Go 1.18, March 2022)
2. ❌ False (type inference is optional)
3. ❌ False (only functions and types can have type parameters)
4. ✅ True
5. ✅ True
6. ❌ False (interface{} still has uses)

**Multiple Choice:**
1. **B** - Approximate (includes underlying type)
2. **C** - constraints.Ordered
3. **B** - 42 (type inference works)
4. **B** - No, only functions (receiver type is fixed)

**Code Challenge:**
```go
type Pair[T, U any] struct {
	first  T
	second U
}

func NewPair[T, U any](first T, second U) Pair[T, U] {
	return Pair[T, U]{first: first, second: second}
}

func (p Pair[T, U]) First() T {
	return p.first
}

func (p Pair[T, U]) Second() U {
	return p.second
}

func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{first: p.second, second: p.first}
}

// Usage
p := NewPair(42, "hello")
fmt.Println(p.First())   // 42
fmt.Println(p.Second())  // hello

swapped := p.Swap()
fmt.Println(swapped.First())   // hello
fmt.Println(swapped.Second())  // 42
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ Generics basics and syntax  
✅ Type constraints (any, comparable, union, custom)  
✅ Generic data structures (Stack, Queue, Tree, LinkedList)  
✅ Generic functions (Map, Filter, Reduce, Find)  
✅ constraints package (Ordered, Signed, etc.)  
✅ Best practices and when to use generics  

**Next Steps:**
- Continue to [Chapter 11: Testing →](11-testing.md)
- Practice creating generic functions
- Complete the Pair type challenge above

---

**Navigation:** [← Concurrency](9-concurrency.md) | [README](./README.md) | [Next: Testing →](11-testing.md)