package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func fetchData(apiURL string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("error loading data: %s\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("fetching done")
}

func main() {
	var wg sync.WaitGroup
	startTime := time.Now() // Correct time function

	apiURL := "https://jsonplaceholder.typicode.com/todos/1"

	// Add one goroutine to the WaitGroup
	wg.Add(1)
	go fetchData(apiURL, &wg)

	// Start main flow
	fmt.Println("Main flow starts")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Main task %d\n", i)
		time.Sleep(1 * time.Second)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("Main flow ends")

	// Calculate and print time taken
	endTime := time.Now()
	fmt.Printf("Time taken: %v\n", endTime.Sub(startTime))
}
