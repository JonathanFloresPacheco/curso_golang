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
//
//	func say1(text string, c chan<- string) {
//		c <- text
//	}
func message(text string, c chan string) {
	c <- text
}
func Ex21() {
	fmt.Println("Range, Close y Select en channels")
	// c, chan cuantos datos simultaneanos va a manejar este canal
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"
	//LEN tiene un canal CAP capacidad de 2
	fmt.Println(len(c), cap(c))

	//Range y close
	//Close: decirle al runtime de go cerrar el channel
	//       lo ideal es cerrar el canal una vez que se a ocupado

	close(c)
	// c <- "Mensaje 3"

	//Range: nos permite realizar un recorrido de los mensajes de ese channel

	for message := range c {
		fmt.Println(message)
	}
	// Select : cuando empezamos a manejar multiples canales es cuando vamos a  ocupar select
	email := make(chan string)
	email2 := make(chan string)
	go message("msj1", email)
	go message("msj2", email2)
	//Como vamos a saber que chanel va a responder primero para eso utilizamos select

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email:
			fmt.Println("Email recibido de emial1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}
}
