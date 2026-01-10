# 9. Concurrency

> Master goroutines, channels, Context, WaitGroup, Mutex, and concurrent programming in Go.

**Navigation:** [← Error Handling](8-error-handling.md) | [README](./README.md) | [Next: Generics →](10-generics.md)

---

## Table of Contents

- [9.1 Goroutines](#91-goroutines)
- [9.2 Channels](#92-channels)
- [9.3 Select Statement](#93-select-statement)
- [9.4 Channel Patterns](#94-channel-patterns)
- [9.5 Context](#95-context)
- [9.6 WaitGroup](#96-waitgroup)
- [9.7 Mutex](#97-mutex)
- [9.8 Advanced Patterns](#98-advanced-patterns)
- [9.9 Race Conditions](#99-race-conditions)
- [Practice Questions - goroutines & channels](#practice-questions---goroutines--channels)
- [Practice Questions - Context, WaitGroup & Patterns](#practice-questions---context-waitgroup--patterns)

---

## 9.1 Goroutines

A **goroutine** is a lightweight thread managed by the Go runtime.

**Key Features:**
- Extremely lightweight (~2KB initial stack)
- Thousands/millions can run concurrently
- Scheduled by Go runtime, not OS

### Creating Goroutines

```go
func sayHello() {
	fmt.Println("Hello from goroutine")
}

func main() {
	go sayHello()  // Launches goroutine
	
	time.Sleep(time.Second)  // Wait for goroutine
	fmt.Println("Main function")
}
```

### Anonymous Goroutines

```go
func main() {
	go func() {
		fmt.Println("Anonymous goroutine")
	}()
	
	time.Sleep(time.Second)
}
```

### Goroutines with Parameters

```go
func printNumber(n int) {
	fmt.Println(n)
}

func main() {
	for i := 0; i < 5; i++ {
		go printNumber(i)
	}
	
	time.Sleep(time.Second)
}
```

### Common Pitfall: Loop Variables

**❌ Wrong:**
```go
for i := 0; i < 5; i++ {
	go func() {
		fmt.Println(i)  // All might print 5!
	}()
}
```

**✅ Correct:**
```go
// Method 1: Pass as parameter
for i := 0; i < 5; i++ {
	go func(n int) {
		fmt.Println(n)
	}(i)
}

// Method 2: Capture in local variable
for i := 0; i < 5; i++ {
	i := i  // Create new variable
	go func() {
		fmt.Println(i)
	}()
}
```

### Goroutine Execution

```go
func main() {
	go func() {
		fmt.Println("Goroutine 1")
	}()
	
	go func() {
		fmt.Println("Goroutine 2")
	}()
	
	// Without waiting, main exits and goroutines are killed
	time.Sleep(time.Second)
}
```

**Important:** Main function exiting kills all goroutines!

---

## 9.2 Channels

**Channels** allow goroutines to communicate safely.

**Think of channels as:**
- Pipes for sending/receiving data
- Thread-safe queues
- Synchronization primitives

### Channel Declaration

```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel (capacity 3)
ch := make(chan int, 3)

// Channel types
var ch chan int        // Bidirectional
var sendCh chan<- int  // Send-only
var recvCh <-chan int  // Receive-only
```

### Sending and Receiving

```go
ch := make(chan int)

// Send value
ch <- 42

// Receive value
value := <-ch

// Receive and discard
<-ch
```

### Unbuffered Channels (Synchronous)

**Unbuffered channels block until both sender and receiver are ready:**

```go
func main() {
	ch := make(chan string)
	
	go func() {
		ch <- "Hello"  // Blocks until received
	}()
	
	msg := <-ch  // Blocks until sent
	fmt.Println(msg)
}
```

**Visual:**
```
Sender                  Receiver
  |                        |
  | ch <- "Hello"          |
  |-------[WAIT]---------->|
  |                        | msg := <-ch
  |<------[SYNC]-----------|
  |                        |
```

### Buffered Channels (Asynchronous)

**Buffered channels only block when full (send) or empty (receive):**

```go
ch := make(chan int, 2)  // Buffer size 2

ch <- 1  // Doesn't block
ch <- 2  // Doesn't block
// ch <- 3  // Would block (buffer full)

fmt.Println(<-ch)  // 1
fmt.Println(<-ch)  // 2
// fmt.Println(<-ch)  // Would block (buffer empty)
```

### Closing Channels

```go
ch := make(chan int)

go func() {
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)  // Signal no more values
}()

// Receive until closed
for val := range ch {
	fmt.Println(val)
}
// Output: 1 2 3
```

**Key Points:**
- Only **senders** should close channels
- Receiving from closed channel returns zero value + false
- Sending to closed channel **panics**

### Checking if Channel is Closed

```go
ch := make(chan int)
close(ch)

// Comma-ok idiom
val, ok := <-ch
if ok {
	fmt.Println("Received:", val)
} else {
	fmt.Println("Channel closed")
}
```

### Real-World Example: Worker Pool

```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	
	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Collect results
	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
}
```

---

## 9.3 Select Statement

**Select** lets you wait on multiple channel operations.

### Basic Select

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() {
	time.Sleep(1 * time.Second)
	ch1 <- "one"
}()

go func() {
	time.Sleep(2 * time.Second)
	ch2 <- "two"
}()

// Wait for first available
select {
case msg1 := <-ch1:
	fmt.Println("Received:", msg1)
case msg2 := <-ch2:
	fmt.Println("Received:", msg2)
}
```

### Select with Default

**Non-blocking operations:**

```go
ch := make(chan int)

select {
case val := <-ch:
	fmt.Println("Received:", val)
default:
	fmt.Println("No value available")
}
```

### Timeout Pattern

```go
ch := make(chan string)

go func() {
	time.Sleep(2 * time.Second)
	ch <- "result"
}()

select {
case res := <-ch:
	fmt.Println(res)
case <-time.After(1 * time.Second):
	fmt.Println("Timeout!")
}
```

### Multiple Channels

```go
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		for {
			ch1 <- "Every 500ms"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	
	go func() {
		for {
			ch2 <- "Every 2s"
			time.Sleep(2 * time.Second)
		}
	}()
	
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Channel 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Channel 2:", msg2)
		}
	}
}
```

---

## 9.4 Channel Patterns

### Pipeline Pattern

**Chain processing stages:**

```go
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Pipeline: generator -> square
	nums := generator(2, 3, 4, 5)
	squares := square(nums)
	
	for sq := range squares {
		fmt.Println(sq)  // 4, 9, 16, 25
	}
}
```

### Fan-Out, Fan-In Pattern

**Multiple workers, single result collector:**

```go
func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func worker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	
	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}
	
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

func main() {
	in := producer(1, 2, 3, 4, 5, 6, 7, 8)
	
	// Fan-out to 3 workers
	w1 := worker(in)
	w2 := worker(in)
	w3 := worker(in)
	
	// Fan-in results
	for result := range merge(w1, w2, w3) {
		fmt.Println(result)
	}
}
```

### Done Channel Pattern

**Signal cancellation:**

```go
func worker(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Worker stopping")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	done := make(chan bool)
	
	go worker(done)
	
	time.Sleep(2 * time.Second)
	close(done)  // Signal stop
	
	time.Sleep(time.Second)
}
```

### Quit Channel Pattern

```go
func generator(quit <-chan bool) <-chan int {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				i++
			case <-quit:
				close(ch)
				return
			}
		}
	}()
	return ch
}

func main() {
	quit := make(chan bool)
	ch := generator(quit)
	
	// Get first 5 numbers
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	
	close(quit)
}
```

### Rate Limiting Pattern

```go
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	// Process 1 request per 500ms
	limiter := time.Tick(500 * time.Millisecond)
	
	for req := range requests {
		<-limiter  // Wait for tick
		fmt.Printf("Request %d at %v\n", req, time.Now())
	}
}
```

---

## 9.5 Context

**Context** carries deadlines, cancellation signals, and request-scoped values across API boundaries.

### Why Context?

**Solves:**
- Cancellation propagation
- Timeout management
- Request-scoped values (user ID, trace ID)

### Creating Contexts

```go
import "context"

// Background context (never canceled)
ctx := context.Background()

// TODO context (placeholder)
ctx := context.TODO()

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// With deadline
deadline := time.Now().Add(5 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// With value
ctx := context.WithValue(context.Background(), "userID", 123)
```

### Context Cancellation

```go
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: stopping\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	
	// Start workers
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}
	
	time.Sleep(2 * time.Second)
	cancel()  // Cancel all workers
	
	time.Sleep(time.Second)
}
```

### Context with Timeout

```go
func slowOperation(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()  // context.DeadlineExceeded
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	err := slowOperation(ctx)
	if err != nil {
		fmt.Println("Error:", err)  // context deadline exceeded
	}
}
```

### Context with Values

```go
type contextKey string

const userIDKey contextKey = "userID"

func processRequest(ctx context.Context) {
	userID := ctx.Value(userIDKey)
	if userID != nil {
		fmt.Printf("Processing request for user: %v\n", userID)
	}
}

func main() {
	ctx := context.WithValue(context.Background(), userIDKey, 123)
	processRequest(ctx)
}
```

**⚠️ Warning:** Only use context values for request-scoped data, not for passing optional parameters.

### Real-World Example: HTTP Request with Context

```go
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	// Add timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
	result, err := fetchData(ctx)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			http.Error(w, "Request timeout", http.StatusRequestTimeout)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Write([]byte(result))
}

func fetchData(ctx context.Context) (string, error) {
	select {
	case <-time.After(3 * time.Second):
		return "data", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
```

### Context Best Practices

**✅ Do:**
- Pass context as first parameter
- Always call cancel() (use defer)
- Pass context down the call stack
- Use for cancellation and timeouts

**❌ Don't:**
- Store context in structs
- Pass nil context (use context.TODO())
- Use for optional parameters
- Use for values that can be passed as function parameters

---

## 9.6 WaitGroup

**WaitGroup** waits for a collection of goroutines to finish.

### Basic Usage

```go
import "sync"

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()  // Decrement counter when done
	
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)  // Increment counter
		go worker(i, &wg)
	}
	
	wg.Wait()  // Block until counter is 0
	fmt.Println("All workers done")
}
```

### WaitGroup Methods

```go
var wg sync.WaitGroup

wg.Add(1)    // Increment counter
wg.Done()    // Decrement counter (same as Add(-1))
wg.Wait()    // Block until counter is 0
```

### Real-World Example: Parallel Processing

```go
func processURLs(urls []string) []string {
	var wg sync.WaitGroup
	results := make([]string, len(urls))
	
	for i, url := range urls {
		wg.Add(1)
		go func(index int, u string) {
			defer wg.Done()
			results[index] = fetch(u)
		}(i, url)
	}
	
	wg.Wait()
	return results
}

func fetch(url string) string {
	// Simulate HTTP request
	time.Sleep(500 * time.Millisecond)
	return "Response from " + url
}
```

### Common Mistake: Add Inside Goroutine

**❌ Wrong:**
```go
for i := 0; i < 3; i++ {
	go func() {
		wg.Add(1)  // Race condition!
		defer wg.Done()
		// work
	}()
}
wg.Wait()  // Might return before Add is called
```

**✅ Correct:**
```go
for i := 0; i < 3; i++ {
	wg.Add(1)  // Add before spawning goroutine
	go func() {
		defer wg.Done()
		// work
	}()
}
wg.Wait()
```

---

## 9.7 Mutex

**Mutex** (mutual exclusion) protects shared resources from concurrent access.

### sync.Mutex

```go
import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
```

### Read-Write Mutex (sync.RWMutex)

**Allows multiple readers OR single writer:**

```go
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()  // Read lock
	defer c.mu.RUnlock()
	
	val, ok := c.data[key]
	return val, ok
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()  // Write lock
	defer c.mu.Unlock()
	
	c.data[key] = value
}
```

### Real-World Example: Thread-Safe Map

```go
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

func (sm *SafeMap) Delete(key string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}
```

### Mutex Best Practices

**✅ Do:**
- Always unlock (use defer)
- Keep critical sections small
- Use RWMutex for read-heavy workloads
- Prefer channels over mutexes when possible

**❌ Don't:**
- Lock inside defer
- Hold locks across I/O operations
- Forget to unlock (causes deadlock)

---

## 9.8 Advanced Patterns

### Worker Pool with WaitGroup

```go
func workerPool(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	var wg sync.WaitGroup
	
	// Start workers
	for w := 0; w < 3; w++ {
		wg.Add(1)
		go workerPool(jobs, results, &wg)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Wait and close results
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Println(result)
	}
}
```

### Semaphore Pattern

**Limit concurrent operations:**

```go
type Semaphore chan struct{}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

func NewSemaphore(max int) Semaphore {
	return make(Semaphore, max)
}

// Usage
func main() {
	sem := NewSemaphore(3)  // Max 3 concurrent
	
	for i := 1; i <= 10; i++ {
		sem.Acquire()
		go func(id int) {
			defer sem.Release()
			fmt.Printf("Worker %d running\n", id)
			time.Sleep(time.Second)
		}(i)
	}
	
	time.Sleep(5 * time.Second)
}
```

### Once Pattern (sync.Once)

**Run exactly once:**

```go
import "sync"

var (
	instance *Database
	once     sync.Once
)

func GetDatabase() *Database {
	once.Do(func() {
		instance = &Database{}
		instance.Connect()
	})
	return instance
}
```

### Broadcast Pattern

**Notify multiple goroutines:**

```go
func main() {
	broadcast := make(chan struct{})
	
	// Start listeners
	for i := 1; i <= 3; i++ {
		go func(id int) {
			<-broadcast
			fmt.Printf("Listener %d received broadcast\n", id)
		}(i)
	}
	
	time.Sleep(time.Second)
	close(broadcast)  // Broadcast to all
	
	time.Sleep(time.Second)
}
```

---

## 9.9 Race Conditions

A **race condition** occurs when multiple goroutines access shared data concurrently and at least one modifies it.

### Race Condition Example

**❌ Race condition:**
```go
var counter int

func increment() {
	counter++  // Not atomic!
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	
	time.Sleep(time.Second)
	fmt.Println(counter)  // Likely < 1000 due to race
}
```

### Detecting Race Conditions

**Use the race detector:**
```bash
go run -race main.go
go test -race
go build -race
```

### Fixing Race Conditions

**Option 1: Mutex**
```go
var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}
```

**Option 2: Channels**
```go
func main() {
	ch := make(chan int)
	counter := 0
	
	// Goroutine owns counter
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

**Option 3: atomic Package**
```go
import "sync/atomic"

var counter int64

func increment() {
	atomic.AddInt64(&counter, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	
	time.Sleep(time.Second)
	fmt.Println(atomic.LoadInt64(&counter))  // 1000
}
```

---

## Practice Questions - goroutines & channels

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. A goroutine is started using the __________ keyword.
2. Unbuffered channels are __________ (synchronous/asynchronous).
3. Only the __________ should close a channel.
4. The __________ statement waits on multiple channel operations.
5. Receiving from a closed channel returns the __________ value.
6. Sending to a closed channel causes a __________.

### True/False

1. ⬜ Goroutines are OS threads
2. ⬜ Buffered channels can hold multiple values without blocking
3. ⬜ The main function waits for all goroutines to finish
4. ⬜ Select chooses a random case if multiple are ready
5. ⬜ Closing a channel is mandatory
6. ⬜ Range over a channel stops when the channel is closed

### Multiple Choice

1. What happens if main() exits?
   - A) Goroutines continue running
   - B) Program waits for goroutines
   - C) All goroutines are killed
   - D) Compile error

2. What is the capacity of this channel?
   ```go
   ch := make(chan int, 5)
   ```
   - A) 0
   - B) 1
   - C) 5
   - D) Unlimited

3. What does select with default do?
   - A) Blocks forever
   - B) Non-blocking operation
   - C) Panics
   - D) Returns error

4. Which is true about unbuffered channels?
   - A) Send blocks until receive
   - B) Can store multiple values
   - C) Never blocks
   - D) Faster than buffered

### Code Challenge

Create a function `processNumbers(nums []int) []int` that:
- Processes numbers concurrently using 3 workers
- Each worker squares the number
- Returns results in any order
- Uses channels for communication

---

### Answers - goroutines & channels

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. go
2. synchronous
3. sender
4. select
5. zero
6. panic

**True/False:**
1. ❌ False (lightweight threads managed by Go runtime)
2. ✅ True
3. ❌ False (main exits, goroutines are killed)
4. ✅ True (pseudo-random selection)
5. ❌ False (optional, but good practice)
6. ✅ True

**Multiple Choice:**
1. **C** - All goroutines are killed
2. **C** - 5 (buffer capacity)
3. **B** - Non-blocking operation
4. **A** - Send blocks until receive (synchronous)

**Code Challenge:**
```go
func processNumbers(nums []int) []int {
	jobs := make(chan int, len(nums))
	results := make(chan int, len(nums))
	
	// Start 3 workers
	for w := 0; w < 3; w++ {
		go func() {
			for num := range jobs {
				results <- num * num
			}
		}()
	}
	
	// Send jobs
	for _, num := range nums {
		jobs <- num
	}
	close(jobs)
	
	// Collect results
	output := make([]int, len(nums))
	for i := range output {
		output[i] = <-results
	}
	
	return output
}

// Usage
nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
results := processNumbers(nums)
fmt.Println(results)  // [1 4 9 16 25 36 49 64 81] (order may vary)
```

</details>

</details>

---

## Practice Questions - Context, WaitGroup & Patterns

<details>
<summary><strong>View Questions</strong></summary>

### Fill in the Blanks

1. Context is used to carry __________, cancellation signals, and request-scoped values.
2. __________ waits for a collection of goroutines to finish.
3. A __________ protects shared resources from concurrent access.
4. sync.RWMutex allows multiple __________ or a single writer.
5. The race detector is run with the __________ flag.
6. __________ runs a function exactly once.

### True/False

1. ⬜ Context should be stored in structs
2. ⬜ WaitGroup.Add() should be called before spawning goroutines
3. ⬜ Mutex.Lock() blocks if already locked
4. ⬜ RWMutex allows multiple concurrent writers
5. ⬜ Race conditions are compile-time errors
6. ⬜ sync.Once guarantees execution exactly once

### Multiple Choice

1. When should you call cancel() for a context?
   - A) Never
   - B) In a defer statement
   - C) After the operation completes
   - D) Both B and C

2. What does WaitGroup.Wait() do?
   - A) Waits for 1 second
   - B) Blocks until counter is 0
   - C) Cancels goroutines
   - D) Returns immediately

3. Which is best for read-heavy workloads?
   - A) sync.Mutex
   - B) sync.RWMutex
   - C) Channels
   - D) atomic

4. How do you detect race conditions?
   - A) go run -race
   - B) go build -race
   - C) go test -race
   - D) All of the above

### Code Challenge

Create a thread-safe counter that supports:
- Increment()
- Decrement()
- Value() int
- Multiple goroutines safely accessing it

---

### Answers - Context, WaitGroup & Patterns

<details>
<summary><strong>View Answers</strong></summary>

**Fill in the Blanks:**
1. deadlines
2. WaitGroup (sync.WaitGroup)
3. mutex
4. readers
5. -race
6. sync.Once

**True/False:**
1. ❌ False (contexts should be passed as parameters)
2. ✅ True
3. ✅ True
4. ❌ False (single writer only)
5. ❌ False (runtime errors, detected with -race flag)
6. ✅ True

**Multiple Choice:**
1. **D** - Both B and C (preferably in defer)
2. **B** - Blocks until counter is 0
3. **B** - sync.RWMutex (multiple readers)
4. **D** - All of the above

**Code Challenge:**
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

func (c *Counter) Decrement() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value--
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Usage
func main() {
	counter := &Counter{}
	var wg sync.WaitGroup
	
	// 100 increments
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	
	// 50 decrements
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}
	
	wg.Wait()
	fmt.Println("Final value:", counter.Value())  // 50
}
```

</details>

</details>

---

## Summary

In this section, you learned:

✅ Goroutines (lightweight threads)  
✅ Channels (communication between goroutines)  
✅ Buffered vs unbuffered channels  
✅ Channel closing and range  
✅ Select statement (multiplexing channels)  
✅ Common channel patterns (pipeline, fan-out/fan-in, done)  
✅ Context (cancellation, timeouts, values)  
✅ WaitGroup (waiting for goroutines)  
✅ Mutex and RWMutex (protecting shared data)  
✅ Advanced patterns (worker pool, semaphore, once, broadcast)  
✅ Race conditions and detection  
✅ Concurrency best practices  

**Next Steps:**
- Continue to [Chapter 10: Generics →](10-generics.md)
- Practice concurrent programming
- Complete the thread-safe counter challenge

---

**Navigation:** [← Error Handling](8-error-handling.md) | [README](./README.md) | [Next: Generics →](10-generics.md)