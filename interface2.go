package main

import "fmt"

type russian struct {}
type english struct {}

type greeter interface {
	greet(string)string
}

func (r russian) greet(name string) string {
	return fmt.Sprintf("Привет %v", name)
}

func (e english) greet(name string) string {
	return fmt.Sprintf("Hello %v", name)
}

func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}

func main() {
	fmt.Println(russian{}.greet("Aidar"))
	//sayHello(&russian{}, "Aidar")
	sayHello(&english{}, "John")
}
