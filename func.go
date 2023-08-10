package main

import (
	"fmt"
	"reflect"
	"user"
)

type User struct {
	id   int
	name string
	age  int
}

func (u User) show() {
	fmt.Println(u)
}

var (
	f = func() {
		fmt.Println("Anonymous function")
	}
)

func main() {
	// showAMessage("Hello", "World", "Again", "and", "Again")
	// f()
	rabbit := User{id: 1, name: "Rabbit", age: 20}
	rabbit.show()
	rabbit.id = 2
	rabbit.show()
	user.Student()

}

func showAMessage(msg ...string) {
	fmt.Println(reflect.TypeOf(msg))
	fmt.Println(msg)
}
