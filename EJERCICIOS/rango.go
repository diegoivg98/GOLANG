package main

import "fmt"

func main() {

	//declarar variables
	var a float64

	//Pedir ingreso de datos
	fmt.Print("Ingrese un numero: ")
    fmt.Scan(&a)

    if(a>=300 && a<=500){
    	fmt.Println("DENTRO DEL RANGO")
    }else{
        fmt.Println("FUERA DEL RANGO")
    }
}