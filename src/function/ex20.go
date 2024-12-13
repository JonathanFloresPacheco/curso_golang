package function

import (
	"fmt"
)

// Se pone al principio con mayusculas para que sea publico
// Concurrencia:
// Concurrencia te permite estar pendiente de varios procesos, comienzas uno,
// empiezas otro, ves si el anterior ya terminó, luego crear otro así
// El paralelismo, es usar varios hilos del procesador para hacer varios procesos a la vez,
// pero siempre estas esperando que la tarea termine.
// func say(text string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(text)
// }

// Channels: pasar informacion entre los goroutines
// Solo entrada de datos del channel
func say0(text string, c chan<- string) {
	c <- text
}
func Ex20() {
	fmt.Println("Channels: La forma de organizar las goroutines")
	// c, chan cuantos datos simultaneanos va a manejar este canal
	c := make(chan string, 1)
	fmt.Println("Hello")
	go say0("Bye", c)
	fmt.Println(<-c)

}
