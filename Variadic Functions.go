package main

import "fmt"

func sum(nums ...int) {
	//fmt.Println(nums," ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(5,9,4)
    sum(33,23)
}
