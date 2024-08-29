package main

import (
	"fmt"
	"sync"
)

func main() {

	wg:=&sync.WaitGroup{}
	var scores = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("this is the first goroutine")
		scores=append(scores, 1)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("this is the second goroutine")
		scores=append(scores, 2)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		fmt.Println("this is the third goroutine")
		scores=append(scores, 3)
		wg.Done()

	}(wg)
wg.Wait()
fmt.Println(scores)
}
