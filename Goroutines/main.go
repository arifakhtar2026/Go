package main

import (
	"fmt"
	"time"
	
)
func greet (s string ){
	fmt.Print(s)
	time.Sleep(1* time.Second)
	fmt.Println()
	fmt.Println("Done 1/n")
}
func greet2 (s string ){
	fmt.Printf(s)
	time.Sleep(2 * time.Second)
	fmt.Println()
	fmt.Println("Done 2")
}


func main() {
	go greet("Arif");
	 greet2("Akhtar");
}

