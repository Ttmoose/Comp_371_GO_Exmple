package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"
)

const (
	defaultLimit = 100
	maxLimit     = 1_000_000
)

func main() {
	limit, err := parseLimit(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "Usage: go run . [limit]")
		os.Exit(2)
	}

	start := time.Now()
	primes := findPrimes(limit)
	fmt.Printf("Found %d primes up to %d using a concurrent worker pool in %s:\n",
		len(primes), limit, time.Since(start).Round(time.Microsecond))
	fmt.Println(primes)
}

func parseLimit(args []string) (int, error) {
	if len(args) > 1 {
		return 0, fmt.Errorf("expected at most one argument")
	}
	if len(args) == 0 {
		return defaultLimit, nil
	}

	limit, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, fmt.Errorf("invalid limit %q: must be an integer", args[0])
	}
	if limit < 2 || limit > maxLimit {
		return 0, fmt.Errorf("limit must be between 2 and %d", maxLimit)
	}
	return limit, nil
}

func findPrimes(limit int) []int {
	workerCount := runtime.NumCPU()
	if workerCount > limit-1 {
		workerCount = limit - 1
	}

	jobs := make(chan int)
	primes := make(chan int)
	var workers sync.WaitGroup
	workers.Add(workerCount)

	// 1. Worker Pool: Spawn N goroutines matching available logical CPUs
	for i := 0; i < workerCount; i++ {
		go func() {
			defer workers.Done()
			for candidate := range jobs {
				if isPrime(candidate) {
					primes <- candidate
				}
			}
		}()
	}

	// 2. Producer: Stream candidate numbers to the jobs channel
	go func() {
		for candidate := 2; candidate <= limit; candidate++ {
			jobs <- candidate
		}
		close(jobs)
	}()

	// 3. Coordinator: Wait for all workers to finish, then close results channel
	go func() {
		workers.Wait()
		close(primes)
	}()

	// 4. Collector: Drain primes channel into slice
	var found []int
	for prime := range primes {
		found = append(found, prime)
	}
	sort.Ints(found)
	return found
}

func isPrime(candidate int) bool {
	for divisor := 2; divisor*divisor <= candidate; divisor++ {
		if candidate%divisor == 0 {
			return false
		}
	}
	return true
}
