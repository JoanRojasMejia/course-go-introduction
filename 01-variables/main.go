package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var myIntVar int;
	myIntVar = -42;
	fmt.Println("My variable is set to", myIntVar);
	fmt.Println("myIntVar's address is", &myIntVar);

	var myBoolVar bool;
	myBoolVar = true;
	fmt.Println("My variable is set to", myBoolVar);
	fmt.Println("myBoolVar's address is", &myBoolVar);

	var myStringVar string;
	myStringVar = "Hello, World!";
	fmt.Println("My variable is set to", myStringVar);
	fmt.Println("myStringVar's address is", &myStringVar);

	var myFloatVar float64;
	myFloatVar = 22.2;
	fmt.Println("My variable is set to", myFloatVar);
	fmt.Println("myFloatVar's address is", &myFloatVar);

	var myComplexVar complex128;
	myComplexVar = 1 + 2i;
	fmt.Println("My variable is set to", myComplexVar);
	fmt.Println("myComplexVar's address is", &myComplexVar);

	myIntVar2 := 43;
	fmt.Println("My variable is set to", myIntVar2);
	fmt.Println("myIntVar2's address is", &myIntVar2);

	myBoolVar2 := false;
	fmt.Println("My variable is set to", myBoolVar2);
	fmt.Println("myBoolVar2's address is", &myBoolVar2);

	myStringVar2 := "Hello, World!";
	fmt.Println("My variable is set to", myStringVar2);
	fmt.Println("myStringVar2's address is", &myStringVar2);

	const myConstVar int = 44;
	fmt.Println("My variable is set to", myConstVar);

	/*
		int8	8 bits	-128 to 127
		int16	16 bits	-2^15 to 2^15 -1
		int32	32 bits	-2^31 to 2^31 -1
		int64	64 bits	-2^63 to 2^63 -1
		int	Platform dependent	Platform dependent
	*/

	var myInt8Var int8 = 127;
	fmt.Printf("My variable is set to %d and its address is %p\n", myInt8Var, &myInt8Var);
}