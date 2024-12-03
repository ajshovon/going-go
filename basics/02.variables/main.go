package main

import "fmt"

func main() {
	var intNumber uint8 = 201
	fmt.Println(intNumber)

	var floatNumber float32 = 1.7
	fmt.Println(floatNumber)

	var intNumberSum uint8 = uint8(floatNumber) + intNumber
	fmt.Println(intNumberSum)

	var floatNumberSum float32 = float32(intNumber) + floatNumber
	fmt.Println(floatNumberSum)

	var stringText string = "Hello, World!"
	fmt.Println(stringText)

	var boolValue bool = true
	fmt.Println(boolValue)

	var runeValue rune = 'A'
	fmt.Println(runeValue)

	intNumber2 := 201
	fmt.Println(intNumber2)

	const myConst = 3.1416
	fmt.Println(myConst)
}
