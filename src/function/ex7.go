package function

import (
	"fmt"
	"log"
	"strconv"
)

func ifConditional(v1, v2 int) bool {
	if v1 < v2 && v2 == 2 {
		return true
	} else {
		return false
	}
}
func ifConditionalOR(v1, v2, v3 int) bool {
	if v1 < v2 || v2 < v3 && v2 == 2 {
		return true
	} else {
		return false
	}
}

// Convertir un numero de texto a numero entero
func numbertoText(v1 string) int {
	value, err := strconv.Atoi(v1)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
func isPar(v1 int) bool {
	//Modulo resultado de una division, el remanente, si queremos saber si el numero es par o impar
	result := v1 % 2
	fmt.Println("Modulo: ", result)
	if result == 0 {
		return true
	} else {
		return false
	}
}

func Ex7() {
	v1 := 1
	v2 := 2
	v3 := 50
	n1 := "1"
	fmt.Printf("%v es menor que %v\n", v1, v2)
	fmt.Println("El resultado es: ", ifConditional(v1, v2))
	fmt.Printf("%d es menor que %d o si %d \n", v1, v2, v3)
	fmt.Println("El resultado es: ", ifConditionalOR(v1, v2, v3))
	fmt.Printf("El numero a convertir en texto es %d\n", v1)
	fmt.Println("El resultado es: ", numbertoText(n1))
	fmt.Printf("El numero es par %d\n", v3)
	fmt.Println("El resultado es: ", isPar(v3))
}
