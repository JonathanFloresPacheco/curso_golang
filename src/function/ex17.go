package function

import "fmt"

//
type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area", f.area())
}

// Se pone al principio con mayusculas para que sea publico
// Interface: un metodo que se puede compartir en diferentes metodos y ese punto de entrada es la interface
func Ex17() {
	fmt.Println(" interface")
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}
	calcular(myCuadrado)
	calcular(myRectangulo)
	// Lista de interfaces creada personalizada
	myInterface := []interface{}{"Hola", 12, 4.90, 10}
	fmt.Println(myInterface...)

}
