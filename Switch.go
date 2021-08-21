package main

import (
	"fmt"
	"time"
)

func main()  {
	i := 2
	fmt.Println("write","as")
	switch i {
	case 1 :
		fmt.Println("one")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 24 :
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")

	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool :
			fmt.Println("i'm a bool")
		case int :
			fmt.Println("i'm a int")
		default:
			fmt.Println("don't know type %T\n",t)
		}
	}
    whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
