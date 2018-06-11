package main

import ("fmt"
		"time"
		"sync")

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovering from: ", r)
	}
}

func say(s string) {
	defer cleanup()
	for i := 0; i<3; i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Ah ha, a 2!")
		}
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