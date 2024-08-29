package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("this a demo for channels in golang")
	mycha := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// fmt.Println(<-mycha)
	// mycha <- 5
	wg.Add(2)
	//recieve only channel
	//so you can't put close(mycha) here cause doesn't make sense
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, IschannelsOpen:= <-mycha
		fmt.Println(IschannelsOpen)
		fmt.Println(val)
		//fmt.Println("hello channels", <-mycha)
		//fmt.Println("hello channels", <-mycha)
		wg.Done()
	}(mycha, wg)
	//send only channel
	go func(ch chan int, wg *sync.WaitGroup) {
		close(mycha)
		//mycha <-0
		// mycha <-6
		//close(mycha)
		wg.Done()
	}(mycha, wg)

	wg.Wait()
}
//finally end of the videos lectures
//now start building projects