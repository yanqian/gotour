package main

import "fmt"

func foo(c chan int, somevalue int) {
	c <- somevalue * 5
}

func main() {
	c := make(chan int)
	go foo(c, 3)
	go foo(c, 5)
	v1, v2 := <- c, <- c

	fmt.Println(v1)
	fmt.Println(v2)

}