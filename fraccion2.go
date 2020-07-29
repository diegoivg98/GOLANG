package main

import "fmt"

func main() {

	//declarar variables
	var a float64
	var b float64
	var res1 float64
    var res2 float64
    var suma float64

	//Pedir ingreso de datos
	fmt.Print("Ingrese un numero: ")
    fmt.Scan(&a)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&b)
  

    res1 = (a+3)/2
    res2 = (b*8)/3
    suma = res1+res2

    fmt.Println("Resultado:",suma)
}