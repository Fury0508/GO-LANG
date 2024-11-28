package main
import "runtime"
func main() {
	// var customerHeight int = 180
	// customerAge :=20

	// switch  {
	// case customerHeight >=150 || customerAge >18:
	// 	println("can access ride")
	// case customerHeight >=120:
	// 	println("can access childer rides")
	// default:
	// 	println("cannot access rides")
	// }

	switch os:= runtime.GOOS; os {
	case "darwin":
		println("os x")
	case "linux":
		println("linux machine")
	default:
		println("somthing ekse")
	}
}