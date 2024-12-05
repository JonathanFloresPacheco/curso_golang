package function

import "fmt"

func Ex9() {
	//Defer
	defer fmt.Println("Hola") //Va a ejecutar esto hasta que todo lo demas muera
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//Continue (Se utiliza cuando deseas que continue el algoritmo)
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

}
