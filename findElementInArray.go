package main

import "fmt"

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(findElementInArray(array, 3))
}

func findElementInArray(array [5]int, element int) int {
	for index, value := range array {
		if value == element {
			return index
		}
	}
	return -1
}
