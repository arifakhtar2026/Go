package main

import (
	"fmt"
	"sync"
	// "time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {

	for i:=0;i<10;i++{
		go sol()
		wg.Add(1)
		// time.Sleep(1*time.Second)
	}
	wg.Wait()
	print(counter)
}
func sol(){
	defer wg.Done()
	mutex.Lock()
	counter++
	fmt.Println(counter)
	mutex.Unlock()
}