package main

import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)
	
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 20
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 = 9
	var intNum2 = 5
	fmt.Println(intNum1/intNum2) // will always return an integer and round down
	fmt.Println(intNum1%intNum2) // returns the remainder

	// var myString = "Hello \nWorld"
	var myString = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(len("Test"))
	// len gives me the amount of bytes in the string; unicode/utf8 import plus RuneCountInString returns the amount of characters in string
	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'b'
	fmt.Println(myRune)

	 var myBoolean bool = false
	 fmt.Println(myBoolean)
	 // can initialize multiple variables at once, and := shorthand allows me to drop the "var"
	 var1, var2 := 1, 2
	 fmt.Println(var1, var2)
	
	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415
	fmt.Println(pi)
}
