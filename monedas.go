package main

import "fmt"

func main() {

	//declarar variables
	var clp float64
    var op int
    var res1 float64
    var res2 float64
    var res3 float64


	//Pedir ingreso de datos
    fmt.Print("INGRESE 1 (CLP A EUROS), 2(CLP A DOLAR), 3(CLP A YEN): ")
    fmt.Scan(&op)  
    fmt.Print("Ingrese una cantidad de pesos chilenos: ")
    fmt.Scan(&clp)   

    if (op == 1){
        res1 = clp * 0.0011
        fmt.Print("RESULTADO: ",res1," EUROS")
    }
    if (op == 2){
        res2 = clp * 0.0012
        fmt.Print("RESULTADO: ",res2," DOLARES")

    }
    if (op == 3){
        res3 = clp * 0.13
        fmt.Print("RESULTADO: ",res3," YENES")

    }
}