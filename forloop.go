package main

import "fmt"

func main() {
	ages := []int{42, 45, 18, 72}

	for i := 0; i < len(ages); i++ {
		fmt.Println(i, ages[i])
	}

	for i := 0; i < len(ages); {
		fmt.Println(i, ages[i])
		i++
	}

	// for{
	// 	break
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	for index, value := range ages {
		fmt.Println(index)
		fmt.Println(value)
	}
}
