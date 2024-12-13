package function

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
}

// Se pone al principio con mayusculas para que sea publico
func Ex15() {
	fmt.Println("Strucs y punteros")
	a := 50
	b := &a //Direccion de memoria de la variable
	fmt.Println("a & b:Direccion de memoria", a, b)
	fmt.Println("Acceder al valor de memoria * acceder al valor de la direccion de memoria")
	fmt.Println("b & *b:Valor de Direccion de memoria", b, *b)
	fmt.Println("Ahora lo que vamos a realizar es cambiar el valor de direccion de memoria")
	*b = 100
	fmt.Println(a)
	//Los punteros se utilizan para llevar de manera más eficiente de una libreria a un paquete
	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println("My pc: ", myPC)
	myPC.ping()
	// Vamos a utilizar los punteros para modificar los valores del struct
	fmt.Println("myPC: ", myPC)
	myPC.duplicateRam()
	fmt.Println("myPC: ", myPC)
	myPC.duplicateRam()
	fmt.Println("myPC: ", myPC)

}
