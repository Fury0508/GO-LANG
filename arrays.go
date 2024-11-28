package main

import "fmt"

func main() {
	planets := [4]string{"mercury", "venus", "earth", "saturn"}
	fmt.Println(planets)

	planetslice := [4]string{"mercury", "venus", "earth", "saturn"}
	fmt.Println(planetslice)

	var planetsliceVerbose []string
	fmt.Println(planetsliceVerbose)
	planetsliceVerbose = append(planetsliceVerbose, "mercury")
	fmt.Println(planetsliceVerbose)
}
