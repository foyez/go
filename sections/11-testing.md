# 11. Testing

> Master testing, benchmarking, and test-driven development in Go.

**Navigation:** [← Generics](10-generics.md) | [README](./README.md) | [Next: Web Development →](12-web-development.md)

---

## Table of Contents

- [11.1 Testing Basics](#111-testing-basics)
- [11.2 Table-Driven Tests](#112-table-driven-tests)
- [11.3 Test Helpers & Subtests](#113-test-helpers--subtests)
- [11.4 Benchmarking](#114-benchmarking)
- [11.5 Test Coverage](#115-test-coverage)
- [11.6 Mocking & Interfaces](#116-mocking--interfaces)
- [11.7 Testing Packages](#117-testing-packages)
- [Practice Questions](#practice-questions)

---

## 11.1 Testing Basics

Go has built-in testing support via the `testing` package.

### Test File Structure

```
myproject/
├── math.go
└── math_test.go
```

**Rules:**
- Test files end with `_test.go`
- Test functions start with `Test`
- Test functions take `*testing.T` parameter

### Writing a Simple Test

**math.go:**
```go
package math

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}
```

**math_test.go:**
```go
package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}
```

### Running Tests

```bash
# Run all tests in current directory
go test

# Run with verbose output
go test -v

# Run specific test
go test -run TestAdd

# Run tests in all subdirectories
go test ./...

# Run tests with race detector
go test -race
```

### Test Methods

```go
func TestExample(t *testing.T) {
	// Report error but continue
	t.Error("something went wrong")
	t.Errorf("got %d, want %d", 5, 10)
	
	// Report error and stop test
	t.Fatal("critical error")
	t.Fatalf("got %d, want %d", 5, 10)
	
	// Log information (only shown with -v)
	t.Log("debug info")
	t.Logf("value: %d", 42)
	
	// Mark test as failed but continue
	t.Fail()
	
	// Skip test
	t.Skip("skipping this test")
	t.Skipf("skipping because %s", "reason")
}
```

### Testing with Setup/Teardown

```go
func TestMain(m *testing.M) {
	// Setup
	fmt.Println("Setting up tests")
	
	// Run tests
	code := m.Run()
	
	// Teardown
	fmt.Println("Cleaning up")
	
	os.Exit(code)
}

func TestSomething(t *testing.T) {
	// Individual test
}
```

---

## 11.2 Table-Driven Tests

**Best practice for testing multiple cases.**

### Basic Table-Driven Test

```go
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
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

### Table Test with Error Cases

```go
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name      string
		a, b      float64
		expected  float64
		expectErr bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative result", -10, 2, -5, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			
			if tt.expectErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			
			if result != tt.expected {
				t.Errorf("got %f, want %f", result, tt.expected)
			}
		})
	}
}
```

### Table Test Pattern (Advanced)

```go
func TestValidateEmail(t *testing.T) {
	tests := map[string]struct {
		email     string
		wantValid bool
	}{
		"valid email":           {"user@example.com", true},
		"missing @":             {"userexample.com", false},
		"missing domain":        {"user@", false},
		"missing username":      {"@example.com", false},
		"multiple @":            {"user@@example.com", false},
		"with subdomain":        {"user@mail.example.com", true},
		"with plus":             {"user+tag@example.com", true},
	}
	
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			valid := ValidateEmail(tt.email)
			if valid != tt.wantValid {
				t.Errorf("ValidateEmail(%q) = %v; want %v", 
					tt.email, valid, tt.wantValid)
			}
		})
	}
}
```

---

## 11.3 Test Helpers & Subtests

### Test Helpers

```go
func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()  // Marks this as helper function
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t *testing.T, err error, wantErr bool) {
	t.Helper()
	if (err != nil) != wantErr {
		t.Errorf("error = %v, wantErr = %v", err, wantErr)
	}
}

func TestSomething(t *testing.T) {
	result := Add(2, 3)
	assertEqual(t, result, 5)
	
	_, err := Divide(10, 0)
	assertError(t, err, true)
}
```

### Subtests

```go
func TestMath(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		result := Add(2, 3)
		if result != 5 {
			t.Errorf("got %d, want 5", result)
		}
	})
	
	t.Run("Subtraction", func(t *testing.T) {
		result := Subtract(5, 3)
		if result != 2 {
			t.Errorf("got %d, want 2", result)
		}
	})
	
	t.Run("Division", func(t *testing.T) {
		t.Run("Normal", func(t *testing.T) {
			result, err := Divide(10, 2)
			if err != nil {
				t.Fatal(err)
			}
			if result != 5 {
				t.Errorf("got %f, want 5.0", result)
			}
		})
		
		t.Run("ByZero", func(t *testing.T) {
			_, err := Divide(10, 0)
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	})
}
```

### Parallel Tests

```go
func TestParallel(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"test1", 1},
		{"test2", 2},
		{"test3", 3},
	}
	
	for _, tt := range tests {
		tt := tt  // Capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()  // Run in parallel
			
			// Test logic
			time.Sleep(time.Second)
			fmt.Println(tt.n)
		})
	}
}
```

---

## 11.4 Benchmarking

**Measure performance of functions.**

### Basic Benchmark

```go
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
```

**Run benchmarks:**
```bash
go test -bench=.
go test -bench=BenchmarkAdd
go test -bench=. -benchmem  # Include memory stats
```

### Benchmark with Setup

```go
func BenchmarkStringConcat(b *testing.B) {
	b.ResetTimer()  // Reset timer after setup
	
	for i := 0; i < b.N; i++ {
		result := ""
		for j := 0; j < 100; j++ {
			result += "a"
		}
	}
}
```

### Comparing Implementations

```go
// Using + operator
func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := ""
		for j := 0; j < 100; j++ {
			result += "a"
		}
	}
}

// Using strings.Builder
func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 100; j++ {
			builder.WriteString("a")
		}
		_ = builder.String()
	}
}
```

### Benchmark Table Tests

```go
func BenchmarkFibonacci(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{"Fib10", 10},
		{"Fib20", 20},
		{"Fib30", 30},
	}
	
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Fibonacci(bm.input)
			}
		})
	}
}
```

---

## 11.5 Test Coverage

**Measure which code is tested.**

### Running Coverage

```bash
# Generate coverage
go test -cover

# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out

# Coverage for all packages
go test -cover ./...
```

### Coverage Output

```bash
$ go test -cover
PASS
coverage: 85.7% of statements
ok      mypackage       0.123s
```

### Example Coverage Report

```go
// math.go
func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	if a > 100 {  // Not covered
		return a - b - 1
	}
	return a - b
}
```

```bash
$ go test -coverprofile=coverage.out
$ go tool cover -func=coverage.out
mypackage/math.go:3:    Add             100.0%
mypackage/math.go:7:    Subtract        66.7%
total:                  (statements)    83.3%
```

---

## 11.6 Mocking & Interfaces

**Use interfaces to enable testing.**

### Interface-Based Mocking

```go
// Production code
type UserStore interface {
	GetUser(id int) (*User, error)
	SaveUser(user *User) error
}

type UserService struct {
	store UserStore
}

func (s *UserService) UpdateEmail(userID int, email string) error {
	user, err := s.store.GetUser(userID)
	if err != nil {
		return err
	}
	
	user.Email = email
	return s.store.SaveUser(user)
}

// Test code
type MockUserStore struct {
	users map[int]*User
}

func (m *MockUserStore) GetUser(id int) (*User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserStore) SaveUser(user *User) error {
	m.users[user.ID] = user
	return nil
}

func TestUpdateEmail(t *testing.T) {
	mock := &MockUserStore{
		users: map[int]*User{
			1: {ID: 1, Email: "old@example.com"},
		},
	}
	
	service := &UserService{store: mock}
	
	err := service.UpdateEmail(1, "new@example.com")
	if err != nil {
		t.Fatal(err)
	}
	
	user, _ := mock.GetUser(1)
	if user.Email != "new@example.com" {
		t.Errorf("email = %s, want new@example.com", user.Email)
	}
}
```

### Using testify/mock

```go
import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserStore struct {
	mock.Mock
}

func (m *MockUserStore) GetUser(id int) (*User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserStore) SaveUser(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}

func TestUpdateEmailWithMock(t *testing.T) {
	mockStore := new(MockUserStore)
	user := &User{ID: 1, Email: "old@example.com"}
	
	mockStore.On("GetUser", 1).Return(user, nil)
	mockStore.On("SaveUser", user).Return(nil)
	
	service := &UserService{store: mockStore}
	err := service.UpdateEmail(1, "new@example.com")
	
	if err != nil {
		t.Fatal(err)
	}
	
	mockStore.AssertExpectations(t)
}
```

---

## 11.7 Testing Packages

### testify/assert

```go
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithAssert(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result)
	assert.NotEqual(t, 10, result)
	
	err := doSomething()
	assert.NoError(t, err)
	
	user := getUser()
	assert.NotNil(t, user)
	assert.Contains(t, user.Name, "Alice")
}
```

### testify/require

**Stops test on failure:**

```go
import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWithRequire(t *testing.T) {
	result, err := Divide(10, 2)
	require.NoError(t, err)  // Stops here if error
	require.Equal(t, 5.0, result)
}
```

### httptest Package

```go
import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World"))
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	
	handler(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	
	body := w.Body.String()
	if body != "Hello, World" {
		t.Errorf("body = %s, want Hello, World", body)
	}
}
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. Test files must end with __________.
2. Test functions must start with __________ and take *testing.T.
3. The __________ method marks a function as a test helper.
4. To run benchmarks, use the __________ flag.
5. The -cover flag shows test __________.
6. t.__________ runs subtests in parallel.

### True/False

1. ⬜ Test files must be in a separate package
2. ⬜ t.Error() stops the test immediately
3. ⬜ Table-driven tests are a Go best practice
4. ⬜ Benchmarks run the test code exactly once
5. ⬜ Test coverage measures code quality
6. ⬜ Interfaces enable easier testing through mocking

### Multiple Choice

1. Which runs all tests?
   - A) `go test`
   - B) `go test .`
   - C) `go test all`
   - D) `go run test`

2. What does t.Helper() do?
   - A) Runs the test
   - B) Marks function as helper (better error reporting)
   - C) Stops the test
   - D) Skips the test

3. How do you run only tests matching a pattern?
   - A) `go test -match Pattern`
   - B) `go test -run Pattern`
   - C) `go test -filter Pattern`
   - D) `go test -only Pattern`

4. What does -benchmem show?
   - A) Test coverage
   - B) Memory allocations
   - C) Execution time
   - D) CPU usage

### Code Challenge

Write tests for this function using table-driven tests:
```go
func IsEven(n int) bool {
	return n%2 == 0
}
```

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. _test.go
2. Test
3. Helper (t.Helper())
4. -bench
5. coverage
6. Parallel

**True/False:**
1. ❌ False (same package, just _test.go suffix)
2. ❌ False (t.Fatal() stops, t.Error() continues)
3. ✅ True
4. ❌ False (runs b.N times, determined by runtime)
5. ❌ False (coverage measures tested code, not quality)
6. ✅ True

**Multiple Choice:**
1. **A** - `go test` (runs all tests in current package)
2. **B** - Marks function as helper (better error reporting)
3. **B** - `go test -run Pattern`
4. **B** - Memory allocations

**Code Challenge:**
```go
func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"zero", 0, true},
		{"positive even", 4, true},
		{"positive odd", 5, false},
		{"negative even", -4, true},
		{"negative odd", -5, false},
		{"large even", 1000, true},
		{"large odd", 1001, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; want %v", 
					tt.input, result, tt.expected)
			}
		})
	}
}
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ Testing basics and test file structure  
✅ Table-driven tests (Go best practice)  
✅ Test helpers and subtests  
✅ Benchmarking performance  
✅ Test coverage measurement  
✅ Mocking with interfaces  
✅ Testing packages (testify, httptest)  

**Next Steps:**
- Continue to [Chapter 12: Web Development →](12-web-development.md)
- Practice writing tests for existing code
- Complete the IsEven test challenge above

---

**Navigation:** [← Generics](10-generics.md) | [README](./README.md) | [Next: Web Development →](12-web-development.md)