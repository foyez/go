**1. Concurrent URL Crawler:**

<details>
<summary>View contents</summary>

### What this crawler will do

* Take a list of URLs
* Fetch them concurrently
* Limit concurrency (so we don’t overload)
* Collect results safely
* Stop cleanly

---

## Stage 0: Sequential version (baseline)

```go
package main

import (
	"fmt"
	"net/http"
)

func fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(url, "->", resp.Status)
	return nil
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
	}

	for _, url := range urls {
		if err := fetch(url); err != nil {
			fmt.Println("error:", err)
		}
	}
}
```

### What you learn

* One URL at a time
* Slow, but predictable

---

## Stage 1: Goroutines (naive concurrency)

Now let’s make it concurrent **the wrong way** (on purpose).

```go
for _, url := range urls {
	go fetch(url)
}
```

Run it — **nothing prints or prints randomly**.

### Why?

* `main()` exits before goroutines finish
* Goroutines are cheap, but not magical

---

## Stage 2: WaitGroup (basic coordination)

We need to **wait** for all goroutines.

```go
import "sync"

var wg sync.WaitGroup

for _, url := range urls {
	wg.Add(1)
	go func(u string) {
		defer wg.Done()
		fetch(u)
	}(url)
}

wg.Wait()
```

### Key concepts

* `WaitGroup` = “don’t exit yet”
* Always pass loop variables explicitly (`u string`)

✅ Now all URLs fetch concurrently
❌ But still **unlimited concurrency**

---

## Stage 3: Worker Pool (real-world concurrency control)

This is where Go starts to shine.

### Idea

* Fixed number of workers
* URLs go into a channel
* Workers pull from the channel

### Architecture

```
URLs → jobs channel → workers → results
```

---

### Step 1: Jobs channel

```go
jobs := make(chan string)
```

---

### Step 2: Worker function

```go
func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d fetching %s\n", id, url)
		fetch(url)
	}
}
```

---

### Step 3: Start workers

```go
numWorkers := 3
var wg sync.WaitGroup

for i := 1; i <= numWorkers; i++ {
	wg.Add(1)
	go worker(i, jobs, &wg)
}
```

---

### Step 4: Send jobs

```go
for _, url := range urls {
	jobs <- url
}
close(jobs)

wg.Wait()
```

### What you just learned

* **Channels coordinate work**
* **Workers limit concurrency**
* This is how production crawlers work

---

## Stage 4: Results channel (no shared state)

Instead of printing inside workers, return data safely.

```go
type Result struct {
	URL    string
	Status string
	Err    error
}
```

### Worker sends results

```go
func worker(id int, jobs <-chan string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d fetching %s\n", id, url)

		client := http.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Get(url)
		if err != nil {
			results <- Result{URL: url, Err: err}
			continue
		}
		resp.Body.Close()

		results <- Result{
			URL:    url,
			Status: resp.Status,
		}
	}
}
```

---

### Collect results in main

```go
results := make(chan Result)

go func() {
	wg.Wait()
	close(results)
}()

for r := range results {
	if r.Err != nil {
		fmt.Println("error:", r.URL, r.Err)
	} else {
		fmt.Println(r.URL, "->", r.Status)
	}
}
```

### Why this matters

* No mutexes
* No shared memory
* **Channels = ownership transfer**

---

## Final Project

📁 `main.go`

```go
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result represents the outcome of fetching a URL
type Result struct {
	URL    string
	Status string
	Err    error
}

// worker pulls URLs from jobs channel and sends results
func worker(
	id int,
	jobs <-chan string,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for url := range jobs {
		fmt.Printf("Worker %d fetching %s\n", id, url)

		client := http.Client{
			Timeout: 5 * time.Second,
		}

		resp, err := client.Get(url)
		if err != nil {
			results <- Result{URL: url, Err: err}
			continue
		}

		resp.Body.Close()

		results <- Result{
			URL:    url,
			Status: resp.Status,
		}
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
		"https://httpbin.org/status/404",
		"https://invalid-url",
	}

	numWorkers := 3

	jobs := make(chan string)
	results := make(chan Result)

	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Close results when workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		if result.Err != nil {
			fmt.Printf("❌ %s error: %v\n", result.URL, result.Err)
			continue
		}
		fmt.Printf("✅ %s -> %s\n", result.URL, result.Status)
	}

	fmt.Println("Crawling finished")
}
```

---

## Why add `context.Context`?

Right now:

* If one URL hangs → it waits until timeout
* If main decides to stop → workers **keep running**
* No way to cancel everything at once

`context.Context` gives you:

* Global cancellation signal
* Timeouts / deadlines
* Clean shutdown of *all* goroutines

Think of it as a **broadcast “STOP” button**.

---

## Step 1: Create a cancellable context in `main`

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

This creates:

* `ctx` → passed to all workers
* `cancel()` → stops everything immediately

---

## Step 2: Use `http.NewRequestWithContext`

Instead of:

```go
client.Get(url)
```

Use:

```go
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
resp, err := client.Do(req)
```

Now:

* If `ctx` is cancelled
* HTTP request is aborted instantly

---

## Step 3: Workers must *listen* for cancellation

Use `select`:

```go
select {
case <-ctx.Done():
	return
case url := <-jobs:
	// work
}
```

---

## Full version: Worker pool with context cancellation

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Worker int
	URL    string
	Status string
	Err    error
}

func fetch(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan string,
	results chan<- Result,
	client *http.Client,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return

		case url, ok := <-jobs:
			if !ok {
				return
			}

			resp, err := fetch(ctx, client, url)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}

			results <- Result{
				Worker: id,
				URL:    url,
				Status: resp.Status,
			}
		}
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
		"https://httpbin.org/status/404",
		"https://invalid-url",
	}

	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second) // Entire crawl → max 30s
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second} // Each request → max 5s

	jobs := make(chan string)
	results := make(chan Result)

	numWorkers := 3
	var wg sync.WaitGroup

	// Start worker pool
	for i := range numWorkers {
		wg.Add(1)
		go worker(ctx, i+1, jobs, results, client, &wg)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Close results when workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consume results
	for r := range results {
		if r.Err != nil {
			fmt.Printf("❌ Worker %d %s error: %v\n", r.Worker, r.URL, r.Err)
			// example: cancel on first fatal error
			// cancel()
			continue
		}
		fmt.Printf("✅ Worker %d %s -> %s\n", r.Worker, r.URL, r.Status)
	}

	// example: cancel everything after 2 seconds
	// time.AfterFunc(1*time.Second, cancel)

	fmt.Println("Crawling finished")
}
```

---

## Mental model

* `context.Context` = **shared cancellation signal**
* `cancel()` = press the red button
* `select` = workers constantly check if they should stop
* HTTP requests respect context automatically

---

## Common mistakes (avoid these)

❌ Creating a new context inside each worker
❌ Forgetting to pass context to HTTP requests
❌ Not checking `ctx.Done()` in long loops

---

## When to use what

| Tool                  | Purpose                            |
| --------------------- | ---------------------------------- |
| `Client.Timeout`      | Max duration of a request          |
| `context.WithTimeout` | Cancel *everything* after deadline |
| `context.WithCancel`  | Manual cancellation                |

They **work together**, not instead of each other.

</details>

---

**2. Add rate limiting with time.Ticker:**

<details>
<summary>View contents</summary>

#### Goal

* Limit the number of HTTP requests per second (or per time interval)
* Keep workers concurrent but **polite** to the server
* Combine with **context cancellation** and **results channel**

---

#### Concept: `time.Ticker` as a rate limiter

```go
ticker := time.NewTicker(time.Second / 2) // 2 requests per second
defer ticker.Stop()
```

* Each tick = permission to send **one request**
* Workers **wait for a tick** before fetching

---

#### How to integrate with worker

Inside the worker:

```go
select {
case <-ctx.Done():
    return
case <-ticker.C:
    // allowed to fetch
}
```

This ensures **no more than X requests per second** globally.

---

#### Full example: Worker pool with context, results, and rate limiting

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	Worker int
	URL    string
	Status string
	Err    error
}

func worker(ctx context.Context, id int, jobs <-chan string, results chan<- Result, client *http.Client, wg *sync.WaitGroup, ticker <-chan time.Time) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled\n", id)
			return

		case url, ok := <-jobs:
			if !ok {
				return
			}

			// Rate limiting: wait for a tick
			select {
			case <-ctx.Done():
				return
			case <-ticker:
				// allowed to fetch
			}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}

			resp, err := client.Do(req)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}
			resp.Body.Close()

			results <- Result{Worker: id, URL: url, Status: resp.Status}
		}
	}
}

func main() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/get",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second}

	numWorkers := 3
	jobs := make(chan string)
	results := make(chan Result)
	var wg sync.WaitGroup

	// Rate limiter: 1 request per second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i+1, jobs, results, client, &wg, ticker.C)
	}

	// Send jobs
	go func() {
		for _, url := range urls {
			jobs <- url
		}
		close(jobs)
	}()

	// Close results when done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for r := range results {
		if r.Err != nil {
			fmt.Printf("❌ Worker %d %s error: %v\n", r.Worker, r.URL, r.Err)
			continue
		}
		fmt.Printf("✅ Worker %d %s -> %s\n", r.Worker, r.URL, r.Status)
	}

	fmt.Println("All done")
}
```

---

#### How it works

1. `ticker := time.NewTicker(1 * time.Second)`

   * produces **1 tick per second**

2. Each worker waits on `<-ticker` before sending a request

   * ensures **global request rate**
   * multiple workers share the same ticker → smooth traffic

3. `ctx.Done()` ensures clean cancellation

4. Results are sent via `results` channel — safe & observable

---

#### Mental model

```
jobs --> worker (waits for tick) --> fetch --> results
        ↑ ctx controls cancellation
```

* `ticker` = “permission slip” to fetch
* `context` = “STOP button”
* `results` = “what happened”

</details>

---

**3. Retry with exponential backoff:**

<details>
<summary>View contents</summary>

#### What is exponential backoff?

Exponential backoff is a retry strategy where:

* The wait time **doubles** after each failure
* Optional **jitter/randomness** to avoid thundering herd problems
* Stops after **max retries**

Example sequence:

```
retry 1 → wait 1s  
retry 2 → wait 2s  
retry 3 → wait 4s  
retry 4 → fail
```

---

#### Implementing a `fetchWithRetry` function

```go
func fetchWithRetry(ctx context.Context, client *http.Client, url string, maxRetries int) (*http.Response, error) {
	var resp *http.Response
	var err error
	backoff := 1 * time.Second

	for i := 0; i <= maxRetries; i++ {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err = client.Do(req)
		if err == nil {
			return resp, nil
		}

		// Retry only if context is not done
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(backoff):
			// increase backoff exponentially
			backoff *= 2
		}
	}

	return nil, err
}
```

✅ Notes:

* `ctx` ensures **cancellation propagates**
* `time.After(backoff)` sleeps between retries
* `backoff *= 2` doubles wait each retry
* `maxRetries` limits the total attempts

---

#### Integrate with worker

```go
func worker(
	ctx context.Context,
	id int,
	jobs <-chan string,
	results chan<- Result,
	client *http.Client,
	wg *sync.WaitGroup,
	ticker <-chan time.Time,
	maxRetries int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case url, ok := <-jobs:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case <-ticker: // rate limiting
			}

			resp, err := fetchWithRetry(ctx, client, url, maxRetries)
			if err != nil {
				results <- Result{Worker: id, URL: url, Err: err}
				continue
			}
			resp.Body.Close()

			results <- Result{Worker: id, URL: url, Status: resp.Status}
		}
	}
}
```

---

#### Update `main` to pass `maxRetries`

```go
maxRetries := 3

for i := 0; i < numWorkers; i++ {
	wg.Add(1)
	go worker(ctx, i+1, jobs, results, client, &wg, ticker.C, maxRetries)
}
```

---

#### How it works together

```
jobs -> worker (wait for tick) -> fetchWithRetry -> results
        ↑ ctx cancels      ↑ maxRetries + backoff
```

* Workers respect **rate limiting** (`ticker`)
* Workers respect **cancellation** (`ctx`)
* Temporary failures are retried **without overwhelming the server**
* Results channel collects successes/errors

---

#### Optional improvement: add jitter

Exponential backoff + jitter prevents **all workers from retrying at the exact same time**:

```go
backoff := time.Duration(rand.Int63n(int64(base))) + base
```

* `rand.Int63n(base)` → random 0–base
* `base` doubles after each retry

</details>

---