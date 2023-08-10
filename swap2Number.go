package main

import "fmt"

func main() {
	// var num1, num2 int
	// fmt.Print("Enter two numbers: ")
	// fmt.Scan(&num1, &num2)
	num1, num2 := 10, 20
	// fmt.Println("Before swapping: ", num1, num2)
	// fmt.Println("After swapping: ", swaUseTemp(num1, num2))
	// swapUsePointer(&num1, &num2)
	fmt.Println("Before swapping: ", num1, num2, &num1, &num2)
	num1, num2 = num2, num1
	fmt.Println("After swapping: ", num1, num2, &num1, &num2)

	// fmt.Println("After swapping: ", num1, num2)
}

func swaUseTemp(num1, num2 int) []int {
	var temp int
	temp = num1
	num1 = num2
	num2 = temp
	return []int{num1, num2}
}

func swapUsePointer(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
	fmt.Println("After swapping: ", num1, num2)
}
