package main

import "fmt"

var a string = "global 'a' variable."

func main() {
	var x string = "Learning about variables in go."
	fmt.Println(x)

	p := "first"
	fmt.Println(p)

	var y string
	y = p + " Second"
	fmt.Println(y)

	y += " Third"
	fmt.Println(y)

	var z string = "first"
	fmt.Println(y == z)
	fmt.Println(p == z)

	fmt.Println(a)
	f()
}

func f() {

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println("You are in function f.")
	fmt.Println(a, b, c)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
