package main

import "fmt"

func main() {
		/*IMPRIMIR LOS 10 PRIMEROS TERMINOS DE LA SERIE DE FIBOACCI
		0-1-2-3-5-8-13-21-34-55*/

	var x int = 0
	var y int = 1
	var z int = 0

	for i := 0; i < 9; i++ {
		z = x+y
		x = y
		y = z
	fmt.Println(z)
}
}