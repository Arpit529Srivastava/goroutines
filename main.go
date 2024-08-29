package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Greeter(s string) {
	defer wg.Done() // Decrease the counter by one when the goroutine is done
	result, err := http.Get(s)
	if err != nil {
		fmt.Println("this is not something I want, change it")
	} else {
		fmt.Println("Hell yeah... status code is", result.StatusCode, s)
	}
}

func main() {
	start := time.Now()
	sites := []string{
		"https://github.com",
		"https://linkedin.com",
		"https://google.com",
		"https://apple.com",
	}
	for _, i := range sites {
		wg.Add(1) // Increment the count before starting the goroutine
		go Greeter(i) // Launch the Greeter function as a goroutine
	}
	wg.Wait() // Blocks the main method until the waitgroup counter reaches zero
	timeElapsed := time.Since(start)
	fmt.Println("This for loop took this much time with goroutines:", timeElapsed)
}
