package main

import ("fmt"
		"time"
		"sync")

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i<3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	
	
}

func main() {
	wg.Add(1)
	go say("Hey there")
	wg.Add(1)
	go say("Yan Qiang")

	wg.Wait()

	wg.Add(1)
	go say("What is going on")
	wg.Wait()
}