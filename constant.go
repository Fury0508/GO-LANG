package main

import (
	"fmt"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n float64 = 2
	const d float64 = 7 / n

	fmt.Println(float32(d))

}
