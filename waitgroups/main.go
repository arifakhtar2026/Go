package main

import (
	"fmt"
	// "time"
	"net/http"
	"sync"
	// "io"
)
var wg sync.WaitGroup//basically it is a pointer to the waitgroup
// func greet (s string ){
// 	fmt.Print(s)
// 	time.Sleep(1* time.Second)
// 	fmt.Println()
// 	fmt.Println("Done 1/n")
// }
// func greet2 (s string ){
// 	fmt.Printf(s)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println()
// 	fmt.Println("Done 2")
// }


func main() {
	website:=[]string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
		"https://www.youtube.com",
	}
	go getWebsite(website[0])
	wg.Add(1)//add is used to add the number of goroutines to the waitgroup

	for _,website:=range website{
		 go getWebsite(website)
		 go getWebsite(website)
		 wg.Add(2)//add is used to add the number of goroutines to the waitgroup
	}
	wg.Wait()//do not end main execution until all the goroutines are done
}


func getWebsite(website string){
	defer wg.Done()//done is used to tell the waitgroup that the goroutine is done
	response,err:=http.Get(website)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("response and body code is = > ",response.StatusCode,response.Body)
	}
	defer response.Body.Close()
	
}