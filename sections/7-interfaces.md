# 7. Interfaces & Polymorphism

> Master interfaces, polymorphism, and Go's approach to abstraction.

**Navigation:** [← Data Structures](6-data-structures-part2.md) | [README](./README.md) | [Next: Error Handling →](8-error-handling.md)

---

## Table of Contents

- [7.1 Interface Basics](#71-interface-basics)
- [7.2 Interface Implementation](#72-interface-implementation)
- [7.3 Empty Interface](#73-empty-interface)
- [7.4 Type Assertions & Type Switches](#74-type-assertions--type-switches)
- [7.5 Common Interfaces](#75-common-interfaces)
- [7.6 Interface Best Practices](#76-interface-best-practices)
- [Practice Questions](#practice-questions)

---

## 7.1 Interface Basics

An **interface** is a type that defines a **set of method signatures**.

**Think of interfaces as:**
- Contracts that types must fulfill
- Behavior definitions (not data)
- Go's approach to polymorphism

### Interface Declaration

```go
type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}
```

**Key Points:**
- Interface defines **what** methods, not **how** they're implemented
- No data fields (only methods)
- Type names often end with "-er" (Writer, Reader, Stringer)

### Why Interfaces?

**Without interfaces:**
```go
func printRectangleArea(r Rectangle) {
	fmt.Println(r.Area())
}

func printCircleArea(c Circle) {
	fmt.Println(c.Area())
}

// Need separate function for each shape type!
```

**With interfaces:**
```go
type Shape interface {
	Area() float64
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

// Works with ANY type that implements Area()
printArea(rect)
printArea(circle)
printArea(triangle)
```

---

## 7.2 Interface Implementation

**Go uses implicit implementation** - no explicit "implements" keyword!

### Implicit Implementation

```go
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle implements Shape (implicitly)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Circle implements Shape (implicitly)
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
```

**Using the interface:**
```go
func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

rect := Rectangle{Width: 10, Height: 5}
circle := Circle{Radius: 7}

printArea(rect)   // Area: 50.00
printArea(circle) // Area: 153.94
```

### Multiple Interfaces

A type can implement multiple interfaces:

```go
type Shape interface {
	Area() float64
}

type Drawable interface {
	Draw()
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Implements Shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implements Drawable
func (r Rectangle) Draw() {
	fmt.Println("Drawing rectangle")
}

// Rectangle implements both Shape and Drawable!
```

### Interface Embedding

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter embeds both Reader and Writer
type ReadWriter interface {
	Reader
	Writer
}

// ReadWriteCloser embeds all three
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
```

### Real-World Example: Payment Processing

```go
type PaymentProcessor interface {
	ProcessPayment(amount float64) error
	RefundPayment(transactionID string) error
}

type CreditCard struct {
	Number string
}

func (cc CreditCard) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via credit card\n", amount)
	// Credit card processing logic
	return nil
}

func (cc CreditCard) RefundPayment(transactionID string) error {
	fmt.Printf("Refunding transaction %s\n", transactionID)
	return nil
}

type PayPal struct {
	Email string
}

func (pp PayPal) ProcessPayment(amount float64) error {
	fmt.Printf("Processing $%.2f via PayPal\n", amount)
	// PayPal processing logic
	return nil
}

func (pp PayPal) RefundPayment(transactionID string) error {
	fmt.Printf("Refunding PayPal transaction %s\n", transactionID)
	return nil
}

// One function works with all payment methods
func checkout(processor PaymentProcessor, amount float64) error {
	return processor.ProcessPayment(amount)
}

// Usage
cc := CreditCard{Number: "1234-5678"}
pp := PayPal{Email: "user@example.com"}

checkout(cc, 99.99)
checkout(pp, 49.99)
```

---

## 7.3 Empty Interface

The **empty interface** `interface{}` has **zero methods**, so **any type** implements it.

### Empty Interface Basics

```go
// Can hold any type
var i interface{}

i = 42
fmt.Println(i)  // 42

i = "hello"
fmt.Println(i)  // hello

i = true
fmt.Println(i)  // true

i = struct{ Name string }{"Alice"}
fmt.Println(i)  // {Alice}
```

### Common Uses

**1. Generic containers:**
```go
func printAnything(v interface{}) {
	fmt.Println(v)
}

printAnything(42)
printAnything("hello")
printAnything([]int{1, 2, 3})
```

**2. Heterogeneous slices:**
```go
things := []interface{}{
	42,
	"hello",
	true,
	3.14,
}

for _, thing := range things {
	fmt.Printf("%v (%T)\n", thing, thing)
}
// Output:
// 42 (int)
// hello (string)
// true (bool)
// 3.14 (float64)
```

**3. JSON unmarshaling:**
```go
var data interface{}
json.Unmarshal([]byte(`{"name":"Alice","age":30}`), &data)

// data is now map[string]interface{}
```

### Modern Alternative: `any`

Go 1.18+ introduced `any` as an alias for `interface{}`:

```go
// These are equivalent
var x interface{}
var y any

// Prefer "any" (more readable)
func printValue(v any) {
	fmt.Println(v)
}
```

---

## 7.4 Type Assertions & Type Switches

### Type Assertion

**Extract concrete type from interface:**

```go
var i interface{} = "hello"

// Type assertion (unsafe)
s := i.(string)
fmt.Println(s)  // "hello"

// Wrong type causes panic!
// n := i.(int)  // ❌ panic: interface conversion

// Safe type assertion (comma-ok idiom)
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

### Type Switch

**Check multiple types:**

```go
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	case []int:
		fmt.Printf("Int slice: %v\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

describe(42)           // Integer: 42
describe("hello")      // String: hello
describe(true)         // Boolean: true
describe([]int{1, 2})  // Int slice: [1 2]
describe(3.14)         // Unknown type: float64
```

### Real-World Example: JSON Processing

```go
func processJSON(data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		// JSON object
		fmt.Println("Object with keys:", len(v))
		for key, val := range v {
			fmt.Printf("  %s: %v\n", key, val)
		}
	case []interface{}:
		// JSON array
		fmt.Println("Array with", len(v), "elements")
		for i, val := range v {
			fmt.Printf("  [%d]: %v\n", i, val)
		}
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Number:", v)
	case bool:
		fmt.Println("Boolean:", v)
	case nil:
		fmt.Println("Null")
	default:
		fmt.Println("Unknown type")
	}
}

var data interface{}
json.Unmarshal([]byte(`{"name":"Alice","age":30}`), &data)
processJSON(data)
```

---

## 7.5 Common Interfaces

### fmt.Stringer

**Control string representation:**

```go
type Stringer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

p := Person{Name: "Alice", Age: 30}
fmt.Println(p)  // Alice (30 years old)
```

### error Interface

```go
type error interface {
	Error() string
}

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Usage
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Message: "must be non-negative",
		}
	}
	return nil
}
```

### io.Reader and io.Writer

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// Example: Copy from reader to writer
func copy(dst Writer, src Reader) error {
	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		
		_, err = dst.Write(buf[:n])
		if err != nil {
			return err
		}
	}
	return nil
}
```

### sort.Interface

```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Usage
people := []Person{
	{"Bob", 31},
	{"Alice", 25},
	{"Carol", 28},
}

sort.Sort(ByAge(people))
fmt.Println(people)
// [{Alice 25} {Carol 28} {Bob 31}]
```

---

## 7.6 Interface Best Practices

### Accept Interfaces, Return Structs

**✅ Good:**
```go
// Accept interface (flexible)
func SaveUser(w io.Writer, user User) error {
	data, _ := json.Marshal(user)
	_, err := w.Write(data)
	return err
}

// Return concrete type (clear)
func NewUser(name string) *User {
	return &User{Name: name}
}
```

**❌ Avoid:**
```go
// Returning interface (less clear)
func NewUser(name string) UserInterface {
	return &User{Name: name}
}
```

### Small Interfaces

**✅ Good (focused):**
```go
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}
```

**❌ Avoid (too broad):**
```go
type FileSystem interface {
	Open(name string) (*File, error)
	Create(name string) (*File, error)
	Remove(name string) error
	Rename(oldpath, newpath string) error
	Stat(name string) (FileInfo, error)
	// ... 10 more methods
}
```

**Rob Pike's Proverb:**
> "The bigger the interface, the weaker the abstraction."

### Define Interfaces at Point of Use

**✅ Good (consumer defines):**
```go
// In package A (consumer)
type UserStore interface {
	GetUser(id int) (*User, error)
}

func ProcessUser(store UserStore, id int) error {
	user, err := store.GetUser(id)
	// ...
}

// In package B (provider)
type Database struct{}

func (db *Database) GetUser(id int) (*User, error) {
	// Implementation
}

// Database implements UserStore (implicitly)
```

**❌ Avoid (provider defines):**
```go
// In package B (provider)
type DatabaseInterface interface {
	GetUser(id int) (*User, error)
	SaveUser(user *User) error
	DeleteUser(id int) error
	// Many methods consumer doesn't need
}
```

### Real-World Example: Testing with Interfaces

**Production code:**
```go
type EmailSender interface {
	Send(to, subject, body string) error
}

type UserService struct {
	emailer EmailSender
}

func (s *UserService) RegisterUser(email string) error {
	// ... create user
	
	err := s.emailer.Send(email, "Welcome", "Welcome to our service!")
	return err
}
```

**Real implementation:**
```go
type SMTPEmailer struct {
	host string
	port int
}

func (s *SMTPEmailer) Send(to, subject, body string) error {
	// Real SMTP logic
	return nil
}
```

**Test mock:**
```go
type MockEmailer struct {
	SentEmails []string
}

func (m *MockEmailer) Send(to, subject, body string) error {
	m.SentEmails = append(m.SentEmails, to)
	return nil
}

// Test
func TestRegisterUser(t *testing.T) {
	mock := &MockEmailer{}
	service := &UserService{emailer: mock}
	
	service.RegisterUser("test@example.com")
	
	if len(mock.SentEmails) != 1 {
		t.Error("Expected 1 email")
	}
}
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. Interfaces in Go are implemented __________ (implicitly/explicitly).
2. The empty interface is written as __________ or `any` in Go 1.18+.
3. Type assertions use the syntax __________.
4. A type can implement __________ (one/multiple) interfaces.
5. The __________ interface has a single method: Error() string.
6. Interfaces should be defined at the __________ of use, not the provider.

### True/False

1. ⬜ Go requires the "implements" keyword to implement interfaces
2. ⬜ The empty interface can hold values of any type
3. ⬜ Type assertions always succeed without errors
4. ⬜ A type can implement multiple interfaces
5. ⬜ Interfaces can contain both methods and data fields
6. ⬜ Smaller interfaces provide stronger abstractions

### Multiple Choice

1. What does this code do?
   ```go
   var i interface{} = 42
   s := i.(string)
   ```
   - A) Converts 42 to "42"
   - B) Compiles but panics at runtime
   - C) Compile error
   - D) Returns empty string

2. Which is the correct way to safely assert a type?
   - A) `s := i.(string)`
   - B) `s, ok := i.(string)`
   - C) `s = string(i)`
   - D) `s := i.string()`

3. What makes a type implement an interface?
   - A) Using the "implements" keyword
   - B) Inheriting from the interface
   - C) Having all the interface's methods
   - D) Explicitly declaring implementation

4. Which interface has zero methods?
   - A) `error`
   - B) `interface{}`
   - C) `io.Reader`
   - D) `fmt.Stringer`

### Code Challenge

Create an interface `Animal` with a `Speak() string` method. Implement it for `Dog` and `Cat` types, then write a function that takes an `Animal` and prints what it says.

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. implicitly
2. interface{}
3. `i.(Type)` or `v := i.(Type)`
4. multiple
5. error
6. point (consumer/caller)

**True/False:**
1. ❌ False (implicit implementation)
2. ✅ True
3. ❌ False (can panic if wrong type)
4. ✅ True
5. ❌ False (only methods, no fields)
6. ✅ True (smaller is better)

**Multiple Choice:**
1. **B** - Compiles but panics at runtime (42 is int, not string)
2. **B** - `s, ok := i.(string)` (safe with comma-ok idiom)
3. **C** - Having all the interface's methods
4. **B** - `interface{}` (empty interface)

**Code Challenge:**
```go
type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func makeSound(a Animal) {
	fmt.Printf("%s says: %s\n", 
		reflect.TypeOf(a).Name(), 
		a.Speak())
}

// Usage
dog := Dog{Name: "Buddy"}
cat := Cat{Name: "Whiskers"}

makeSound(dog)  // Dog says: Woof!
makeSound(cat)  // Cat says: Meow!

// Alternative without reflection:
func makeSound2(a Animal, name string) {
	fmt.Printf("%s says: %s\n", name, a.Speak())
}
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ Interface basics and declaration  
✅ Implicit interface implementation  
✅ Interface embedding and composition  
✅ Empty interface (`interface{}` / `any`)  
✅ Type assertions and type switches  
✅ Common standard library interfaces  
✅ Interface best practices (accept interfaces, return structs)  

**Next Steps:**
- Continue to [Chapter 8: Error Handling →](8-error-handling.md)
- Practice implementing interfaces
- Complete the Animal interface challenge above

---

**Navigation:** [← Data Structures](6-data-structures-part2.md) | [README](./README.md) | [Next: Error Handling →](8-error-handling.md)