# 13. Best Practices

> Master Go idioms, design patterns, and professional development practices.

**Navigation:** [← Web Development](12-web-development.md) | [README](./README.md) | [Next: Interview Questions →](14-interview-questions.md)

---

## Table of Contents

- [13.1 Code Organization](#131-code-organization)
- [13.2 Error Handling](#132-error-handling)
- [13.3 Naming Conventions](#133-naming-conventions)
- [13.4 Performance](#134-performance)
- [13.5 Concurrency Patterns](#135-concurrency-patterns)
- [13.6 Security](#136-security)
- [13.7 Documentation](#137-documentation)
- [Practice Questions](#practice-questions)

---

## 13.1 Code Organization

### Project Structure

**Standard Go project layout:**

```
myproject/
├── cmd/                    # Main applications
│   └── myapp/
│       └── main.go
├── internal/               # Private application code
│   ├── handlers/
│   ├── models/
│   └── services/
├── pkg/                    # Public library code
│   └── utils/
├── api/                    # API definitions (OpenAPI, protobuf)
├── web/                    # Web assets
├── configs/                # Configuration files
├── scripts/                # Build and deployment scripts
├── test/                   # Additional test data
├── docs/                   # Documentation
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### Package Organization

**✅ Good:**
```go
// Package user provides user management functionality
package user

import (
	"myproject/internal/database"
	"myproject/pkg/validation"
)

type Service struct {
	db database.DB
}

func NewService(db database.DB) *Service {
	return &Service{db: db}
}
```

**❌ Bad:**
```go
// Everything in one package
package main

// Hundreds of types and functions...
```

### Circular Dependencies

**❌ Avoid:**
```
package a imports package b
package b imports package a
```

**✅ Solution:**
- Extract shared code to new package
- Use interfaces
- Restructure dependencies

---

## 13.2 Error Handling

### Error Wrapping

**✅ Good:**
```go
func processUser(id int) error {
	user, err := getUser(id)
	if err != nil {
		return fmt.Errorf("process user %d: %w", id, err)
	}
	
	if err := validateUser(user); err != nil {
		return fmt.Errorf("validate user %d: %w", id, err)
	}
	
	return nil
}
```

**❌ Bad:**
```go
func processUser(id int) error {
	user, err := getUser(id)
	if err != nil {
		return err  // Lost context
	}
	return nil
}
```

### Early Returns

**✅ Good:**
```go
func validateUser(user User) error {
	if user.Name == "" {
		return errors.New("name required")
	}
	
	if user.Email == "" {
		return errors.New("email required")
	}
	
	if user.Age < 0 {
		return errors.New("invalid age")
	}
	
	return nil
}
```

**❌ Bad:**
```go
func validateUser(user User) error {
	if user.Name != "" {
		if user.Email != "" {
			if user.Age >= 0 {
				return nil
			}
			return errors.New("invalid age")
		}
		return errors.New("email required")
	}
	return errors.New("name required")
}
```

### Don't Panic

**✅ Use errors:**
```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
```

**❌ Don't panic for expected errors:**
```go
func divide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")  // Bad!
	}
	return a / b
}
```

---

## 13.3 Naming Conventions

### Variable Names

**✅ Good:**
```go
// Short, clear names in small scopes
for i := 0; i < 10; i++ {
	// i is clear here
}

// Descriptive names in larger scopes
userRepository := NewRepository()
```

**❌ Bad:**
```go
var usrRepo Repository  // Abbreviated unnecessarily
var theUserRepositoryForTheApplication Repository  // Too long
```

### Function Names

**✅ Good:**
```go
func GetUser(id int) (*User, error)
func ValidateEmail(email string) bool
func ParseJSON(data []byte) (interface{}, error)
```

**❌ Bad:**
```go
func get_user(id int) (*User, error)  // Snake case
func GETUSER(id int) (*User, error)   // All caps
```

### Interface Names

**✅ Good:**
```go
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Stringer interface {
	String() string
}
```

**Pattern:** Single-method interfaces often end in `-er`

### Package Names

**✅ Good:**
```go
package http
package json
package user
```

**❌ Bad:**
```go
package httputils  // Don't use "utils" suffix
package UserPackage  // Don't capitalize
package user_service  // No underscores
```

---

## 13.4 Performance

### Slice Preallocation

**✅ Good:**
```go
// Known size
users := make([]User, 0, 100)
for i := 0; i < 100; i++ {
	users = append(users, User{ID: i})
}
```

**❌ Bad:**
```go
var users []User  // Reallocates multiple times
for i := 0; i < 100; i++ {
	users = append(users, User{ID: i})
}
```

### String Concatenation

**✅ Good:**
```go
var builder strings.Builder
for i := 0; i < 1000; i++ {
	builder.WriteString("a")
}
result := builder.String()
```

**❌ Bad:**
```go
result := ""
for i := 0; i < 1000; i++ {
	result += "a"  // Creates new string each time
}
```

### Map Preallocation

**✅ Good:**
```go
m := make(map[string]int, 100)  // Preallocate
```

**❌ Bad:**
```go
m := make(map[string]int)  // Will grow multiple times
```

### Struct Field Ordering

**✅ Good (aligned):**
```go
type Good struct {
	a int64   // 8 bytes
	b int64   // 8 bytes
	c int32   // 4 bytes
	d int32   // 4 bytes
}
// Total: 24 bytes
```

**❌ Bad (padding):**
```go
type Bad struct {
	a int64   // 8 bytes
	c int32   // 4 bytes + 4 padding
	b int64   // 8 bytes
	d int32   // 4 bytes + 4 padding
}
// Total: 32 bytes (33% larger!)
```

### Pointer vs Value

**Use pointers when:**
- Struct is large (>64 bytes)
- Need to modify receiver
- Need to share state

**Use values when:**
- Struct is small
- Immutability desired
- Avoiding heap allocations

```go
// Small struct - use value
type Point struct {
	X, Y int
}

// Large struct - use pointer
type User struct {
	ID        int
	Name      string
	Email     string
	Profile   Profile
	Settings  Settings
	// ... many fields
}

func (u *User) UpdateEmail(email string) {
	u.Email = email  // Needs pointer to modify
}
```

---

## 13.5 Concurrency Patterns

### Channel Direction

**✅ Good:**
```go
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int) {
	for val := range in {
		fmt.Println(val)
	}
}
```

### Don't Leak Goroutines

**❌ Bad:**
```go
func search(query string) Result {
	ch := make(chan Result)
	go func() {
		ch <- doSearch(query)
	}()
	
	select {
	case result := <-ch:
		return result
	case <-time.After(time.Second):
		return Result{}  // Goroutine leaks!
	}
}
```

**✅ Good:**
```go
func search(query string) Result {
	ch := make(chan Result, 1)  // Buffered
	go func() {
		ch <- doSearch(query)  // Won't block
	}()
	
	select {
	case result := <-ch:
		return result
	case <-time.After(time.Second):
		return Result{}
	}
}
```

### Use Context for Cancellation

**✅ Good:**
```go
func worker(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// Do work
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	worker(ctx)
}
```

---

## 13.6 Security

### Input Validation

**✅ Always validate:**
```go
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	// Validate
	if user.Email == "" {
		http.Error(w, "Email required", http.StatusBadRequest)
		return
	}
	
	if !isValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	
	// Process...
}
```

### SQL Injection Prevention

**✅ Use parameterized queries:**
```go
func getUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	
	var user User
	err := row.Scan(&user.ID, &user.Name)
	return &user, err
}
```

**❌ Never concatenate:**
```go
func getUser(db *sql.DB, id string) (*User, error) {
	query := "SELECT * FROM users WHERE id = " + id  // SQL injection!
	// ...
}
```

### Password Hashing

**✅ Use bcrypt:**
```go
import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
```

### Environment Variables for Secrets

**✅ Good:**
```go
dbPassword := os.Getenv("DB_PASSWORD")
apiKey := os.Getenv("API_KEY")
```

**❌ Bad:**
```go
const dbPassword = "super-secret-password"  // Don't hardcode!
```

---

## 13.7 Documentation

### Package Documentation

```go
// Package user provides user management functionality.
// It includes authentication, authorization, and profile management.
package user
```

### Function Documentation

```go
// GetUser retrieves a user by ID from the database.
// It returns an error if the user is not found or if a database error occurs.
//
// Example:
//   user, err := GetUser(db, 123)
//   if err != nil {
//       log.Fatal(err)
//   }
func GetUser(db *sql.DB, id int) (*User, error) {
	// ...
}
```

### Type Documentation

```go
// User represents a user in the system.
type User struct {
	// ID is the unique identifier for the user
	ID int
	
	// Name is the user's full name
	Name string
	
	// Email is the user's email address
	Email string
}
```

### Examples

```go
func ExampleGetUser() {
	db := setupTestDB()
	user, err := GetUser(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Name)
	// Output: Alice
}
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. The __________ directory contains private application code.
2. Error messages should provide __________ about what went wrong.
3. Interface names with one method often end in __________.
4. Preallocating slices avoids multiple __________.
5. Use __________ to prevent SQL injection.
6. Package documentation appears before the __________ declaration.

### True/False

1. ⬜ Package names should use snake_case
2. ⬜ Panic should be used for expected errors
3. ⬜ Early returns reduce nesting
4. ⬜ String concatenation with + is efficient for loops
5. ⬜ Buffered channels can prevent goroutine leaks
6. ⬜ Secrets should be hardcoded in source files

### Multiple Choice

1. Which is better for large structs?
   - A) Value receivers
   - B) Pointer receivers
   - C) No receivers
   - D) Doesn't matter

2. How should errors be wrapped?
   - A) return err
   - B) return errors.New(err.Error())
   - C) return fmt.Errorf("context: %w", err)
   - D) panic(err)

3. What's the Go convention for private fields?
   - A) _fieldName
   - B) fieldName (lowercase)
   - C) FIELDNAME
   - D) FieldName_

4. Which prevents SQL injection?
   - A) String concatenation
   - B) Parameterized queries
   - C) Input sanitization only
   - D) None of the above

### Code Review

**Identify issues in this code:**
```go
package Utils

type user_data struct {
	name string
	AGE int
}

func Get_User(ID string) user_data {
	query := "SELECT * FROM users WHERE id = " + ID
	// ... execute query
	
	var u user_data
	return u
}
```

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. internal
2. context
3. -er (e.g., Reader, Writer)
4. reallocations
5. parameterized queries
6. package

**True/False:**
1. ❌ False (use lowercase, no underscores)
2. ❌ False (use errors for expected failures)
3. ✅ True
4. ❌ False (use strings.Builder)
5. ✅ True
6. ❌ False (use environment variables)

**Multiple Choice:**
1. **B** - Pointer receivers (avoid copying)
2. **C** - `return fmt.Errorf("context: %w", err)`
3. **B** - fieldName (lowercase)
4. **B** - Parameterized queries

**Code Review Issues:**
1. ❌ Package name "Utils" - should be lowercase "utils"
2. ❌ Type name "user_data" - should be "UserData" (or "userData" if private)
3. ❌ Field names inconsistent - "name" lowercase, "AGE" uppercase
4. ❌ Function name "Get_User" - should be "GetUser"
5. ❌ SQL injection vulnerability - concatenating ID directly
6. ❌ No error handling

**Fixed version:**
```go
package user

type User struct {
	Name string
	Age  int
}

func GetUser(db *sql.DB, id int) (*User, error) {
	query := "SELECT name, age FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	
	var u User
	err := row.Scan(&u.Name, &u.Age)
	if err != nil {
		return nil, fmt.Errorf("get user %d: %w", id, err)
	}
	
	return &u, nil
}
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ Code organization and project structure  
✅ Error handling best practices  
✅ Go naming conventions  
✅ Performance optimization techniques  
✅ Concurrency patterns and pitfalls  
✅ Security best practices  
✅ Documentation standards  

**Key Takeaways:**
- Accept interfaces, return structs
- Handle errors, don't ignore them
- Use clear, idiomatic names
- Preallocate when size is known
- Don't leak goroutines
- Validate all inputs
- Document public APIs

**Next Steps:**
- Continue to [Chapter 14: Interview Questions →](14-interview-questions.md)
- Review Go proverbs
- Practice code reviews

---

**Navigation:** [← Web Development](12-web-development.md) | [README](./README.md) | [Next: Interview Questions →](14-interview-questions.md)