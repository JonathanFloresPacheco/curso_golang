package function

import "fmt"

func Ex1() {
	const pi float64 = 3.14
	const pi2 = 3.1416
	fmt.Println("Pi", pi)
	fmt.Println("Pi2", pi2)
	//Declaracion de variables enteras
	// := variable que no a sido declara anteriormente
	base := 12
	var altura int = 14
	var area int
	fmt.Println("base", base)
	fmt.Println("altura", altura)
	fmt.Println("area", area)

	fmt.Println("(base * altura / 2)", base*altura/2)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("a, b, c, d", a, b, c, d)

	//Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("areaCuadrado", areaCuadrado)
}
