package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var scores = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {

		fmt.Println("this is the first goroutine")
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {

		fmt.Println("this is the second goroutine")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {

		fmt.Println("this is the third goroutine")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	wg.Wait()
	fmt.Println(scores)
}
