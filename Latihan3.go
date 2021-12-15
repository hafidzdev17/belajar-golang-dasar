package main

import "fmt"

func main() {

	//	TODO: Type Declaration
	type NUMBER int

	// TODO: Arithmetic Operators
	var number1 NUMBER = 10
	var number2 NUMBER = 10

	fmt.Println("\n")
	fmt.Println(number1 + number2)
	fmt.Println(number1 - number2)
	fmt.Println(number1 * number2)
	fmt.Println(number1 / number2)
	fmt.Println(number1 % number2)

	// TODO: Relational Operators
	var value1 NUMBER = 50
	var value2 NUMBER = 25

	fmt.Println("\n")
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)

	// TODO: Logical Operator
	var check NUMBER = 99
	var test NUMBER = 88

	fmt.Println("")
	fmt.Println(check == test && check >= test)
	fmt.Println(check == test || check != test)

	// TODO: Bitwise Operator
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	fmt.Println("")

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << b /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> b /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - Value of c is %d\n", c)


	// TODO: Assignment Operators
	var i NUMBER = 10

	fmt.Println("")
	i++
	i+=10
	i-=10
	i*=10

	fmt.Println(i)


}
