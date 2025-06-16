package main

import(
	"fmt"
	"sync"
)

func hello_world(s string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Hello!")
	fmt.Println("Good phrase: ", s)
}

func main(){
	fmt.Println("Welcome,")
	var wg sync.WaitGroup
	wg.Add(1)
	go hello_world("double in the okati", &wg)
	wg.Wait()
}
