package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var intNum int = 10
	fmt.Println(intNum)

	const name string = "Ani"
	fmt.Println(name)

	fmt.Println(printMe(name))
}

func printMe(printVal string) string {
	fmt.Println(printVal)
	return printVal
}
