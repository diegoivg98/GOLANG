package main

import "fmt"

func main() {

	//declarar variables
	var n1 float64
	var n2 float64
	var div float64 

	//Pedir ingreso de datos
	fmt.Print("Ingrese un numero: ")
    fmt.Scan(&n1)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&n2)

    div = n1/n2

    fmt.Println("Resultado:",div)
}