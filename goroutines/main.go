package main

import (
	"fmt"
	
	"sync"
)
func main(){
	wg :=&sync.WaitGroup{}
	mt := &sync.Mutex{}
	var score =[]int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, mt *sync.Mutex,){
		defer wg.Done()
		fmt.Println("Adding 1 R")
		mt.Lock()
		score = append(score, 1)
		mt.Unlock()
	}(wg, mt)
	go func(wg *sync.WaitGroup, mt *sync.Mutex,){
		defer wg.Done()
		fmt.Println("Adding 2 R")
		mt.Lock()
		score = append(score, 2)
		mt.Unlock()
	}(wg, mt)
	go func(wg *sync.WaitGroup, mt *sync.Mutex,){
		defer wg.Done()
		fmt.Println("Adding 3 R")
		mt.Lock()
		score = append(score, 3)
		mt.Unlock()
	}(wg, mt)
	wg.Wait()
	fmt.Println(score)
}