package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

func timSONT(x int) bool {
	if x < 2 {
		return false
	} else {
		for i := 2; i < x; i++ {
			if x%i == 0 {
				return false
			}
		}
	}
	return true

}

func initList() *list.List {
	l := list.New()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		// Sinh số ngẫu nhiên trong khoảng từ 0 đến 20
		randomNumber := rand.Intn(21)
		l.PushBack(randomNumber)
	}
	return l

}
func main() {
	l := initList()
	fmt.Println("Số nguyên tố có trong list: ")
	for e := l.Front(); e != nil; e = e.Next() {
		if timSONT(e.Value.(int)) {
			fmt.Printf("%d\t", e.Value)
		}
	}
}
