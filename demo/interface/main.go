package main

import "fmt"

type animal interface {
	eat()
}

type human int
type cat int
type dog int

func (h human) eat() {
	fmt.Println("Human eat")
}

func (c cat) eat() {
	fmt.Println("Cat eat")
}

func (d dog) eat() {
	fmt.Println("Dog eat")
}

func callEat(a animal) {
	a.eat()
}

func callEat2(i interface{}) {
	i.(animal).eat()
}

func main() {
	var h human
	var c cat
	callEat2(h)
	callEat2(c)
}
