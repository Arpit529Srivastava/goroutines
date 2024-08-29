package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup //pointer
func Greeter(s string){
	defer wg.Done()
	result, err := http.Get(s)
	if err!=nil{
		fmt.Println("this is not somthing i want change it")
	}else{
		fmt.Println("Hell yeah...status code is", result.StatusCode, s)
	}
	
}
func main(){
	sites:=[]string{
		"https://github.com",
		"https://linkedin.com",
		"https://google.com",
		"https://apple.com",
	}
	for _, i:= range sites{
		go Greeter(i)
		wg.Add(1)
	}
wg.Wait()
}
