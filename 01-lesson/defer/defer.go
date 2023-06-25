package main

import "fmt"

// LIFO (Last In First Out)
func multipleDefer() {
	fmt.Println("=========multiple defer============")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

}

func main() {

	// defer the execution of Println() function
	defer fmt.Println("Three")

	fmt.Println("One")
	fmt.Println("Two")

	multipleDefer()

}
