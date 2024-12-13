// Anteriormente en GO se tenía que trabajar el código dentro del GOPATH. Este GOPATH es una de las variables de
// entorno que declaramos al principio del curso (export GOPATH=$HOME/go) y no es más que la ruta a nuestro entorno
// de trabajo en donde encontraremos/crearemos la carpeta source que es donde trabajamos todo nuestro código.

// Afortunadamente desde la versión 1.11 podemos realizar nuestros proyectos de Go desde donde queramos gracias a los módulos.

// Veamos un ejemplo.

// Cree un proyecto llamando 0.0-fundamentals en la siguiente ruta /home/est14/my-own-path/15-go-course/0.0-fundamentals como
// se puede apreciar estoy fuera de mi GOPATH (/home/est14/go/source).

// Ahora cree mi go.mod aquí, en la raíz de 0.0-fundamentals con el siguiente comando.

// go mod init est14/0.0-fundamentals
// Esto me creo el archivo go.mod con el siguiente contenido :

// module est14/0.0-fundamentals

// go 1.16  // Esta es la version de Go
// Luego cree un archivo llamado main.go, al mismo nivel de este cree un directorio llamado greetings,
// dentro de greetings cree un archivo llamado greetings.go y finalmente cree una función llamada Hola.

// Mi estructura quedo asi:

// ├── go.mod
// ├── greetings
// │ └── greetings.go // Aquí cree mi funcion Hola
// └── main.go

// Ahora mi intención era traer la función llamada Hola que estaba en greetings.go a mi archivo main.go y
// eso lo hice importando mi paquete “est14/0.0-fundamentals/grettings”

// Este es el contenido de mi archivo main.go
package function

import "fmt"

// Se pone al principio con mayusculas para que sea publico
func Ex14() {
	fmt.Println("First Explication with module")

}

// Y listo finalmente pude importar mi paquete y utilizar mi función Hola.
//  Hay mucho detrás de todo esto y sería bueno detenerse un poco aquí ir por
//  un helado luego leer un poco de documentación y seguir ;).
