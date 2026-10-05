// Command main showcases some of Go's strengths: simple static typing,
// interfaces, generics, goroutines/channels, and explicit error handling.
package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
)

// Shape demonstrates implicit interface satisfaction.
type Shape interface {
	Area() float64
}

type Circle struct{ R float64 }
type Rect struct{ W, H float64 }

func (c Circle) Area() float64 { return math.Pi * c.R * c.R }
func (r Rect) Area() float64   { return r.W * r.H }

// Map applies f to every element (generics).
func Map[T, U any](in []T, f func(T) U) []U {
	out := make([]U, 0, len(in))
	for _, v := range in {
		out = append(out, f(v))
	}
	return out
}

// Divide demonstrates explicit error handling.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// SquareAll squares numbers concurrently using goroutines and channels.
// Results are returned in input order.
func SquareAll(nums []int) []int {
	results := make([]int, len(nums))
	var wg sync.WaitGroup
	for i, n := range nums {
		wg.Add(1)
		go func(i, n int) {
			defer wg.Done()
			results[i] = n * n
		}(i, n)
	}
	wg.Wait()
	return results
}

// Producer streams values over a channel.
func Producer(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	fmt.Println("== Interfaces ==")
	shapes := []Shape{Circle{R: 2}, Rect{W: 3, H: 4}}
	for _, s := range shapes {
		fmt.Printf("%T area = %.2f\n", s, s.Area())
	}

	fmt.Println("== Generics ==")
	fmt.Println(Map([]int{1, 2, 3}, func(i int) string { return fmt.Sprint(i * 10) }))

	fmt.Println("== Error handling ==")
	if _, err := Divide(1, 0); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("== Concurrency ==")
	fmt.Println(SquareAll([]int{1, 2, 3, 4, 5}))
	sum := 0
	for v := range Producer(10) {
		sum += v
	}
	fmt.Println("sum of 1..10 from channel =", sum)
}
