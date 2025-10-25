package main

import "fmt"

type person struct{ age int }

func (p *person) unsafeBirthday() {
	p.age++
}
func (p *person) safeBirthday() {
	if p == nil {
		return
	}
	p.age++
}

func main() {
	var (
		baby   person
		nobody *person
	)
	fmt.Println(baby)
	fmt.Println(nobody)
	// Won't panic because the person is initialized, just with an age of 0
	baby.unsafeBirthday()
	// Will panic because the person pointer is pointing to nobody
	// nobody.unsafeBirthday()
	// Won't panic because if contains nil checks within it
	nobody.safeBirthday()
	fmt.Println(baby)
	fmt.Println(nobody)
}
