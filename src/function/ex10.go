package function

import "fmt"

func Ex10() {
	var array [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println("Array: ", array)
	fmt.Println("Cuantos elementos existen: ", len(array))
	fmt.Println("Cual es la capacidad maxima: ", cap(array))
	//Slice No se le indica la cantidad de valores que va a tener
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("slice: ", slice)
	fmt.Println("Cuantos elementos existen: ", len(slice))
	fmt.Println("Cual es la capacidad maxima: ", cap(slice))

	//Metodos en el slice
	fmt.Println(slice[0])   //Solo posicion exacta
	fmt.Println(slice[:3])  //Solo a esta posicion
	fmt.Println(slice[2:4]) //Ranggo
	fmt.Println(slice[4:])  // de este espacio en adelante

	//Append
	slice = append(slice, 11)
	fmt.Println(slice)
	// Cuando queremos agregar una lista append nueva list
	// Se le agregar nuevos elementos al array
	newslice := []int{12, 13, 14}
	slice = append(slice, newslice...)
	fmt.Println(slice)
}
