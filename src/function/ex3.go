package function

import "fmt"

func Ex3() {
	//Paquete FMT
	// Declaración de variables
	helloMessage := "!Hola"
	worldMessage := "Mundo!"

	// Println: Hace un salto de linea al terminar el mensaje
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)
	fmt.Println("")

	// Printf: Imprime con un formato
	nombre := "Platzi"
	cursos := 700
	// Cuando sabemos que tipo de variables son podemos usar
	/*
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	*/
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)

	// Si desconocemos el tipo de variable en el formato
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf: Genera un string pero no lo imprime en consola
	var message string = fmt.Sprintf("%s es un string", nombre)
	fmt.Println(message)
	fmt.Println("")

	// Imprimir el tipo de variable
	fmt.Printf("HelloMessage es tipo: %T\n", helloMessage)
	fmt.Printf("Cursos es tipo: %T\n", cursos)
	fmt.Println("")

	// Si queremos ignorar el % del formato podemos usar doble %%
	fmt.Printf("%s es de tipo %T y para usarlo en el Printf se hace uso del %%s\n", nombre, nombre)

}
