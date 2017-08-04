package main

import (
	"fmt"
	"sync"
	"errors"
	"time"
)

var timeOutError = errors.New("timeout")
var wg sync.WaitGroup 
var errChan = make(chan error, 3)


func doSth(i int) <-chan struct{} {
	resultCh := make(chan struct{})
	go func() {
		<-time.After(time.Duration(i) * time.Second)
		resultCh <- struct{}{}
	}()
	return resultCh
}

func f1(){
	defer wg.Done()

	select {
		case <- doSth(1):
			fmt.Printf("f1 success\n")
			return
		case <- time.After(1500*time.Millisecond):
			fmt.Printf("f1 error\n")
			errChan <- timeOutError
	}
	
}

func f2(){
	defer wg.Done()
	select {
		case <- doSth(2):
			fmt.Printf("f2 success\n")
			return
		case <- time.After(1*time.Second):
			fmt.Printf("f2 error\n")
			errChan <- timeOutError
	}
}

func f3(){
	defer wg.Done()
	select {
		case <- doSth(3):
			fmt.Printf("f3 success\n")
		case <- time.After(5*time.Second):
			fmt.Printf("f3 error\n")
			errChan <- timeOutError
	}
}

func main() {
	wg.Add(1)
	go f1()
	wg.Add(1)
	go f2()
	wg.Add(1)
	go f3()
	
	// closer waits for all workers to finish, and notify main go routine.
	go func() {
		wg.Wait()
		close(errChan)
	}()
	
	for err := range errChan {
		// return early in case of error in one of the go routine
		if err != nil {
			return
		}
	}
}

