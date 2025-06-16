package main

import(
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

func writer(ch chan string,wg *sync.WaitGroup){
	defer wg.Done()
	ch <- strconv.Itoa(rand.Int())
}

func reader(ch chan string,wg *sync.WaitGroup){
	defer wg.Done()
	msg := <- ch
	fmt.Println(msg)
}

func main(){
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go reader(ch, &wg)
	wg.Add(1)
	go writer(ch, &wg)
	wg.Wait()
}
