package function

import "fmt"

func Ex12() {
	//Los mapas en Go son estructuras que almacenan datos en pares de clave-valor, donde:
	// Clave (key):
	// Es el identificador único que usas para acceder a un valor.
	// Todas las claves en un mapa deben ser del mismo tipo (por ejemplo, string, int, etc.).
	// Valor (value):
	// Es el dato asociado con la clave.
	// Todos los valores en un mapa también deben ser del mismo tipo (por ejemplo, int, string, etc.).

	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)
	//Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}
	//Obtener uno de los valores
	// Cuando accedemos a una llave no asiganda ella te pone el valor del 0 value, el value de int es 0
	value, ok := m["Jose"]
	fmt.Println(value, ok)
	// Como sabemos que esa llave este dentro del did
}
