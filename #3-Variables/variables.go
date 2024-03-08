package main

import "fmt"

func main() {
	fmt.Println("This is a file about variables")
	//we either have to initialize the variable for go to infer its type
	//or we need to declare the variable with the type without initializing it
	//we use the var keyword to create variables
	//we can create multiple variables using the var keyword

	//single initialized and declared variable
	var a int = 2
	fmt.Println(a)

	//same type variables
	var b, c int = 2, 3
	fmt.Println(b)
	fmt.Println(c)

	//uninitialized declared variables init-to 0 by default
	//variable with_name this and with type this and is = value;
	var fl float32
	fmt.Println(fl)

	var bl bool
	fmt.Println(bl)

	var str string
	fmt.Println(str)

	var integer int
	fmt.Println(integer)

	//initialized and undeclared(inferred)
	var int1 = 16
	var str1 = "xyz"
	var bool1 = false
	var float1 = 3.33

	fmt.Println(int1)
	fmt.Println(str1)
	fmt.Println(bool1)
	fmt.Println(float1)

}

//default bool --> false
//default int/float32 --> 0 and default string --> "" empty string
