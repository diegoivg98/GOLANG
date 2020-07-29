package main

import "fmt"

func main() {

	//declarar variables
	var a float64
	var b float64
	var c float64

	//Pedir ingreso de datos
	fmt.Print("Ingrese un numero: ")
    fmt.Scan(&a)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&b)
    fmt.Print("Ingrese otro numero: ")
    fmt.Scan(&c)

    if(a>b && a>c){//"&&"->AND
    	fmt.Println("El mayor es:",a)
    }
    if (b>a && b>c){
    	fmt.Println("El mayor es:",b)
    }
    if(c>a && c>b){
    	fmt.Println("El mayor es:",c)
    }
}