package main

import "fmt"

func main() {

	message := make(chan string,2)

	message <- "mamd"
	message <- "hamed"
	fmt.Println(<-message)
	fmt.Println(<-message)

}
