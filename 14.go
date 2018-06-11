package main

import ("fmt")

func main() {
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Sam"] = 92
	grades["yan"] = 98.5
	grades["noone"] = 0.0
	fmt.Println(grades)
	fmt.Println("deleting no one!!!")
	delete(grades, "noone")

	for k, v := range grades{
		fmt.Println(k, ":", v)
	}
}

