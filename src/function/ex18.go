package function

import "fmt"

// Se pone al principio con mayusculas para que sea publico
// Concurrencia:
// Concurrencia te permite estar pendiente de varios procesos, comienzas uno,
// empiezas otro, ves si el anterior ya terminó, luego crear otro así
// El paralelismo, es usar varios hilos del procesador para hacer varios procesos a la vez,
// pero siempre estas esperando que la tarea termine.
func Ex18() {
	fmt.Println("Concurrencia")

}
