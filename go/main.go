package main

import (
	"fmt"
)

//const append = 5

func main() {
	fmt.Println("value of one() is ", one()) // print 1
	append := 1
	fmt.Println(append)
	{
		fmt.Println(append)
		append := 2
		fmt.Println(append)
		{
			fmt.Println(append)
			append := 3
			fmt.Println(append)
		}
		fmt.Println(append)
	}
	fmt.Println(append)
	one := one()
	fmt.Println("value of one is ", one) // print 1
	{
		one := 5
		fmt.Println("value of one is ", one) // print 5
	}
	// two := one() + 1;
	// Error: cannot call non-function one (type int)

	two := one + 1
	fmt.Println("value of two is ", two) // print 2

	x := 4
	fmt.Println("factorial of ", x, " is ", factorial(x))

	//array := make([]int, 0, 10)
	//array = append(array, 1)
}

func one() int {
	return 1
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	factorial := func() int { return factorial(x-1) * x }()
	return factorial
}
