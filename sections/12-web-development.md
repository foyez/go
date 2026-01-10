# 12. Web Development

> Master web development with Go's standard library and popular frameworks.

**Navigation:** [← Testing](11-testing.md) | [README](./README.md) | [Next: Best Practices →](13-best-practices.md)

---

## Table of Contents

- [12.1 HTTP Basics](#121-http-basics)
- [12.2 Routing & Handlers](#122-routing--handlers)
- [12.3 Middleware](#123-middleware)
- [12.4 JSON & REST APIs](#124-json--rest-apis)
- [12.5 Templates](#125-templates)
- [12.6 Gin Framework](#126-gin-framework)
- [Practice Questions](#practice-questions)

---

## 12.1 HTTP Basics

Go's `net/http` package provides powerful HTTP client and server capabilities.

### Simple HTTP Server

```go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

### Handler Function Signature

```go
func(w http.ResponseWriter, r *http.Request)
```

**Parameters:**
- `w` - Write response (headers, status, body)
- `r` - Read request (method, URL, headers, body)

### Request Methods

```go
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET request")
	case http.MethodPost:
		fmt.Fprintf(w, "POST request")
	case http.MethodPut:
		fmt.Fprintf(w, "PUT request")
	case http.MethodDelete:
		fmt.Fprintf(w, "DELETE request")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
```

### Response Writing

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// Set status code
	w.WriteHeader(http.StatusOK)
	
	// Set headers
	w.Header().Set("Content-Type", "text/plain")
	
	// Write body
	w.Write([]byte("Hello"))
	
	// Or use fmt.Fprintf
	fmt.Fprintf(w, "World")
}
```

### Reading Request Data

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// URL path
	path := r.URL.Path
	
	// Query parameters
	name := r.URL.Query().Get("name")
	
	// Headers
	userAgent := r.Header.Get("User-Agent")
	
	// Request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	
	fmt.Fprintf(w, "Path: %s, Name: %s", path, name)
}
```

---

## 12.2 Routing & Handlers

### Multiple Routes

```go
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact Page")
}
```

### ServeMux (Custom Router)

```go
func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/create", createUserHandler)
	
	http.ListenAndServe(":8080", mux)
}
```

### Handler Interface

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// Custom handler type
type HelloHandler struct {
	message string
}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, h.message)
}

func main() {
	mux := http.NewServeMux()
	
	mux.Handle("/hello", HelloHandler{message: "Hello, World!"})
	
	http.ListenAndServe(":8080", mux)
}
```

### Static File Server

```go
func main() {
	// Serve files from ./static directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", homeHandler)
	
	http.ListenAndServe(":8080", nil)
}
```

---

## 12.3 Middleware

**Middleware** wraps handlers to add functionality.

### Basic Middleware

```go
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	
	// Wrap with middleware
	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
```

### Chaining Middleware

```go
func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware 1 - Before")
		next.ServeHTTP(w, r)
		log.Println("Middleware 1 - After")
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware 2 - Before")
		next.ServeHTTP(w, r)
		log.Println("Middleware 2 - After")
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	
	// Chain middleware
	handler := middleware1(middleware2(mux))
	http.ListenAndServe(":8080", handler)
}
```

### Authentication Middleware

```go
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		
		if token != "secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Protected Resource")
}

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", homeHandler)
	mux.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))
	
	http.ListenAndServe(":8080", mux)
}
```

### CORS Middleware

```go
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
```

---

## 12.4 JSON & REST APIs

### Encoding JSON

```go
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
```

### Decoding JSON

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	
	// Save user to database...
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
```

### REST API Example

```go
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// GET /users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func getUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	
	http.Error(w, "User not found", http.StatusNotFound)
}

// POST /users
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	user.ID = len(users) + 1
	users = append(users, user)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id}
func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	
	http.Error(w, "User not found", http.StatusNotFound)
}

// DELETE /users/{id}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	
	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, r)
		case http.MethodPost:
			createUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUser(w, r)
		case http.MethodPut:
			updateUser(w, r)
		case http.MethodDelete:
			deleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	
	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
```

---

## 12.5 Templates

Go's `html/template` package provides safe HTML templating.

### Basic Template

```go
import "html/template"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("home").Parse(`
		<!DOCTYPE html>
		<html>
		<head><title>{{.Title}}</title></head>
		<body>
			<h1>{{.Heading}}</h1>
			<p>{{.Message}}</p>
		</body>
		</html>
	`))
	
	data := struct {
		Title   string
		Heading string
		Message string
	}{
		Title:   "Home Page",
		Heading: "Welcome",
		Message: "Hello, World!",
	}
	
	tmpl.Execute(w, data)
}
```

### Template from File

**templates/home.html:**
```html
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
	<h1>{{.Heading}}</h1>
	<p>{{.Message}}</p>
</body>
</html>
```

**Go code:**
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	
	data := struct {
		Title   string
		Heading string
		Message string
	}{
		Title:   "Home Page",
		Heading: "Welcome",
		Message: "Hello from Go!",
	}
	
	tmpl.Execute(w, data)
}
```

### Template Actions

```html
<!-- Variables -->
{{.Name}}

<!-- Conditionals -->
{{if .IsActive}}
	<p>Active</p>
{{else}}
	<p>Inactive</p>
{{end}}

<!-- Range -->
{{range .Items}}
	<li>{{.}}</li>
{{end}}

<!-- With -->
{{with .User}}
	<p>Name: {{.Name}}</p>
{{end}}
```

---

## 12.6 Gin Framework

**Gin** is a high-performance web framework.

### Installation

```bash
go get -u github.com/gin-gonic/gin
```

### Basic Gin Server

```go
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	
	r.Run(":8080")
}
```

### Gin Routes

```go
func main() {
	r := gin.Default()
	
	// GET
	r.GET("/users", getUsers)
	
	// POST
	r.POST("/users", createUser)
	
	// PUT
	r.PUT("/users/:id", updateUser)
	
	// DELETE
	r.DELETE("/users/:id", deleteUser)
	
	r.Run(":8080")
}
```

### Gin Path Parameters

```go
r.GET("/users/:id", func(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"user_id": id,
	})
})
```

### Gin Query Parameters

```go
r.GET("/search", func(c *gin.Context) {
	name := c.Query("name")
	age := c.DefaultQuery("age", "0")
	
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
})
```

### Gin JSON Binding

```go
type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=0,lte=130"`
}

r.POST("/users", func(c *gin.Context) {
	var user User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// Save user...
	
	c.JSON(201, user)
})
```

### Gin Middleware

```go
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		
		// Before request
		c.Next()
		
		// After request
		latency := time.Since(t)
		log.Printf("%s %s - %v", c.Request.Method, c.Request.URL.Path, latency)
	}
}

func main() {
	r := gin.New()
	
	// Global middleware
	r.Use(Logger())
	r.Use(gin.Recovery())
	
	// Route-specific middleware
	r.GET("/admin", AuthMiddleware(), adminHandler)
	
	r.Run(":8080")
}
```

### Gin Groups

```go
func main() {
	r := gin.Default()
	
	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", getUsers)
		v1.POST("/users", createUser)
	}
	
	// API v2
	v2 := r.Group("/api/v2")
	{
		v2.GET("/users", getUsersV2)
		v2.POST("/users", createUserV2)
	}
	
	r.Run(":8080")
}
```

---

## Practice Questions

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. HTTP handler functions take __________ and *http.Request parameters.
2. The __________ package provides HTTP server functionality.
3. Middleware wraps handlers to add __________.
4. JSON struct tags use the __________ format.
5. Templates from files are parsed with __________.
6. Gin uses c.__________ to bind JSON data.

### True/False

1. ⬜ http.HandleFunc requires a ServeMux
2. ⬜ Middleware can be chained in Go
3. ⬜ JSON encoding requires manual conversion
4. ⬜ Templates automatically escape HTML
5. ⬜ Gin is part of Go's standard library
6. ⬜ Path parameters in Gin use the :param syntax

### Multiple Choice

1. Which sets the response status code?
   - A) w.Status(200)
   - B) w.WriteHeader(200)
   - C) w.SetStatus(200)
   - D) w.Code(200)

2. How do you read query parameters?
   - A) r.Query.Get("name")
   - B) r.URL.Query().Get("name")
   - C) r.GetQuery("name")
   - D) r.Param("name")

3. Which Gin method handles GET requests?
   - A) r.Get()
   - B) r.GET()
   - C) r.HandleGet()
   - D) r.GetHandler()

4. What does json.NewEncoder(w).Encode(data) do?
   - A) Reads JSON
   - B) Writes JSON to response
   - C) Validates JSON
   - D) Parses JSON

### Code Challenge

Create a simple REST API endpoint that:
- Accepts POST requests to /tasks
- Reads JSON: {"title": "Task name", "done": false}
- Returns the task with a generated ID
- Sets proper content type and status code

---

### Answers

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. http.ResponseWriter
2. net/http
3. functionality (or behavior)
4. `json:"field_name"`
5. template.ParseFiles
6. ShouldBindJSON

**True/False:**
1. ❌ False (http.HandleFunc uses DefaultServeMux)
2. ✅ True
3. ❌ False (json.NewEncoder does it automatically)
4. ✅ True (html/template escapes by default)
5. ❌ False (third-party framework)
6. ✅ True

**Multiple Choice:**
1. **B** - w.WriteHeader(200)
2. **B** - r.URL.Query().Get("name")
3. **B** - r.GET()
4. **B** - Writes JSON to response

**Code Challenge:**
```go
package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	tasks  []Task
	nextID = 1
	mu     sync.Mutex
)

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	
	mu.Lock()
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	mu.Unlock()
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func main() {
	http.HandleFunc("/tasks", createTaskHandler)
	http.ListenAndServe(":8080", nil)
}
```

</details>

</details>

---

## Summary

In this chapter, you learned:

✅ HTTP basics with net/http package  
✅ Routing and handler patterns  
✅ Middleware for cross-cutting concerns  
✅ JSON encoding/decoding for REST APIs  
✅ HTML templates for server-side rendering  
✅ Gin framework for rapid development  

**Next Steps:**
- Continue to [Chapter 13: Best Practices →](13-best-practices.md)
- Build a complete REST API
- Complete the task endpoint challenge above

---

**Navigation:** [← Testing](11-testing.md) | [README](./README.md) | [Next: Best Practices →](13-best-practices.md)