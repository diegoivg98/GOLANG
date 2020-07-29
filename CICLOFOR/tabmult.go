package main

import "fmt"

func main() {

	var n int
	var res int
	fmt.Print("Ingrese un numero: ")
    fmt.Scan(&n)

    //IMPRIMIR UNA TABLA DE MULTIPLICACION
	for i := 1; i < 11; i++ {
		res = n*i
		fmt.Println(n,"X",i," = ",res)
	}
}