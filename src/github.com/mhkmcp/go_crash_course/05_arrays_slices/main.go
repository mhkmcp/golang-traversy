package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [3]string

	// fruitArr[0] = "Mango"
	// fruitArr[1] = "Apple"
	// fruitArr[2] = "Litchi"

	fruitArr := [3]string{"Mango", "Apple", "Litchi"}

	fruitslice := []string{"Mango", "Apple", "Litchi"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[2])
	fmt.Println(fruitslice)
	fmt.Println(fruitslice[0:2])
	fmt.Println(len(fruitslice))
}
