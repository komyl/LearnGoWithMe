package main

import "fmt"

func main(){
	message := make(chan string)

	go func() {message <- "mamad"}()

	msg := <- message
	fmt.Println(msg)
}
