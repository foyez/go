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
			// fmt.Printf("Worker %d cancelled\n", id)
			return

		case url, ok := <-jobs:
			if !ok {
				return
			}

			// fmt.Printf("Worker %d fetching %s\n", id, url)
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := &http.Client{Timeout: 5 * time.Second}

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

// func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for url := range jobs {
// 		fmt.Printf("Worker %d fetching %s\n", id, url)
// 		fetchUrl(url)
// 	}
// }

// func main() {
// 	urls := []string{
// 		"https://example.com",
// 		"https://golang.org",
// 		"https://httpbin.org/get",
// 	}

// 	// numWorkers := runtime.NumCPU() * 2
// 	numWorkers := 3
// 	jobs := make(chan string)
// 	var wg sync.WaitGroup

// 	for i := range numWorkers {
// 		wg.Add(1)
// 		go worker(i+1, jobs, &wg)
// 	}

// 	for _, url := range urls {
// 		jobs <- url
// 	}
// 	close(jobs)

// 	wg.Wait()
// }
