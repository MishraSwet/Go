package main

import "fmt"

func main() {
	//similiar to for loop for a block code {} are involved but not ()
	if 6%2 == 0 {
		fmt.Println("divisible")
	} else {
		fmt.Println("indivisible")
	}

	// both if and else should be followed with a {} where the } ends for if the else statement should start on the same line

	if num := 9; num > 0 {
		fmt.Println("+ve")
	} else if num < 0 {
		fmt.Println("-ve")
	} else {
		fmt.Println("0")
	}

	//to init use var_name:=val ; condn
}
