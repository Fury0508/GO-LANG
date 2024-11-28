package main

import "fmt"

func HelloWorld(name string, age, height int) {
	fmt.Println("Hello", name)
	fmt.Println("Age", age)
	fmt.Println("Height", height)
}

func AddTotal(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	HelloWorld("dhananjay", 26, 172)
	total, negativeTotal := AddTotal(2, 3)
	fmt.Println(total, negativeTotal)
}
