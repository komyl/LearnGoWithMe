package main

import (
	"fmt"
	"math"
)
const a string= "komeyl"

func main()  {
	fmt.Println(a)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Cos(n))


}

