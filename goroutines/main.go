package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup
var mt sync.Mutex
var signals = []string{"test"}
func main() {
	fmt.Println("Go routines for corrcurency")
	webList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}
	for _, web := range webList {
		
		wg.Add(1)
		go getStatus(web)
	}
	wg.Wait()
	fmt.Println(signals)
}
func getStatus(s string){
	defer wg.Done()
	res,err :=http.Get(s)
	if err != nil{
		fmt.Println("Error")
	}else{
		mt.Lock()
		signals = append(signals, s)
		mt.Unlock()
        fmt.Printf("%dStatus code is %s\n",res.StatusCode, s)
	}
}