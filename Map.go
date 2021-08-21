package main

import "fmt"

func main()  {
	m:= make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:",len(m))

	delete(m,"k2")
	fmt.Println("map:",m)

	_ , prs := m["k2"]
	fmt.Println("prs:",prs)

	n := map[string]int{"mamd1" : 3 , "mamad2" : 5}
	fmt.Println("newmap:",n)

	m["mamd1"] = 8
	//you can't use mamd1 = 3 in this case , you can write again count mamad1 and print
	v2 := m["mamd1"]
	fmt.Println("mamad1:",v2)
}
