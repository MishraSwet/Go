package main

import "fmt"

func main() {
	//all loops must-have a body inside {} initial value ==> for i:=;;{}
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	//this is similiar to while
	j := 5
	for j >= 0 {
		fmt.Println(j)
		j = j - 1
	}

	//similiar to in range upto range-1
	for i := range 6 {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

//there are no parenthesis involved in block statements in go!
