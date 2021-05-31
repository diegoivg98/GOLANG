package main

import "fmt"

func main() {

	//declarar variables
	var n1 int
	var n2 int 
	var suma int 

	//Pedir ingreso de datos
    fmt.Print("Ingrese un numero: ")
    fmt.Scan(&n1)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&n2)

    suma = n1+n2

    fmt.Println("La suma es: ",suma)
}
