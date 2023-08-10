package main

import "fmt"

func main() {
	var i int = 0
	for i = 0; i < 10; i++ {
		var num int
		fmt.Print("Enter a number: ")
		// Lấy giá trị từ bàn phím
		fmt.Scan(&num)
		if isOdd(num) {
			fmt.Println(num, "is odd")
		} else if isEven(num) {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is neither odd nor even")
		}

		if isOddWithAND(num) {
			fmt.Println(num, "is odd")
		}
		if isEvenWithAND(num) {
			fmt.Println(num, "is even")
		}
	}
}

func isOdd(num int) bool {
	return num%2 != 0
}

func isEven(num int) bool {
	return num%2 == 0
}

func isOddWithAND(num int) bool {
	return num&1 != 0
}

func isEvenWithAND(num int) bool {
	return num&1 == 0
}
