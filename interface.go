package main

import "fmt"

type animals interface {
	makeSound()
}

type cat struct {}

type dog struct {}

func (c *cat)makeSound(){
	fmt.Println("meow")
}

func (d *dog)makeSound(){
	fmt.Println("woof")
}

func main(){

	var c, d animals = &cat{}, &dog{}
	d.makeSound()
	c.makeSound()


}

