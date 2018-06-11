package main 

import ("fmt"
		"math"
		"math/rand")

func main() {
	fmt.Println("The square root of  4 is ", math.Sqrt(4))
	rand.Seed(42)
	fmt.Println("Random number is  ", rand.Intn(100))
}