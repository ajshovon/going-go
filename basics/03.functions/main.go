package main

import "fmt"

func main() {
	printHelloWorld()
	printHelloUser("Eren")
	a, b := 5, 7
	var ans int = sumIntNumbers(a, b)
	fmt.Println(ans)
}

func printHelloWorld() {
	fmt.Println("Hello, World!")
}

func printHelloUser(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func sumIntNumbers(a, b int) int {
	return a + b
}
