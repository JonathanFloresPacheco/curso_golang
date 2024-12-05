package function

import "fmt"

func forElement(numberElem int) {
	for i := 0; i < numberElem; i++ {
		fmt.Println(i)
	}
}
func forWhileElement(numberElem int) {
	counter := 0
	for counter < numberElem {
		fmt.Println(counter)
		counter++
	}
}
func forForever() {
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
func Ex5() {
	// for, For while y for forever
	fmt.Println("For")
	forElement(5)
	fmt.Println("For While")
	forWhileElement(10)
	// fmt.Println("For forever")
	// forForever()
}
