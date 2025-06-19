// This code leads to non deterministic behaviour since the variable i is being tried to overwrite by multiple go routines running concurrently. 
package main

import (
	"fmt"
	"time"
	"sync"
)

func demo(i *int64, wg *sync.WaitGroup){
	defer wg.Done()
	t := time.Now()
	fmt.Printf("%dY %dM %dD %dh %dm %ds %dns - %d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), *i)
	*i += int64(1)
}

func main(){
	var wg sync.WaitGroup
	var i int64
	wg.Add(1)
	go demo(&i, &wg)
	go demo(&i, &wg)
	go demo(&i, &wg)
	go demo(&i, &wg)
	wg.Wait()
}

// One misunderstanding that I had is that, "Well, I am using a WaitGroup here so my go routines must be synchronized." and by synchronized I assumed that the four routines will run one after the other as in the order that I have created them.
// But all a WaitGroup does it maintain a counter which keeps track of how many routines are running and the wg.Wait() will not let the main function terminate until the counter becomes 0
// Meaning, it does not control the order of execution of routines. Think about it that's the whole point of concurrency order must not matter, the routines must be independent and must run concurrently. Hence, here two routines might be trying to access i at the same time or i might be blocked by one routine when the other might be trying to access etc.
