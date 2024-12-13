package function

import (
	"fmt"
	"sync"
	"time"
)

// Se pone al principio con mayusculas para que sea publico
// Concurrencia:
// Concurrencia te permite estar pendiente de varios procesos, comienzas uno,
// empiezas otro, ves si el anterior ya terminó, luego crear otro así
// El paralelismo, es usar varios hilos del procesador para hacer varios procesos a la vez,
// pero siempre estas esperando que la tarea termine.
func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

// Debemos colocar el time.sleep pero no es recomendable
// Ya que debe de esperar
func Ex19() {
	fmt.Println("Primer contacto con las Goroutines")
	//Acumula un conjunto de goroutines y las va liberando poco a poco
	var wg sync.WaitGroup
	fmt.Println("Hello")
	// Agregamos 1 goroutines al waitgroup
	wg.Add(1)
	// Que imprima world y le mandamos el puntero de wait group
	go say("World", &wg)
	// le decirmos que espere hasta que haga el done de cada goroutine que fue enviada
	wg.Wait()
	// Funcion anonima: tiene su vida dentro de la misma funcion
	// esta esta como goroutine
	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)

	wg.Add(1)
	go say("Si si ya vete", &wg)
	wg.Wait()

}
