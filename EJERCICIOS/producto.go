package main

import "fmt"

func main() {

	//declarar variables
	var n1 int
	var n2 int 
	var mult int 

	//Pedir ingreso de datos
	fmt.Print("Ingrese un numero: ")
    fmt.Scan(&n1)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&n2)

    mult = n1*n2

    fmt.Println("El producto es:",mult)
}