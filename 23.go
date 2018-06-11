package main

import ("fmt"
		"sync")

var wg sync.WaitGroup

func foo(c chan int, somevalue int) {
	defer wg.Done()
	c <- somevalue * 5
}

func main() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(c, i)
	}
	// before closing the channel, wait
	wg.Wait()
	close(c)
	
	for item := range c {
		fmt.Println(item)
	}
}