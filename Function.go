package main

import "fmt"

func plus(a int,b int )int  {
	return a+b

}
func plusplus(d,e,f int)int {
	return d+e+f
}
func main()  {
	res := plus(1,9)
	fmt.Println("Result:", res)
	res2 := plusplus(70,60,90)
	fmt.Println("result2:",res2)
}
