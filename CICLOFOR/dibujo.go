package main

import "fmt"

func main() {
	/*IMPRIMIR EL SIGUIENTE DIBUJO
	  111111
	  222222
	  333333
	  444444
	  555555
	  666666*/
	for i := 1; i < 7; i++ {
		fmt.Println(i*111111)
	}
}