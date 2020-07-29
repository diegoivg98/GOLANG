package main

import "fmt"

func main() {

	//declarar variables
	var a float64
	var b float64
	var c float64
	var res float64

	//Pedir ingreso de datos
	fmt.Print("Ingrese un numero: ")
    fmt.Scan(&a)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&b)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&c)

    res = (a+b)*c/2

    fmt.Println("Resultado: ",res)
}