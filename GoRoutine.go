package main

import "fmt"

func f(from string) {
	for i := 0 ; i < 3 ; i++ {
		fmt.Println(from, ";",i)
	}
}

func main() {
	f("heloo")

	go f("mamad")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
