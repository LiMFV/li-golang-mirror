package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Print("What is your name?...: ")
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// reader = bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello, " + name)
}
