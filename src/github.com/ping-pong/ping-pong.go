package main

// Ex 9.5 "The Go Programming Language"

import (
	"time"
	"fmt"
)

var chan1 = make(chan string)
var chan2 = make(chan string)
var count int = 2000000

func ping(){
	start := time.Now()
	defer func (){
		fmt.Printf("time take %s", time.Since(start))
	}()
	for i:=0; i <count; i++{
		//fmt.Println("1.write ping")
		chan1 <- "ping"
		//fmt.Println("1.recv pong")
		<-chan2

	}

}

func pong(){
	for i:=0; i <count; i++{
		//fmt.Println("2.read ping")
		<-chan1
		//fmt.Println("2.write pong")
		chan2<-"pong"


	}

}


func main() {

	go pong()
	ping()

}

