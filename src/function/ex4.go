package function

import "fmt"

func normalFunction(msg string) {
	fmt.Println(msg)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func Ex4() {
	// F
	normalFunction("Hola mundo")
	tripeArgument(1, 2, "Hola")
	value := returnValue(2)
	fmt.Println("Value:", value)
	value1, value2 := doubleReturn(2)
	fmt.Println("value1 y value2", value1, value2)
}
