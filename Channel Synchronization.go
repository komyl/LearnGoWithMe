package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//this code have bug becuse , if you send true code is run and if you sen false code run !
	done <- true
}

func main() {
	done := make(chan bool,1)
	go worker(done)

	<-done

}
