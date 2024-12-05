package function

import (
	"fmt"
	"strings"
)

func isPalindromo(txt string) bool {
	var txtReverse string
	for i := len(txt) - 1; i >= 0; i-- {
		// temp := string(txt[i])
		txtReverse += string(txt[i])
		//unicode.ToLower(temp)
	}
	if txtReverse == txt {
		return true
	} else {
		return false
	}
}
func Ex11() {
	slice := []string{"hola", "que", "hace"}
	for _, valor := range slice {
		fmt.Println(valor)
	}
	// Recibe un texto si es un palindromo
	// ama  es igual de leer de derecha a izquerda
	fmt.Printf("Es un palindromo %s:\n", "ama")
	//strings.ToLower convierte todo en minusculas
	fmt.Println(isPalindromo(strings.ToLower("Ama")))
}
