package function

import "fmt"

type pc1 struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc1) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

// Se pone al principio con mayusculas para que sea publico
func Ex16() {
	fmt.Println(" Stringers: personalizar el output de Structs")
	myPC := pc1{ram: 16, disk: 200, brand: "intel"}
	fmt.Println("myPC", myPC)
	// sE PERSONALIZA EL STRUCT de salida  con el sprintf
}
