package main

import "fmt"

// Định nghĩa interface
type Speaker interface {
	Speak()
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func MakeItSpeak(s Speaker) {
	s.Speak()
}

func main() {
	cat := Cat{}
	MakeItSpeak(cat)
}
