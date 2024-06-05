package main

import "fmt"

// Định nghĩa struct Person
type Person struct {
	Name string
	Age  int
}

// Định nghĩa method cho struct Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	person.Greet()
}
