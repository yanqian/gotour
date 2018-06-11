package main 

import "fmt"

func add(a,b float32) float32 {
	return a + b	
}
func multiple(word1, word2 string)(string, string){
	return word2, word1
}

func main() {
	var a, b float32 = 5.6, 8.9
	fmt.Println(add(a, b))
	word1, word2 := "Hey", "there"
	fmt.Println(multiple(word1, word2))

}