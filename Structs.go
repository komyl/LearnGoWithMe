package main

import "fmt"

type person struct {
	name string
	age int
}


func main() {
	//fmt.Println(person{name: "komeyl",age: 21})
	//fmt.Println(&person{name: "komeyl",age: 21})
	s := person{name: "ALI", age: 42}
	//fmt.Println(s.name)
	sp := &s
	//fmt.Println(sp.age)
	sp.age = 32
	fmt.Println(sp.age)

}
