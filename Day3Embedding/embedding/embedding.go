package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("Some generic animal sound")
}

type Dog struct {
	Animal // Embedding struct Animal
	Breed  string
}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func main() {
	dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
	dog.Speak()
	dog.Animal.Speak() // Gọi method từ struct embedded
}
