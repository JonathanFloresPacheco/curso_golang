package function

import "fmt"

func Ex2() {
	// Suma
	x := 10
	y := 50
	fmt.Println("Suma:", x+y)
	//Resta
	result := y - x
	fmt.Println("Resta: ", result)
	//Multiplicacion
	result = y * x
	fmt.Println("Multiplicacion: ", result)
	//Division
	result = y / x
	fmt.Println("Division: ", result)
	//Modulo resultado de una division, el remanente, si queremos saber si el numero es par o impar
	result = y % x
	fmt.Println("Modulo: ", result)
	//Incremental
	x++
	fmt.Println("Incremental", x)
	// Decremental
	x--
	fmt.Println("Decremental", x)

}
