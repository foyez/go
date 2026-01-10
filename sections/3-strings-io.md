# 3. Strings & I/O

> Master string manipulation and input/output operations in Go.

**Navigation:** [← Fundamentals](2-fundamentals.md) | [README](./README.md) | [Next: Control Flow →](4-control-flow.md)

---

## Table of Contents

- [3.1 String Basics](#31-string-basics)
- [3.2 String Operations](#32-string-operations)
- [3.3 Input/Output](#33-inputoutput)
- [3.4 File Operations](#34-file-operations)
- [Practice Questions](#practice-questions)

---

## 3.1 String Basics

In Go, strings are:
- **Immutable sequences of bytes**
- **UTF-8 encoded**
- Internally a **read-only byte slice**

### String Declaration

```go
s := "Hello, 世界"

// Multi-line strings (raw string literals)
s3 := `This is a
multi-line
string`
```

### String Length

**`len(string)` counts bytes, not characters:**

```go
len("hello") // 5 bytes
len("ক")     // 3 bytes (Bengali character)
len("世界")   // 6 bytes (2 Chinese characters, 3 bytes each)
```

**Why?**
- Strings are UTF-8 encoded
- Non-ASCII characters use multiple bytes

### String Indexing

**Indexing returns bytes:**

```go
s := "Hello"
fmt.Println(s[0])           // 72 (byte value of 'H')
fmt.Println(string(s[0]))   // "H" (convert to string)
```

### String Immutability

❌ **Cannot modify strings:**

```go
s := "hello"
s[0] = 'H' // ❌ compile-time error
```

✅ **Create new string instead:**

```go
s = "Hello"  // Assigns new string
```

### String Concatenation

```go
s1 := "Hello"
s2 := " World"
s3 := s1 + s2  // "Hello World"

// Using fmt.Sprintf
s4 := fmt.Sprintf("%s %s", s1, s2)
```

### Iterating Over Strings

**Wrong way (iterates over bytes):**
```go
s := "Hello, 世界"
for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i])  // Breaks on Unicode
}
// Output: H e l l o ,   ä ¸  ç  ...
```

**Correct way (iterates over runes):**
```go
s := "Hello, 世界"
for i, r := range s {
    fmt.Printf("%d: %c\n", i, r)
}
// Output:
// 0: H
// 1: e
// 2: l
// 3: l
// 4: o
// 5: ,
// 6:  
// 7: 世
// 10: 界
```

**Note:** `range` iterates over **runes**, not bytes - correct for Unicode!

### Rune vs Byte

```go
s := "Hello, 世界"

// Rune count (actual characters)
runeCount := utf8.RuneCountInString(s)
fmt.Println(runeCount)  // 9

// Byte count
byteCount := len(s)
fmt.Println(byteCount)  // 13
```

---

## 3.2 String Operations

### strings Package

```go
import "strings"

// Contains
strings.Contains("hello", "ll")        // true

// Prefix/Suffix
strings.HasPrefix("hello", "he")       // true
strings.HasSuffix("hello", "lo")       // true

// Split/Join
parts := strings.Split("a,b,c", ",")   // ["a", "b", "c"]
joined := strings.Join([]string{"a", "b"}, "-")  // "a-b"

// Case conversion
strings.ToUpper("hello")               // "HELLO"
strings.ToLower("HELLO")               // "hello"
strings.Title("hello world")           // "Hello World" (deprecated)

// Replace
strings.Replace("fooo", "o", "O", 2)   // "fOOo" (replace 2)
strings.Replace("fooo", "o", "O", -1)  // "fOOO" (replace all)
strings.ReplaceAll("fooo", "o", "O")   // "fOOO" (Go 1.12+)

// Trim
strings.TrimSpace("  hello  ")         // "hello"
strings.Trim("--hello--", "-")         // "hello"
strings.TrimPrefix("hello", "he")      // "llo"
strings.TrimSuffix("hello", "lo")      // "hel"

// Count/Index
strings.Count("test", "t")             // 2
strings.Index("test", "t")             // 0 (first occurrence)
strings.LastIndex("test", "t")         // 3 (last occurrence)

// Repeat
strings.Repeat("Go", 3)                // "GoGoGo"

// Fields (split on whitespace)
strings.Fields("  hello   world  ")    // ["hello", "world"]
```

### String Builder (Efficient Concatenation)

**❌ Inefficient (creates many strings):**
```go
result := ""
for i := 0; i < 1000; i++ {
    result += "a"  // Creates new string each time
}
```

**✅ Efficient (single allocation):**
```go
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("a")
}
result := builder.String()
```

**Real-World Example:**
```go
func buildHTML(items []string) string {
    var builder strings.Builder
    builder.WriteString("<ul>\n")
    for _, item := range items {
        builder.WriteString("  <li>")
        builder.WriteString(item)
        builder.WriteString("</li>\n")
    }
    builder.WriteString("</ul>")
    return builder.String()
}
```

### strconv Package (String Conversions)

```go
import "strconv"

// Int to String
s := strconv.Itoa(42)                  // "42"

// String to Int
i, err := strconv.Atoi("42")           // 42, nil
if err != nil {
    // Handle error
}

// Parse with base
i, err := strconv.ParseInt("1010", 2, 64)  // 10 (binary to int)

// Float conversions
s := strconv.FormatFloat(3.14, 'f', 2, 64)  // "3.14"
f, err := strconv.ParseFloat("3.14", 64)     // 3.14, nil

// Bool conversions
b, err := strconv.ParseBool("true")    // true, nil
s := strconv.FormatBool(true)          // "true"
```

### fmt Package (Formatting)

```go
import "fmt"

// Sprintf (format to string)
s := fmt.Sprintf("Name: %s, Age: %d", "Alice", 30)

// Printf (format to stdout)
fmt.Printf("Pi: %.2f\n", 3.14159)  // Pi: 3.14

// Common format verbs
%v    // Default format
%+v   // Struct with field names
%#v   // Go-syntax representation
%T    // Type
%d    // Integer
%f    // Float
%s    // String
%q    // Quoted string
%t    // Boolean
%p    // Pointer
```

---

## 3.3 Input/Output

### Output (fmt Package)

**Print functions:**
```go
fmt.Print("Hello")           // No newline
fmt.Println("Hello")         // With newline
fmt.Printf("Age: %d\n", 30)  // Formatted
```

**Format to string:**
```go
msg := fmt.Sprintf("User: %s, Age: %d", "Alice", 30)
```

**Real-World Example:**
```go
type User struct {
    Name string
    Age  int
}

u := User{Name: "Alice", Age: 30}

fmt.Printf("%v\n", u)   // {Alice 30}
fmt.Printf("%+v\n", u)  // {Name:Alice Age:30}
fmt.Printf("%#v\n", u)  // main.User{Name:"Alice", Age:30}
fmt.Printf("%T\n", u)   // main.User
```

### Input (fmt Package)

**Basic input:**
```go
var name string
fmt.Print("Enter your name: ")
fmt.Scan(&name)  // Reads until whitespace
fmt.Println("Hello,", name)
```

**Multiple values:**
```go
var a, b int
fmt.Print("Enter two numbers: ")
fmt.Scanf("%d %d", &a, &b)
fmt.Println("Sum:", a+b)
```

### Input (bufio Package)

**Reading full lines:**
```go
import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Enter text: ")
    text, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    text = strings.TrimSpace(text)  // Remove newline
    fmt.Println("You entered:", text)
}
```

### Command-Line Arguments

```go
import (
    "fmt"
    "os"
)

func main() {
    // os.Args[0] is program name
    // os.Args[1:] are actual arguments
    
    if len(os.Args) < 2 {
        fmt.Println("Usage: program <arg>")
        os.Exit(1)
    }
    
    arg := os.Args[1]
    fmt.Println("Argument:", arg)
}
```

**Run:**
```bash
go run main.go hello
# Output: Argument: hello
```

### Logging (log Package)

**Basic logging:**
```go
import "log"

log.Print("app started")
log.Println("user logged in")
log.Printf("user %s logged in", name)
```

**Fatal and Panic:**
```go
log.Fatal("error")  // Prints + os.Exit(1)
log.Panic("error")  // Prints + panic()
```

**Custom logger:**
```go
log.SetPrefix("INFO: ")
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
log.Println("Custom log")
// Output: INFO: 2024/01/10 15:04:05 main.go:10: Custom log
```

**Log to file:**
```go
file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

log.SetOutput(file)
log.Println("Logging to file")
```

---

## 3.4 File Operations

### Reading Files

**Read entire file:**
```go
import "os"

// ReadFile (simple, loads entire file)
data, err := os.ReadFile("config.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(data))
```

**Read with bufio (line by line):**
```go
import (
    "bufio"
    "os"
)

file, err := os.Open("data.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

### Writing Files

**Write entire file:**
```go
data := []byte("Hello, World!")
err := os.WriteFile("output.txt", data, 0644)
if err != nil {
    log.Fatal(err)
}
```

**Write with file handle:**
```go
file, err := os.Create("output.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

file.WriteString("Line 1\n")
file.WriteString("Line 2\n")
```

**Append to file:**
```go
file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    log.Fatal(err)
}
defer file.Close()

file.WriteString("New log entry\n")
```

### File Information

```go
info, err := os.Stat("file.txt")
if err != nil {
    if os.IsNotExist(err) {
        fmt.Println("File does not exist")
    }
    return
}

fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size(), "bytes")
fmt.Println("Modified:", info.ModTime())
fmt.Println("Is Directory:", info.IsDir())
fmt.Println("Permissions:", info.Mode())
```

### File Operations

```go
// Check if exists
if _, err := os.Stat("file.txt"); err == nil {
    fmt.Println("File exists")
}

// Rename/Move
err := os.Rename("old.txt", "new.txt")

// Delete
err := os.Remove("file.txt")

// Create directory
err := os.Mkdir("mydir", 0755)

// Create nested directories
err := os.MkdirAll("path/to/mydir", 0755)

// Remove directory
err := os.Remove("mydir")

// Remove directory and contents
err := os.RemoveAll("mydir")
```

### Real-World Example (Save/Load Data)

```go
package main

import (
    "fmt"
    "os"
    "strings"
)

type Names []string

func (n Names) toString() string {
    return strings.Join(n, ",")
}

func (n Names) saveToFile(filename string) error {
    return os.WriteFile(filename, []byte(n.toString()), 0666)
}

func readNamesFromFile(filename string) (Names, error) {
    bs, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    return strings.Split(string(bs), ","), nil
}

func main() {
    names := Names{"Alice", "Bob", "Carol"}
    
    // Save
    if err := names.saveToFile("names.txt"); err != nil {
        fmt.Println("Error saving:", err)
        return
    }
    
    // Load
    loaded, err := readNamesFromFile("names.txt")
    if err != nil {
        fmt.Println("Error loading:", err)
        return
    }
    
    fmt.Println("Loaded names:", loaded)
    
    // Clean up
    os.Remove("names.txt")
}
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. In Go, strings are __________ (can/cannot) be modified after creation.
2. The `len()` function returns the number of __________ in a string, not characters.
3. To iterate over Unicode characters in a string, use __________ loop.
4. The __________ package provides efficient string concatenation.
5. The function __________ reads an entire file into memory.
6. File permissions in Go are specified in __________ format (e.g., 0644).

### True/False

1. ⬜ Strings in Go are mutable
2. ⬜ The `len()` function returns the number of Unicode characters
3. ⬜ `range` iterates over runes (Unicode characters), not bytes
4. ⬜ String concatenation with `+` is efficient for large strings
5. ⬜ `os.ReadFile` reads the entire file into memory
6. ⬜ `defer file.Close()` should be called immediately after opening a file

### Multiple Choice

1. What does `len("世界")` return?
   - A) 2
   - B) 4
   - C) 6
   - D) 8

2. Which is the most efficient way to build a large string?
   - A) Using `+` operator
   - B) Using `fmt.Sprintf`
   - C) Using `strings.Builder`
   - D) Using string concatenation in a loop

3. How do you read a file line by line efficiently?
   - A) `os.ReadFile`
   - B) `bufio.Scanner`
   - C) `fmt.Scan`
   - D) `io.ReadAll`

4. What does `strings.TrimSpace("  hello  ")` return?
   - A) "  hello  "
   - B) "hello  "
   - C) "  hello"
   - D) "hello"

### Code Output

```go
s := "Hello"
for i, r := range s {
    if i == 1 {
        fmt.Println(r)
    }
}
```

What is printed?

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. cannot
2. bytes
3. range
4. strings.Builder
5. os.ReadFile
6. octal

**True/False:**
1. ❌ False (strings are immutable)
2. ❌ False (returns bytes, not characters)
3. ✅ True
4. ❌ False (inefficient, creates many intermediate strings)
5. ✅ True
6. ✅ True (ensures file is closed even if errors occur)

**Multiple Choice:**
1. **C** - 6 (each Chinese character is 3 bytes in UTF-8)
2. **C** - Using `strings.Builder` (most efficient)
3. **B** - `bufio.Scanner` (reads line by line without loading entire file)
4. **D** - "hello" (removes leading and trailing whitespace)

**Code Output:**
```
101
```
Explanation: `r` is the rune value (byte value) of 'e', which is 101 in ASCII.

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ String basics (immutability, UTF-8 encoding, indexing)  
✅ String operations (strings package, Builder, strconv)  
✅ Input/Output (fmt, bufio, log)  
✅ File operations (read, write, delete, stat)  
✅ Efficient string handling techniques  

**Next Steps:**
- Continue to [Chapter 4: Control Flow →](4-control-flow.md)
- Practice file I/O operations
- Build a simple text file processor

---

**Navigation:** [← Fundamentals](2-fundamentals.md) | [README](./README.md) | [Next: Control Flow →](4-control-flow.md)