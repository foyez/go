# 6. Data Structures

> Master arrays, slices, maps, structs, pointers, and their operations in Go.

**Navigation:** [← Functions](5-functions.md) | [README](./README.md) | [Next: Part 2 - Maps & Structs →](7-interfaces.md)

---

## Table of Contents

- [6.1 Arrays](#61-arrays)
- [6.2 Slices](#62-slices)
- [6.3 Maps](#63-maps)
- [6.4 Structs](#64-structs)
- [6.5 Pointers](#65-pointers)
- [Practice Questions - Array & Slices](#practice-questions---array--slices)
- [Practice Questions - Maps, Structs & Pointers](#practice-questions---maps-structs--pointers)

---

## 6.1 Arrays

An **array** is a **fixed-size** collection of elements of the **same type**.

### Array Declaration

```go
// Declare with size
var nums [5]int  // [0 0 0 0 0]

// Initialize with values
var nums = [5]int{1, 2, 3, 4, 5}

// Short declaration
nums := [5]int{1, 2, 3, 4, 5}

// Let compiler count
nums := [...]int{1, 2, 3, 4, 5}

// Partial initialization (rest are zero values)
nums := [5]int{1, 2}  // [1 2 0 0 0]

// Index-based initialization
nums := [5]int{0: 10, 2: 30, 4: 50}  // [10 0 30 0 50]
```

### Accessing Elements

```go
nums := [5]int{10, 20, 30, 40, 50}

fmt.Println(nums[0])  // 10
fmt.Println(nums[4])  // 50

nums[2] = 100
fmt.Println(nums[2])  // 100
```

### Array Properties

```go
nums := [5]int{1, 2, 3, 4, 5}

// Length
fmt.Println(len(nums))  // 5

// Arrays are value types (copied when assigned)
nums2 := nums  // Creates a copy
nums2[0] = 100
fmt.Println(nums[0])   // 1 (unchanged)
fmt.Println(nums2[0])  // 100
```

### Iterating Over Arrays

```go
nums := [5]int{10, 20, 30, 40, 50}

// Using index
for i := 0; i < len(nums); i++ {
	fmt.Println(nums[i])
}

// Using range
for i, v := range nums {
	fmt.Printf("Index: %d, Value: %d\n", i, v)
}

// Ignore index
for _, v := range nums {
	fmt.Println(v)
}
```

### Multidimensional Arrays

```go
// 2D array
var matrix [3][3]int

matrix[0][0] = 1
matrix[1][1] = 5
matrix[2][2] = 9

// Initialize 2D array
matrix := [3][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

// Iterate 2D array
for i := 0; i < len(matrix); i++ {
	for j := 0; j < len(matrix[i]); j++ {
		fmt.Printf("%d ", matrix[i][j])
	}
	fmt.Println()
}
```

### Array Limitations

**❌ Fixed size:**
```go
nums := [3]int{1, 2, 3}
// Cannot add or remove elements
// Cannot change size
```

**❌ Size is part of the type:**
```go
var a [3]int
var b [5]int

// a = b  // ❌ Compile error: different types!
```

**Why arrays are rarely used:**
- Fixed size (can't grow/shrink)
- Size is part of the type
- **Slices are more flexible** (used 99% of the time)

---

## 6.2 Slices

A **slice** is a **dynamic, flexible view** into an array.

**Think of slices as:**
- Dynamic arrays (can grow/shrink)
- References to underlying arrays
- Most commonly used collection type in Go

### Slice Declaration

```go
// Nil slice (zero value)
var nums []int  // nil

// Make with length
nums := make([]int, 5)  // [0 0 0 0 0]

// Make with length and capacity
nums := make([]int, 5, 10)  // length=5, capacity=10

// Literal initialization
nums := []int{1, 2, 3, 4, 5}

// Empty slice (not nil)
nums := []int{}
```

### Slice vs Array

```go
// Array (fixed size)
arr := [5]int{1, 2, 3, 4, 5}

// Slice (dynamic size)
slice := []int{1, 2, 3, 4, 5}

// Key difference: no size in brackets
```

### Length and Capacity

**Length:** Number of elements in slice  
**Capacity:** Number of elements in underlying array

```go
nums := make([]int, 3, 5)

fmt.Println(len(nums))  // 3 (length)
fmt.Println(cap(nums))  // 5 (capacity)

// Visual representation:
// [0 0 0 _ _]
//  ^^^^^      length = 3
//  ^^^^^^^^^  capacity = 5
```

### Appending to Slices

```go
nums := []int{1, 2, 3}

// Append single element
nums = append(nums, 4)
fmt.Println(nums)  // [1 2 3 4]

// Append multiple elements
nums = append(nums, 5, 6, 7)
fmt.Println(nums)  // [1 2 3 4 5 6 7]

// Append another slice
more := []int{8, 9}
nums = append(nums, more...)
fmt.Println(nums)  // [1 2 3 4 5 6 7 8 9]
```

**Important:** `append` returns a new slice (may point to new array if capacity exceeded).

### Slicing Operations

**Syntax:** `slice[low:high]` (low inclusive, high exclusive)

```go
nums := []int{0, 1, 2, 3, 4, 5}

fmt.Println(nums[1:4])   // [1 2 3]
fmt.Println(nums[:3])    // [0 1 2] (from start)
fmt.Println(nums[3:])    // [3 4 5] (to end)
fmt.Println(nums[:])     // [0 1 2 3 4 5] (entire slice)

// Negative indices NOT supported
// nums[-1]  // ❌ Compile error
```

### Slice Memory Sharing

**Slices share underlying array:**

```go
nums := []int{0, 1, 2, 3, 4, 5}
slice1 := nums[1:4]  // [1 2 3]
slice2 := nums[2:5]  // [2 3 4]

slice1[1] = 100

fmt.Println(nums)    // [0 1 100 3 4 5]
fmt.Println(slice1)  // [1 100 3]
fmt.Println(slice2)  // [100 3 4]
```

**Visual:**
```
nums:   [0  1  2  3  4  5]
             ^  ^  ^
slice1:     [1  2  3]
                ^  ^  ^
slice2:        [2  3  4]

Modifying slice1[1] (value 2) affects:
- nums[2]
- slice1[1]
- slice2[0]
```

### Copy Slice (avoiding shared memory)

```go
original := []int{1, 2, 3, 4, 5}

// Create independent copy
copied := make([]int, len(original))
copy(copied, original)

copied[0] = 100

fmt.Println(original)  // [1 2 3 4 5] (unchanged)
fmt.Println(copied)    // [100 2 3 4 5]
```

### Capacity and Reallocation

```go
nums := make([]int, 0, 3)
fmt.Printf("len=%d cap=%d %v\n", len(nums), cap(nums), nums)
// len=0 cap=3 []

nums = append(nums, 1)
fmt.Printf("len=%d cap=%d %v\n", len(nums), cap(nums), nums)
// len=1 cap=3 [1]

nums = append(nums, 2, 3)
fmt.Printf("len=%d cap=%d %v\n", len(nums), cap(nums), nums)
// len=3 cap=3 [1 2 3]

nums = append(nums, 4)  // Capacity exceeded!
fmt.Printf("len=%d cap=%d %v\n", len(nums), cap(nums), nums)
// len=4 cap=6 [1 2 3 4] (capacity doubled)
```

**When capacity is exceeded:**
- Go allocates a new, larger array
- Copies existing elements
- Returns slice pointing to new array

### Removing Elements

**Remove by index:**
```go
nums := []int{1, 2, 3, 4, 5}

// Remove element at index 2 (value 3)
index := 2
nums = append(nums[:index], nums[index+1:]...)
fmt.Println(nums)  // [1 2 4 5]
```

**Remove first element:**
```go
nums := []int{1, 2, 3, 4, 5}
nums = nums[1:]
fmt.Println(nums)  // [2 3 4 5]
```

**Remove last element:**
```go
nums := []int{1, 2, 3, 4, 5}
nums = nums[:len(nums)-1]
fmt.Println(nums)  // [1 2 3 4]
```

### Iterating Over Slices

```go
nums := []int{10, 20, 30, 40, 50}

// Index and value
for i, v := range nums {
	fmt.Printf("nums[%d] = %d\n", i, v)
}

// Value only
for _, v := range nums {
	fmt.Println(v)
}

// Index only
for i := range nums {
	fmt.Println(i)
}

// Traditional loop
for i := 0; i < len(nums); i++ {
	fmt.Println(nums[i])
}
```

### Nil vs Empty Slice

```go
// Nil slice
var s1 []int
fmt.Println(s1 == nil)  // true
fmt.Println(len(s1))    // 0
fmt.Println(cap(s1))    // 0

// Empty slice (not nil)
s2 := []int{}
fmt.Println(s2 == nil)  // false
fmt.Println(len(s2))    // 0
fmt.Println(cap(s2))    // 0

// Both can be used with append
s1 = append(s1, 1)  // Works
s2 = append(s2, 1)  // Works
```

**Recommendation:** Use `nil` slices unless you specifically need an empty but non-nil slice.

### Multi-Dimensional Slices

```go
// 2D slice
matrix := [][]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

fmt.Println(matrix[1][2])  // 6

// Jagged arrays (different row lengths)
jagged := [][]int{
	{1},
	{2, 3},
	{4, 5, 6},
}
```

### Common Slice Patterns

**Pre-allocate capacity:**
```go
// ❌ Inefficient (many reallocations)
var nums []int
for i := 0; i < 1000; i++ {
	nums = append(nums, i)
}

// ✅ Efficient (pre-allocated)
nums := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
	nums = append(nums, i)
}
```

**Filter:**
```go
nums := []int{1, 2, 3, 4, 5, 6}
var evens []int

for _, n := range nums {
	if n%2 == 0 {
		evens = append(evens, n)
	}
}
fmt.Println(evens)  // [2 4 6]
```

**Map (transform):**
```go
nums := []int{1, 2, 3, 4, 5}
squared := make([]int, len(nums))

for i, n := range nums {
	squared[i] = n * n
}
fmt.Println(squared)  // [1 4 9 16 25]
```

**Reverse:**
```go
nums := []int{1, 2, 3, 4, 5}

for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
	nums[i], nums[j] = nums[j], nums[i]
}
fmt.Println(nums)  // [5 4 3 2 1]
```

**Check if element exists:**
```go
func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

nums := []int{1, 2, 3, 4, 5}
fmt.Println(contains(nums, 3))  // true
fmt.Println(contains(nums, 10)) // false
```

### Real-World Example: Dynamic User List

```go
type User struct {
	ID   int
	Name string
}

type UserList struct {
	users []User
}

func (ul *UserList) Add(user User) {
	ul.users = append(ul.users, user)
}

func (ul *UserList) Remove(id int) {
	for i, user := range ul.users {
		if user.ID == id {
			ul.users = append(ul.users[:i], ul.users[i+1:]...)
			return
		}
	}
}

func (ul *UserList) Find(id int) *User {
	for _, user := range ul.users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

// Usage
list := &UserList{}
list.Add(User{ID: 1, Name: "Alice"})
list.Add(User{ID: 2, Name: "Bob"})

user := list.Find(1)
fmt.Println(user.Name)  // Alice

list.Remove(1)
```

---

## 6.3 Maps

A **map** is an unordered collection of **key-value pairs**.

**Think of maps as:**
- Hash tables / dictionaries
- Key → Value lookups
- Fast O(1) average access time

### Map Declaration

```go
// Nil map (cannot add elements)
var m map[string]int  // nil

// Create with make
m := make(map[string]int)

// Literal initialization
m := map[string]int{
	"Alice": 90,
	"Bob":   85,
	"Carol": 95,
}

// Empty map (not nil)
m := map[string]int{}
```

### Adding and Accessing Elements

```go
scores := make(map[string]int)

// Add elements
scores["Alice"] = 90
scores["Bob"] = 85
scores["Carol"] = 95

// Access elements
fmt.Println(scores["Alice"])  // 90

// Access non-existent key (returns zero value)
fmt.Println(scores["Dave"])   // 0 (not an error!)
```

### Checking if Key Exists

```go
scores := map[string]int{
	"Alice": 90,
	"Bob":   0,  // Bob has score of 0
}

// Wrong way (ambiguous with zero value)
score := scores["Bob"]
fmt.Println(score)  // 0 (but is Bob in the map?)

// Correct way (comma-ok idiom)
score, exists := scores["Bob"]
if exists {
	fmt.Println("Bob's score:", score)
} else {
	fmt.Println("Bob not found")
}

// Check without storing value
if _, exists := scores["Dave"]; exists {
	fmt.Println("Dave exists")
} else {
	fmt.Println("Dave not found")
}
```

### Deleting Elements

```go
scores := map[string]int{
	"Alice": 90,
	"Bob":   85,
}

delete(scores, "Bob")

// Delete non-existent key (no error)
delete(scores, "Carol")  // Safe, does nothing
```

### Iterating Over Maps

```go
scores := map[string]int{
	"Alice": 90,
	"Bob":   85,
	"Carol": 95,
}

// Key and value
for name, score := range scores {
	fmt.Printf("%s: %d\n", name, score)
}

// Keys only
for name := range scores {
	fmt.Println(name)
}

// Values only (uncommon)
for _, score := range scores {
	fmt.Println(score)
}
```

**Important:** Map iteration order is **random** and not guaranteed!

### Map Length

```go
scores := map[string]int{
	"Alice": 90,
	"Bob":   85,
}

fmt.Println(len(scores))  // 2
```

### Maps are Reference Types

```go
m1 := map[string]int{"a": 1, "b": 2}
m2 := m1  // m2 points to same map as m1

m2["a"] = 100

fmt.Println(m1["a"])  // 100 (modified!)
fmt.Println(m2["a"])  // 100
```

### Nested Maps

```go
// Map of maps
users := map[string]map[string]string{
	"user1": {
		"name":  "Alice",
		"email": "alice@example.com",
	},
	"user2": {
		"name":  "Bob",
		"email": "bob@example.com",
	},
}

fmt.Println(users["user1"]["name"])  // Alice
```

### Map with Struct Values

```go
type User struct {
	Name  string
	Email string
	Age   int
}

users := map[int]User{
	1: {Name: "Alice", Email: "alice@example.com", Age: 30},
	2: {Name: "Bob", Email: "bob@example.com", Age: 25},
}

user := users[1]
fmt.Println(user.Name)  // Alice
```

### Real-World Example: Counting Words

```go
func countWords(text string) map[string]int {
	words := strings.Fields(text)
	counts := make(map[string]int)
	
	for _, word := range words {
		word = strings.ToLower(word)
		counts[word]++
	}
	
	return counts
}

text := "hello world hello go world"
counts := countWords(text)
fmt.Println(counts)  // map[go:1 hello:2 world:2]
```

### Real-World Example: Grouping Data

```go
type User struct {
	Name string
	City string
}

func groupByCity(users []User) map[string][]User {
	groups := make(map[string][]User)
	
	for _, user := range users {
		groups[user.City] = append(groups[user.City], user)
	}
	
	return groups
}

users := []User{
	{Name: "Alice", City: "NYC"},
	{Name: "Bob", City: "LA"},
	{Name: "Carol", City: "NYC"},
}

groups := groupByCity(users)
fmt.Println(groups["NYC"])  // [{Alice NYC} {Carol NYC}]
```

---

## 6.4 Structs

A **struct** is a collection of fields with different types.

**Think of structs as:**
- Custom data types
- Objects (without methods initially)
- Records or data containers

### Struct Declaration

```go
type Person struct {
	Name string
	Age  int
	City string
}

// Create struct
p := Person{
	Name: "Alice",
	Age:  30,
	City: "NYC",
}

// Short form (positional)
p := Person{"Alice", 30, "NYC"}

// Partial initialization (rest are zero values)
p := Person{Name: "Alice"}

// Zero value (all fields get zero values)
var p Person
```

### Accessing Fields

```go
p := Person{Name: "Alice", Age: 30, City: "NYC"}

// Access fields
fmt.Println(p.Name)  // Alice
fmt.Println(p.Age)   // 30

// Modify fields
p.Age = 31
p.City = "SF"
```

### Anonymous Structs

```go
// Declare and initialize
person := struct {
	Name string
	Age  int
}{
	Name: "Alice",
	Age:  30,
}

fmt.Println(person.Name)  // Alice
```

**Use cases:**
- One-time use structures
- Unmarshaling JSON
- Testing

### Nested Structs

```go
type Address struct {
	Street string
	City   string
	ZIP    string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

p := Person{
	Name: "Alice",
	Age:  30,
	Address: Address{
		Street: "123 Main St",
		City:   "NYC",
		ZIP:    "10001",
	},
}

fmt.Println(p.Address.City)  // NYC
```

### Embedded Structs (Composition)

```go
type Address struct {
	Street string
	City   string
}

type Person struct {
	Name string
	Age  int
	Address  // Embedded (anonymous field)
}

p := Person{
	Name: "Alice",
	Age:  30,
	Address: Address{
		Street: "123 Main St",
		City:   "NYC",
	},
}

// Can access embedded fields directly
fmt.Println(p.City)    // NYC (promoted field)
fmt.Println(p.Street)  // 123 Main St

// Or explicitly
fmt.Println(p.Address.City)  // NYC
```

### Struct Tags (Metadata)

```go
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// Used by encoding/json package
user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}
jsonData, _ := json.Marshal(user)
fmt.Println(string(jsonData))
// {"id":1,"name":"Alice","email":"alice@example.com"}
```

**Common tags:**
- `json` - JSON encoding/decoding
- `xml` - XML encoding/decoding
- `db` - Database column mapping
- `validate` - Validation rules

### Structs are Value Types

```go
p1 := Person{Name: "Alice", Age: 30}
p2 := p1  // Creates a copy

p2.Age = 31

fmt.Println(p1.Age)  // 30 (unchanged)
fmt.Println(p2.Age)  // 31
```

### Comparing Structs

```go
p1 := Person{Name: "Alice", Age: 30}
p2 := Person{Name: "Alice", Age: 30}

fmt.Println(p1 == p2)  // true

// Can only compare if all fields are comparable
// Structs with slices, maps, or functions can't be compared with ==
```

### Methods on Structs

```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
fmt.Println(rect.Area())  // 50

rect.Scale(2)
fmt.Println(rect.Area())  // 200
```

### Real-World Example: User Management

```go
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func NewUser(name, email string) *User {
	return &User{
		ID:        generateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func (u *User) UpdateEmail(email string) error {
	if !isValidEmail(email) {
		return errors.New("invalid email")
	}
	u.Email = email
	return nil
}

func (u User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s}", u.ID, u.Name)
}

// Usage
user := NewUser("Alice", "alice@example.com")
user.UpdateEmail("alice.new@example.com")
fmt.Println(user)
```

---

## 6.5 Pointers

A **pointer** stores the **memory address** of a value.

**Why use pointers:**
- Modify values in functions
- Avoid copying large structs
- Share data across functions

### Pointer Basics

```go
var x int = 10
var p *int = &x  // p is a pointer to x

fmt.Println(x)   // 10 (value)
fmt.Println(&x)  // 0xc0000... (address)
fmt.Println(p)   // 0xc0000... (same address)
fmt.Println(*p)  // 10 (dereference: value at address)
```

**Operators:**
- `&` - Address-of operator (get pointer)
- `*` - Dereference operator (get value)

### Modifying Through Pointers

```go
func increment(n *int) {
	*n++  // Modify value at address
}

x := 10
increment(&x)
fmt.Println(x)  // 11
```

### Pointers vs Values

```go
// Pass by value (copy)
func incrementValue(n int) {
	n++  // Modifies copy
}

// Pass by pointer (reference)
func incrementPointer(n *int) {
	*n++  // Modifies original
}

x := 10
incrementValue(x)
fmt.Println(x)  // 10 (unchanged)

incrementPointer(&x)
fmt.Println(x)  // 11 (changed)
```

### nil Pointers

```go
var p *int
fmt.Println(p)  // nil

// Dereferencing nil pointer causes panic
// fmt.Println(*p)  // ❌ panic: runtime error

// Check before dereferencing
if p != nil {
	fmt.Println(*p)
}
```

### new() Function

```go
// Allocates memory and returns pointer
p := new(int)
fmt.Println(p)   // 0xc0000... (address)
fmt.Println(*p)  // 0 (zero value)

*p = 42
fmt.Println(*p)  // 42
```

### Pointers to Structs

```go
type Person struct {
	Name string
	Age  int
}

// Using &
p := &Person{Name: "Alice", Age: 30}

// Using new
p := new(Person)
p.Name = "Alice"
p.Age = 30

// Access fields (Go auto-dereferences)
fmt.Println(p.Name)  // Alice (Go does p->Name automatically)

// Explicit dereference (same result)
fmt.Println((*p).Name)  // Alice
```

### When to Use Pointers

**Use pointers when:**
1. Need to modify the receiver
2. Struct is large (avoid copying)
3. Need to indicate "optional" or "may not exist" (nil)

**Don't use pointers when:**
1. Value is small (int, bool, small structs)
2. Value types (maps, slices, channels are already references)
3. Immutability is desired

### Pointer Receivers vs Value Receivers

```go
type Counter struct {
	Value int
}

// Value receiver (copy)
func (c Counter) Get() int {
	return c.Value
}

// Pointer receiver (reference)
func (c *Counter) Increment() {
	c.Value++
}

counter := Counter{Value: 0}
counter.Increment()
fmt.Println(counter.Get())  // 1
```

**Google Style Guide:**
- Use pointer receivers when modifying state
- Use pointer receivers for large structs
- Be consistent (if one method uses pointer, all should)

### Real-World Example: Linked List

```go
type Node struct {
	Value int
	Next  *Node  // Pointer to next node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(value int) {
	newNode := &Node{Value: value}
	
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

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

// Usage
list := &LinkedList{}
list.Add(1)
list.Add(2)
list.Add(3)
list.Print()  // 1 -> 2 -> 3 -> nil
```

---

## Practice Questions - Array & Slices

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. Arrays in Go have a __________ size that cannot be changed.
2. The size of an array is part of its __________.
3. A slice is a __________ view into an array.
4. The __________ function adds elements to a slice.
5. Slices have both a __________ and a capacity.
6. When a slice's capacity is exceeded, Go allocates a __________ array.

### True/False

1. ⬜ Arrays can grow dynamically in Go
2. ⬜ Slices are passed by reference (share underlying array)
3. ⬜ `[3]int` and `[5]int` are the same type
4. ⬜ A nil slice has length 0
5. ⬜ The `copy` function creates an independent copy of a slice
6. ⬜ Slicing creates a new independent array

### Multiple Choice

1. What is the output?
   ```go
   nums := []int{1, 2, 3, 4, 5}
   fmt.Println(nums[1:4])
   ```
   - A) [1 2 3 4]
   - B) [2 3 4]
   - C) [1 2 3]
   - D) [2 3 4 5]

2. What happens when you exceed a slice's capacity?
   - A) Compile error
   - B) Runtime panic
   - C) Go allocates a larger array and copies elements
   - D) Elements are silently dropped

3. What does this print?
   ```go
   s := make([]int, 3, 5)
   fmt.Println(len(s), cap(s))
   ```
   - A) 3 3
   - B) 5 5
   - C) 3 5
   - D) 5 3

4. How do you create an independent copy of a slice?
   - A) `copied := original`
   - B) `copied := original[:]`
   - C) `copied := make([]int, len(original)); copy(copied, original)`
   - D) `copied := append([]int{}, original)`

### Code Challenge

Write a function that removes duplicate values from a slice of integers while preserving order.

---

### Answers - Array & Slices

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. fixed
2. type
3. dynamic
4. append
5. length
6. new (larger)

**True/False:**
1. ❌ False (arrays have fixed size; use slices for dynamic size)
2. ✅ True
3. ❌ False (different types)
4. ✅ True
5. ✅ True
6. ❌ False (shares underlying array)

**Multiple Choice:**
1. **B** - [2 3 4] (indices 1, 2, 3; high index is exclusive)
2. **C** - Go allocates a larger array and copies elements
3. **C** - 3 5 (length 3, capacity 5)
4. **C** - `copied := make([]int, len(original)); copy(copied, original)` (most explicit; D also works)

**Code Challenge:**
```go
func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	
	return result
}

// Usage
nums := []int{1, 2, 2, 3, 4, 3, 5}
unique := removeDuplicates(nums)
fmt.Println(unique)  // [1 2 3 4 5]
```

</details>

</details>

---

## Practice Questions - Maps, Structs & Pointers

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. Maps in Go are __________ (ordered/unordered) collections of key-value pairs.
2. The __________ function removes a key from a map.
3. Structs are __________ types (value/reference).
4. The `&` operator gets the __________ of a variable.
5. The `*` operator __________ a pointer.
6. Map iteration order is __________ (guaranteed/random).

### True/False

1. ⬜ Accessing a non-existent map key returns an error
2. ⬜ Structs with slices cannot be compared with ==
3. ⬜ Pointers allow functions to modify original values
4. ⬜ Maps are passed by value (copied)
5. ⬜ Embedded structs promote their fields to the parent struct
6. ⬜ Dereferencing a nil pointer is safe

### Multiple Choice

1. What does this print?
   ```go
   m := map[string]int{"a": 1}
   fmt.Println(m["b"])
   ```
   - A) Error
   - B) nil
   - C) 0
   - D) Panic

2. How do you create a pointer to an integer with value 42?
   - A) `p := &42`
   - B) `x := 42; p := &x`
   - C) `p := *42`
   - D) `p := new(42)`

3. What is the output?
   ```go
   type Point struct{ X, Y int }
   p1 := Point{1, 2}
   p2 := p1
   p2.X = 10
   fmt.Println(p1.X)
   ```
   - A) 10
   - B) 1
   - C) 0
   - D) Error

4. Which is true about maps?
   - A) Keys must be strings
   - B) Values must be comparable
   - C) Iteration order is random
   - D) They are value types

### Code Challenge

Write a function that merges two maps, with values from the second map overwriting values from the first for duplicate keys.

---

### Answers - Maps, Structs & Pointers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. unordered
2. delete
3. value
4. address
5. dereferences
6. random

**True/False:**
1. ❌ False (returns zero value)
2. ✅ True
3. ✅ True
4. ❌ False (maps are reference types)
5. ✅ True
6. ❌ False (causes panic)

**Multiple Choice:**
1. **C** - 0 (zero value for int)
2. **B** - `x := 42; p := &x`
3. **B** - 1 (structs are value types, p2 is a copy)
4. **C** - Iteration order is random

**Code Challenge:**
```go
func mergeMaps(m1, m2 map[string]int) map[string]int {
	result := make(map[string]int)
	
	// Copy m1
	for k, v := range m1 {
		result[k] = v
	}
	
	// Merge m2 (overwrites duplicates)
	for k, v := range m2 {
		result[k] = v
	}
	
	return result
}

// Usage
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 20, "c": 3}
merged := mergeMaps(m1, m2)
fmt.Println(merged)  // map[a:1 b:20 c:3]
```

</details>

</details>

---

## Summary

In this section, you learned:

✅ Arrays (fixed-size, value types)  
✅ Array declaration and initialization  
✅ Slices (dynamic, flexible views)  
✅ Slice operations (append, slicing, copy)  
✅ Length vs capacity  
✅ Memory sharing between slices  
✅ Common slice patterns  
✅ Maps (key-value pairs, unordered)  
✅ Map operations (add, access, delete, iterate)  
✅ Structs (custom types, fields, methods)  
✅ Struct composition and embedding  
✅ Pointers (addresses, dereferencing)  
✅ When to use pointers vs values  

**Next Steps:**
- Continue to [Chapter 7: Interfaces →](7-interfaces.md)
- Practice map and struct operations
- Complete the merge maps challenge

---

**Navigation:** [← Functions](5-functions.md) | [README](./README.md) | [Next: Interfaces →](7-interfaces.md)