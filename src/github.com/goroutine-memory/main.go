package main

import (
	"fmt"
	"time"
	"runtime"
)

var count = 22000000

func fun ( in chan int, out chan int){
	for i := range in {
		out <- i+1
	}
}

func profiling(){
	for {
		select{
			case <-time.After(5*time.Second):
				fmt.Printf("num:%d\n", runtime.NumGoroutine())
		}
	}
}

func main(){
	start := time.Now()
	defer func (){
		fmt.Printf("time take %s", time.Since(start))
	}()
	go profiling()
	in := make (chan int)
	go func(c chan int){
		c <- 0
	}(in)
	var out chan int
	for i:=0; i < count; i++ {
		out = make(chan int)
		go fun(in, out)
		in = out
	}
	final := <-out
	fmt.Printf("final:%d\n", final)
}
