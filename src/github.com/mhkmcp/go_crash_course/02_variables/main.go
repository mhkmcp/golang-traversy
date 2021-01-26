package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	var name string = "Humayun"
	var age int = 27
	// const
	// const isCool bool = false
	// isCool = true

	// Shorthand
	lastName := "Kabir"

	fName, lName := "Humayun", "Kabir"

	fmt.Println("Name: " + name + " " + lastName)
	fmt.Println("Age: ", age)
	fmt.Printf("%T\n", age)
	fmt.Println(fName, lName)

}
