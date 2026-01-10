# 8. Error Handling

> Master Go's error handling patterns and best practices.

**Navigation:** [← Interfaces](7-interfaces.md) | [README](./README.md) | [Next: Concurrency →](9-concurrency.md)

---

## Table of Contents

- [8.1 Error Basics](#81-error-basics)
- [8.2 Creating Errors](#82-creating-errors)
- [8.3 Error Wrapping](#83-error-wrapping)
- [8.4 Custom Errors](#84-custom-errors)
- [8.5 Error Handling Patterns](#85-error-handling-patterns)
- [8.6 Panic and Recover](#86-panic-and-recover)
- [Practice Questions](#practice-questions)

---

## 8.1 Error Basics

Go uses **explicit error values** instead of exceptions.

### The error Interface

```go
type error interface {
	Error() string
}
```

**Any type with an `Error() string` method implements the error interface.**

### Basic Error Handling

```go
import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println("Result:", result)
```

**Key Pattern:**
- Functions return `(value, error)`
- Check `if err != nil` immediately
- Handle error before using value

### Multiple Return Values

```go
func getUser(id int) (User, error) {
	if id < 0 {
		return User{}, errors.New("invalid user ID")
	}
	// ... fetch user
	return user, nil
}

// Usage
user, err := getUser(5)
if err != nil {
	log.Fatal(err)
}
fmt.Println(user)
```

### Error Checking is Mandatory

```go
// ❌ Bad: Ignoring error
data, _ := os.ReadFile("config.json")

// ✅ Good: Handling error
data, err := os.ReadFile("config.json")
if err != nil {
	return fmt.Errorf("failed to read config: %w", err)
}
```

---

## 8.2 Creating Errors

### errors.New()

**Simple error messages:**

```go
import "errors"

err := errors.New("something went wrong")
```

### fmt.Errorf()

**Formatted error messages:**

```go
import "fmt"

func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("invalid age: %d", age)
	}
	if age > 150 {
		return fmt.Errorf("age %d is unrealistic", age)
	}
	return nil
}
```

### Sentinel Errors

**Predefined errors for comparison:**

```go
import "errors"

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalidInput = errors.New("invalid input")
)

func getUser(id int) (User, error) {
	if id < 0 {
		return User{}, ErrInvalidInput
	}
	
	user, found := database[id]
	if !found {
		return User{}, ErrNotFound
	}
	
	return user, nil
}

// Usage
user, err := getUser(5)
if err == ErrNotFound {
	fmt.Println("User not found")
} else if err == ErrInvalidInput {
	fmt.Println("Invalid ID")
} else if err != nil {
	fmt.Println("Unknown error:", err)
}
```

**Standard library examples:**
```go
io.EOF           // End of file
sql.ErrNoRows    // No rows in result set
os.ErrNotExist   // File does not exist
```

---

## 8.3 Error Wrapping

**Go 1.13+ provides error wrapping** to add context while preserving the original error.

### Wrapping Errors with %w

```go
import "fmt"

func readConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}
	
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}
	
	return config, nil
}

// Error chain:
// "parse config: invalid character '}' looking for beginning of value"
```

### Unwrapping Errors

```go
import "errors"

// Check if error is a specific type
if errors.Is(err, os.ErrNotExist) {
	fmt.Println("File does not exist")
}

// Extract specific error type
var pathErr *os.PathError
if errors.As(err, &pathErr) {
	fmt.Println("Path error:", pathErr.Path)
}
```

### errors.Is() vs ==

```go
var ErrNotFound = errors.New("not found")

func getUser(id int) error {
	return fmt.Errorf("user lookup failed: %w", ErrNotFound)
}

err := getUser(5)

// ❌ Wrong: Direct comparison fails
if err == ErrNotFound {
	// This doesn't work!
}

// ✅ Correct: Use errors.Is
if errors.Is(err, ErrNotFound) {
	fmt.Println("User not found")  // This works!
}
```

### errors.As() for Type Assertions

```go
import (
	"errors"
	"os"
)

func processFile(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open: %w", err)
	}
	return nil
}

err := processFile("missing.txt")

// Extract specific error type
var pathErr *os.PathError
if errors.As(err, &pathErr) {
	fmt.Println("Operation:", pathErr.Op)
	fmt.Println("Path:", pathErr.Path)
	fmt.Println("Error:", pathErr.Err)
}
```

### Real-World Example: Error Chain

```go
func loadUserProfile(userID int) error {
	user, err := getUser(userID)
	if err != nil {
		return fmt.Errorf("load profile for user %d: %w", userID, err)
	}
	
	profile, err := getProfile(user.ProfileID)
	if err != nil {
		return fmt.Errorf("load profile %d: %w", user.ProfileID, err)
	}
	
	// ... use profile
	return nil
}

// Error output:
// load profile for user 123: load profile 456: database connection failed
```

---

## 8.4 Custom Errors

### Custom Error Types

```go
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{
			Field:   "email",
			Message: "must contain @ symbol",
		}
	}
	return nil
}

// Usage
err := validateEmail("invalid-email")
if err != nil {
	var validErr *ValidationError
	if errors.As(err, &validErr) {
		fmt.Printf("Field: %s, Message: %s\n", validErr.Field, validErr.Message)
	}
}
```

### Error with Additional Context

```go
type APIError struct {
	StatusCode int
	Message    string
	Endpoint   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Endpoint, e.Message)
}

func (e *APIError) Temporary() bool {
	return e.StatusCode >= 500
}

func callAPI(endpoint string) error {
	// ... make request
	return &APIError{
		StatusCode: 503,
		Message:    "service unavailable",
		Endpoint:   endpoint,
	}
}

// Usage
err := callAPI("/users")
if err != nil {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		if apiErr.Temporary() {
			// Retry logic
			fmt.Println("Temporary error, retrying...")
		}
	}
}
```

### Multiple Errors

```go
import "errors"

type MultiError []error

func (m MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

func validateUser(user User) error {
	var errs MultiError
	
	if user.Name == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if user.Email == "" {
		errs = append(errs, errors.New("email is required"))
	}
	if user.Age < 0 {
		errs = append(errs, errors.New("age must be positive"))
	}
	
	if len(errs) > 0 {
		return errs
	}
	return nil
}
```

---

## 8.5 Error Handling Patterns

### Early Return Pattern

```go
func processUser(id int) error {
	user, err := getUser(id)
	if err != nil {
		return err  // Early return
	}
	
	if !user.Active {
		return errors.New("user inactive")  // Early return
	}
	
	if user.Age < 18 {
		return errors.New("user too young")  // Early return
	}
	
	// Happy path - no nesting
	return saveUser(user)
}
```

### Error Wrapping Pattern

```go
func saveUserToDatabase(user User) error {
	err := db.Insert(user)
	if err != nil {
		return fmt.Errorf("save user %s: %w", user.Name, err)
	}
	return nil
}
```

### Retry Pattern

```go
func retryOperation(operation func() error, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		
		// Check if error is temporary
		if temp, ok := err.(interface{ Temporary() bool }); ok && temp.Temporary() {
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		
		// Non-temporary error, don't retry
		return err
	}
	return fmt.Errorf("max retries exceeded: %w", err)
}

// Usage
err := retryOperation(func() error {
	return callExternalAPI()
}, 3)
```

### Cleanup Pattern

```go
func processFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()  // Cleanup on all exit paths
	
	// Process file
	// Any return or panic will still call Close()
	
	return nil
}
```

### Error Logging Pattern

```go
import "log"

func handleRequest(w http.ResponseWriter, r *http.Request) {
	err := processRequest(r)
	if err != nil {
		log.Printf("Error processing request: %v", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	
	w.WriteHeader(200)
}
```

---

## 8.6 Panic and Recover

**panic** stops normal execution and **unwinds the stack**.  
**recover** catches a panic and resumes execution.

### When to Use Panic

**Use panic for:**
- Programming errors (bugs)
- Unrecoverable errors
- Initialization failures

**Don't use panic for:**
- Expected errors
- Business logic errors
- Input validation

### Panic Basics

```go
func mustConnect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)  // Unrecoverable - can't proceed
	}
	return db
}

// Program terminates if connection fails
```

### Recover from Panic

```go
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	
	// Code that might panic
	panic("something went wrong")
	
	fmt.Println("This won't execute")
}

safeFunction()
fmt.Println("Program continues")

// Output:
// Recovered from panic: something went wrong
// Program continues
```

### Real-World Example: HTTP Server Recovery

```go
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				log.Printf("Stack trace: %s", debug.Stack())
				http.Error(w, "Internal Server Error", 500)
			}
		}()
		
		next.ServeHTTP(w, r)
	})
}
```

### Panic vs Error

```go
// ✅ Use error for expected failures
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// ❌ Don't panic for expected errors
func badDivide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")  // Bad!
	}
	return a / b
}

// ✅ Panic for programming errors
func mustPositive(n int) {
	if n < 0 {
		panic("n must be positive")  // OK - programming error
	}
}
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. In Go, errors are values that implement the __________ interface.
2. The __________ verb in fmt.Errorf wraps an error.
3. Use __________ to check if an error is a specific sentinel error.
4. Use __________ to extract a specific error type from an error chain.
5. __________ stops normal execution and unwinds the stack.
6. __________ catches a panic and resumes execution.

### True/False

1. ⬜ Go uses exceptions for error handling
2. ⬜ Ignoring errors is a compile-time error in Go
3. ⬜ errors.Is() can unwrap error chains
4. ⬜ Panic should be used for expected errors
5. ⬜ defer functions run even when a panic occurs
6. ⬜ Multiple errors can be returned from a function

### Multiple Choice

1. What is the correct way to wrap an error?
   - A) `fmt.Errorf("context: %v", err)`
   - B) `fmt.Errorf("context: %w", err)`
   - C) `errors.Wrap(err, "context")`
   - D) `err.Wrap("context")`

2. Which is true about sentinel errors?
   - A) They are predefined error variables
   - B) They can be compared with ==
   - C) They should be compared with errors.Is()
   - D) All of the above

3. When should you use panic?
   - A) For input validation errors
   - B) For network failures
   - C) For programming errors
   - D) For business logic errors

4. What does recover() return if there's no panic?
   - A) error
   - B) nil
   - C) false
   - D) panic

### Code Challenge

Write a function `safeDivide(a, b float64) (float64, error)` that:
- Returns an error for division by zero
- Wraps any unexpected panics into an error
- Returns the result on success

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. error
2. %w
3. errors.Is()
4. errors.As()
5. panic
6. recover

**True/False:**
1. ❌ False (Go uses explicit error values)
2. ❌ False (ignored errors are allowed but discouraged)
3. ✅ True
4. ❌ False (use errors for expected failures)
5. ✅ True
6. ❌ False (only one error returned, but can contain multiple errors)

**Multiple Choice:**
1. **B** - `fmt.Errorf("context: %w", err)` (wraps error)
2. **D** - All of the above
3. **C** - For programming errors (bugs, unrecoverable states)
4. **B** - nil

**Code Challenge:**
```go
func safeDivide(a, b float64) (result float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic during division: %v", r)
		}
	}()
	
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	
	return a / b, nil
}

// Usage
result, err := safeDivide(10, 2)
if err != nil {
	log.Println("Error:", err)
} else {
	fmt.Println("Result:", result)
}

// Test with zero
result, err = safeDivide(10, 0)
if err != nil {
	log.Println("Error:", err)  // division by zero
}
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ Error interface and error handling basics  
✅ Creating errors with errors.New() and fmt.Errorf()  
✅ Sentinel errors for predefined error values  
✅ Error wrapping with %w (Go 1.13+)  
✅ errors.Is() and errors.As() for error checking  
✅ Custom error types with additional context  
✅ Error handling patterns (early return, wrapping, retry)  
✅ Panic and recover for exceptional situations  

**Next Steps:**
- Continue to [Chapter 9: Concurrency →](9-concurrency.md)
- Practice error wrapping and custom errors
- Complete the safe divide challenge above

---

**Navigation:** [← Interfaces](7-interfaces.md) | [README](./README.md) | [Next: Concurrency →](9-concurrency.md)